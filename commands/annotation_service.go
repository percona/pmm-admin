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

type annotationServiceCommand struct {
	Text        string
	Tags        string
	ServiceName string
}

// register command
var (
	AnnotationService  = new(annotationServiceCommand)
	AnnotationServiceC = AnnotationC.Command("service", "Add an service annotation to Grafana charts")
)

func init() {
	AnnotationServiceC.Arg("text", "Text of annotation").Required().StringVar(&AnnotationService.Text)
	AnnotationServiceC.Flag("tags", "Tags to filter annotations. Multiple tags are separated by a comma").StringVar(&AnnotationService.Tags)
	AnnotationServiceC.Flag("service-name", "Name of service for annotate").StringVar(&AnnotationService.ServiceName)
}
