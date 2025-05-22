package stream

import "reflect"

// Map 将一个数组转换成为map
func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}

func ToMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		dataMap[v] = struct{}{}
	}
	return dataMap
}

// DiffSet 用于计算两个集合的之间的差集
func DiffSet[T comparable](src, dst []T) []T {
	srcMap := ToMap[T](src)

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

// IntersectSet 求两个集合之间的交集
func IntersectSet[T comparable](src, dst []T) []T {
	srcMap := ToMap[T](src)
	dstMap := ToMap[T](dst)

	var ret = make([]T, 0, len(srcMap))
	for key := range srcMap {
		if _, ok := dstMap[key]; ok {
			ret = append(ret, key)
		}
	}
	return ret
}

// Filter 用于根据条件过滤出来对应的数组
func Filter[T any](array []T, fn func(value T) bool) []T {
	var newArray = make([]T, 0, len(array))
	for index, value := range array {
		if fn(value) {
			newArray = append(newArray, array[index])
		}
	}
	return newArray
}

// StructMap 这个函数主要是为了解决ddd 分层的时候我们通常需要把dao 层中的对象值转换成为domain 层的对象
func StructMap[Src any, Dst any](src *Src, dst *Dst) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
	}

	srcType := srcVal.Type()
	dstType := dstVal.Type()

	// 遍历所有的源filed
	l := srcVal.NumField()
	for i := 0; i < l; i++ {
		//根据索引获取field
		srcField := srcType.Field(i)
		// 首先在目标 field 中找到对应的元素
		if dstFiled, ok := dstType.FieldByName(srcField.Name); ok {
			if srcField.Type == dstFiled.Type {
				dstVal.FieldByName(srcField.Name).Set(srcVal.Field(i))
			}
		}
	}

	return nil
}
