package main

import (
	// "os"

	// "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "rest-gin-gorm/internal"
	"rest-gin-gorm/internal/one"
	"rest-gin-gorm/pkg/invoice"
	"rest-gin-gorm/pkg/product"
	"rest-gin-gorm/pkg/user"
	"rest-gin-gorm/server"
)

func initDB() *gorm.DB {

	// internal.InternalPing()
	one.OnePing()

	db, err := gorm.Open("mysql", "root:Qaz1@345@tcp(127.0.0.1:3306)/alpha?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&invoice.Invoice{})
	db.AutoMigrate(&user.User{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	s := server.Setup(db)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
