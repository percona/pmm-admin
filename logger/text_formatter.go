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

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// TextFormatter formats logs into text.
type TextFormatter struct {
	*logrus.TextFormatter
}

// Format renders a single log entry.
func (f *TextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	levelText := strings.ToUpper(entry.Level.String())
	if !f.DisableLevelTruncation {
		levelText = levelText[0:4]
	}
	bytesArray, err := f.TextFormatter.Format(entry)
	if err != nil {
		return bytesArray, err
	}
	bytesArray = bytes.TrimPrefix(bytesArray, []byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, levelText)))

	bytesArray = bytes.TrimPrefix(bytesArray, []byte(" "))

	return bytesArray, nil
}
