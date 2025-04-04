// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeploymentsExternalV1(params *GetDeploymentsExternalV1Params, opts ...ClientOption) (*GetDeploymentsExternalV1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDeploymentsExternalV1 gets deployment resources by ids
*/
func (a *Client) GetDeploymentsExternalV1(params *GetDeploymentsExternalV1Params, opts ...ClientOption) (*GetDeploymentsExternalV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentsExternalV1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeploymentsExternalV1",
		Method:             "GET",
		PathPattern:        "/deployment-coordinator/entities/deployments/external/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentsExternalV1Reader{formats: a.formats},
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
	success, ok := result.(*GetDeploymentsExternalV1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeploymentsExternalV1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
