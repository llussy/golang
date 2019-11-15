package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func DBInit() {
	var err error
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		"root",
		"root",
		"192.168.1.1:3306",
		"gogin",
		true,
		"Local")

	DB, err = gorm.Open("mysql", config)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("connect mysql success!")
}
