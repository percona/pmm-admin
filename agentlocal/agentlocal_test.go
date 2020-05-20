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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRawStatus(t *testing.T) {
	type args struct {
		ctx                context.Context
		requestNetworkInfo NetworkInfo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "emptyContext",
			args: args{},
			want: "Post http://127.0.0.1:7777/local/Status: context deadline exceeded",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetRawStatus(tt.args.ctx, tt.args.requestNetworkInfo)
			if assert.Error(t, err) {
				assert.Equal(t, tt.want, err.Error())
			}
		})
	}
}
