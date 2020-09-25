// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new agents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddExternalExporter(params *AddExternalExporterParams) (*AddExternalExporterOK, error)

	AddMongoDBExporter(params *AddMongoDBExporterParams) (*AddMongoDBExporterOK, error)

	AddMySQLdExporter(params *AddMySQLdExporterParams) (*AddMySQLdExporterOK, error)

	AddNodeExporter(params *AddNodeExporterParams) (*AddNodeExporterOK, error)

	AddPMMAgent(params *AddPMMAgentParams) (*AddPMMAgentOK, error)

	AddPostgresExporter(params *AddPostgresExporterParams) (*AddPostgresExporterOK, error)

	AddProxySQLExporter(params *AddProxySQLExporterParams) (*AddProxySQLExporterOK, error)

	AddQANMongoDBProfilerAgent(params *AddQANMongoDBProfilerAgentParams) (*AddQANMongoDBProfilerAgentOK, error)

	AddQANMySQLPerfSchemaAgent(params *AddQANMySQLPerfSchemaAgentParams) (*AddQANMySQLPerfSchemaAgentOK, error)

	AddQANMySQLSlowlogAgent(params *AddQANMySQLSlowlogAgentParams) (*AddQANMySQLSlowlogAgentOK, error)

	AddQANPostgreSQLPgStatMonitorAgent(params *AddQANPostgreSQLPgStatMonitorAgentParams) (*AddQANPostgreSQLPgStatMonitorAgentOK, error)

	AddQANPostgreSQLPgStatementsAgent(params *AddQANPostgreSQLPgStatementsAgentParams) (*AddQANPostgreSQLPgStatementsAgentOK, error)

	AddRDSExporter(params *AddRDSExporterParams) (*AddRDSExporterOK, error)

	ChangeExternalExporter(params *ChangeExternalExporterParams) (*ChangeExternalExporterOK, error)

	ChangeMongoDBExporter(params *ChangeMongoDBExporterParams) (*ChangeMongoDBExporterOK, error)

	ChangeMySQLdExporter(params *ChangeMySQLdExporterParams) (*ChangeMySQLdExporterOK, error)

	ChangeNodeExporter(params *ChangeNodeExporterParams) (*ChangeNodeExporterOK, error)

	ChangePostgresExporter(params *ChangePostgresExporterParams) (*ChangePostgresExporterOK, error)

	ChangeProxySQLExporter(params *ChangeProxySQLExporterParams) (*ChangeProxySQLExporterOK, error)

	ChangeQANMongoDBProfilerAgent(params *ChangeQANMongoDBProfilerAgentParams) (*ChangeQANMongoDBProfilerAgentOK, error)

	ChangeQANMySQLPerfSchemaAgent(params *ChangeQANMySQLPerfSchemaAgentParams) (*ChangeQANMySQLPerfSchemaAgentOK, error)

	ChangeQANMySQLSlowlogAgent(params *ChangeQANMySQLSlowlogAgentParams) (*ChangeQANMySQLSlowlogAgentOK, error)

	ChangeQANPostgreSQLPgStatMonitorAgent(params *ChangeQANPostgreSQLPgStatMonitorAgentParams) (*ChangeQANPostgreSQLPgStatMonitorAgentOK, error)

	ChangeQANPostgreSQLPgStatementsAgent(params *ChangeQANPostgreSQLPgStatementsAgentParams) (*ChangeQANPostgreSQLPgStatementsAgentOK, error)

	ChangeRDSExporter(params *ChangeRDSExporterParams) (*ChangeRDSExporterOK, error)

	GetAgent(params *GetAgentParams) (*GetAgentOK, error)

	ListAgents(params *ListAgentsParams) (*ListAgentsOK, error)

	RemoveAgent(params *RemoveAgentParams) (*RemoveAgentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddExternalExporter adds external exporter adds external exporter agent
*/
func (a *Client) AddExternalExporter(params *AddExternalExporterParams) (*AddExternalExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddExternalExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddExternalExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddExternalExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddExternalExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddExternalExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddExternalExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddMongoDBExporter adds mongo DB exporter adds mongodb exporter agent
*/
func (a *Client) AddMongoDBExporter(params *AddMongoDBExporterParams) (*AddMongoDBExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMongoDBExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMongoDBExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddMongoDBExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMongoDBExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddMongoDBExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMongoDBExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddMySQLdExporter adds my s q ld exporter adds mysqld exporter agent
*/
func (a *Client) AddMySQLdExporter(params *AddMySQLdExporterParams) (*AddMySQLdExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySQLdExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQLdExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddMySQLdExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySQLdExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddMySQLdExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMySQLdExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddNodeExporter adds node exporter adds node exporter agent
*/
func (a *Client) AddNodeExporter(params *AddNodeExporterParams) (*AddNodeExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNodeExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddNodeExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddNodeExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddNodeExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddNodeExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddNodeExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddPMMAgent adds PMM agent adds pmm agent agent
*/
func (a *Client) AddPMMAgent(params *AddPMMAgentParams) (*AddPMMAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPMMAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddPMMAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddPMMAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPMMAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPMMAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddPMMAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddPostgresExporter adds postgres exporter adds postgres exporter agent
*/
func (a *Client) AddPostgresExporter(params *AddPostgresExporterParams) (*AddPostgresExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPostgresExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddPostgresExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddPostgresExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPostgresExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPostgresExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddPostgresExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddProxySQLExporter adds proxy SQL exporter adds proxysql exporter agent
*/
func (a *Client) AddProxySQLExporter(params *AddProxySQLExporterParams) (*AddProxySQLExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProxySQLExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddProxySQLExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddProxySQLExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddProxySQLExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddProxySQLExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddProxySQLExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddQANMongoDBProfilerAgent adds QAN mongo DB profiler agent adds mongo DB profiler QAN agent
*/
func (a *Client) AddQANMongoDBProfilerAgent(params *AddQANMongoDBProfilerAgentParams) (*AddQANMongoDBProfilerAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddQANMongoDBProfilerAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddQANMongoDBProfilerAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddQANMongoDBProfilerAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddQANMongoDBProfilerAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddQANMongoDBProfilerAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddQANMongoDBProfilerAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddQANMySQLPerfSchemaAgent adds QAN my SQL perf schema agent adds my SQL perf schema QAN agent
*/
func (a *Client) AddQANMySQLPerfSchemaAgent(params *AddQANMySQLPerfSchemaAgentParams) (*AddQANMySQLPerfSchemaAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddQANMySQLPerfSchemaAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddQANMySQLPerfSchemaAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddQANMySQLPerfSchemaAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddQANMySQLPerfSchemaAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddQANMySQLPerfSchemaAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddQANMySQLPerfSchemaAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddQANMySQLSlowlogAgent adds QAN my SQL slowlog agent adds my SQL perf schema QAN agent
*/
func (a *Client) AddQANMySQLSlowlogAgent(params *AddQANMySQLSlowlogAgentParams) (*AddQANMySQLSlowlogAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddQANMySQLSlowlogAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddQANMySQLSlowlogAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddQANMySQLSlowlogAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddQANMySQLSlowlogAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddQANMySQLSlowlogAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddQANMySQLSlowlogAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddQANPostgreSQLPgStatMonitorAgent adds QAN postgre SQL pg stat monitor agent adds postgre SQL pg stat monitor QAN agent
*/
func (a *Client) AddQANPostgreSQLPgStatMonitorAgent(params *AddQANPostgreSQLPgStatMonitorAgentParams) (*AddQANPostgreSQLPgStatMonitorAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddQANPostgreSQLPgStatMonitorAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddQANPostgreSQLPgStatMonitorAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddQANPostgreSQLPgStatMonitorAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddQANPostgreSQLPgStatMonitorAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddQANPostgreSQLPgStatMonitorAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddQANPostgreSQLPgStatMonitorAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddQANPostgreSQLPgStatementsAgent adds QAN postgre SQL pg statements agent adds postgre SQL pg stat statements QAN agent
*/
func (a *Client) AddQANPostgreSQLPgStatementsAgent(params *AddQANPostgreSQLPgStatementsAgentParams) (*AddQANPostgreSQLPgStatementsAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddQANPostgreSQLPgStatementsAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddQANPostgreSQLPgStatementsAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddQANPostgreSQLPgStatementsAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddQANPostgreSQLPgStatementsAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddQANPostgreSQLPgStatementsAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddQANPostgreSQLPgStatementsAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AddRDSExporter adds RDS exporter adds rds exporter agent
*/
func (a *Client) AddRDSExporter(params *AddRDSExporterParams) (*AddRDSExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRDSExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddRDSExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddRDSExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRDSExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddRDSExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddRDSExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeExternalExporter changes external exporter changes external exporter agent
*/
func (a *Client) ChangeExternalExporter(params *ChangeExternalExporterParams) (*ChangeExternalExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeExternalExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeExternalExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeExternalExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeExternalExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeExternalExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeExternalExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeMongoDBExporter changes mongo DB exporter changes mongodb exporter agent
*/
func (a *Client) ChangeMongoDBExporter(params *ChangeMongoDBExporterParams) (*ChangeMongoDBExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeMongoDBExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeMongoDBExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeMongoDBExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeMongoDBExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeMongoDBExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeMongoDBExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeMySQLdExporter changes my s q ld exporter changes mysqld exporter agent
*/
func (a *Client) ChangeMySQLdExporter(params *ChangeMySQLdExporterParams) (*ChangeMySQLdExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeMySQLdExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeMySQLdExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeMySQLdExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeMySQLdExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeMySQLdExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeMySQLdExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeNodeExporter changes node exporter changes node exporter agent
*/
func (a *Client) ChangeNodeExporter(params *ChangeNodeExporterParams) (*ChangeNodeExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeNodeExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeNodeExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeNodeExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeNodeExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeNodeExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeNodeExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangePostgresExporter changes postgres exporter changes postgres exporter agent
*/
func (a *Client) ChangePostgresExporter(params *ChangePostgresExporterParams) (*ChangePostgresExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePostgresExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangePostgresExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangePostgresExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangePostgresExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangePostgresExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangePostgresExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeProxySQLExporter changes proxy SQL exporter changes proxysql exporter agent
*/
func (a *Client) ChangeProxySQLExporter(params *ChangeProxySQLExporterParams) (*ChangeProxySQLExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeProxySQLExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeProxySQLExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeProxySQLExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeProxySQLExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeProxySQLExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeProxySQLExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeQANMongoDBProfilerAgent changes QAN mongo DB profiler agent changes mongo DB profiler QAN agent
*/
func (a *Client) ChangeQANMongoDBProfilerAgent(params *ChangeQANMongoDBProfilerAgentParams) (*ChangeQANMongoDBProfilerAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeQANMongoDBProfilerAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeQANMongoDBProfilerAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeQANMongoDBProfilerAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeQANMongoDBProfilerAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeQANMongoDBProfilerAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeQANMongoDBProfilerAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeQANMySQLPerfSchemaAgent changes QAN my SQL perf schema agent changes my SQL perf schema QAN agent
*/
func (a *Client) ChangeQANMySQLPerfSchemaAgent(params *ChangeQANMySQLPerfSchemaAgentParams) (*ChangeQANMySQLPerfSchemaAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeQANMySQLPerfSchemaAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeQANMySQLPerfSchemaAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeQANMySQLPerfSchemaAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeQANMySQLPerfSchemaAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeQANMySQLPerfSchemaAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeQANMySQLPerfSchemaAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeQANMySQLSlowlogAgent changes QAN my SQL slowlog agent changes my SQL perf schema QAN agent
*/
func (a *Client) ChangeQANMySQLSlowlogAgent(params *ChangeQANMySQLSlowlogAgentParams) (*ChangeQANMySQLSlowlogAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeQANMySQLSlowlogAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeQANMySQLSlowlogAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeQANMySQLSlowlogAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeQANMySQLSlowlogAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeQANMySQLSlowlogAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeQANMySQLSlowlogAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeQANPostgreSQLPgStatMonitorAgent changes QAN postgre SQL pg stat monitor agent changes postgre SQL pg stat monitor QAN agent
*/
func (a *Client) ChangeQANPostgreSQLPgStatMonitorAgent(params *ChangeQANPostgreSQLPgStatMonitorAgentParams) (*ChangeQANPostgreSQLPgStatMonitorAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeQANPostgreSQLPgStatMonitorAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeQANPostgreSQLPgStatMonitorAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeQANPostgreSQLPgStatMonitorAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeQANPostgreSQLPgStatMonitorAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeQANPostgreSQLPgStatMonitorAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeQANPostgreSQLPgStatMonitorAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeQANPostgreSQLPgStatementsAgent changes QAN postgre SQL pg statements agent changes postgre SQL pg stat statements QAN agent
*/
func (a *Client) ChangeQANPostgreSQLPgStatementsAgent(params *ChangeQANPostgreSQLPgStatementsAgentParams) (*ChangeQANPostgreSQLPgStatementsAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeQANPostgreSQLPgStatementsAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeQANPostgreSQLPgStatementsAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeQANPostgreSQLPgStatementsAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeQANPostgreSQLPgStatementsAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeQANPostgreSQLPgStatementsAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeQANPostgreSQLPgStatementsAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeRDSExporter changes RDS exporter changes rds exporter agent
*/
func (a *Client) ChangeRDSExporter(params *ChangeRDSExporterParams) (*ChangeRDSExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeRDSExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeRDSExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeRDSExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeRDSExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeRDSExporterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeRDSExporterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAgent gets agent returns a single agent by ID
*/
func (a *Client) GetAgent(params *GetAgentParams) (*GetAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAgents lists agents returns a list of all agents
*/
func (a *Client) ListAgents(params *ListAgentsParams) (*ListAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAgents",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAgentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RemoveAgent removes agent removes agent
*/
func (a *Client) RemoveAgent(params *RemoveAgentParams) (*RemoveAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/Remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveAgentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
