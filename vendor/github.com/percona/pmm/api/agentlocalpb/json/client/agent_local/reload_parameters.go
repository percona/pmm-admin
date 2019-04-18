// Code generated by go-swagger; DO NOT EDIT.

package agent_local

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewReloadParams creates a new ReloadParams object
// with the default values initialized.
func NewReloadParams() *ReloadParams {
	var ()
	return &ReloadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReloadParamsWithTimeout creates a new ReloadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReloadParamsWithTimeout(timeout time.Duration) *ReloadParams {
	var ()
	return &ReloadParams{

		timeout: timeout,
	}
}

// NewReloadParamsWithContext creates a new ReloadParams object
// with the default values initialized, and the ability to set a context for a request
func NewReloadParamsWithContext(ctx context.Context) *ReloadParams {
	var ()
	return &ReloadParams{

		Context: ctx,
	}
}

// NewReloadParamsWithHTTPClient creates a new ReloadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReloadParamsWithHTTPClient(client *http.Client) *ReloadParams {
	var ()
	return &ReloadParams{
		HTTPClient: client,
	}
}

/*ReloadParams contains all the parameters to send to the API endpoint
for the reload operation typically these are written to a http.Request
*/
type ReloadParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reload params
func (o *ReloadParams) WithTimeout(timeout time.Duration) *ReloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reload params
func (o *ReloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reload params
func (o *ReloadParams) WithContext(ctx context.Context) *ReloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reload params
func (o *ReloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reload params
func (o *ReloadParams) WithHTTPClient(client *http.Client) *ReloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reload params
func (o *ReloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reload params
func (o *ReloadParams) WithBody(body interface{}) *ReloadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reload params
func (o *ReloadParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ReloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
