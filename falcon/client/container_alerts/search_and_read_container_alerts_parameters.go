// Code generated by go-swagger; DO NOT EDIT.

package container_alerts

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

// NewSearchAndReadContainerAlertsParams creates a new SearchAndReadContainerAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAndReadContainerAlertsParams() *SearchAndReadContainerAlertsParams {
	return &SearchAndReadContainerAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAndReadContainerAlertsParamsWithTimeout creates a new SearchAndReadContainerAlertsParams object
// with the ability to set a timeout on a request.
func NewSearchAndReadContainerAlertsParamsWithTimeout(timeout time.Duration) *SearchAndReadContainerAlertsParams {
	return &SearchAndReadContainerAlertsParams{
		timeout: timeout,
	}
}

// NewSearchAndReadContainerAlertsParamsWithContext creates a new SearchAndReadContainerAlertsParams object
// with the ability to set a context for a request.
func NewSearchAndReadContainerAlertsParamsWithContext(ctx context.Context) *SearchAndReadContainerAlertsParams {
	return &SearchAndReadContainerAlertsParams{
		Context: ctx,
	}
}

// NewSearchAndReadContainerAlertsParamsWithHTTPClient creates a new SearchAndReadContainerAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAndReadContainerAlertsParamsWithHTTPClient(client *http.Client) *SearchAndReadContainerAlertsParams {
	return &SearchAndReadContainerAlertsParams{
		HTTPClient: client,
	}
}

/*
SearchAndReadContainerAlertsParams contains all the parameters to send to the API endpoint

	for the search and read container alerts operation.

	Typically these are written to a http.Request.
*/
type SearchAndReadContainerAlertsParams struct {

	/* Filter.

	     Search Container Alerts using a query in Falcon Query Language (FQL). Supported filter fields:
	- `cid`
	- `container_id`
	- `last_seen`
	- `name`
	- `severity`
	*/
	Filter *string

	/* Limit.

	   The upper-bound on the number of records to retrieve.

	   Default: 100
	*/
	Limit *int64

	/* Offset.

	   The offset from where to begin.
	*/
	Offset *int64

	/* Sort.

	   The fields to sort the records on.
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search and read container alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAndReadContainerAlertsParams) WithDefaults() *SearchAndReadContainerAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search and read container alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAndReadContainerAlertsParams) SetDefaults() {
	var (
		limitDefault = int64(100)
	)

	val := SearchAndReadContainerAlertsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithTimeout(timeout time.Duration) *SearchAndReadContainerAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithContext(ctx context.Context) *SearchAndReadContainerAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithHTTPClient(client *http.Client) *SearchAndReadContainerAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithFilter(filter *string) *SearchAndReadContainerAlertsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithLimit(limit *int64) *SearchAndReadContainerAlertsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithOffset(offset *int64) *SearchAndReadContainerAlertsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSort adds the sort to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) WithSort(sort *string) *SearchAndReadContainerAlertsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search and read container alerts params
func (o *SearchAndReadContainerAlertsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAndReadContainerAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
