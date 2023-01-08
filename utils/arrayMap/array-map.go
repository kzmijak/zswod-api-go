package arraymap

import (
	"github.com/samber/lo"
)

func CreateArrayMapper[TSource, TDestination any](
	schema func(src TSource) TDestination,
	) func(src []TSource) []TDestination {
	return func(src []TSource) []TDestination {
		return lo.Map(src, func(s TSource, _ int) TDestination { return schema(s) })
	}
}