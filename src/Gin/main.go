package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// @func: test1
// @date: 2023-10-04 17:07:13
// @brief: Gin入门-基本操作
// @author: Kewin Li
func test1() {
	//1. 获取默认Web引擎(engine)
	server := gin.Default()

	//2. 进行路由注册、接入middleware中间件等核心功能
	server.GET("/hello/:name/", func(ctx *gin.Context) {

		// 2.1 参数路由
		name := ctx.Param("name")
		fmt.Printf("param routine: name[%s] \n", name)

		// 2.2 查询参数
		id := ctx.Query("id")
		fmt.Printf("query param: id[%s] \n", id)

		// 2.3 返回响应
		ctx.String(http.StatusOK, fmt.Sprintf("hello welcom to go Web! name[%s], id[%s]\n", name, id))
	})

	//3. 在指定的[IP:端口上]开启Web 服务器
	server.Run(":8080")
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("input xxx.exe [number] \n")
		return
	}

	switch args[1][0] {
	case '1':
		test1()
	}

}
