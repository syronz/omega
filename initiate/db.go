package initiate

import (
	"omega/config"
	"omega/pkg/invoice"
	"omega/pkg/product"
	"omega/pkg/user"

	"github.com/jinzhu/gorm"
)

func initDB(c config.CFG, dbType string, url string) *gorm.DB {

	db, err := gorm.Open(dbType, url)
	if err != nil {
		c.Log.Fatalln(err)
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&invoice.Invoice{})
	db.AutoMigrate(&user.User{})

	return db
}
