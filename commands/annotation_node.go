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
	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-submodules/sources/pmm-admin/src/github.com/percona/pmm-admin/commands"
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

func (cmd *annotationNodeCommand) Run() (commands.Result, error) {
	if cmd.NodeName == "" {
		_, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}

		// params := &nodes.ListNodesParams{
		// 	Body: nodes.ListNodesBody{
		// 		NodeID: status.NodeID,
		// 	},
		// 	Context: commands.Ctx,
		// }

		// result, err := client.
		// if err != nil {
		// 	return nil, err
		// }

		// params := &nodes.ListNodesParams{
		// 	Body:    nodes.ListNodesBody{NodeType: status.nodeType},
		// 	Context: commands.Ctx,
		// }
		// result, err := client.Default.Nodes.ListNodes(params)
		// if err != nil {
		// 	return nil, err
		// }

		// var nodesList []listResultNode
		// for _, n := range result.Payload.Generic {
		// 	nodesList = append(nodesList, listResultNode{
		// 		NodeType: types.NodeTypeGenericNode,
		// 		NodeName: n.NodeName,
		// 		Address:  n.Address,
		// 		NodeID:   n.NodeID,
		// 	})
		// }
	}

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
