package article

import (
	"strings"
)

func SanitizeTitle(input string) string {
	input = strings.Replace(input, " ", "-", -1)

	return input
}