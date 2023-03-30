/*
 * @file:perimeter_test.go
 * @brief: 结构体+"方法" + "接口"的测试学习
 * @author: Kewin Li
 * @date:2023-03-31
 */

package main

import (
	"math"
	"testing"
)

// Shape @brief: 形状接口(类似C++的虚函数)
type Shape interface {
	Area() float64
}

// Rectangle @brief: 长方形结构体定义
type Rectangle struct {
	width  float64
	length float64
}

// @func: Area
// @brief: 长方形面积方法
// @author: Kewin Li
// @receiver: Rectangle r
// @return float64
func (r Rectangle) Area() float64 {
	return r.width * r.length
}

// Circles @brief: 圆形结构体定义
type Circles struct {
	radius float64
}

// @func: Area
// @brief: 圆形面积方法
// @author: Kewin Li
// @receiver: Circles c
// @return float64
func (c Circles) Area() float64 {
	return c.radius * c.radius * math.Pi
}

// Triangle @brief: 三角形结构体定义
type Triangle struct {
	base   float64
	height float64
}

// @func: Area
// @brief: 三角形面积方法
// @author: Kewin Li
// @receiver: Triangle t
// @return float64
func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}

// @func: TestArea
// @brief: 面积测试
// @author: Kewin Li
// @param: *testing.T t
func TestArea(t *testing.T) {

	//一个匿名结构体
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle Test", shape: Rectangle{width: 3.0, length: 4.0}, want: 12.0}, //长方形
		{name: "Circles Test", shape: Circles{radius: 1.0}, want: math.Pi},              // 圆形
		{name: "Triangle Test", shape: Triangle{base: 12.0, height: 6.0}, want: 36.0},   //三角形
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()

			if got != test.want {
				t.Errorf("%#v  got=%.2f  want=%.2f \n", test, got, test.want)
			}
		})

	}

}

// @func: TestPerimeter
// @brief: 周长测试
// @author: Kewin Li
// @param: *testing.T t
func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.00

	if got != want {
		t.Errorf("got=%.2f want=%.2f \n", got, want)
	}

}

// @func: Perimeter
// @breif: 计算长方形周长
// @author: Kewin Li
// @param: Rectangle rect
// @return float64
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.length)
}

// @func: Area
// @breif: 计算长方形面积
// @author: Kewin Li
// @param: Rectangle rect
// @return float64
func Area(rect Rectangle) float64 {
	return rect.width * rect.length
}
