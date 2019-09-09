// Code generated by go-swagger; DO NOT EDIT.

package postgre_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new postgre sql API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for postgre sql API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddPostgreSQL adds postgre SQL adds postgre SQL service and starts postgres exporter it automatically adds a service to inventory which is running on provided node id then adds postgres exporter with provided pmm agent id and other parameters
*/
func (a *Client) AddPostgreSQL(params *AddPostgreSQLParams) (*AddPostgreSQLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPostgreSQLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddPostgreSQL",
		Method:             "POST",
		PathPattern:        "/v0/management/PostgreSQL/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPostgreSQLReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPostgreSQLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddPostgreSQLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
