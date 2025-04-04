// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// NewListObjectsByVersionParams creates a new ListObjectsByVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListObjectsByVersionParams() *ListObjectsByVersionParams {
	return &ListObjectsByVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListObjectsByVersionParamsWithTimeout creates a new ListObjectsByVersionParams object
// with the ability to set a timeout on a request.
func NewListObjectsByVersionParamsWithTimeout(timeout time.Duration) *ListObjectsByVersionParams {
	return &ListObjectsByVersionParams{
		timeout: timeout,
	}
}

// NewListObjectsByVersionParamsWithContext creates a new ListObjectsByVersionParams object
// with the ability to set a context for a request.
func NewListObjectsByVersionParamsWithContext(ctx context.Context) *ListObjectsByVersionParams {
	return &ListObjectsByVersionParams{
		Context: ctx,
	}
}

// NewListObjectsByVersionParamsWithHTTPClient creates a new ListObjectsByVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewListObjectsByVersionParamsWithHTTPClient(client *http.Client) *ListObjectsByVersionParams {
	return &ListObjectsByVersionParams{
		HTTPClient: client,
	}
}

/*
ListObjectsByVersionParams contains all the parameters to send to the API endpoint

	for the list objects by version operation.

	Typically these are written to a http.Request.
*/
type ListObjectsByVersionParams struct {

	/* CollectionName.

	   The name of the collection
	*/
	CollectionName string

	/* CollectionVersion.

	   The version of the collection
	*/
	CollectionVersion string

	/* End.

	   The end key to end listing to
	*/
	End string

	/* Limit.

	   The limit of results to return
	*/
	Limit int64

	/* Start.

	   The start key to start listing from
	*/
	Start string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list objects by version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListObjectsByVersionParams) WithDefaults() *ListObjectsByVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list objects by version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListObjectsByVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list objects by version params
func (o *ListObjectsByVersionParams) WithTimeout(timeout time.Duration) *ListObjectsByVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list objects by version params
func (o *ListObjectsByVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list objects by version params
func (o *ListObjectsByVersionParams) WithContext(ctx context.Context) *ListObjectsByVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list objects by version params
func (o *ListObjectsByVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list objects by version params
func (o *ListObjectsByVersionParams) WithHTTPClient(client *http.Client) *ListObjectsByVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list objects by version params
func (o *ListObjectsByVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionName adds the collectionName to the list objects by version params
func (o *ListObjectsByVersionParams) WithCollectionName(collectionName string) *ListObjectsByVersionParams {
	o.SetCollectionName(collectionName)
	return o
}

// SetCollectionName adds the collectionName to the list objects by version params
func (o *ListObjectsByVersionParams) SetCollectionName(collectionName string) {
	o.CollectionName = collectionName
}

// WithCollectionVersion adds the collectionVersion to the list objects by version params
func (o *ListObjectsByVersionParams) WithCollectionVersion(collectionVersion string) *ListObjectsByVersionParams {
	o.SetCollectionVersion(collectionVersion)
	return o
}

// SetCollectionVersion adds the collectionVersion to the list objects by version params
func (o *ListObjectsByVersionParams) SetCollectionVersion(collectionVersion string) {
	o.CollectionVersion = collectionVersion
}

// WithEnd adds the end to the list objects by version params
func (o *ListObjectsByVersionParams) WithEnd(end string) *ListObjectsByVersionParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the list objects by version params
func (o *ListObjectsByVersionParams) SetEnd(end string) {
	o.End = end
}

// WithLimit adds the limit to the list objects by version params
func (o *ListObjectsByVersionParams) WithLimit(limit int64) *ListObjectsByVersionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list objects by version params
func (o *ListObjectsByVersionParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithStart adds the start to the list objects by version params
func (o *ListObjectsByVersionParams) WithStart(start string) *ListObjectsByVersionParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list objects by version params
func (o *ListObjectsByVersionParams) SetStart(start string) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListObjectsByVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection_name
	if err := r.SetPathParam("collection_name", o.CollectionName); err != nil {
		return err
	}

	// path param collection_version
	if err := r.SetPathParam("collection_version", o.CollectionVersion); err != nil {
		return err
	}

	// query param end
	qrEnd := o.End
	qEnd := qrEnd

	if err := r.SetQueryParam("end", qEnd); err != nil {
		return err
	}

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)

	if err := r.SetQueryParam("limit", qLimit); err != nil {
		return err
	}

	// query param start
	qrStart := o.Start
	qStart := qrStart

	if err := r.SetQueryParam("start", qStart); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
