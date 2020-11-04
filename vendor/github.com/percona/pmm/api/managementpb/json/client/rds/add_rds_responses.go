// Code generated by go-swagger; DO NOT EDIT.

package rds

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

// AddRDSReader is a Reader for the AddRDS structure.
type AddRDSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRDSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRDSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddRDSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddRDSOK creates a AddRDSOK with default headers values
func NewAddRDSOK() *AddRDSOK {
	return &AddRDSOK{}
}

/*AddRDSOK handles this case with default header values.

A successful response.
*/
type AddRDSOK struct {
	Payload *AddRDSOKBody
}

func (o *AddRDSOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/RDS/Add][%d] addRdsOk  %+v", 200, o.Payload)
}

func (o *AddRDSOK) GetPayload() *AddRDSOKBody {
	return o.Payload
}

func (o *AddRDSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRDSOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRDSDefault creates a AddRDSDefault with default headers values
func NewAddRDSDefault(code int) *AddRDSDefault {
	return &AddRDSDefault{
		_statusCode: code,
	}
}

/*AddRDSDefault handles this case with default header values.

An unexpected error response.
*/
type AddRDSDefault struct {
	_statusCode int

	Payload *AddRDSDefaultBody
}

// Code gets the status code for the add RDS default response
func (o *AddRDSDefault) Code() int {
	return o._statusCode
}

func (o *AddRDSDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/RDS/Add][%d] AddRDS default  %+v", o._statusCode, o.Payload)
}

func (o *AddRDSDefault) GetPayload() *AddRDSDefaultBody {
	return o.Payload
}

func (o *AddRDSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRDSDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRDSBody add RDS body
swagger:model AddRDSBody
*/
type AddRDSBody struct {

	// AWS region.
	Region string `json:"region,omitempty"`

	// AWS availability zone.
	Az string `json:"az,omitempty"`

	// AWS instance ID.
	InstanceID string `json:"instance_id,omitempty"`

	// AWS instance class.
	NodeModel string `json:"node_model,omitempty"`

	// Address used to connect to it.
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// DiscoverRDSEngine describes supported RDS instance engines.
	// Enum: [DISCOVER_RDS_ENGINE_INVALID DISCOVER_RDS_MYSQL]
	Engine *string `json:"engine,omitempty"`

	// Unique across all Nodes user-defined name. Defaults to AWS instance ID.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all Services user-defined name. Defaults to AWS instance ID.
	ServiceName string `json:"service_name,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Password for scraping metrics.
	Password string `json:"password,omitempty"`

	// AWS Access key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

	// AWS Secret key.
	AWSSecretKey string `json:"aws_secret_key,omitempty"`

	// If true, adds rds_exporter.
	RDSExporter bool `json:"rds_exporter,omitempty"`

	// If true, adds qan-mysql-perfschema-agent.
	QANMysqlPerfschema bool `json:"qan_mysql_perfschema,omitempty"`

	// Custom user-assigned labels for Node and Service.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disable_query_examples,omitempty"`

	// Tablestats group collectors will be disabled if there are more than that number of tables.
	// If zero, server's default value is used.
	// Use negative value to disable them.
	TablestatsGroupTableLimit int32 `json:"tablestats_group_table_limit,omitempty"`

	// Disable basic metrics.
	DisableBasicMetrics bool `json:"disable_basic_metrics,omitempty"`

	// Disable enhanced metrics.
	DisableEnhancedMetrics bool `json:"disable_enhanced_metrics,omitempty"`
}

// Validate validates this add RDS body
func (o *AddRDSBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addRdsBodyTypeEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISCOVER_RDS_ENGINE_INVALID","DISCOVER_RDS_MYSQL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addRdsBodyTypeEnginePropEnum = append(addRdsBodyTypeEnginePropEnum, v)
	}
}

const (

	// AddRDSBodyEngineDISCOVERRDSENGINEINVALID captures enum value "DISCOVER_RDS_ENGINE_INVALID"
	AddRDSBodyEngineDISCOVERRDSENGINEINVALID string = "DISCOVER_RDS_ENGINE_INVALID"

	// AddRDSBodyEngineDISCOVERRDSMYSQL captures enum value "DISCOVER_RDS_MYSQL"
	AddRDSBodyEngineDISCOVERRDSMYSQL string = "DISCOVER_RDS_MYSQL"
)

// prop value enum
func (o *AddRDSBody) validateEngineEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addRdsBodyTypeEnginePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddRDSBody) validateEngine(formats strfmt.Registry) error {

	if swag.IsZero(o.Engine) { // not required
		return nil
	}

	// value enum
	if err := o.validateEngineEnum("body"+"."+"engine", "body", *o.Engine); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSBody) UnmarshalBinary(b []byte) error {
	var res AddRDSBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSDefaultBody add RDS default body
swagger:model AddRDSDefaultBody
*/
type AddRDSDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add RDS default body
func (o *AddRDSDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddRDS default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddRDSDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSOKBody add RDS OK body
swagger:model AddRDSOKBody
*/
type AddRDSOKBody struct {

	// Actual table count at the moment of adding.
	TableCount int32 `json:"table_count,omitempty"`

	// mysql
	Mysql *AddRDSOKBodyMysql `json:"mysql,omitempty"`

	// mysqld exporter
	MysqldExporter *AddRDSOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`

	// node
	Node *AddRDSOKBodyNode `json:"node,omitempty"`

	// qan mysql perfschema
	QANMysqlPerfschema *AddRDSOKBodyQANMysqlPerfschema `json:"qan_mysql_perfschema,omitempty"`

	// rds exporter
	RDSExporter *AddRDSOKBodyRDSExporter `json:"rds_exporter,omitempty"`
}

// Validate validates this add RDS OK body
func (o *AddRDSOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysql(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANMysqlPerfschema(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRDSExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRDSOKBody) validateMysql(formats strfmt.Registry) error {

	if swag.IsZero(o.Mysql) { // not required
		return nil
	}

	if o.Mysql != nil {
		if err := o.Mysql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsOk" + "." + "mysql")
			}
			return err
		}
	}

	return nil
}

func (o *AddRDSOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsOk" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddRDSOKBody) validateNode(formats strfmt.Registry) error {

	if swag.IsZero(o.Node) { // not required
		return nil
	}

	if o.Node != nil {
		if err := o.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsOk" + "." + "node")
			}
			return err
		}
	}

	return nil
}

func (o *AddRDSOKBody) validateQANMysqlPerfschema(formats strfmt.Registry) error {

	if swag.IsZero(o.QANMysqlPerfschema) { // not required
		return nil
	}

	if o.QANMysqlPerfschema != nil {
		if err := o.QANMysqlPerfschema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsOk" + "." + "qan_mysql_perfschema")
			}
			return err
		}
	}

	return nil
}

func (o *AddRDSOKBody) validateRDSExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.RDSExporter) { // not required
		return nil
	}

	if o.RDSExporter != nil {
		if err := o.RDSExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRdsOk" + "." + "rds_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSOKBody) UnmarshalBinary(b []byte) error {
	var res AddRDSOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSOKBodyMysql MySQLService represents a generic MySQL instance.
swagger:model AddRDSOKBodyMysql
*/
type AddRDSOKBodyMysql struct {

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

// Validate validates this add RDS OK body mysql
func (o *AddRDSOKBodyMysql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSOKBodyMysql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSOKBodyMysql) UnmarshalBinary(b []byte) error {
	var res AddRDSOKBodyMysql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL Service metrics.
swagger:model AddRDSOKBodyMysqldExporter
*/
type AddRDSOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Tablestats group collectors are disabled if there are more than that number of tables.
	// 0 means tablestats group collectors are always enabled (no limit).
	// Negative value means tablestats group collectors are always disabled.
	TablestatsGroupTableLimit int32 `json:"tablestats_group_table_limit,omitempty"`

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

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if tablestats group collectors are currently disabled.
	TablestatsGroupDisabled bool `json:"tablestats_group_disabled,omitempty"`
}

// Validate validates this add RDS OK body mysqld exporter
func (o *AddRDSOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addRdsOkBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addRdsOkBodyMysqldExporterTypeStatusPropEnum = append(addRdsOkBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddRDSOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddRDSOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddRDSOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	AddRDSOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// AddRDSOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	AddRDSOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// AddRDSOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	AddRDSOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// AddRDSOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	AddRDSOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// AddRDSOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	AddRDSOKBodyMysqldExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddRDSOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addRdsOkBodyMysqldExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddRDSOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addRdsOk"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res AddRDSOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSOKBodyNode RemoteRDSNode represents remote RDS Node. Agents can't run on Remote RDS Nodes.
swagger:model AddRDSOKBodyNode
*/
type AddRDSOKBodyNode struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add RDS OK body node
func (o *AddRDSOKBodyNode) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSOKBodyNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSOKBodyNode) UnmarshalBinary(b []byte) error {
	var res AddRDSOKBodyNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSOKBodyQANMysqlPerfschema QANMySQLPerfSchemaAgent runs within pmm-agent and sends MySQL Query Analytics data to the PMM Server.
swagger:model AddRDSOKBodyQANMysqlPerfschema
*/
type AddRDSOKBodyQANMysqlPerfschema struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for getting performance data.
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

// Validate validates this add RDS OK body QAN mysql perfschema
func (o *AddRDSOKBodyQANMysqlPerfschema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addRdsOkBodyQanMysqlPerfschemaTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addRdsOkBodyQanMysqlPerfschemaTypeStatusPropEnum = append(addRdsOkBodyQanMysqlPerfschemaTypeStatusPropEnum, v)
	}
}

const (

	// AddRDSOKBodyQANMysqlPerfschemaStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddRDSOKBodyQANMysqlPerfschemaStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddRDSOKBodyQANMysqlPerfschemaStatusSTARTING captures enum value "STARTING"
	AddRDSOKBodyQANMysqlPerfschemaStatusSTARTING string = "STARTING"

	// AddRDSOKBodyQANMysqlPerfschemaStatusRUNNING captures enum value "RUNNING"
	AddRDSOKBodyQANMysqlPerfschemaStatusRUNNING string = "RUNNING"

	// AddRDSOKBodyQANMysqlPerfschemaStatusWAITING captures enum value "WAITING"
	AddRDSOKBodyQANMysqlPerfschemaStatusWAITING string = "WAITING"

	// AddRDSOKBodyQANMysqlPerfschemaStatusSTOPPING captures enum value "STOPPING"
	AddRDSOKBodyQANMysqlPerfschemaStatusSTOPPING string = "STOPPING"

	// AddRDSOKBodyQANMysqlPerfschemaStatusDONE captures enum value "DONE"
	AddRDSOKBodyQANMysqlPerfschemaStatusDONE string = "DONE"
)

// prop value enum
func (o *AddRDSOKBodyQANMysqlPerfschema) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addRdsOkBodyQanMysqlPerfschemaTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddRDSOKBodyQANMysqlPerfschema) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addRdsOk"+"."+"qan_mysql_perfschema"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSOKBodyQANMysqlPerfschema) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSOKBodyQANMysqlPerfschema) UnmarshalBinary(b []byte) error {
	var res AddRDSOKBodyQANMysqlPerfschema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRDSOKBodyRDSExporter RDSExporter runs on Generic or Container Node and exposes RemoteRDS Node metrics.
swagger:model AddRDSOKBodyRDSExporter
*/
type AddRDSOKBodyRDSExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Node identifier.
	NodeID string `json:"node_id,omitempty"`

	// AWS Access Key.
	AWSAccessKey string `json:"aws_access_key,omitempty"`

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

	// Listen port for scraping metrics (the same for several configurations).
	ListenPort int64 `json:"listen_port,omitempty"`

	// Basic metrics are disabled.
	BasicMetricsDisabled bool `json:"basic_metrics_disabled,omitempty"`

	// Enhanced metrics are disabled.
	EnhancedMetricsDisabled bool `json:"enhanced_metrics_disabled,omitempty"`
}

// Validate validates this add RDS OK body RDS exporter
func (o *AddRDSOKBodyRDSExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addRdsOkBodyRdsExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addRdsOkBodyRdsExporterTypeStatusPropEnum = append(addRdsOkBodyRdsExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddRDSOKBodyRDSExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddRDSOKBodyRDSExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddRDSOKBodyRDSExporterStatusSTARTING captures enum value "STARTING"
	AddRDSOKBodyRDSExporterStatusSTARTING string = "STARTING"

	// AddRDSOKBodyRDSExporterStatusRUNNING captures enum value "RUNNING"
	AddRDSOKBodyRDSExporterStatusRUNNING string = "RUNNING"

	// AddRDSOKBodyRDSExporterStatusWAITING captures enum value "WAITING"
	AddRDSOKBodyRDSExporterStatusWAITING string = "WAITING"

	// AddRDSOKBodyRDSExporterStatusSTOPPING captures enum value "STOPPING"
	AddRDSOKBodyRDSExporterStatusSTOPPING string = "STOPPING"

	// AddRDSOKBodyRDSExporterStatusDONE captures enum value "DONE"
	AddRDSOKBodyRDSExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddRDSOKBodyRDSExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addRdsOkBodyRdsExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddRDSOKBodyRDSExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addRdsOk"+"."+"rds_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRDSOKBodyRDSExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRDSOKBodyRDSExporter) UnmarshalBinary(b []byte) error {
	var res AddRDSOKBodyRDSExporter
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
