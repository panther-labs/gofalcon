// Code generated by go-swagger; DO NOT EDIT.

package a_s_p_m

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
	"github.com/go-openapi/swag"
)

// NewGetExecutorNodesMetadataParams creates a new GetExecutorNodesMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExecutorNodesMetadataParams() *GetExecutorNodesMetadataParams {
	return &GetExecutorNodesMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExecutorNodesMetadataParamsWithTimeout creates a new GetExecutorNodesMetadataParams object
// with the ability to set a timeout on a request.
func NewGetExecutorNodesMetadataParamsWithTimeout(timeout time.Duration) *GetExecutorNodesMetadataParams {
	return &GetExecutorNodesMetadataParams{
		timeout: timeout,
	}
}

// NewGetExecutorNodesMetadataParamsWithContext creates a new GetExecutorNodesMetadataParams object
// with the ability to set a context for a request.
func NewGetExecutorNodesMetadataParamsWithContext(ctx context.Context) *GetExecutorNodesMetadataParams {
	return &GetExecutorNodesMetadataParams{
		Context: ctx,
	}
}

// NewGetExecutorNodesMetadataParamsWithHTTPClient creates a new GetExecutorNodesMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExecutorNodesMetadataParamsWithHTTPClient(client *http.Client) *GetExecutorNodesMetadataParams {
	return &GetExecutorNodesMetadataParams{
		HTTPClient: client,
	}
}

/*
GetExecutorNodesMetadataParams contains all the parameters to send to the API endpoint

	for the get executor nodes metadata operation.

	Typically these are written to a http.Request.
*/
type GetExecutorNodesMetadataParams struct {

	/* ExecutorNodeIds.

	   executor node ids
	*/
	ExecutorNodeIds []string

	/* ExecutorNodeNames.

	   executor node names
	*/
	ExecutorNodeNames []string

	/* ExecutorNodeStates.

	   executor node states
	*/
	ExecutorNodeStates []int64

	/* ExecutorNodeTypes.

	   executor node types
	*/
	ExecutorNodeTypes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get executor nodes metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExecutorNodesMetadataParams) WithDefaults() *GetExecutorNodesMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get executor nodes metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExecutorNodesMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithTimeout(timeout time.Duration) *GetExecutorNodesMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithContext(ctx context.Context) *GetExecutorNodesMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithHTTPClient(client *http.Client) *GetExecutorNodesMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutorNodeIds adds the executorNodeIds to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithExecutorNodeIds(executorNodeIds []string) *GetExecutorNodesMetadataParams {
	o.SetExecutorNodeIds(executorNodeIds)
	return o
}

// SetExecutorNodeIds adds the executorNodeIds to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetExecutorNodeIds(executorNodeIds []string) {
	o.ExecutorNodeIds = executorNodeIds
}

// WithExecutorNodeNames adds the executorNodeNames to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithExecutorNodeNames(executorNodeNames []string) *GetExecutorNodesMetadataParams {
	o.SetExecutorNodeNames(executorNodeNames)
	return o
}

// SetExecutorNodeNames adds the executorNodeNames to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetExecutorNodeNames(executorNodeNames []string) {
	o.ExecutorNodeNames = executorNodeNames
}

// WithExecutorNodeStates adds the executorNodeStates to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithExecutorNodeStates(executorNodeStates []int64) *GetExecutorNodesMetadataParams {
	o.SetExecutorNodeStates(executorNodeStates)
	return o
}

// SetExecutorNodeStates adds the executorNodeStates to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetExecutorNodeStates(executorNodeStates []int64) {
	o.ExecutorNodeStates = executorNodeStates
}

// WithExecutorNodeTypes adds the executorNodeTypes to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) WithExecutorNodeTypes(executorNodeTypes []string) *GetExecutorNodesMetadataParams {
	o.SetExecutorNodeTypes(executorNodeTypes)
	return o
}

// SetExecutorNodeTypes adds the executorNodeTypes to the get executor nodes metadata params
func (o *GetExecutorNodesMetadataParams) SetExecutorNodeTypes(executorNodeTypes []string) {
	o.ExecutorNodeTypes = executorNodeTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetExecutorNodesMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExecutorNodeIds != nil {

		// binding items for executor_node_ids
		joinedExecutorNodeIds := o.bindParamExecutorNodeIds(reg)

		// query array param executor_node_ids
		if err := r.SetQueryParam("executor_node_ids", joinedExecutorNodeIds...); err != nil {
			return err
		}
	}

	if o.ExecutorNodeNames != nil {

		// binding items for executor_node_names
		joinedExecutorNodeNames := o.bindParamExecutorNodeNames(reg)

		// query array param executor_node_names
		if err := r.SetQueryParam("executor_node_names", joinedExecutorNodeNames...); err != nil {
			return err
		}
	}

	if o.ExecutorNodeStates != nil {

		// binding items for executor_node_states
		joinedExecutorNodeStates := o.bindParamExecutorNodeStates(reg)

		// query array param executor_node_states
		if err := r.SetQueryParam("executor_node_states", joinedExecutorNodeStates...); err != nil {
			return err
		}
	}

	if o.ExecutorNodeTypes != nil {

		// binding items for executor_node_types
		joinedExecutorNodeTypes := o.bindParamExecutorNodeTypes(reg)

		// query array param executor_node_types
		if err := r.SetQueryParam("executor_node_types", joinedExecutorNodeTypes...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetExecutorNodesMetadata binds the parameter executor_node_ids
func (o *GetExecutorNodesMetadataParams) bindParamExecutorNodeIds(formats strfmt.Registry) []string {
	executorNodeIdsIR := o.ExecutorNodeIds

	var executorNodeIdsIC []string
	for _, executorNodeIdsIIR := range executorNodeIdsIR { // explode []string

		executorNodeIdsIIV := executorNodeIdsIIR // string as string
		executorNodeIdsIC = append(executorNodeIdsIC, executorNodeIdsIIV)
	}

	// items.CollectionFormat: "csv"
	executorNodeIdsIS := swag.JoinByFormat(executorNodeIdsIC, "csv")

	return executorNodeIdsIS
}

// bindParamGetExecutorNodesMetadata binds the parameter executor_node_names
func (o *GetExecutorNodesMetadataParams) bindParamExecutorNodeNames(formats strfmt.Registry) []string {
	executorNodeNamesIR := o.ExecutorNodeNames

	var executorNodeNamesIC []string
	for _, executorNodeNamesIIR := range executorNodeNamesIR { // explode []string

		executorNodeNamesIIV := executorNodeNamesIIR // string as string
		executorNodeNamesIC = append(executorNodeNamesIC, executorNodeNamesIIV)
	}

	// items.CollectionFormat: "csv"
	executorNodeNamesIS := swag.JoinByFormat(executorNodeNamesIC, "csv")

	return executorNodeNamesIS
}

// bindParamGetExecutorNodesMetadata binds the parameter executor_node_states
func (o *GetExecutorNodesMetadataParams) bindParamExecutorNodeStates(formats strfmt.Registry) []string {
	executorNodeStatesIR := o.ExecutorNodeStates

	var executorNodeStatesIC []string
	for _, executorNodeStatesIIR := range executorNodeStatesIR { // explode []int64

		executorNodeStatesIIV := swag.FormatInt64(executorNodeStatesIIR) // int64 as string
		executorNodeStatesIC = append(executorNodeStatesIC, executorNodeStatesIIV)
	}

	// items.CollectionFormat: "csv"
	executorNodeStatesIS := swag.JoinByFormat(executorNodeStatesIC, "csv")

	return executorNodeStatesIS
}

// bindParamGetExecutorNodesMetadata binds the parameter executor_node_types
func (o *GetExecutorNodesMetadataParams) bindParamExecutorNodeTypes(formats strfmt.Registry) []string {
	executorNodeTypesIR := o.ExecutorNodeTypes

	var executorNodeTypesIC []string
	for _, executorNodeTypesIIR := range executorNodeTypesIR { // explode []string

		executorNodeTypesIIV := executorNodeTypesIIR // string as string
		executorNodeTypesIC = append(executorNodeTypesIC, executorNodeTypesIIV)
	}

	// items.CollectionFormat: "csv"
	executorNodeTypesIS := swag.JoinByFormat(executorNodeTypesIC, "csv")

	return executorNodeTypesIS
}
