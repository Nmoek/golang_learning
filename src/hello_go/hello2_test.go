package main

import (
	"testing"
)

const test_str = "hello "

func TestHello(t *testing.T) {

	/*增加断言assert*/

	t.Run("say hello to all peolpe", func(t *testing.T) {
		got := Hello("ljk1")
		want := "hello ljk1"

		if got != want {
			t.Errorf("test1 got='%q' want='%q' \n", got, want)
		}
	})

	t.Run("say hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello world"

		if got != want {
			t.Errorf("test2 got='%q' want='%q' \n", got, want)
		}
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
