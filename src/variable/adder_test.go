package integers

import (
	"fmt"
	"testing"
)

/*
 * ExampleXXX 测试法
 * $go test -v  xxx_test.go
 * 注意: 输出下面必须有"Output: <期望结果>"
 */
func ExampleAdd() {

	sum := Add(1, 1)
	fmt.Printf("%d\n", sum)
	// output: 2
}

/*
 * TestXXX 测试法
 * $go test xxx_test.go
 */
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("sum= '%d' expect='%d'\n", sum, expect)
	}

}

func Add(num1 int, num2 int) int {
	return num1 + num2
}
