// Code generated by go-swagger; DO NOT EDIT.

package postgre_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddPostgreSQLReader is a Reader for the AddPostgreSQL structure.
type AddPostgreSQLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPostgreSQLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPostgreSQLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPostgreSQLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPostgreSQLOK creates a AddPostgreSQLOK with default headers values
func NewAddPostgreSQLOK() *AddPostgreSQLOK {
	return &AddPostgreSQLOK{}
}

/*AddPostgreSQLOK handles this case with default header values.

A successful response.
*/
type AddPostgreSQLOK struct {
	Payload *AddPostgreSQLOKBody
}

func (o *AddPostgreSQLOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/PostgreSQL/Add][%d] addPostgreSqlOk  %+v", 200, o.Payload)
}

func (o *AddPostgreSQLOK) GetPayload() *AddPostgreSQLOKBody {
	return o.Payload
}

func (o *AddPostgreSQLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPostgreSQLDefault creates a AddPostgreSQLDefault with default headers values
func NewAddPostgreSQLDefault(code int) *AddPostgreSQLDefault {
	return &AddPostgreSQLDefault{
		_statusCode: code,
	}
}

/*AddPostgreSQLDefault handles this case with default header values.

An unexpected error response.
*/
type AddPostgreSQLDefault struct {
	_statusCode int

	Payload *AddPostgreSQLDefaultBody
}

// Code gets the status code for the add postgre SQL default response
func (o *AddPostgreSQLDefault) Code() int {
	return o._statusCode
}

func (o *AddPostgreSQLDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/PostgreSQL/Add][%d] AddPostgreSQL default  %+v", o._statusCode, o.Payload)
}

func (o *AddPostgreSQLDefault) GetPayload() *AddPostgreSQLDefaultBody {
	return o.Payload
}

func (o *AddPostgreSQLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddPostgreSQLBody add postgre SQL body
swagger:model AddPostgreSQLBody
*/
type AddPostgreSQLBody struct {

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"node_id,omitempty"`

	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node and Service access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `json:"address,omitempty"`

	// Service Access port.
	// Port is required when the address present.
	Port int64 `json:"port,omitempty"`

	// Service Access socket.
	// Address (and port) or socket is required.
	Socket string `json:"socket,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// PostgreSQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// If true, adds qan-postgresql-pgstatements-agent for provided service.
	QANPostgresqlPgstatementsAgent bool `json:"qan_postgresql_pgstatements_agent,omitempty"`

	// If true, adds qan-postgresql-pgstatmonitor-agent for provided service.
	QANPostgresqlPgstatmonitorAgent bool `json:"qan_postgresql_pgstatmonitor_agent,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disable_query_examples,omitempty"`

	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// MetricsMode defines desired metrics mode for agent,
	// it can be pull, push or auto mode chosen by server.
	// Enum: [AUTO PULL PUSH]
	MetricsMode *string `json:"metrics_mode,omitempty"`

	// add node
	AddNode *AddPostgreSQLParamsBodyAddNode `json:"add_node,omitempty"`
}

// Validate validates this add postgre SQL body
func (o *AddPostgreSQLBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetricsMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlBodyTypeMetricsModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AUTO","PULL","PUSH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlBodyTypeMetricsModePropEnum = append(addPostgreSqlBodyTypeMetricsModePropEnum, v)
	}
}

const (

	// AddPostgreSQLBodyMetricsModeAUTO captures enum value "AUTO"
	AddPostgreSQLBodyMetricsModeAUTO string = "AUTO"

	// AddPostgreSQLBodyMetricsModePULL captures enum value "PULL"
	AddPostgreSQLBodyMetricsModePULL string = "PULL"

	// AddPostgreSQLBodyMetricsModePUSH captures enum value "PUSH"
	AddPostgreSQLBodyMetricsModePUSH string = "PUSH"
)

// prop value enum
func (o *AddPostgreSQLBody) validateMetricsModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addPostgreSqlBodyTypeMetricsModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLBody) validateMetricsMode(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateMetricsModeEnum("body"+"."+"metrics_mode", "body", *o.MetricsMode); err != nil {
		return err
	}

	return nil
}

func (o *AddPostgreSQLBody) validateAddNode(formats strfmt.Registry) error {

	if swag.IsZero(o.AddNode) { // not required
		return nil
	}

	if o.AddNode != nil {
		if err := o.AddNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "add_node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLDefaultBody add postgre SQL default body
swagger:model AddPostgreSQLDefaultBody
*/
type AddPostgreSQLDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add postgre SQL default body
func (o *AddPostgreSQLDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddPostgreSQL default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBody add postgre SQL OK body
swagger:model AddPostgreSQLOKBody
*/
type AddPostgreSQLOKBody struct {

	// postgres exporter
	PostgresExporter *AddPostgreSQLOKBodyPostgresExporter `json:"postgres_exporter,omitempty"`

	// qan postgresql pgstatements agent
	QANPostgresqlPgstatementsAgent *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent `json:"qan_postgresql_pgstatements_agent,omitempty"`

	// qan postgresql pgstatmonitor agent
	QANPostgresqlPgstatmonitorAgent *AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent `json:"qan_postgresql_pgstatmonitor_agent,omitempty"`

	// service
	Service *AddPostgreSQLOKBodyService `json:"service,omitempty"`
}

// Validate validates this add postgre SQL OK body
func (o *AddPostgreSQLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePostgresExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANPostgresqlPgstatementsAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANPostgresqlPgstatmonitorAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLOKBody) validatePostgresExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.PostgresExporter) { // not required
		return nil
	}

	if o.PostgresExporter != nil {
		if err := o.PostgresExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "postgres_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddPostgreSQLOKBody) validateQANPostgresqlPgstatementsAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANPostgresqlPgstatementsAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatementsAgent != nil {
		if err := o.QANPostgresqlPgstatementsAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "qan_postgresql_pgstatements_agent")
			}
			return err
		}
	}

	return nil
}

func (o *AddPostgreSQLOKBody) validateQANPostgresqlPgstatmonitorAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANPostgresqlPgstatmonitorAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatmonitorAgent != nil {
		if err := o.QANPostgresqlPgstatmonitorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "qan_postgresql_pgstatmonitor_agent")
			}
			return err
		}
	}

	return nil
}

func (o *AddPostgreSQLOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyPostgresExporter PostgresExporter runs on Generic or Container Node and exposes PostgreSQL Service metrics.
swagger:model AddPostgreSQLOKBodyPostgresExporter
*/
type AddPostgreSQLOKBodyPostgresExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter use pull metrics mode.
	PushMetricsDisabled bool `json:"push_metrics_disabled,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`
}

// Validate validates this add postgre SQL OK body postgres exporter
func (o *AddPostgreSQLOKBodyPostgresExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum = append(addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddPostgreSQLOKBodyPostgresExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddPostgreSQLOKBodyPostgresExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddPostgreSQLOKBodyPostgresExporterStatusSTARTING captures enum value "STARTING"
	AddPostgreSQLOKBodyPostgresExporterStatusSTARTING string = "STARTING"

	// AddPostgreSQLOKBodyPostgresExporterStatusRUNNING captures enum value "RUNNING"
	AddPostgreSQLOKBodyPostgresExporterStatusRUNNING string = "RUNNING"

	// AddPostgreSQLOKBodyPostgresExporterStatusWAITING captures enum value "WAITING"
	AddPostgreSQLOKBodyPostgresExporterStatusWAITING string = "WAITING"

	// AddPostgreSQLOKBodyPostgresExporterStatusSTOPPING captures enum value "STOPPING"
	AddPostgreSQLOKBodyPostgresExporterStatusSTOPPING string = "STOPPING"

	// AddPostgreSQLOKBodyPostgresExporterStatusDONE captures enum value "DONE"
	AddPostgreSQLOKBodyPostgresExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddPostgreSQLOKBodyPostgresExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLOKBodyPostgresExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addPostgreSqlOk"+"."+"postgres_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyPostgresExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyPostgresExporter) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyPostgresExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent QANPostgreSQLPgStatementsAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent
*/
type AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for getting pg stat statements data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this add postgre SQL OK body QAN postgresql pgstatements agent
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum = append(addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING captures enum value "STARTING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING string = "STARTING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING captures enum value "RUNNING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING string = "RUNNING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusWAITING captures enum value "WAITING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusWAITING string = "WAITING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING captures enum value "STOPPING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING string = "STOPPING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusDONE captures enum value "DONE"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addPostgreSqlOk"+"."+"qan_postgresql_pgstatements_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent QANPostgreSQLPgStatMonitorAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent
*/
type AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// PostgreSQL username for getting pg stat monitor data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"query_examples_disabled,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`
}

// Validate validates this add postgre SQL OK body QAN postgresql pgstatmonitor agent
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum = append(addPostgreSqlOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusSTARTING captures enum value "STARTING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusSTARTING string = "STARTING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusRUNNING captures enum value "RUNNING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusRUNNING string = "RUNNING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusWAITING captures enum value "WAITING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusWAITING string = "WAITING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusSTOPPING captures enum value "STOPPING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusSTOPPING string = "STOPPING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusDONE captures enum value "DONE"
	AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addPostgreSqlOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addPostgreSqlOk"+"."+"qan_postgresql_pgstatmonitor_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyQANPostgresqlPgstatmonitorAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyService PostgreSQLService represents a generic PostgreSQL instance.
swagger:model AddPostgreSQLOKBodyService
*/
type AddPostgreSQLOKBodyService struct {

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

// Validate validates this add postgre SQL OK body service
func (o *AddPostgreSQLOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLParamsBodyAddNode AddNodeParams is a params to add new node to inventory while adding new service.
swagger:model AddPostgreSQLParamsBodyAddNode
*/
type AddPostgreSQLParamsBodyAddNode struct {

	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_INVALID GENERIC_NODE CONTAINER_NODE REMOTE_NODE REMOTE_RDS_NODE]
	NodeType *string `json:"node_type,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels for Node.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add postgre SQL params body add node
func (o *AddPostgreSQLParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE","REMOTE_RDS_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum = append(addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddPostgreSQLParamsBodyAddNodeNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	AddPostgreSQLParamsBodyAddNodeNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeGENERICNODE string = "GENERIC_NODE"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeREMOTENODE string = "REMOTE_NODE"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeREMOTERDSNODE captures enum value "REMOTE_RDS_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeREMOTERDSNODE string = "REMOTE_RDS_NODE"
)

// prop value enum
func (o *AddPostgreSQLParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"add_node"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLParamsBodyAddNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
