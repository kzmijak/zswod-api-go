package models

type ILogger interface {
	Trace(...interface{})
	Error(...interface{})
}