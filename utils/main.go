package utils

import (
	"fmt"
	"strings"
)

func MapToString(m map[string]string, delimiter string) string {
	tokens := make([]string, len(m), len(m))
	for key, val := range m {
		tokens = append(tokens, fmt.Sprintf("%v=%v", key, val))
	}

	str := strings.Join(tokens, delimiter)
	return str
}
