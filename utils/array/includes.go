package array

func Includes[TElement comparable](array []TElement, testedElement TElement) bool {
	for _, element := range array {
		if testedElement == element {
			return true
		}
	}
	return false
}