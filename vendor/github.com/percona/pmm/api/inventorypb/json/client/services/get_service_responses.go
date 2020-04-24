// Code generated by go-swagger; DO NOT EDIT.

package services

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

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*GetServiceOK handles this case with default header values.

A successful response.
*/
type GetServiceOK struct {
	Payload *GetServiceOKBody
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/Get][%d] getServiceOk  %+v", 200, o.Payload)
}

func (o *GetServiceOK) GetPayload() *GetServiceOKBody {
	return o.Payload
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceDefault creates a GetServiceDefault with default headers values
func NewGetServiceDefault(code int) *GetServiceDefault {
	return &GetServiceDefault{
		_statusCode: code,
	}
}

/*GetServiceDefault handles this case with default header values.

An unexpected error response
*/
type GetServiceDefault struct {
	_statusCode int

	Payload *GetServiceDefaultBody
}

// Code gets the status code for the get service default response
func (o *GetServiceDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/Get][%d] GetService default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceDefault) GetPayload() *GetServiceDefaultBody {
	return o.Payload
}

func (o *GetServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetServiceBody get service body
swagger:model GetServiceBody
*/
type GetServiceBody struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this get service body
func (o *GetServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceBody) UnmarshalBinary(b []byte) error {
	var res GetServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceDefaultBody get service default body
swagger:model GetServiceDefaultBody
*/
type GetServiceDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get service default body
func (o *GetServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("GetService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBody get service OK body
swagger:model GetServiceOKBody
*/
type GetServiceOKBody struct {

	// external
	External *GetServiceOKBodyExternal `json:"external,omitempty"`

	// mongodb
	Mongodb *GetServiceOKBodyMongodb `json:"mongodb,omitempty"`

	// mysql
	Mysql *GetServiceOKBodyMysql `json:"mysql,omitempty"`

	// postgresql
	Postgresql *GetServiceOKBodyPostgresql `json:"postgresql,omitempty"`

	// proxysql
	Proxysql *GetServiceOKBodyProxysql `json:"proxysql,omitempty"`
}

// Validate validates this get service OK body
func (o *GetServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMongodb(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProxysql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetServiceOKBody) validateExternal(formats strfmt.Registry) error {

	if swag.IsZero(o.External) { // not required
		return nil
	}

	if o.External != nil {
		if err := o.External.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "external")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validateMongodb(formats strfmt.Registry) error {

	if swag.IsZero(o.Mongodb) { // not required
		return nil
	}

	if o.Mongodb != nil {
		if err := o.Mongodb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "mongodb")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	if o.Mysql != nil {
		if err := o.Mysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "mysql")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validatePostgresql(formats strfmt.Registry) error {

	if swag.IsZero(o.Postgresql) { // not required
		return nil
	}

	if o.Postgresql != nil {
		if err := o.Postgresql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "postgresql")
			}
			return err
		}
	}

	return nil
}

func (o *GetServiceOKBody) validateProxysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Proxysql) { // not required
		return nil
	}

	if o.Proxysql != nil {
		if err := o.Proxysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getServiceOk" + "." + "proxysql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBody) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyExternal ExternalService represents a generic External service instance.
swagger:model GetServiceOKBodyExternal
*/
type GetServiceOKBodyExternal struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this service instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get service OK body external
func (o *GetServiceOKBodyExternal) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyExternal) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyExternal) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyExternal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyMongodb MongoDBService represents a generic MongoDB instance.
swagger:model GetServiceOKBodyMongodb
*/
type GetServiceOKBodyMongodb struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get service OK body mongodb
func (o *GetServiceOKBodyMongodb) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyMongodb) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyMongodb) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyMongodb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyMysql MySQLService represents a generic MySQL instance.
swagger:model GetServiceOKBodyMysql
*/
type GetServiceOKBodyMysql struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Access unix socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get service OK body mysql
func (o *GetServiceOKBodyMysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyMysql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyPostgresql PostgreSQLService represents a generic PostgreSQL instance.
swagger:model GetServiceOKBodyPostgresql
*/
type GetServiceOKBodyPostgresql struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get service OK body postgresql
func (o *GetServiceOKBodyPostgresql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyPostgresql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyPostgresql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyPostgresql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetServiceOKBodyProxysql ProxySQLService represents a generic ProxySQL instance.
swagger:model GetServiceOKBodyProxysql
*/
type GetServiceOKBodyProxysql struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this get service OK body proxysql
func (o *GetServiceOKBodyProxysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetServiceOKBodyProxysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetServiceOKBodyProxysql) UnmarshalBinary(b []byte) error {
	var res GetServiceOKBodyProxysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
