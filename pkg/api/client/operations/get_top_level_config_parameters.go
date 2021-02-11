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
)

// NewGetTopLevelConfigParams creates a new GetTopLevelConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTopLevelConfigParams() *GetTopLevelConfigParams {
	return &GetTopLevelConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTopLevelConfigParamsWithTimeout creates a new GetTopLevelConfigParams object
// with the ability to set a timeout on a request.
func NewGetTopLevelConfigParamsWithTimeout(timeout time.Duration) *GetTopLevelConfigParams {
	return &GetTopLevelConfigParams{
		timeout: timeout,
	}
}

// NewGetTopLevelConfigParamsWithContext creates a new GetTopLevelConfigParams object
// with the ability to set a context for a request.
func NewGetTopLevelConfigParamsWithContext(ctx context.Context) *GetTopLevelConfigParams {
	return &GetTopLevelConfigParams{
		Context: ctx,
	}
}

// NewGetTopLevelConfigParamsWithHTTPClient creates a new GetTopLevelConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTopLevelConfigParamsWithHTTPClient(client *http.Client) *GetTopLevelConfigParams {
	return &GetTopLevelConfigParams{
		HTTPClient: client,
	}
}

/* GetTopLevelConfigParams contains all the parameters to send to the API endpoint
   for the get top level config operation.

   Typically these are written to a http.Request.
*/
type GetTopLevelConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get top level config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTopLevelConfigParams) WithDefaults() *GetTopLevelConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get top level config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTopLevelConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get top level config params
func (o *GetTopLevelConfigParams) WithTimeout(timeout time.Duration) *GetTopLevelConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get top level config params
func (o *GetTopLevelConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get top level config params
func (o *GetTopLevelConfigParams) WithContext(ctx context.Context) *GetTopLevelConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get top level config params
func (o *GetTopLevelConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get top level config params
func (o *GetTopLevelConfigParams) WithHTTPClient(client *http.Client) *GetTopLevelConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get top level config params
func (o *GetTopLevelConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTopLevelConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
