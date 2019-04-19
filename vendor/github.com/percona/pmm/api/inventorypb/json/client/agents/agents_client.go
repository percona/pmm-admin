// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new agents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddExternalExporter adds external exporter adds external agent
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
	return result.(*AddExternalExporterOK), nil

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
	return result.(*AddMongoDBExporterOK), nil

}

/*
AddMySqldExporter adds my sqld exporter adds mysqld exporter agent
*/
func (a *Client) AddMySqldExporter(params *AddMySqldExporterParams) (*AddMySqldExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySqldExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddMySQLdExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/AddMySQLdExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySqldExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddMySqldExporterOK), nil

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
	return result.(*AddNodeExporterOK), nil

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
	return result.(*AddPMMAgentOK), nil

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
	return result.(*AddPostgresExporterOK), nil

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
	return result.(*AddQANMongoDBProfilerAgentOK), nil

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
	return result.(*AddQANMySQLPerfSchemaAgentOK), nil

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
	return result.(*AddQANMySQLSlowlogAgentOK), nil

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
	return result.(*AddRDSExporterOK), nil

}

/*
ChangeExternalExporter changes external exporter changes external agent
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
	return result.(*ChangeExternalExporterOK), nil

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
	return result.(*ChangeMongoDBExporterOK), nil

}

/*
ChangeMySqldExporter changes my sqld exporter changes mysqld exporter agent
*/
func (a *Client) ChangeMySqldExporter(params *ChangeMySqldExporterParams) (*ChangeMySqldExporterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeMySqldExporterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeMySQLdExporter",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangeMySQLdExporter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeMySqldExporterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeMySqldExporterOK), nil

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
	return result.(*ChangeNodeExporterOK), nil

}

/*
ChangePMMAgent changes PMM agent changes pmm agent agent
*/
func (a *Client) ChangePMMAgent(params *ChangePMMAgentParams) (*ChangePMMAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePMMAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangePMMAgent",
		Method:             "POST",
		PathPattern:        "/v1/inventory/Agents/ChangePMMAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangePMMAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangePMMAgentOK), nil

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
	return result.(*ChangePostgresExporterOK), nil

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
	return result.(*ChangeQANMongoDBProfilerAgentOK), nil

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
	return result.(*ChangeQANMySQLPerfSchemaAgentOK), nil

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
	return result.(*ChangeQANMySQLSlowlogAgentOK), nil

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
	return result.(*ChangeRDSExporterOK), nil

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
	return result.(*GetAgentOK), nil

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
	return result.(*ListAgentsOK), nil

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
	return result.(*RemoveAgentOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
