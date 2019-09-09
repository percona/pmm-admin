// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// VersionReader is a Reader for the Version structure.
type VersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionOK creates a VersionOK with default headers values
func NewVersionOK() *VersionOK {
	return &VersionOK{}
}

/*VersionOK handles this case with default header values.

A successful response.
*/
type VersionOK struct {
	Payload *VersionOKBody
}

func (o *VersionOK) Error() string {
	return fmt.Sprintf("[GET /v1/version][%d] versionOk  %+v", 200, o.Payload)
}

func (o *VersionOK) GetPayload() *VersionOKBody {
	return o.Payload
}

func (o *VersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionDefault creates a VersionDefault with default headers values
func NewVersionDefault(code int) *VersionDefault {
	return &VersionDefault{
		_statusCode: code,
	}
}

/*VersionDefault handles this case with default header values.

An error response.
*/
type VersionDefault struct {
	_statusCode int

	Payload *VersionDefaultBody
}

// Code gets the status code for the version default response
func (o *VersionDefault) Code() int {
	return o._statusCode
}

func (o *VersionDefault) Error() string {
	return fmt.Sprintf("[GET /v1/version][%d] Version default  %+v", o._statusCode, o.Payload)
}

func (o *VersionDefault) GetPayload() *VersionDefaultBody {
	return o.Payload
}

func (o *VersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(VersionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*VersionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model VersionDefaultBody
*/
type VersionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this version default body
func (o *VersionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *VersionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionDefaultBody) UnmarshalBinary(b []byte) error {
	var res VersionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionOKBody version OK body
swagger:model VersionOKBody
*/
type VersionOKBody struct {

	// managed
	Managed *VersionOKBodyManaged `json:"managed,omitempty"`

	// server
	Server *VersionOKBodyServer `json:"server,omitempty"`

	// PMM Server version.
	Version string `json:"version,omitempty"`
}

// Validate validates this version OK body
func (o *VersionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateManaged(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionOKBody) validateManaged(formats strfmt.Registry) error {

	if swag.IsZero(o.Managed) { // not required
		return nil
	}

	if o.Managed != nil {
		if err := o.Managed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionOk" + "." + "managed")
			}
			return err
		}
	}

	return nil
}

func (o *VersionOKBody) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(o.Server) { // not required
		return nil
	}

	if o.Server != nil {
		if err := o.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("versionOk" + "." + "server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionOKBody) UnmarshalBinary(b []byte) error {
	var res VersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionOKBodyManaged VersionInfo describes component version, or PMM Server as a whole.
swagger:model VersionOKBodyManaged
*/
type VersionOKBodyManaged struct {

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// User-visible version.
	Version string `json:"version,omitempty"`
}

// Validate validates this version OK body managed
func (o *VersionOKBodyManaged) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionOKBodyManaged) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("versionOk"+"."+"managed"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionOKBodyManaged) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionOKBodyManaged) UnmarshalBinary(b []byte) error {
	var res VersionOKBodyManaged
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*VersionOKBodyServer VersionInfo describes component version, or PMM Server as a whole.
swagger:model VersionOKBodyServer
*/
type VersionOKBodyServer struct {

	// Full version for debugging.
	FullVersion string `json:"full_version,omitempty"`

	// Build or release date.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// User-visible version.
	Version string `json:"version,omitempty"`
}

// Validate validates this version OK body server
func (o *VersionOKBodyServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VersionOKBodyServer) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("versionOk"+"."+"server"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *VersionOKBodyServer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VersionOKBodyServer) UnmarshalBinary(b []byte) error {
	var res VersionOKBodyServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
