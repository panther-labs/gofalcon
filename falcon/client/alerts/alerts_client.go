// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAggregateV2(params *GetAggregateV2Params, opts ...ClientOption) (*GetAggregateV2OK, error)

	GetQueriesAlertsV1(params *GetQueriesAlertsV1Params, opts ...ClientOption) (*GetQueriesAlertsV1OK, error)

	GetV2(params *GetV2Params, opts ...ClientOption) (*GetV2OK, error)

	PatchEntitiesAlertsV2(params *PatchEntitiesAlertsV2Params, opts ...ClientOption) (*PatchEntitiesAlertsV2OK, error)

	PostAggregatesAlertsV1(params *PostAggregatesAlertsV1Params, opts ...ClientOption) (*PostAggregatesAlertsV1OK, error)

	PostCombinedAlertsV1(params *PostCombinedAlertsV1Params, opts ...ClientOption) (*PostCombinedAlertsV1OK, error)

	PostEntitiesAlertsV1(params *PostEntitiesAlertsV1Params, opts ...ClientOption) (*PostEntitiesAlertsV1OK, error)

	QueryV2(params *QueryV2Params, opts ...ClientOption) (*QueryV2OK, error)

	UpdateV3(params *UpdateV3Params, opts ...ClientOption) (*UpdateV3OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetAggregateV2 retrieves aggregate values for alerts across all c i ds
*/
func (a *Client) GetAggregateV2(params *GetAggregateV2Params, opts ...ClientOption) (*GetAggregateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAggregateV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAggregateV2",
		Method:             "POST",
		PathPattern:        "/alerts/aggregates/alerts/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAggregateV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAggregateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAggregateV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetQueriesAlertsV1 deprecateds please use version v2 of this endpoint retrieves all alerts ids that match a given query
*/
func (a *Client) GetQueriesAlertsV1(params *GetQueriesAlertsV1Params, opts ...ClientOption) (*GetQueriesAlertsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueriesAlertsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetQueriesAlertsV1",
		Method:             "GET",
		PathPattern:        "/alerts/queries/alerts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQueriesAlertsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetQueriesAlertsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetQueriesAlertsV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetV2 retrieves all alerts given their composite ids
*/
func (a *Client) GetV2(params *GetV2Params, opts ...ClientOption) (*GetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV2",
		Method:             "POST",
		PathPattern:        "/alerts/entities/alerts/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchEntitiesAlertsV2 deprecateds please use version v3 of this endpoint perform actions on alerts identified by composite ID s in request each action has a name and a description which describes what the action does if a request adds and removes tag in a single request the order of processing would be to remove tags before adding new ones in
*/
func (a *Client) PatchEntitiesAlertsV2(params *PatchEntitiesAlertsV2Params, opts ...ClientOption) (*PatchEntitiesAlertsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEntitiesAlertsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PatchEntitiesAlertsV2",
		Method:             "PATCH",
		PathPattern:        "/alerts/entities/alerts/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchEntitiesAlertsV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchEntitiesAlertsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchEntitiesAlertsV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAggregatesAlertsV1 deprecateds please use version v2 of this endpoint retrieves aggregate values for alerts across all c i ds
*/
func (a *Client) PostAggregatesAlertsV1(params *PostAggregatesAlertsV1Params, opts ...ClientOption) (*PostAggregatesAlertsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAggregatesAlertsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAggregatesAlertsV1",
		Method:             "POST",
		PathPattern:        "/alerts/aggregates/alerts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAggregatesAlertsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAggregatesAlertsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAggregatesAlertsV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostCombinedAlertsV1 retrieves all alerts that match a particular f q l filter this API is intended for retrieval of large amounts of alerts 10k using a pagination based on a after token if you need to use offset pagination consider using g e t alerts queries alerts and p o s t alerts entities alerts a p is
*/
func (a *Client) PostCombinedAlertsV1(params *PostCombinedAlertsV1Params, opts ...ClientOption) (*PostCombinedAlertsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCombinedAlertsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostCombinedAlertsV1",
		Method:             "POST",
		PathPattern:        "/alerts/combined/alerts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCombinedAlertsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCombinedAlertsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostCombinedAlertsV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostEntitiesAlertsV1 deprecateds please use version v2 of this endpoint retrieves all alerts given their ids
*/
func (a *Client) PostEntitiesAlertsV1(params *PostEntitiesAlertsV1Params, opts ...ClientOption) (*PostEntitiesAlertsV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEntitiesAlertsV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostEntitiesAlertsV1",
		Method:             "POST",
		PathPattern:        "/alerts/entities/alerts/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEntitiesAlertsV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostEntitiesAlertsV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostEntitiesAlertsV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryV2 retrieves all alerts ids that match a given query
*/
func (a *Client) QueryV2(params *QueryV2Params, opts ...ClientOption) (*QueryV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QueryV2",
		Method:             "GET",
		PathPattern:        "/alerts/queries/alerts/v2",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for QueryV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateV3 performs actions on alerts identified by composite ID s in request each action has a name and a description which describes what the action does if a request adds and removes tag in a single request the order of processing would be to remove tags before adding new ones in
*/
func (a *Client) UpdateV3(params *UpdateV3Params, opts ...ClientOption) (*UpdateV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateV3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateV3",
		Method:             "PATCH",
		PathPattern:        "/alerts/entities/alerts/v3",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateV3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateV3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
