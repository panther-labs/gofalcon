// Code generated by go-swagger; DO NOT EDIT.

package container_images

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

// NewCombinedImageIssuesSummaryParams creates a new CombinedImageIssuesSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCombinedImageIssuesSummaryParams() *CombinedImageIssuesSummaryParams {
	return &CombinedImageIssuesSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCombinedImageIssuesSummaryParamsWithTimeout creates a new CombinedImageIssuesSummaryParams object
// with the ability to set a timeout on a request.
func NewCombinedImageIssuesSummaryParamsWithTimeout(timeout time.Duration) *CombinedImageIssuesSummaryParams {
	return &CombinedImageIssuesSummaryParams{
		timeout: timeout,
	}
}

// NewCombinedImageIssuesSummaryParamsWithContext creates a new CombinedImageIssuesSummaryParams object
// with the ability to set a context for a request.
func NewCombinedImageIssuesSummaryParamsWithContext(ctx context.Context) *CombinedImageIssuesSummaryParams {
	return &CombinedImageIssuesSummaryParams{
		Context: ctx,
	}
}

// NewCombinedImageIssuesSummaryParamsWithHTTPClient creates a new CombinedImageIssuesSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewCombinedImageIssuesSummaryParamsWithHTTPClient(client *http.Client) *CombinedImageIssuesSummaryParams {
	return &CombinedImageIssuesSummaryParams{
		HTTPClient: client,
	}
}

/*
CombinedImageIssuesSummaryParams contains all the parameters to send to the API endpoint

	for the combined image issues summary operation.

	Typically these are written to a http.Request.
*/
type CombinedImageIssuesSummaryParams struct {

	/* Cid.

	   CS Customer ID
	*/
	Cid string

	/* IncludeBaseImageVuln.

	   Include base image vulnerabilities.
	*/
	IncludeBaseImageVuln *bool

	/* Registry.

	   Registry
	*/
	Registry string

	/* Repository.

	   Repository name
	*/
	Repository string

	/* Tag.

	   Tag name
	*/
	Tag string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the combined image issues summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedImageIssuesSummaryParams) WithDefaults() *CombinedImageIssuesSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the combined image issues summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CombinedImageIssuesSummaryParams) SetDefaults() {
	var (
		includeBaseImageVulnDefault = bool(false)
	)

	val := CombinedImageIssuesSummaryParams{
		IncludeBaseImageVuln: &includeBaseImageVulnDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithTimeout(timeout time.Duration) *CombinedImageIssuesSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithContext(ctx context.Context) *CombinedImageIssuesSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithHTTPClient(client *http.Client) *CombinedImageIssuesSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithCid(cid string) *CombinedImageIssuesSummaryParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetCid(cid string) {
	o.Cid = cid
}

// WithIncludeBaseImageVuln adds the includeBaseImageVuln to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithIncludeBaseImageVuln(includeBaseImageVuln *bool) *CombinedImageIssuesSummaryParams {
	o.SetIncludeBaseImageVuln(includeBaseImageVuln)
	return o
}

// SetIncludeBaseImageVuln adds the includeBaseImageVuln to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetIncludeBaseImageVuln(includeBaseImageVuln *bool) {
	o.IncludeBaseImageVuln = includeBaseImageVuln
}

// WithRegistry adds the registry to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithRegistry(registry string) *CombinedImageIssuesSummaryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithRepository adds the repository to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithRepository(repository string) *CombinedImageIssuesSummaryParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTag adds the tag to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) WithTag(tag string) *CombinedImageIssuesSummaryParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the combined image issues summary params
func (o *CombinedImageIssuesSummaryParams) SetTag(tag string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *CombinedImageIssuesSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param cid
	qrCid := o.Cid
	qCid := qrCid
	if qCid != "" {

		if err := r.SetQueryParam("cid", qCid); err != nil {
			return err
		}
	}

	if o.IncludeBaseImageVuln != nil {

		// query param include_base_image_vuln
		var qrIncludeBaseImageVuln bool

		if o.IncludeBaseImageVuln != nil {
			qrIncludeBaseImageVuln = *o.IncludeBaseImageVuln
		}
		qIncludeBaseImageVuln := swag.FormatBool(qrIncludeBaseImageVuln)
		if qIncludeBaseImageVuln != "" {

			if err := r.SetQueryParam("include_base_image_vuln", qIncludeBaseImageVuln); err != nil {
				return err
			}
		}
	}

	// query param registry
	qrRegistry := o.Registry
	qRegistry := qrRegistry
	if qRegistry != "" {

		if err := r.SetQueryParam("registry", qRegistry); err != nil {
			return err
		}
	}

	// query param repository
	qrRepository := o.Repository
	qRepository := qrRepository
	if qRepository != "" {

		if err := r.SetQueryParam("repository", qRepository); err != nil {
			return err
		}
	}

	// query param tag
	qrTag := o.Tag
	qTag := qrTag
	if qTag != "" {

		if err := r.SetQueryParam("tag", qTag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
