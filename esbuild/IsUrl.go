package esbuild

import "strings"

func IsUrl(str string) bool {
	return strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://")
}
