

package main

import "fmt"
import "testing"

func TestHello(t *testing.T) {

	got := Hello("ljk")
	want := "hello ljk"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func Hello(name string) string {
	return "hello " + name
}

func main() {

	fmt.Printf("test over!\n")
}
