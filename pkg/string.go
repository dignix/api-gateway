package pkg

import "strings"

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-1]
	}
	return s
}
