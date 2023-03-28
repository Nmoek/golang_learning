package main

import (
	"math"
	"testing"
)

type A struct {
	m_var int
}

func (a A) test_var() int {
	return a.m_var
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
 * @breif: 测试返回长方形/圆形面积
 */
func TestArea(t *testing.T) {

	/* 此处解决重命名的方案有两个
	 * 1.在新的包中重新定Aera(c Circles)
	 * 2. 定义"方法"。方法不等同于函数，函数可以任意时刻任意地方调用;
	 *    方法必须依赖于一个对象进行调用
	 */
	t.Run("return Rectanle Area", func(t *testing.T) {
		rect := Rectangle{3.0, 4.0}
		// got := Area(rect)  //重命名点1
		got := rect.Area()
		want := 12.0

		if got != want {
			t.Errorf("got=%.2f want=%.2f \n", got, want)
		}
	})

	t.Run("return Circles Area", func(t *testing.T) {
		cir := Circles{1.0}
		// got := Area(cir) //重命名点2
		got := cir.Area()
		want := math.Pi

		if got != want {
			t.Errorf("got=%.2f want=%.2f \n", got, want)
		}
	})

	t.Run("return A var", func(t *testing.T) {
		a := A{1}
		got := a.test_var()
		want := 1

		if got != want {
			t.Errorf("got=%d want=%d \n", got, want)
		}
	})
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
