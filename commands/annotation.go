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
	"strings"

	"github.com/percona/pmm/api/managementpb/json/client"
	"github.com/percona/pmm/api/managementpb/json/client/annotation"
	"gopkg.in/alecthomas/kingpin.v2"
)

// annotationResult is a result of annotation command.
type annotationResult struct {
	Message string `json:"message"`
}

// Result is a command run result.
func (res *annotationResult) Result() {}

// String stringifies command result.
func (res *annotationResult) String() string {
	return res.Message
}

type annotationCommand struct {
	Text string
	Tags string
}

// Run runs annotation command.
func (cmd *annotationCommand) Run() (Result, error) {
	_, err := client.Default.Annotation.AddAnnotation(&annotation.AddAnnotationParams{
		Body: annotation.AddAnnotationBody{
			Text: cmd.Text,
			Tags: strings.Split(cmd.Tags, ","),
		},
		Context: Ctx,
	})
	if err != nil {
		return nil, err
	}

	return &annotationResult{Message: "Annotation added."}, nil
}

// register command
var (
	Annotation  = new(annotationCommand)
	AnnotationC = kingpin.Command("annotation", "Add an annotation to Grafana charts")
)

func init() {
	AnnotationC.Arg("text", "Text of annotation").Required().StringVar(&Annotation.Text)
	AnnotationC.Flag("tags", "Tags to filter annotations. Multiple tags are separated by a comma.").StringVar(&Annotation.Tags)
}
