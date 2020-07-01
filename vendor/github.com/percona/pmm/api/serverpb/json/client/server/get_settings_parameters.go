// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewGetSettingsParams creates a new GetSettingsParams object
// with the default values initialized.
func NewGetSettingsParams() *GetSettingsParams {
	var ()
	return &GetSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSettingsParamsWithTimeout creates a new GetSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSettingsParamsWithTimeout(timeout time.Duration) *GetSettingsParams {
	var ()
	return &GetSettingsParams{

		timeout: timeout,
	}
}

// NewGetSettingsParamsWithContext creates a new GetSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSettingsParamsWithContext(ctx context.Context) *GetSettingsParams {
	var ()
	return &GetSettingsParams{

		Context: ctx,
	}
}

// NewGetSettingsParamsWithHTTPClient creates a new GetSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSettingsParamsWithHTTPClient(client *http.Client) *GetSettingsParams {
	var ()
	return &GetSettingsParams{
		HTTPClient: client,
	}
}

/*GetSettingsParams contains all the parameters to send to the API endpoint
for the get settings operation typically these are written to a http.Request
*/
type GetSettingsParams struct {

	/*Body*/
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get settings params
func (o *GetSettingsParams) WithTimeout(timeout time.Duration) *GetSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get settings params
func (o *GetSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get settings params
func (o *GetSettingsParams) WithContext(ctx context.Context) *GetSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get settings params
func (o *GetSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get settings params
func (o *GetSettingsParams) WithHTTPClient(client *http.Client) *GetSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get settings params
func (o *GetSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get settings params
func (o *GetSettingsParams) WithBody(body interface{}) *GetSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get settings params
func (o *GetSettingsParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
