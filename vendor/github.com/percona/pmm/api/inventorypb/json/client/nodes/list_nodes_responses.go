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

// ListNodesReader is a Reader for the ListNodes structure.
type ListNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodesOK creates a ListNodesOK with default headers values
func NewListNodesOK() *ListNodesOK {
	return &ListNodesOK{}
}

/*ListNodesOK handles this case with default header values.

A successful response.
*/
type ListNodesOK struct {
	Payload *ListNodesOKBody
}

func (o *ListNodesOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/List][%d] listNodesOk  %+v", 200, o.Payload)
}

func (o *ListNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListNodesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodesDefault creates a ListNodesDefault with default headers values
func NewListNodesDefault(code int) *ListNodesDefault {
	return &ListNodesDefault{
		_statusCode: code,
	}
}

/*ListNodesDefault handles this case with default header values.

An error response.
*/
type ListNodesDefault struct {
	_statusCode int

	Payload *ListNodesDefaultBody
}

// Code gets the status code for the list nodes default response
func (o *ListNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListNodesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/List][%d] ListNodes default  %+v", o._statusCode, o.Payload)
}

func (o *ListNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListNodesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerItems0 ContainerNode represents a Docker container.
swagger:model ContainerItems0
*/
type ContainerItems0 struct {

	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address string `json:"address,omitempty"`

	// Node availability zone. Auto-detected and auto-updated.
	Az string `json:"az,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	// Auto-detected and auto-updated.
	ContainerID string `json:"container_id,omitempty"`

	// Container name. Auto-detected and auto-updated.
	ContainerName string `json:"container_name,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. Auto-detected and auto-updated.
	// If defined, Generic Node with that machine_id must exist.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Node model. Auto-detected and auto-updated.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Node region. Auto-detected and auto-updated.
	Region string `json:"region,omitempty"`
}

// Validate validates this container items0
func (o *ContainerItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerItems0) UnmarshalBinary(b []byte) error {
	var res ContainerItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GenericItems0 GenericNode represents a bare metal server or virtual machine.
swagger:model GenericItems0
*/
type GenericItems0 struct {

	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address string `json:"address,omitempty"`

	// Node availability zone. Auto-detected and auto-updated.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux distribution name and version. Auto-detected and auto-updated.
	Distro string `json:"distro,omitempty"`

	// Linux machine-id. Auto-detected and auto-updated.
	// Must be unique across all Generic Nodes if specified.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Node model. Auto-detected and auto-updated.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Node region. Auto-detected and auto-updated.
	Region string `json:"region,omitempty"`
}

// Validate validates this generic items0
func (o *GenericItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GenericItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenericItems0) UnmarshalBinary(b []byte) error {
	var res GenericItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListNodesDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ListNodesDefaultBody
*/
type ListNodesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this list nodes default body
func (o *ListNodesDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListNodesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListNodesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListNodesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListNodesOKBody list nodes OK body
swagger:model ListNodesOKBody
*/
type ListNodesOKBody struct {

	// container
	Container []*ContainerItems0 `json:"container"`

	// generic
	Generic []*GenericItems0 `json:"generic"`

	// remote
	Remote []*RemoteItems0 `json:"remote"`

	// remote amazon rds
	RemoteAmazonRDS []*RemoteAmazonRDSItems0 `json:"remote_amazon_rds"`
}

// Validate validates this list nodes OK body
func (o *ListNodesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteAmazonRDS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListNodesOKBody) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(o.Container) { // not required
		return nil
	}

	for i := 0; i < len(o.Container); i++ {
		if swag.IsZero(o.Container[i]) { // not required
			continue
		}

		if o.Container[i] != nil {
			if err := o.Container[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "container" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	for i := 0; i < len(o.Generic); i++ {
		if swag.IsZero(o.Generic[i]) { // not required
			continue
		}

		if o.Generic[i] != nil {
			if err := o.Generic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "generic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	for i := 0; i < len(o.Remote); i++ {
		if swag.IsZero(o.Remote[i]) { // not required
			continue
		}

		if o.Remote[i] != nil {
			if err := o.Remote[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "remote" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateRemoteAmazonRDS(formats strfmt.Registry) error {

	if swag.IsZero(o.RemoteAmazonRDS) { // not required
		return nil
	}

	for i := 0; i < len(o.RemoteAmazonRDS); i++ {
		if swag.IsZero(o.RemoteAmazonRDS[i]) { // not required
			continue
		}

		if o.RemoteAmazonRDS[i] != nil {
			if err := o.RemoteAmazonRDS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "remote_amazon_rds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListNodesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListNodesOKBody) UnmarshalBinary(b []byte) error {
	var res ListNodesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoteAmazonRDSItems0 RemoteAmazonRDSNode represents a Remote Node for Amazon RDS. Agents can't run on Remote Nodes.
swagger:model RemoteAmazonRDSItems0
*/
type RemoteAmazonRDSItems0 struct {

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// DB instance identifier. Unique across all RemoteAmazonRDS Nodes in combination with region. Can be changed.
	Instance string `json:"instance,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all RemoteAmazonRDS Nodes in combination with instance. Can't be changed.
	Region string `json:"region,omitempty"`
}

// Validate validates this remote amazon RDS items0
func (o *RemoteAmazonRDSItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoteAmazonRDSItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoteAmazonRDSItems0) UnmarshalBinary(b []byte) error {
	var res RemoteAmazonRDSItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoteItems0 RemoteNode represents generic remote Node. Agents can't run on Remote Nodes.
swagger:model RemoteItems0
*/
type RemoteItems0 struct {

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this remote items0
func (o *RemoteItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoteItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoteItems0) UnmarshalBinary(b []byte) error {
	var res RemoteItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
