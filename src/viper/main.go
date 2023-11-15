// Package main
// @Description: viper练习测试
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// @func: test1
// @date: 2023-11-15 21:16:37
// @brief: viper用法1
// @author: Kewin Li
func test1() {

	viper.SetConfigName("test_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("viper")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	val := viper.Get("test.key")

	fmt.Printf("%T %v \n", val, val)

}

func test2() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("viper/test_config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	val := viper.Get("test.key")

	fmt.Printf("%T %v \n", val, val)
}

func main() {
	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test1()
	}
}
