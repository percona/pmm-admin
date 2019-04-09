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
