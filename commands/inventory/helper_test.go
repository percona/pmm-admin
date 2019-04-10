package inventory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCustomLabel(t *testing.T) {
	errWrongFormat := fmt.Errorf("wrong custom label format")
	for _, v := range []struct {
		name     string
		input    string
		expected map[string]string
		expErr   error
	}{
		{"simple label", "foo=bar", map[string]string{"foo": "bar"}, nil},
		{"two labels", "foo=bar,bar=foo", map[string]string{"foo": "bar", "bar": "foo"}, nil},
		{"no value", "foo=", map[string]string(nil), errWrongFormat},
		{"no key", "=foo", map[string]string(nil), errWrongFormat},
		{"wrong format", "foo=bar,bar+foo", map[string]string(nil), errWrongFormat},
	} {
		t.Run(v.name, func(tt *testing.T) {
			customLabels, err := parseCustomLabels(v.input)
			assert.Equal(t, v.expected, customLabels)
			assert.Equal(t, v.expErr, err)
		})
	}
}
