package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// @func: test2
// @brief: 读取文件内容并将文件删除
// @author: Kewin Li
func test2() {

	m_f, err1 := os.OpenFile("./test.json", os.O_RDWR, 0666)
	if err1 != nil {
		fmt.Printf("err1=%s \n", err1.Error())
		return
	}

	m_buf := make([]byte, 100)

	n, err2 := m_f.Read(m_buf)
	if err2 != nil {
		fmt.Printf("err2=%s \n", err2.Error())
		return
	}

	m := make(map[string]interface{}, 4)
	err3 := json.Unmarshal(m_buf[:n], &m)
	if err3 != nil {
		fmt.Printf("err3=%s \n", err3.Error())
		return
	}

	fmt.Printf("read json:%+v \n", m)

	m_f.Close()

	err4 := os.Remove(m_f.Name())
	if err4 != nil {
		fmt.Printf("err4=%s \n", err3.Error())
		return
	}
}

// @func: test1
// @brief: 创建文件并写入数据
// @author: Kewin Li
func test1() {

	m := make(map[string]interface{}, 4)
	m["name"] = "ljk"
	m["age"] = 24
	m["sex"] = byte(1)

	jsonBuf, err1 := json.MarshalIndent(m, "", " ")
	if err1 != nil {
		fmt.Printf("err=%s\n", err1.Error())
		return
	}

	m_f, err2 := os.OpenFile("./test.json", os.O_CREATE|os.O_RDWR, 0666)
	if err2 != nil {
		fmt.Printf("err2=%s \n", err2.Error())
		return
	}

	n, err3 := m_f.Write(jsonBuf)
	if err3 != nil {
		fmt.Printf("err3=%s \n", err3.Error())
		return
	}

	fmt.Printf("file: %s write byte %d \n", m_f.Name(), n)
	m_f.Close()
}

func main() {

	args := os.Args

	switch args[1][0] {
	case '1':
		test1()
	case '2':
		test2()
		//TODO: 还有一种读取方式是按行读取，使用包bufio
	}

}
