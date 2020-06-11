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

type annotationNodeCommand struct {
	Text     string
	Tags     string
	NodeName string
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
