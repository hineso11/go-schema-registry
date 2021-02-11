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

// NewListVersionsParams creates a new ListVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVersionsParams() *ListVersionsParams {
	return &ListVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVersionsParamsWithTimeout creates a new ListVersionsParams object
// with the ability to set a timeout on a request.
func NewListVersionsParamsWithTimeout(timeout time.Duration) *ListVersionsParams {
	return &ListVersionsParams{
		timeout: timeout,
	}
}

// NewListVersionsParamsWithContext creates a new ListVersionsParams object
// with the ability to set a context for a request.
func NewListVersionsParamsWithContext(ctx context.Context) *ListVersionsParams {
	return &ListVersionsParams{
		Context: ctx,
	}
}

// NewListVersionsParamsWithHTTPClient creates a new ListVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVersionsParamsWithHTTPClient(client *http.Client) *ListVersionsParams {
	return &ListVersionsParams{
		HTTPClient: client,
	}
}

/* ListVersionsParams contains all the parameters to send to the API endpoint
   for the list versions operation.

   Typically these are written to a http.Request.
*/
type ListVersionsParams struct {

	// Deleted.
	Deleted *bool

	/* Subject.

	   Name of the Subject
	*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionsParams) WithDefaults() *ListVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list versions params
func (o *ListVersionsParams) WithTimeout(timeout time.Duration) *ListVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list versions params
func (o *ListVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list versions params
func (o *ListVersionsParams) WithContext(ctx context.Context) *ListVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list versions params
func (o *ListVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list versions params
func (o *ListVersionsParams) WithHTTPClient(client *http.Client) *ListVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list versions params
func (o *ListVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the list versions params
func (o *ListVersionsParams) WithDeleted(deleted *bool) *ListVersionsParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the list versions params
func (o *ListVersionsParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithSubject adds the subject to the list versions params
func (o *ListVersionsParams) WithSubject(subject string) *ListVersionsParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the list versions params
func (o *ListVersionsParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *ListVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subject
	if err := r.SetPathParam("subject", o.Subject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
