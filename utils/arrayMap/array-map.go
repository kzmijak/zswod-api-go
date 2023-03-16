package arraymap

func CreateArrayMapper[TSource, TDestination any](
	schema func(src TSource) (TDestination, error),
) func(src []TSource) ([]TDestination, error) {
	return func(src []TSource) ([]TDestination, error) {
		mapped := []TDestination{}
		for _, srcObj := range src {
			dstObj, err := schema(srcObj)
			if err != nil {
				return nil, err
			}
			mapped = append(mapped, dstObj)
		}
		return mapped, nil
	}
}