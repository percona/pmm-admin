// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// CheckNodeReader is a Reader for the CheckNode structure.
type CheckNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckNodeOK creates a CheckNodeOK with default headers values
func NewCheckNodeOK() *CheckNodeOK {
	return &CheckNodeOK{}
}

/*CheckNodeOK handles this case with default header values.

A successful response.
*/
type CheckNodeOK struct {
	Payload *CheckNodeOKBody
}

func (o *CheckNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Check][%d] checkNodeOk  %+v", 200, o.Payload)
}

func (o *CheckNodeOK) GetPayload() *CheckNodeOKBody {
	return o.Payload
}

func (o *CheckNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckNodeDefault creates a CheckNodeDefault with default headers values
func NewCheckNodeDefault(code int) *CheckNodeDefault {
	return &CheckNodeDefault{
		_statusCode: code,
	}
}

/*CheckNodeDefault handles this case with default header values.

An unexpected error response
*/
type CheckNodeDefault struct {
	_statusCode int

	Payload *CheckNodeDefaultBody
}

// Code gets the status code for the check node default response
func (o *CheckNodeDefault) Code() int {
	return o._statusCode
}

func (o *CheckNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Check][%d] CheckNode default  %+v", o._statusCode, o.Payload)
}

func (o *CheckNodeDefault) GetPayload() *CheckNodeDefaultBody {
	return o.Payload
}

func (o *CheckNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CheckNodeBody check node body
swagger:model CheckNodeBody
*/
type CheckNodeBody struct {

	// node id
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this check node body
func (o *CheckNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckNodeBody) UnmarshalBinary(b []byte) error {
	var res CheckNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckNodeDefaultBody check node default body
swagger:model CheckNodeDefaultBody
*/
type CheckNodeDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this check node default body
func (o *CheckNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CheckNodeDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CheckNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CheckNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res CheckNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CheckNodeOKBody check node OK body
swagger:model CheckNodeOKBody
*/
type CheckNodeOKBody struct {

	// exists
	Exists bool `json:"exists,omitempty"`
}

// Validate validates this check node OK body
func (o *CheckNodeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckNodeOKBody) UnmarshalBinary(b []byte) error {
	var res CheckNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
