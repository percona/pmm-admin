// pmm-admin
// Copyright (C) 2018 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package commands

import "testing"

func TestListResultString(t *testing.T) {
	type fields struct {
		Services []listResultService
		Agents   []listResultAgent
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "filled",
			fields: fields{
				Services: []listResultService{
					{ServiceType: "mysql", ServiceID: "/service_id/test", ServiceName: "mysql-service"},
				},
				Agents: []listResultAgent{
					{AgentType: "mysql", AgentID: "/service_id/test", ServiceID: "/service_id/test", Status: "running"},
				},
			},
			want: `Service type  Service name         Address and port  Service ID
mysql         mysql-service                          /service_id/test

Agent type                  Status     Agent ID                                        Service ID
mysql                       running    /service_id/test  /service_id/test
`,
		},
		{
			name: "empty",
			fields: fields{
				Services: []listResultService{},
				Agents:   []listResultAgent{},
			},
			want: `Service type  Service name         Address and port  Service ID

Agent type                  Status     Agent ID                                        Service ID
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &listResult{
				Services: tt.fields.Services,
				Agents:   tt.fields.Agents,
			}
			if got := res.String(); got != tt.want {
				t.Errorf("listResult.String() =\n%v, want\n%v", got, tt.want)
			}
		})
	}
}
