package sanitizer

import "strings"

func SanitizeString(target string) string {
	return strings.Replace(target, " ", "-", -1)
}