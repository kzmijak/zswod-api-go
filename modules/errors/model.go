package errors

type Error string

type wrappedError struct {
	cause   error
	message string
}