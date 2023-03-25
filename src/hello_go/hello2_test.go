package main

import (
	"testing"
)

const HELLO_PREFIX = "hello "
const TONGSHI_FIX = ",u wokermate is "
const TONGSHI1 = "zxs"
const TONGSHI2 = "xt"

func TestHello(t *testing.T) {

	/*增加断言assert函数 替代原来的判断*/
	assertCorrectMsg := func(t *testing.T, got, want string) {
		//这个函数的作用:报错时指向调用assertCorrectMsg的位置，而非该函数内部
		t.Helper()
		if got != want {
			t.Errorf("got='%q' want='%q'", got, want)
		}
	}

	/*分组测试1*/
	t.Run("say hello to all peolpe", func(t *testing.T) {
		got := Hello("ljk", "zxs")
		want := "hello ljk,u wokermate is zxs"

		assertCorrectMsg(t, got, want)

	})

	/*分组测试2*/
	t.Run("say hello to all peolpe", func(t *testing.T) {
		got := Hello("ljk", "xt")
		want := "hello ljk,u wokermate is xt"

		assertCorrectMsg(t, got, want)

	})

	/*分组测试3*/
	t.Run("say hello wolrd when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world"

		assertCorrectMsg(t, got, want)

	})

}

func Hello(name string, tongshi string) string {

	switch tongshi {
	case TONGSHI1:
		{
			return HELLO_PREFIX + name + TONGSHI_FIX + tongshi
		}

	case TONGSHI2:
		{
			return HELLO_PREFIX + name + TONGSHI_FIX + tongshi
		}
	}

	return HELLO_PREFIX + "world"
}

func main() {

}
