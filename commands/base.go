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

import (
	"context"
	"fmt"
	"reflect"
)

var Ctx = context.Background()

// FIXME Expand this interface to cover our use cases:
// * Normal results output
// * Live logging / progress output
// * JSON output
// * Exit codes

type Result interface {
	result()
	fmt.Stringer
}

// Command is a common interface for all commands.
type Command interface {
	Run() (Result, error)
}

type ErrorResponse interface {
	error
	Code() int
}

type Error struct {
	Code  int
	Error string
}

func GetError(err ErrorResponse) Error {
	v := reflect.ValueOf(err)
	p := v.Elem().FieldByName("Payload")
	e := p.Elem().FieldByName("Error")
	return Error{
		Code:  err.Code(),
		Error: e.String(),
	}
}
