/*
 * @file: func.go
 * @brief: 练习GO泛型-函数泛型
 * @author: Kewin Li
 * @date: 2023年9月30日
 */

package func_generics

import "fmt"

// @func: Sum
// @date: 2023年9月30日
// @brief: 切片求和
// @author: Kewin Li
// @param: []T vals
// @return T
func Sum[T Number](vals []T) T {
	var sum T

	for _, v := range vals {
		sum += v
	}
	return sum
}

// @func: Max
// @date: 2023年9月30日
// @brief: 切片找出最大值
// @author: Kewin Li
// @param: T vals
// @return T
func Max[T Number](vals []T) T {

	if len(vals) <= 0 {
		panic("slice size=0! \n")
	}

	res := vals[0]
	for _, v := range vals {
		if v > res {
			res = v
		}
	}

	return res
}

// @func: Min
// @date: 2023年9月30日
// @brief: 切片找出最小值
// @author: Kewin Li
// @param: T vals
// @return T
func Min[T Number](vals []T) T {

	if len(vals) <= 0 {
		panic("slice size=0! \n")
	}

	res := vals[0]
	for _, v := range vals {
		if v < res {
			res = v
		}
	}

	return res
}

// @func: Find
// @date: 2023年9月30日
// @brief: 切片的过滤与查找
// @author: Kewin Li
// @param: []T vals
// @param: func(v T) bool filter 设置过滤条件的函数
// @return T
func Find[T any](vals []T, filter func(v T) bool) ([]T, bool) {
	if len(vals) <= 0 {
		panic("slice size=0! \n")
	}

	var res []T
	for _, v := range vals {
		if filter(v) {
			res = append(res, v)
		}
	}

	if len(res) <= 0 {
		return nil, false
	}

	return res, true
}

func Insert[T any](idx int, v T, vals []T) []T {

	if idx < 0 || idx > len(vals) {
		fmt.Printf("idx %d out of range! \n", idx)
		return nil
	}

	res := make([]T, len(vals)+1)
	copy(res, vals[:idx])
	res[idx] = v
	copy(res[idx+1:], vals[idx:])

	return res

}

// Number @brief: 泛型传参限制
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}
