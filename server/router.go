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

	sample4API := initSample4API(c)
	r.GET("/sample4", sample4API.FindAll)
	r.GET("/sample4/:id", sample4API.FindByID)
	r.POST("/sample4", sample4API.Create)
	r.PUT("/sample4/:id", sample4API.Update)
	r.DELETE("/sample4/:id", sample4API.Delete)

	sample5API := initSample5API(c)
	r.GET("/sample5", sample5API.FindAll)
	r.GET("/sample5/:id", sample5API.FindByID)
	r.POST("/sample5", sample5API.Create)
	r.PUT("/sample5/:id", sample5API.Update)
	r.DELETE("/sample5/:id", sample5API.Delete)
}
