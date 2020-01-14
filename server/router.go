package server

import (
	"omega/config"

	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine, c config.CFG) {

	productAPI := initProductAPI(c)
	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	invoiceAPI := initInvoiceAPI(c)
	r.POST("/invoices", invoiceAPI.Create)

	userAPI := initUserAPI(c)
	r.GET("/users", userAPI.FindAll)
	r.GET("/users/:id", userAPI.FindByID)
	r.POST("/users", userAPI.Create)
	r.PUT("/users/:id", userAPI.Update)
	r.DELETE("/users/:id", userAPI.Delete)

}
