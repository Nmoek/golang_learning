/*
 * @file: mocking_test.go
 * @brief: 模拟mocking测试
 * @author: Kewin Li
 * @date: 2023-04-03
 */

package mocking_test

import (
	"reflect"
	"testing"
)

// Sleeper @brief: 自定义睡眠接口
type Sleeper interface {
	Sleep()
}

type SpySleep struct {
	Calls []string
}

// @func: Sleep
// @brief: 模拟睡眠(实际没有睡眠，记录一个操作字符串)
// @author: Kewin Li
// @receiver: *SpySleep s
func (s *SpySleep) Sleep() {
	s.Calls = append(s.Calls, "sleep\n")
}

// @func: Writer
// @brief: 模拟数据写入(实际没有进行写入, 记录一个操作字符串)
// @author: Kewin Li
// @receiver: *SpySleep s
// @param: []byte p
// @return n
// @return err
func (s *SpySleep) Writer(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write\n")
	return
}

func Count(s1 *SpySleep, n int, s2 *SpySleep) {
	for i := n; i > 0; i-- {
		s2.Sleep()
		s1.Writer(nil)
	}

	s2.Sleep()
	s1.Writer(nil)
}

func TestCount(t *testing.T) {

	//输出到指定buffer
	t.Run("out buffer", func(t *testing.T) {

		spy := &SpySleep{}

		Count(spy, 3, spy)

		got := spy.Calls
		want := []string{`sleep
write
sleep
write
sleep
write
sleep
write
`}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got= '%s' want='%s' \n", got, want)
		}

	})

}
