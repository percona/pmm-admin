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

// Package agentlocal provides facilities for accessing local pmm-agent API.
package agentlocal

import (
	"reflect"
	"testing"
	"time"
)

func TestGetStatus(t *testing.T) {
	type args struct {
		requestNetworkInfo NetworkInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *Status
		wantErr bool
	}{
		{
			name: "status",
			args: args{
				requestNetworkInfo: DoNotRequestNetworkInfo,
			},
			want: &Status{
				"/agent_id/89718740-e04c-44eb-965a-99e0ff5e8a7c",
				"/node_id/526f1957-d967-4cff-b71a-8e068100633a",
				"https://admin:admin@127.0.0.1:443/",
				true,
				"2.5.0",
				[]AgentStatus{
					{
						"/agent_id/bfb939ca-9ec0-4c07-b444-cafabc79707c",
						"NODE_EXPORTER",
						"RUNNING",
					},
				},
				true,
				time.Duration(0 * time.Second),
				time.Duration(0 * time.Second),
				"",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStatus(tt.args.requestNetworkInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
