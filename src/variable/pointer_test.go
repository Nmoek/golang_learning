package main

import (
	"errors"
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
func (w *Wallet) withDraw(m Bitcoin) error {

	if m > w.money {
		return errors.New("money not enough!!\n")
	}

	w.money -= m
	return nil
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

	errorMsg := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()

		if w.get() != want {
			t.Errorf("w.get()=%d  want=%d \n", w.get(), want)
		}
	}

	assertErr := func(t *testing.T, err error) {
		t.Helper()

		if err == nil {
			t.Errorf("want an error but not get \n")
		}
	}

	t.Run("存钱", func(t *testing.T) {
		w := Wallet{}

		w.add(Bitcoin(10)) //增加10元存款

		errorMsg(t, w, Bitcoin(10))
	})

	t.Run("取钱", func(t *testing.T) {
		w := Wallet{Bitcoin(30)}

		w.withDraw(Bitcoin(10)) //取出10元存款

		errorMsg(t, w, Bitcoin(20))

	})

	t.Run("超额透支取钱", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}

		err := w.withDraw(Bitcoin(100)) //只有10元却要取出30元

		errorMsg(t, w, Bitcoin(20))
		assertErr(t, err)

	})

}
