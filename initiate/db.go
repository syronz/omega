package initiate

import (
	"omega/config"
	"omega/pkg/invoice"
	"omega/pkg/product"
	"omega/pkg/sample4"
	"omega/pkg/sample5"
	"omega/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func initDB(c config.CFG, dbType string, url string) *gorm.DB {

	db, err := gorm.Open(dbType, url)
	if err != nil {
		c.Log.Fatalln(err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	db.AutoMigrate(&product.Product{})
	db.AutoMigrate(&invoice.Invoice{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&sample4.Sample4{})
	db.AutoMigrate(&sample5.Sample5{})

	return db
}
