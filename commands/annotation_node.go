// pmm-admin
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"errors"

	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/nodes"

	"github.com/percona/pmm-admin/agentlocal"
)

type annotationNodeCommand struct {
	Text     string
	Tags     string
	NodeName string
}

type annotationNodeResult struct {
}

func (res *annotationNodeResult) Result() {}

func (res *annotationNodeResult) String() string {
	return ""
}

func (cmd *annotationNodeCommand) Run() (Result, error) {
	name := cmd.NodeName
	if name == "" {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}

		params := &nodes.GetNodeParams{
			Body: nodes.GetNodeBody{
				NodeID: status.NodeID,
			},
			Context: Ctx,
		}

		result, err := client.Default.Nodes.GetNode(params)
		if err != nil {
			return nil, err
		}

		name = result.Payload.Generic.NodeName
	}

	params := &nodes.CheckNodeParams{
		Body: nodes.CheckNodeBody{
			NodeName: name,
		},
		Context: Ctx,
	}

	result, err := client.Default.Nodes.CheckNode(params)
	if err != nil {
		return nil, err
	}

	if !result.Payload.Exists {
		return nil, errors.New("service name doesnt exists")
	}

	//do request for annotate

	return nil, nil
}

// register command
var (
	AnnotationNode  = new(annotationNodeCommand)
	AnnotationNodeC = AnnotationC.Command("node", "Add an node annotation to Grafana charts")
)

func init() {
	AnnotationNodeC.Arg("text", "Text of annotation").Required().StringVar(&AnnotationNode.Text)
	AnnotationNodeC.Flag("tags", "Tags to filter annotations. Multiple tags are separated by a comma").StringVar(&AnnotationNode.Tags)
	AnnotationNodeC.Flag("node-name", "Name of node for annotate").StringVar(&AnnotationNode.NodeName)
}
