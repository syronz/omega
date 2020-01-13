package initiate

import (
	"rest-gin-gorm/pkg/invoice"
	"rest-gin-gorm/pkg/product"
	"rest-gin-gorm/pkg/user"

	"github.com/jinzhu/gorm"
)

func initDB(dbType string, url string) *gorm.DB {

	// db, err := gorm.Open("mysql", "root:Qaz1@345@tcp(127.0.0.1:3306)/alpha?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dbType, url)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&invoice.Invoice{})
	db.AutoMigrate(&user.User{})

	return db
}
