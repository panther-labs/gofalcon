// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_protection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewNodesByCloudCountParams creates a new NodesByCloudCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodesByCloudCountParams() *NodesByCloudCountParams {
	return &NodesByCloudCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodesByCloudCountParamsWithTimeout creates a new NodesByCloudCountParams object
// with the ability to set a timeout on a request.
func NewNodesByCloudCountParamsWithTimeout(timeout time.Duration) *NodesByCloudCountParams {
	return &NodesByCloudCountParams{
		timeout: timeout,
	}
}

// NewNodesByCloudCountParamsWithContext creates a new NodesByCloudCountParams object
// with the ability to set a context for a request.
func NewNodesByCloudCountParamsWithContext(ctx context.Context) *NodesByCloudCountParams {
	return &NodesByCloudCountParams{
		Context: ctx,
	}
}

// NewNodesByCloudCountParamsWithHTTPClient creates a new NodesByCloudCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodesByCloudCountParamsWithHTTPClient(client *http.Client) *NodesByCloudCountParams {
	return &NodesByCloudCountParams{
		HTTPClient: client,
	}
}

/*
NodesByCloudCountParams contains all the parameters to send to the API endpoint

	for the nodes by cloud count operation.

	Typically these are written to a http.Request.
*/
type NodesByCloudCountParams struct {

	/* Filter.

	     Search Kubernetes nodes using a query in Falcon Query Language (FQL). Supported filter fields:
	- `agent_id`
	- `agent_type`
	- `annotations_list`
	- `cid`
	- `cloud_account_id`
	- `cloud_name`
	- `cloud_region`
	- `cloud_service`
	- `cluster_id`
	- `cluster_name`
	- `container_count`
	- `container_runtime_version`
	- `first_seen`
	- `image_digest`
	- `ipv4`
	- `kac_agent_id`
	- `last_seen`
	- `linux_sensor_coverage`
	- `node_name`
	- `pod_count`
	- `resource_status`
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nodes by cloud count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodesByCloudCountParams) WithDefaults() *NodesByCloudCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nodes by cloud count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodesByCloudCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the nodes by cloud count params
func (o *NodesByCloudCountParams) WithTimeout(timeout time.Duration) *NodesByCloudCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes by cloud count params
func (o *NodesByCloudCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes by cloud count params
func (o *NodesByCloudCountParams) WithContext(ctx context.Context) *NodesByCloudCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes by cloud count params
func (o *NodesByCloudCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes by cloud count params
func (o *NodesByCloudCountParams) WithHTTPClient(client *http.Client) *NodesByCloudCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes by cloud count params
func (o *NodesByCloudCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the nodes by cloud count params
func (o *NodesByCloudCountParams) WithFilter(filter *string) *NodesByCloudCountParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the nodes by cloud count params
func (o *NodesByCloudCountParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *NodesByCloudCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
