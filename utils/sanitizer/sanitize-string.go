package sanitizer

import "strings"

const MAX_TITLE_LENGTH = 64	

func SanitizeString(target string) string {
	stringSanitized := strings.Replace(target, " ", "-", -1)

	if len(stringSanitized) > MAX_TITLE_LENGTH {
		stringSanitized = stringSanitized[:MAX_TITLE_LENGTH - 1]
	}

	return stringSanitized
}