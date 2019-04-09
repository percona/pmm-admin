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

// Package logger provides helpers for logger.
package logger

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// TextFormatter formats logs into text.
type TextFormatter struct {
	*logrus.TextFormatter
}

// Format renders a single log entry.
func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry.Level == logrus.FatalLevel {

		var b *bytes.Buffer
		if entry.Buffer != nil {
			b = entry.Buffer
		} else {
			b = &bytes.Buffer{}
		}
		entry.Message = strings.TrimSuffix(entry.Message, "\n")

		caller := ""

		if entry.HasCaller() {
			funcVal := fmt.Sprintf("%s()", entry.Caller.Function)
			fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)

			if f.CallerPrettyfier != nil {
				funcVal, fileVal = f.CallerPrettyfier(entry.Caller)
			}
			caller = fileVal + " " + funcVal

			fmt.Fprintf(b, "%s %-44s ", caller, entry.Message)
		} else {
			fmt.Fprintf(b, "%-44s ", entry.Message)
		}

		b.WriteByte('\n')
		return b.Bytes(), nil
	}
	return f.TextFormatter.Format(entry)
}
