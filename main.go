package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

//Product is represent of a product
type Product struct {
	ID    int    `json:"id"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

//Result is an array of product
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/go_rest_api_crud?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Oops.. Connection failed", err)
	} else {
		log.Println("Connection established!")
	}

	db.AutoMigrate(&Product{})
}
