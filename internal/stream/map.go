package stream

func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}

func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		dataMap[v] = struct{}{}
	}
	return dataMap
}

func DiffSet[T comparable](src, dst []T) []T {
	srcMap := toMap[T](src)

	//首先根据 dst 删除 srcMap 中的值
	for _, val := range dst {
		delete(srcMap, val)
	}

	var ret = make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}

	return ret
}
