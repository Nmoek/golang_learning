package main

import (
	"testing"
)

// 类型重命名
type Bitcoin int

// Stringer @brief: 字符串转换接口
type Stringer interface {
	String() string
}

// Wallet @brief: 钱包结构体
type Wallet struct {
	money Bitcoin
}

// @func: add
// @brief: 往钱包中存入钱
// @author: Kewin Li
// @receiver: Wallet w 某个钱包
// @param[in]: int32 m 具体金额
func (w *Wallet) add(m Bitcoin) {
	//w.money address=0xc000018128
	w.money += m
}

// @func: withDraw
// @brief: 从钱包中取钱
// @author: Kewin Li
// @receiver: Wallet w
// @param: Bitcoin m
func (w *Wallet) withDraw(m Bitcoin) {
	w.money -= m
}

// @func: get
// @brief: 查看钱包中的钱
// @author: Kewin Li
// @receiver: Wallet w
// @return int32
func (w Wallet) get() Bitcoin {
	return w.money
}

func TestWallet(t *testing.T) {

	t.Run("存钱", func(t *testing.T) {
		w := Wallet{}

		w.add(Bitcoin(10)) //增加10元存款

		//w.money address=0xc000018120
		got := w.get() //查看余额
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got=%d want=%d \n", got, want)
		}
	})

	t.Run("取钱", func(t *testing.T) {
		w := Wallet{Bitcoin(30)}

		w.withDraw(Bitcoin(10)) //取出10元存款

		got := w.get() //查看余额
		want := Bitcoin(20)

		if got != want {
			t.Errorf("got=%d want=%d \n", got, want)
		}
	})

}
