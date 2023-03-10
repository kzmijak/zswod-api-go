package errors

import "fmt"

const ErrSeperator = " -- "

func (s Error) Error() string {
	return string(s)
}

func (s Error) Wrap(err error) error {
	return wrappedError{
		cause:   err,
		message: string(s),
	}
}

func (we wrappedError) Error() string {
	if we.cause != nil {
		return fmt.Sprintf("%s%s%v", we.message, ErrSeperator, we.cause)
	}
	return we.message
}
