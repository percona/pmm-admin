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

	"github.com/percona/pmm-admin/agentlocal"
	"github.com/percona/pmm-admin/commands"
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
)

type annotationServiceCommand struct {
	Text        string
	Tags        string
	ServiceName string
}

type annotationServiceResult struct {
}

func (res *annotationServiceResult) Result() {}

func (res *annotationServiceResult) String() string {
	return ""
}

func (cmd *annotationServiceCommand) Run() (commands.Result, error) {
	var servicesNameList []string
	if cmd.ServiceName == "" {
		status, err := agentlocal.GetStatus(agentlocal.DoNotRequestNetworkInfo)
		if err != nil {
			return nil, err
		}

		params := &services.ListServicesParams{
			Body: services.ListServicesBody{
				NodeID: status.NodeID,
			},
			Context: commands.Ctx,
		}

		result, err := client.Default.Services.ListServices(params)
		if err != nil {
			return nil, err
		}

		for _, s := range result.Payload.Mysql {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.Mongodb {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.Postgresql {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.Proxysql {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
		for _, s := range result.Payload.External {
			servicesNameList = append(servicesNameList, s.ServiceName)
		}
	} else {
		params := &services.CheckServiceParams{
			Body: services.CheckServiceBody{
				ServiceName: cmd.ServiceName,
			},
			Context: commands.Ctx,
		}

		result, err := client.Default.Services.CheckService(params)
		if err != nil {
			return nil, err
		}

		if !result.Payload.Exists {
			return nil, errors.New("service name doesnt exists")
		}

		servicesNameList = append(servicesNameList, cmd.ServiceName)
	}

	//iterate servicesNameList and do request for annotate

	return nil, nil
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
