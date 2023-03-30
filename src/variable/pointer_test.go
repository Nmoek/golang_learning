package main

import (
	"fmt"
	"testing"
)

// Wallet @brief: 钱包结构体
type Wallet struct {
	money int
}

// @func: add
// @brief: 往钱包中存入钱
// @author: Kewin Li
// @receiver: Wallet w 某个钱包
// @param[in]: int32 m 具体金额
func (w Wallet) add(m int) {
	fmt.Printf("w.money address=%p\n", &w.money)
	//w.money address=0xc000018128
	w.money += m
}

// @func: get
// @brief: 从钱包中取出钱
// @author: Kewin Li
// @receiver: Wallet w
// @return int32
func (w Wallet) get() int {
	return w.money
}

func TestWallet(t *testing.T) {

	w := Wallet{0}

	w.add(10) //增加10元存款

	fmt.Printf("w.money address=%p\n", &w.money)
	//w.money address=0xc000018120
	got := w.get() //取出10元
	want := 11

	if got != want {
		t.Errorf("got=%d want=%d \n", got, want)
	}

}
