package errors

type Error string

func (s Error) Error() string {
	return string(s)
}
