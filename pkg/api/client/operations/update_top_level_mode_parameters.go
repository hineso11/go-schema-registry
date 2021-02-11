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

	"github.com/hineso11/go-schema-registry/pkg/api/models"
)

// NewUpdateTopLevelModeParams creates a new UpdateTopLevelModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTopLevelModeParams() *UpdateTopLevelModeParams {
	return &UpdateTopLevelModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTopLevelModeParamsWithTimeout creates a new UpdateTopLevelModeParams object
// with the ability to set a timeout on a request.
func NewUpdateTopLevelModeParamsWithTimeout(timeout time.Duration) *UpdateTopLevelModeParams {
	return &UpdateTopLevelModeParams{
		timeout: timeout,
	}
}

// NewUpdateTopLevelModeParamsWithContext creates a new UpdateTopLevelModeParams object
// with the ability to set a context for a request.
func NewUpdateTopLevelModeParamsWithContext(ctx context.Context) *UpdateTopLevelModeParams {
	return &UpdateTopLevelModeParams{
		Context: ctx,
	}
}

// NewUpdateTopLevelModeParamsWithHTTPClient creates a new UpdateTopLevelModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTopLevelModeParamsWithHTTPClient(client *http.Client) *UpdateTopLevelModeParams {
	return &UpdateTopLevelModeParams{
		HTTPClient: client,
	}
}

/* UpdateTopLevelModeParams contains all the parameters to send to the API endpoint
   for the update top level mode operation.

   Typically these are written to a http.Request.
*/
type UpdateTopLevelModeParams struct {

	/* Body.

	   Update Request
	*/
	Body *models.ModeUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update top level mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTopLevelModeParams) WithDefaults() *UpdateTopLevelModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update top level mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTopLevelModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update top level mode params
func (o *UpdateTopLevelModeParams) WithTimeout(timeout time.Duration) *UpdateTopLevelModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update top level mode params
func (o *UpdateTopLevelModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update top level mode params
func (o *UpdateTopLevelModeParams) WithContext(ctx context.Context) *UpdateTopLevelModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update top level mode params
func (o *UpdateTopLevelModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update top level mode params
func (o *UpdateTopLevelModeParams) WithHTTPClient(client *http.Client) *UpdateTopLevelModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update top level mode params
func (o *UpdateTopLevelModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update top level mode params
func (o *UpdateTopLevelModeParams) WithBody(body *models.ModeUpdateRequest) *UpdateTopLevelModeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update top level mode params
func (o *UpdateTopLevelModeParams) SetBody(body *models.ModeUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTopLevelModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}