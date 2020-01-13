package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func router(r *gin.Engine, db *gorm.DB) {

	productAPI := InitProductAPI(db)
	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	invoiceAPI := InitInvoiceAPI(db)
	r.POST("/invoices", invoiceAPI.Create)

	userAPI := InitUserAPI(db)
	r.GET("/users", userAPI.FindAll)
	r.GET("/users/:id", userAPI.FindByID)
	r.POST("/users", userAPI.Create)
	r.PUT("/users/:id", userAPI.Update)
	r.DELETE("/users/:id", userAPI.Delete)

}
