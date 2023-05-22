package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// IT @brief: !!!注意通过结构体生成Json成员必须首字母大写
// type IT struct {
// 	Company  string   `json:"company"` //将字段Company编为company
// 	SubJects []string `json:"-"`       //代表不进行json编码
// 	Isok     bool     `json:",string"` //将val值编为string类型
// 	Price    float64
// }

type IT struct {
	Company  string
	SubJects []string
	Isok     bool
	Price    float64
}

// @func: test2
// @brief: 通过map生成json/从json解析回map
// @author: Kewin Li
func test2() {

	m := make(map[string]interface{}, 4)

	m["company"] = "6666"
	m["subjects"] = []string{"go", "C++"}
	m["isok"] = true
	m["price"] = 3.1561

	res, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Printf("err=%s \n", err.Error())
		return
	}

	fmt.Printf("res=%+v \n", string(res))

	fmt.Printf("---------------------------\n")

	// 2. 从json解析回map

	jsonBuf := `{
		"company": "6666",
		"subjects": [
		 "go",
		 "C++"
		],
		"isok": true,
		"price": 3.1561
}`
	var parse_map map[string]interface{}

	json.Unmarshal([]byte(jsonBuf), &parse_map)

	fmt.Printf("parse_map=%+v \n", parse_map)

	str := parse_map["company"]
	fmt.Printf("str =%+v \n", str)

}

// @func: test1
// @brief: 通过 结构体生成json/从json解析回结构体
// @author: Kewin Li
func test1() {

	it1 := IT{"test1", []string{"C++", "go", "pyhon"}, true, 1.2156}

	// 1. 进行json数据编码
	buf, err := json.Marshal(it1) //普通编码
	if err != nil {
		fmt.Printf("err=%s \n", err.Error())
		return
	}
	fmt.Printf("buf=%+v \n", string(buf))

	fmt.Printf("--------------------\n")

	buf, err = json.MarshalIndent(it1, "", " ") //格式化编码
	if err != nil {
		fmt.Printf("err=%s \n", err.Error())
		return
	}
	fmt.Printf("buf=%+v \n", string(buf))

	fmt.Printf("-----------------------------\n")

	// 2. 从json解析回结构体
	jsonBuf := `{
		"company": "6666",
		"subjects": [
		 "go",
		 "C++"
		],
		"isok": true,
		"price": 3.1561
}`

	var parse_it IT

	err = json.Unmarshal([]byte(jsonBuf), &parse_it)
	if err != nil {
		fmt.Printf("err=%s \n", err.Error())
		return
	}

	fmt.Printf("parse_it=%+v \n", parse_it)
}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
	}

}
