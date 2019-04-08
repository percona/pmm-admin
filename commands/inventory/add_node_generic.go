package inventory

import (
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/percona/pmm/api/inventory/json/client/nodes"

	"github.com/percona/pmm-admin/commands"
)

var addGenericNodeResultT = commands.ParseTemplate(`
Generic Node added.
Node ID  : {{ .Node.NodeID }}
Node name: {{ .Node.NodeName }}

MachineID    : {{ .Node.MachineID }}
Distro       : {{ .Node.Distro }}
DistroVersion: {{ .Node.DistroVersion }}
Address      : {{ .Node.Address }}
`)

type addNodeGenericResult struct {
	Node *nodes.AddGenericNodeOKBodyGeneric
}

func (res *addNodeGenericResult) Result() {}

func (res *addNodeGenericResult) String() string {
	return commands.RenderTemplate(addGenericNodeResultT, res)
}

type addNodeGenericCommand struct {
	NodeType      string
	NodeName      string
	MachineID     string
	Distro        string
	DistroVersion string
	Address       string
	CustomLabels  map[string]string
}

func (cmd *addNodeGenericCommand) Run() (commands.Result, error) {
	params := &nodes.AddGenericNodeParams{
		Body: nodes.AddGenericNodeBody{
			NodeName:      cmd.NodeName,
			MachineID:     cmd.MachineID,
			Distro:        cmd.Distro,
			DistroVersion: cmd.DistroVersion,
			Address:       cmd.Address,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Nodes.AddGenericNode(params)
	if err != nil {
		return nil, err
	}
	return &addNodeGenericResult{
		Node: resp.Payload.Generic,
	}, nil
}

// register command
var (
	AddNodeGeneric  = new(addNodeGenericCommand)
	AddNodeGenericC = InventoryAddC.Command("node", "Add to inventory commands.")
)

func init() {
	AddNodeGenericC.Arg("name", "Node name").StringVar(&AddNodeGeneric.NodeName)

	AddNodeGenericC.Flag("machine-id", "Linux machine-id.").StringVar(&AddNodeGeneric.MachineID)
	AddNodeGenericC.Flag("distro", "Linux distribution (if any).").StringVar(&AddNodeGeneric.Distro)
	AddNodeGenericC.Flag("distro-version", "Linux distribution version (if any).").StringVar(&AddNodeGeneric.DistroVersion)
	AddNodeGenericC.Flag("address", "Address.").StringVar(&AddNodeGeneric.Address)
}
