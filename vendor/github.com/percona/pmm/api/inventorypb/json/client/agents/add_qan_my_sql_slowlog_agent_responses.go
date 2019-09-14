// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddQANMySQLSlowlogAgentReader is a Reader for the AddQANMySQLSlowlogAgent structure.
type AddQANMySQLSlowlogAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddQANMySQLSlowlogAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddQANMySQLSlowlogAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddQANMySQLSlowlogAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddQANMySQLSlowlogAgentOK creates a AddQANMySQLSlowlogAgentOK with default headers values
func NewAddQANMySQLSlowlogAgentOK() *AddQANMySQLSlowlogAgentOK {
	return &AddQANMySQLSlowlogAgentOK{}
}

/*AddQANMySQLSlowlogAgentOK handles this case with default header values.

A successful response.
*/
type AddQANMySQLSlowlogAgentOK struct {
	Payload *AddQANMySQLSlowlogAgentOKBody
}

func (o *AddQANMySQLSlowlogAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANMySQLSlowlogAgent][%d] addQanMySqlSlowlogAgentOk  %+v", 200, o.Payload)
}

func (o *AddQANMySQLSlowlogAgentOK) GetPayload() *AddQANMySQLSlowlogAgentOKBody {
	return o.Payload
}

func (o *AddQANMySQLSlowlogAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANMySQLSlowlogAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddQANMySQLSlowlogAgentDefault creates a AddQANMySQLSlowlogAgentDefault with default headers values
func NewAddQANMySQLSlowlogAgentDefault(code int) *AddQANMySQLSlowlogAgentDefault {
	return &AddQANMySQLSlowlogAgentDefault{
		_statusCode: code,
	}
}

/*AddQANMySQLSlowlogAgentDefault handles this case with default header values.

An error response.
*/
type AddQANMySQLSlowlogAgentDefault struct {
	_statusCode int

	Payload *AddQANMySQLSlowlogAgentDefaultBody
}

// Code gets the status code for the add QAN my SQL slowlog agent default response
func (o *AddQANMySQLSlowlogAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddQANMySQLSlowlogAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANMySQLSlowlogAgent][%d] AddQANMySQLSlowlogAgent default  %+v", o._statusCode, o.Payload)
}

func (o *AddQANMySQLSlowlogAgentDefault) GetPayload() *AddQANMySQLSlowlogAgentDefaultBody {
	return o.Payload
}

func (o *AddQANMySQLSlowlogAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANMySQLSlowlogAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddQANMySQLSlowlogAgentBody add QAN my SQL slowlog agent body
swagger:model AddQANMySQLSlowlogAgentBody
*/
type AddQANMySQLSlowlogAgentBody struct {

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disable_query_examples,omitempty"`

	// Rotate slowlog file at this size if > 0.
	// Use zero or negative value to disable rotation.
	MaxSlowlogFileSize string `json:"max_slowlog_file_size,omitempty"`

	// MySQL password for getting slowlog data.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MySQL username for getting slowlog data.
	Username string `json:"username,omitempty"`
}

// Validate validates this add QAN my SQL slowlog agent body
func (o *AddQANMySQLSlowlogAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentBody) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLSlowlogAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMySQLSlowlogAgentDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddQANMySQLSlowlogAgentDefaultBody
*/
type AddQANMySQLSlowlogAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add QAN my SQL slowlog agent default body
func (o *AddQANMySQLSlowlogAgentDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLSlowlogAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMySQLSlowlogAgentOKBody add QAN my SQL slowlog agent OK body
swagger:model AddQANMySQLSlowlogAgentOKBody
*/
type AddQANMySQLSlowlogAgentOKBody struct {

	// qan mysql slowlog agent
	QANMysqlSlowlogAgent *AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent `json:"qan_mysql_slowlog_agent,omitempty"`
}

// Validate validates this add QAN my SQL slowlog agent OK body
func (o *AddQANMySQLSlowlogAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANMysqlSlowlogAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANMySQLSlowlogAgentOKBody) validateQANMysqlSlowlogAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlSlowlogAgent) { // not required
		return nil
	}

	if o.QANMysqlSlowlogAgent != nil {
		if err := o.QANMysqlSlowlogAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanMySqlSlowlogAgentOk" + "." + "qan_mysql_slowlog_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLSlowlogAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent QANMySQLSlowlogAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent
*/
type AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Slowlog file is rotated at this size if > 0.
	MaxSlowlogFileSize string `json:"max_slowlog_file_size,omitempty"`

	// MySQL password for getting performance data.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MySQL username for getting performance data.
	Username string `json:"username,omitempty"`
}

// Validate validates this add QAN my SQL slowlog agent OK body QAN mysql slowlog agent
func (o *AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum = append(addQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTARTING captures enum value "STARTING"
	AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTARTING string = "STARTING"

	// AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusRUNNING captures enum value "RUNNING"
	AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusRUNNING string = "RUNNING"

	// AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusWAITING captures enum value "WAITING"
	AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusWAITING string = "WAITING"

	// AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTOPPING captures enum value "STOPPING"
	AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusSTOPPING string = "STOPPING"

	// AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusDONE captures enum value "DONE"
	AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addQanMySqlSlowlogAgentOkBodyQanMysqlSlowlogAgentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addQanMySqlSlowlogAgentOk"+"."+"qan_mysql_slowlog_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent) UnmarshalBinary(b []byte) error {
	var res AddQANMySQLSlowlogAgentOKBodyQANMysqlSlowlogAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
