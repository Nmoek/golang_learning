package main

import (
	"testing"
)

const test_str = "hello "

func TestHello(t *testing.T) {

	/*增加断言assert函数 替代原来的判断*/
	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got='%q' want='%q'", got, want)
		}
	}

	/*分组测试1*/
	t.Run("say hello to all peolpe", func(t *testing.T) {
		got := Hello("ljk1")
		want := "hello ljk"

		assertCorrectMsg(t, got, want)

	})

	/*分组测试2*/
	t.Run("say hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello world"

		assertCorrectMsg(t, got, want)

	})

}

func Hello(name string) string {

	if name == "" {
		return test_str + "world"
	}
	return test_str + name
}

func main() {

}
