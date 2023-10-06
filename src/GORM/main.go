package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Product struct {
	gorm.Model
	Name  string
	Price uint
}

// @func: test1
// @date: 2023-10-07 03:34:40
// @brief: GORM框架官方例子
// @author: Kewin Li
func test1() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("open mysql database err! %s \n", err)
		return
	}

	// 自动迁移(自动创建表格)
	db.AutoMigrate(&Product{})

	// 插入记录
	ctx := db.Create(&Product{Name: "111", Price: 100})
	if ctx.Error != nil {
		fmt.Printf("mysql create err! %s \n", ctx.Error)
		return
	}

	ctx = db.Create(&Product{Name: "222", Price: 200})
	if ctx.Error != nil {
		fmt.Printf("mysql create err! %s \n", ctx.Error)
		return
	}

	readProduct := Product{}
	//// 通过GORM提供的主键查询
	db.First(&readProduct, 1)
	fmt.Printf("read1: %v \n", readProduct)

	readProduct = Product{}
	//// 通过条件查询
	db.Where("name = ?", "222").Find(&readProduct)
	fmt.Printf("read2: %v \n", readProduct)

	//修改
	db.Model(&readProduct).Updates(&Product{Name: "333", Price: 300})

	//删除(记录还存在，被打上删除标记)
	db.Delete(&Product{}, 1)

}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("input xx.exe [number] \n")
		return
	}

	switch args[1][0] {
	case '1':
		test1()
	}

}
