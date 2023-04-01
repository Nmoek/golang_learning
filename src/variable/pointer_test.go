package main

import (
	"errors"
	"fmt"
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

// @func: String
// @brief: 比特币字符串方法
// @author: Kewin Li
// @receiver: Bitcoin b
// @return string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// @func: add
// @brief: 往钱包中存入钱
// @author: Kewin Li
// @receiver: Wallet w 某个钱包
// @param[in]: int32 m 具体金额
func (w *Wallet) add(m Bitcoin) {
	w.money += m
}

// 包内全局错误对象，针对处理钱包错误信息
var InWalletErr = errors.New("cannot withdraw, insufficient funds")

// @func: withDraw
// @brief: 从钱包中取钱
// @author: Kewin Li
// @receiver: Wallet w
// @param: Bitcoin m
func (w *Wallet) withDraw(m Bitcoin) error {

	if m > w.money {
		return InWalletErr
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

// @func: TestWallet
// @brief: 钱包对象测试
// @author: Kewin Li
// @param: *testing.T t
func TestWallet(t *testing.T) {

	t.Run("存钱", func(t *testing.T) {
		w := Wallet{}

		w.add(Bitcoin(10)) //增加10元存款

		assertMoney(t, w, Bitcoin(10))
	})

	t.Run("取钱", func(t *testing.T) {
		w := Wallet{Bitcoin(30)}

		err := w.withDraw(Bitcoin(10)) //取出10元存款

		assertMoney(t, w, Bitcoin(20))
		assertErr(t, err, InWalletErr)

	})

	// 引入error使用
	t.Run("超额透支取钱", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}

		err := w.withDraw(Bitcoin(100)) // 只有20元却要取出100元 取出必定失败

		assertMoney(t, w, Bitcoin(20))
		assertErr(t, err, InWalletErr)
	})
}

// @func: assertMoney
// @brief: 余额不符合预期
// @author: Kewin Li
// @param: *testing.T t
// @param: Wallet w
// @param: Bitcoin want
func assertMoney(t *testing.T, w Wallet, want Bitcoin) {

	if w.get() != want {
		t.Errorf("w.get()=%s  want=%s \n", w.get(), want)
	}
}

// @func: assertNoErr
// @brief: 无错误返回中断
// @author: Kewin Li
// @param: *testing.T t
// @param: error got
// @param: error want
func assertNoErr(t *testing.T, got error) {
	if got != nil {
		t.Fatal("get an error but not request!\n")
	}
}

// @func: assertErr
// @brief: 错误中断
// @author: Kewin Li
// @param: *testing.T t
// @param: error got
// @param: error want
func assertErr(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("want an error but not get \n")
	}

	// 获取到错误信息字符串
	if got != want {
		t.Errorf("got= %s want= %s\n", got, want)
	}

}
