//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"rest-gin-gorm/product"
)

func InitProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}
