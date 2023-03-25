package main

import (
	"testing"
)

const test_str = "hello "

func TestHello(t *testing.T) {

	t.Run("say hello to all peolpe", func(t *testing.T) {
		got := Hello("ljk1")
		want := "hello ljk1"

		if got != want {
			t.Errorf("test1 got='%q' want='%q' \n", got, want)
		}
	})

}

func Hello(name string) string {
	return test_str + name
}

func main() {

}
