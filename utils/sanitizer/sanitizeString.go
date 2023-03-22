package sanitizer

import (
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type StringUnderSanitization struct { string }

const MAX_TITLE_LENGTH = 64	

func SanitizeString(target string) string {
	return builder(target).
		trimmedToLength(MAX_TITLE_LENGTH).
		withReplacedDiacriticalSigns().
		withReplacedSpacesWithDashes().
		asLowerCase().
		string
}

func builder(target string) StringUnderSanitization {
	return StringUnderSanitization{target}
}

func (target StringUnderSanitization)	withReplacedSpacesWithDashes () StringUnderSanitization {
	targetDashed := strings.Replace(target.string, " ", "-", -1)

	return builder(targetDashed)
}

func (target StringUnderSanitization) withReplacedDiacriticalSigns() StringUnderSanitization {

	transformer := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
        return unicode.Is(unicode.Mn, r)
    }), norm.NFC)
	result, _, _ := transform.String(transformer, target.string)

	return builder(result)
}

func (target StringUnderSanitization) trimmedToLength(length int) StringUnderSanitization {
	targetString := target.string;
	if len(targetString) > MAX_TITLE_LENGTH {
		targetString = targetString[:length - 1]
	}
	
	return builder(targetString)
}

func (target StringUnderSanitization) asLowerCase() StringUnderSanitization {
	targetToLower := strings.ToLower(target.string)
	return builder(targetToLower) 
}