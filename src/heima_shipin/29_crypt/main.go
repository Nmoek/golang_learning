package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	str1 := "123456"
	crypt_str1, err := bcrypt.GenerateFromPassword([]byte(str1), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt err! %s \n", err)
		return
	}

	fmt.Printf("%v \n", string(crypt_str1))

	// 值得注意的是, 该种加密方式无法逆解析，
	// 只能通过明文再一次加密后 的加密中间值判断
	err = bcrypt.CompareHashAndPassword(crypt_str1, []byte(str1))
	if err != nil {
		fmt.Printf("CompareHashAndPassword err! %s \n", err)
		return
	}

}
