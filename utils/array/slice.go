package array

func Slice[TElement any](array []TElement, start int, end int) []TElement {
	var subArray []TElement
	if len(array) >= end {
		subArray = array[start:end]
	} else {
		subArray = array[start:]
	}
	return subArray
}