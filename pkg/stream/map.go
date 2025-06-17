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

// ToList 将一个单个 value 转换成为[]T
func ToList[T any](value T) []T {
	return []T{value}
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

// Each 就是一个for 循环
func Each[T any](array []T, fn func(index int, value T)) {
	for index, value := range array {
		fn(index, value)
	}
}

// MapObject 用来做对象的映射
func MapObject[Src any, Dst any](srcPtr *Src) Dst {
	dst := new(Dst)

	srcVal := reflect.ValueOf(srcPtr)
	dstVal := reflect.ValueOf(dst).Elem()

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	srcType := srcVal.Type()
	dstType := dstVal.Type()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcType.Field(i)
		srcFieldValue := srcVal.Field(i)

		if dstField, ok := dstType.FieldByName(srcField.Name); ok {
			if srcField.Type == dstField.Type {
				dstVal.FieldByName(srcField.Name).Set(srcFieldValue)
			}
		}
	}
	return dstVal.Interface().(Dst)
}
