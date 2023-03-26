package main

import (
	"testing"
)

/*
 * @brief: 长方形结构体定义
 */
type Rectangle struct {
	width  float64
	length float64
}

/*
 * @breif: 测试返回长方形面积
 */
func TestArea(t *testing.T) {
	rect := Rectangle{3.0, 4.0}
	got := Area(rect)
	want := 12.0

	if got != want {
		t.Errorf("got=%.2f want=%.2f \n", got, want)
	}

}

/*
 * @breif: 测试返回长方形周长
 */
func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.00

	if got != want {
		t.Errorf("got=%.2f want=%.2f \n", got, want)
	}

}

/*
 * @breif: 计算长方形周长
 */
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.length)
}

func Area(rect Rectangle) float64 {
	return rect.width * rect.length
}
