// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetActionReader is a Reader for the GetAction structure.
type GetActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetActionOK creates a GetActionOK with default headers values
func NewGetActionOK() *GetActionOK {
	return &GetActionOK{}
}

/*GetActionOK handles this case with default header values.

A successful response.
*/
type GetActionOK struct {
	Payload *GetActionOKBody
}

func (o *GetActionOK) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/Get][%d] getActionOk  %+v", 200, o.Payload)
}

func (o *GetActionOK) GetPayload() *GetActionOKBody {
	return o.Payload
}

func (o *GetActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionDefault creates a GetActionDefault with default headers values
func NewGetActionDefault(code int) *GetActionDefault {
	return &GetActionDefault{
		_statusCode: code,
	}
}

/*GetActionDefault handles this case with default header values.

An error response.
*/
type GetActionDefault struct {
	_statusCode int

	Payload *GetActionDefaultBody
}

// Code gets the status code for the get action default response
func (o *GetActionDefault) Code() int {
	return o._statusCode
}

func (o *GetActionDefault) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/Get][%d] GetAction default  %+v", o._statusCode, o.Payload)
}

func (o *GetActionDefault) GetPayload() *GetActionDefaultBody {
	return o.Payload
}

func (o *GetActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetActionBody get action body
swagger:model GetActionBody
*/
type GetActionBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`
}

// Validate validates this get action body
func (o *GetActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetActionBody) UnmarshalBinary(b []byte) error {
	var res GetActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetActionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model GetActionDefaultBody
*/
type GetActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get action default body
func (o *GetActionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetActionOKBody get action OK body
swagger:model GetActionOKBody
*/
type GetActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// True if Action is finished.
	Done bool `json:"done,omitempty"`

	// Error message if Action failed.
	Error string `json:"error,omitempty"`

	// Current Action output; may be partial if Action is still running.
	Output string `json:"output,omitempty"`

	// pmm-agent ID where this Action is running / was run.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this get action OK body
func (o *GetActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetActionOKBody) UnmarshalBinary(b []byte) error {
	var res GetActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
