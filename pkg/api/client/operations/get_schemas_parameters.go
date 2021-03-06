// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetSchemasParams creates a new GetSchemasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSchemasParams() *GetSchemasParams {
	return &GetSchemasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemasParamsWithTimeout creates a new GetSchemasParams object
// with the ability to set a timeout on a request.
func NewGetSchemasParamsWithTimeout(timeout time.Duration) *GetSchemasParams {
	return &GetSchemasParams{
		timeout: timeout,
	}
}

// NewGetSchemasParamsWithContext creates a new GetSchemasParams object
// with the ability to set a context for a request.
func NewGetSchemasParamsWithContext(ctx context.Context) *GetSchemasParams {
	return &GetSchemasParams{
		Context: ctx,
	}
}

// NewGetSchemasParamsWithHTTPClient creates a new GetSchemasParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSchemasParamsWithHTTPClient(client *http.Client) *GetSchemasParams {
	return &GetSchemasParams{
		HTTPClient: client,
	}
}

/* GetSchemasParams contains all the parameters to send to the API endpoint
   for the get schemas operation.

   Typically these are written to a http.Request.
*/
type GetSchemasParams struct {

	// Deleted.
	Deleted *bool

	// LatestOnly.
	LatestOnly *bool

	// Limit.
	//
	// Format: int32
	// Default: -1
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// SubjectPrefix.
	SubjectPrefix *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schemas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchemasParams) WithDefaults() *GetSchemasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schemas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchemasParams) SetDefaults() {
	var (
		deletedDefault = bool(false)

		latestOnlyDefault = bool(false)

		limitDefault = int32(-1)

		offsetDefault = int32(0)
	)

	val := GetSchemasParams{
		Deleted:    &deletedDefault,
		LatestOnly: &latestOnlyDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get schemas params
func (o *GetSchemasParams) WithTimeout(timeout time.Duration) *GetSchemasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schemas params
func (o *GetSchemasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schemas params
func (o *GetSchemasParams) WithContext(ctx context.Context) *GetSchemasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schemas params
func (o *GetSchemasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schemas params
func (o *GetSchemasParams) WithHTTPClient(client *http.Client) *GetSchemasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schemas params
func (o *GetSchemasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the get schemas params
func (o *GetSchemasParams) WithDeleted(deleted *bool) *GetSchemasParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get schemas params
func (o *GetSchemasParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithLatestOnly adds the latestOnly to the get schemas params
func (o *GetSchemasParams) WithLatestOnly(latestOnly *bool) *GetSchemasParams {
	o.SetLatestOnly(latestOnly)
	return o
}

// SetLatestOnly adds the latestOnly to the get schemas params
func (o *GetSchemasParams) SetLatestOnly(latestOnly *bool) {
	o.LatestOnly = latestOnly
}

// WithLimit adds the limit to the get schemas params
func (o *GetSchemasParams) WithLimit(limit *int32) *GetSchemasParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get schemas params
func (o *GetSchemasParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get schemas params
func (o *GetSchemasParams) WithOffset(offset *int32) *GetSchemasParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get schemas params
func (o *GetSchemasParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSubjectPrefix adds the subjectPrefix to the get schemas params
func (o *GetSchemasParams) WithSubjectPrefix(subjectPrefix *string) *GetSchemasParams {
	o.SetSubjectPrefix(subjectPrefix)
	return o
}

// SetSubjectPrefix adds the subjectPrefix to the get schemas params
func (o *GetSchemasParams) SetSubjectPrefix(subjectPrefix *string) {
	o.SubjectPrefix = subjectPrefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	if o.LatestOnly != nil {

		// query param latestOnly
		var qrLatestOnly bool

		if o.LatestOnly != nil {
			qrLatestOnly = *o.LatestOnly
		}
		qLatestOnly := swag.FormatBool(qrLatestOnly)
		if qLatestOnly != "" {

			if err := r.SetQueryParam("latestOnly", qLatestOnly); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.SubjectPrefix != nil {

		// query param subjectPrefix
		var qrSubjectPrefix string

		if o.SubjectPrefix != nil {
			qrSubjectPrefix = *o.SubjectPrefix
		}
		qSubjectPrefix := qrSubjectPrefix
		if qSubjectPrefix != "" {

			if err := r.SetQueryParam("subjectPrefix", qSubjectPrefix); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
