// Package main
// @Description: viper练习测试
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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

// @func: test1
// @date: 2023-11-15 21:16:37
// @brief: viper用法2
// @author: Kewin Li
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

// @func: test3
// @date: 2023-11-18 10:38:24
// @brief: 使用etcd远程配置
// @author: Kewin Li
func test3() {
	err := viper.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", "/kitbook")
	if err != nil {
		panic(err)
	}

	//viper.SetConfigName("test_config.yaml")
	viper.SetConfigType("yaml")

	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	val := viper.Get("test.key")

	fmt.Printf("%T %v \n", val, val)
}

// @func: test4
// @date: 2023-11-18 14:32:08
// @brief: 监听配置变更
// @author: Kewin Li
func test4() {
	viper.SetConfigFile("viper/test_config.yaml")
	viper.SetConfigType("yaml")

	//注意: 调用顺序不能改变
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("test.key= %v \n", viper.GetString("test.key"))
	})

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	for {

	}

}

func main() {
	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	case '3':
		test3()
	case '4':
		test4()
	}
}
