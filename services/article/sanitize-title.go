package article

import (
	"regexp"
	"strings"
)

func SanitizeTitle(input string) string {
	reg, err := regexp.Compile("[ąćęłńóśźż]")
	if err != nil {
		panic(err)
	}
	input = reg.ReplaceAllString(input, "")

	input = strings.Replace(input, " ", "-", -1)

	return input
}