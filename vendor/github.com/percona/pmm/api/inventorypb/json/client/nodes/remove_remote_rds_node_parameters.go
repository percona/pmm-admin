// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewRemoveRemoteRDSNodeParams creates a new RemoveRemoteRDSNodeParams object
// with the default values initialized.
func NewRemoveRemoteRDSNodeParams() *RemoveRemoteRDSNodeParams {
	var ()
	return &RemoveRemoteRDSNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveRemoteRDSNodeParamsWithTimeout creates a new RemoveRemoteRDSNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveRemoteRDSNodeParamsWithTimeout(timeout time.Duration) *RemoveRemoteRDSNodeParams {
	var ()
	return &RemoveRemoteRDSNodeParams{

		timeout: timeout,
	}
}

// NewRemoveRemoteRDSNodeParamsWithContext creates a new RemoveRemoteRDSNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveRemoteRDSNodeParamsWithContext(ctx context.Context) *RemoveRemoteRDSNodeParams {
	var ()
	return &RemoveRemoteRDSNodeParams{

		Context: ctx,
	}
}

// NewRemoveRemoteRDSNodeParamsWithHTTPClient creates a new RemoveRemoteRDSNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveRemoteRDSNodeParamsWithHTTPClient(client *http.Client) *RemoveRemoteRDSNodeParams {
	var ()
	return &RemoveRemoteRDSNodeParams{
		HTTPClient: client,
	}
}

/*RemoveRemoteRDSNodeParams contains all the parameters to send to the API endpoint
for the remove remote RDS node operation typically these are written to a http.Request
*/
type RemoveRemoteRDSNodeParams struct {

	/*Body*/
	Body RemoveRemoteRDSNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) WithTimeout(timeout time.Duration) *RemoveRemoteRDSNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) WithContext(ctx context.Context) *RemoveRemoteRDSNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) WithHTTPClient(client *http.Client) *RemoveRemoteRDSNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) WithBody(body RemoveRemoteRDSNodeBody) *RemoveRemoteRDSNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove remote RDS node params
func (o *RemoveRemoteRDSNodeParams) SetBody(body RemoveRemoteRDSNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveRemoteRDSNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
