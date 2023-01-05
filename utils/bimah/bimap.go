package bimap

type key[K int] string

type BiMap[TEnum comparable] struct {
	keyValue map[TEnum]string
	valueKey map[string]TEnum
}

func NewBiMap[TEnum comparable]() BiMap[TEnum] {
	return BiMap[TEnum]{
		make(map[TEnum]string), make(map[string]TEnum),
	}
}

func (bm BiMap[TEnum]) WithMember(key TEnum, value string) BiMap[TEnum] {
	bm.keyValue[key] = value
	bm.valueKey[value] = key
	return bm
}

func (bm BiMap[TEnum]) GetByKey(key TEnum) string {
	value, exists := bm.keyValue[key]

	if !exists {
		return ""
	}

	return value
}

func (bm BiMap[TEnum]) GetByValue(value string) (TEnum, bool) {
	key, exists := bm.valueKey[value]

	return key, exists
}