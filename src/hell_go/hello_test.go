package main

import "fmt"
import "testing"

func TestHello(t *testing.T) {

	got := Hello()
	want := "hello world"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func Hello() string {
	return "hello test go"
}

func main() {

	fmt.Printf("%s\n", Hello())
}
