package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client"
	"github.com/crowdstrike/gofalcon/falcon/client/event_streams"
	"github.com/crowdstrike/gofalcon/falcon/models"
	"github.com/crowdstrike/gofalcon/falcon/models/streaming_models"
	"github.com/crowdstrike/gofalcon/pkg/falcon_util"
)

func main() {
	clientId := flag.String("client-id", os.Getenv("FALCON_CLIENT_ID"), "Client ID for accessing CrowdStrike Falcon Platform (default taken from FALCON_CLIENT_ID env)")
	clientSecret := flag.String("client-secret", os.Getenv("FALCON_CLIENT_SECRET"), "Client Secret for accessing CrowdStrike Falcon Platform (default taken from FALCON_CLIENT_SECRET)")
	clientCloud := flag.String("cloud", os.Getenv("FALCON_CLOUD"), "Falcon cloud abbreviation (us-1, us-2, eu-1, us-gov-1)")

	flag.Parse()

	if *clientId == "" {
		*clientId = promptUser(`Missing FALCON_CLIENT_ID environment variable. Please provide your OAuth2 API Client ID for authentication with CrowdStrike Falcon platform. Establishing and retrieving OAuth2 API credentials can be performed at https://falcon.crowdstrike.com/support/api-clients-and-keys.
Falcon Client ID`)
	}
	if *clientSecret == "" {
		*clientSecret = promptUser(`Missing FALCON_CLIENT_SECRET environment variable. Please provide your OAuth2 API Client Secret for authentication with CrowdStrike Falcon platform. Establishing and retrieving OAuth2 API credentials can be performed at https://falcon.crowdstrike.com/support/api-clients-and-keys.
Falcon Client Secret`)
	}

	client, err := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     *clientId,
		ClientSecret: *clientSecret,
		Cloud:        falcon.Cloud(*clientCloud),
		Context:      context.Background(),
	})
	if err != nil {
		panic(err)
	}

	json := "json"
	appId := "falcon_event_stream"
	// Step 1: Discover Available Streams
	response, err := client.EventStreams.ListAvailableStreamsOAuth2(&event_streams.ListAvailableStreamsOAuth2Params{
		AppID:   appId,
		Format:  &json,
		Context: context.Background(),
	})
	if err != nil {
		panic(falcon.ErrorExplain(err))
	}
	for _, e := range response.Payload.Errors {
		fmt.Println(e)
	}

	availableStreams := response.Payload.Resources
	for _, availableStream := range availableStreams {
		stream, err := NewStream(context.Background(), client, appId, availableStream)
		if err != nil {
			panic(err)
		}
		defer stream.Close()

		var fatal error
		for fatal == nil {
			select {
			case err := <-stream.Errors:
				if err.Fatal {
					fatal = err.Err
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			case event := <-stream.Events:
				pretty, err := falcon_util.PrettyJson(event)
				if err != nil {
					panic(err)
				}
				fmt.Println(pretty)
			}
		}
		panic(fatal)
	}
}

func promptUser(prompt string) string {
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(s)
}

type StreamingHandle struct {
	ctx    context.Context
	client *client.CrowdStrikeAPISpecification
	appId  string
	stream *models.MainAvailableStreamV2
	Events chan *streaming_models.EventItem
	Errors chan StreamingError
}

func NewStream(ctx context.Context, client *client.CrowdStrikeAPISpecification, appId string, stream *models.MainAvailableStreamV2) (*StreamingHandle, error) {
	sh := &StreamingHandle{
		ctx:    ctx,
		client: client,
		appId:  appId,
		stream: stream,
		Events: make(chan *streaming_models.EventItem),
		Errors: make(chan StreamingError),
	}
	sh.maintainSession()
	err := sh.open()
	if err != nil {
		sh.Close()
		return nil, err
	}
	return sh, nil
}

func (sh *StreamingHandle) Close() {
	close(sh.Errors)
	sh.Errors = nil
}

type StreamingError struct {
	Fatal bool
	Err   error
}

func (e StreamingError) Error() string {
	return e.Err.Error()
}

func (sh *StreamingHandle) maintainSession() {
	ticker := time.NewTicker(time.Duration(*sh.stream.RefreshActiveSessionInterval*9/10) * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-sh.ctx.Done():
				return
			case <-ticker.C:
				_, err := sh.client.EventStreams.RefreshActiveStreamSession(&event_streams.RefreshActiveStreamSessionParams{
					AppID:      sh.appId,
					ActionName: "refresh_active_stream_session",
					Partition:  0,
					Context:    sh.ctx,
				})

				if err != nil {
					sh.Errors <- StreamingError{
						Fatal: true,
						Err:   err,
					}
					return
				}
			}
		}
	}()
}

func (sh *StreamingHandle) open() error {

	req, err := http.NewRequestWithContext(sh.ctx, "GET", *sh.stream.DataFeedURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Authorization", "Token "+*sh.stream.SessionToken.Token)
	req.Header.Add("Connection", "Keep-Alive")
	req.Header.Add("Date", time.Now().Format(time.RFC1123Z))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	sh.Events = make(chan *streaming_models.EventItem)
	go func() {
		defer resp.Body.Close()

		dec := json.NewDecoder(resp.Body)
		for dec.More() {
			var detection streaming_models.EventItem
			//dec.DisallowUnknownFields()
			err := dec.Decode(&detection)
			if err != nil {
				sh.Errors <- StreamingError{Fatal: false, Err: err}
			}
			sh.Events <- &detection
		}
		sh.Errors <- StreamingError{
			Fatal: true,
			Err:   errors.New("Streaming connection closed"),
		}
		close(sh.Events)
	}()

	return nil
}
