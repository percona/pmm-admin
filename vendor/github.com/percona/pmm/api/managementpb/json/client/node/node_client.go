// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new node API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for node API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RegisterNode registers node registers a new node and pmm agent
*/
func (a *Client) RegisterNode(params *RegisterNodeParams) (*RegisterNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterNode",
		Method:             "POST",
		PathPattern:        "/v1/management/Node/Register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegisterNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
