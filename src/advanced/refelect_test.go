/*
 * @file: refelect_test.go
 * @brief: Go的反射机制测试学习(选学)
 * @author: Kewin Li
 * @date:2023-04-06
 */

package reflect_test

import (
	"testing"
)

func walk(x interface{}, fn func(input string)) {
	fn("hello go!")
}

func TestWalk(t *testing.T) {

	want := "ljk"

	var got []string

	x := struct {
		name string
	}{want}

	walk(x, func(input string) {
		got = append(got, input)
	})

	t.Errorf("got=%s \n", got)

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got=%d, want=%d \n", len(got), 1)
	}

}
