package utils

type PaginationQuery struct {
	Offset int `form:"offset"`
	Amount int `form:"amount"`
}
