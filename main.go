package main

import (
	// "os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "rest-gin-gorm/internal"
	"rest-gin-gorm/internal/one"
	"rest-gin-gorm/invoice"
	"rest-gin-gorm/product"
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

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := InitProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	invoiceAPI := InitInvoiceAPI(db)
	r.POST("/invoices", invoiceAPI.Create)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
