// Code generated by go-swagger; DO NOT EDIT.

package correlation_rules

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

// NewCombinedRulesGetV1Params creates a new CombinedRulesGetV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCombinedRulesGetV1Params() *CombinedRulesGetV1Params {
	return &CombinedRulesGetV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCombinedRulesGetV1ParamsWithTimeout creates a new CombinedRulesGetV1Params object
// with the ability to set a timeout on a request.
func NewCombinedRulesGetV1ParamsWithTimeout(timeout time.Duration) *CombinedRulesGetV1Params {
	return &CombinedRulesGetV1Params{
		timeout: timeout,
	}
}

// NewCombinedRulesGetV1ParamsWithContext creates a new CombinedRulesGetV1Params object
// with the ability to set a context for a request.
func NewCombinedRulesGetV1ParamsWithContext(ctx context.Context) *CombinedRulesGetV1Params {
	return &CombinedRulesGetV1Params{
		Context: ctx,
	}
}

// NewCombinedRulesGetV1ParamsWithHTTPClient creates a new CombinedRulesGetV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewCombinedRulesGetV1ParamsWithHTTPClient(client *http.Client) *CombinedRulesGetV1Params {
	return &CombinedRulesGetV1Params{
		HTTPClient: client,
	}
}

/*
CombinedRulesGetV1Params contains all the parameters to send to the API endpoint

	for the combined rules get v1 operation.

	Typically these are written to a http.Request.
*/
type CombinedRulesGetV1Params struct {

	/* Filter.

	   FQL query specifying the filter parameters
	*/
	Filter *string

	/* Limit.

	   Number of IDs to return

	   Default: 100
	*/
	Limit *int64

	/* Offset.

	   Starting index of overall result set from which to return IDs
	*/
	Offset *int64

	/* Q.

	   Match query criteria, which includes all the filter string fields
	*/
	Q *string

	/* Sort.

	   Rule property to sort on

	   Default: "created_on"
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the combined rules get v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedRulesGetV1Params) WithDefaults() *CombinedRulesGetV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the combined rules get v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedRulesGetV1Params) SetDefaults() {
	var (
		limitDefault = int64(100)

		offsetDefault = int64(0)

		sortDefault = string("created_on")
	)

	val := CombinedRulesGetV1Params{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		Sort:   &sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithTimeout(timeout time.Duration) *CombinedRulesGetV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithContext(ctx context.Context) *CombinedRulesGetV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithHTTPClient(client *http.Client) *CombinedRulesGetV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithFilter(filter *string) *CombinedRulesGetV1Params {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithLimit(limit *int64) *CombinedRulesGetV1Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithOffset(offset *int64) *CombinedRulesGetV1Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithQ(q *string) *CombinedRulesGetV1Params {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetQ(q *string) {
	o.Q = q
}

// WithSort adds the sort to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) WithSort(sort *string) *CombinedRulesGetV1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the combined rules get v1 params
func (o *CombinedRulesGetV1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *CombinedRulesGetV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
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
