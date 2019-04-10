package inventory

import (
	"fmt"
	"regexp"
	"strings"
)

func parseCustomLabels(labels string) (map[string]string, error) {
	regex := regexp.MustCompile(`(\w+)=(\w+)`)
	result := make(map[string]string)
	parts := strings.Split(labels, ",")
	for _, part := range parts {
		if !regex.MatchString(part) {
			return nil, fmt.Errorf("wrong custom label format")
		}
		submatches := regex.FindStringSubmatch(part)
		result[submatches[1]] = submatches[2]
	}
	return result, nil
}
