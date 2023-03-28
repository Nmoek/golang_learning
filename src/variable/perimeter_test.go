package main

import (
	"math"
	"testing"
)

/*
 * @brief: 形状接口(类似C++的虚函数)
 */
type Shape interface {
	Area() float64
}

/*
 * @brief: 长方形结构体定义
 */
type Rectangle struct {
	width  float64
	length float64
}

/*
 * @brief: 长方形面积方法
 */
func (r Rectangle) Area() float64 {
	return r.width * r.length
}

/*
 * @brief: 圆形结构体定义
 */
type Circles struct {
	radius float64
}

/*
 * @brief: 圆形面积方法
 */
func (c Circles) Area() float64 {
	return c.radius * c.radius * math.Pi
}

/*
 * @brief: 三角形结构体定义
 */
type Triangle struct {
	base   float64
	height float64
}

/*
 * @brief: 三角形面积方法
 */
func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}

/*
 * @breif: 测试返回长方形/圆形面积
 */
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

/*
 * @breif: 计算长方形面积
 */
func Area(rect Rectangle) float64 {
	return rect.width * rect.length
}
