//+build wireinject
package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"rest-gin-gorm/product"
	"rest-gin-gorm/invoice"
)

func InitProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}

func InitInvoiceAPI(db *gorm.DB) invoice.InvoiceAPI {
	wire.Build(invoice.ProvideInvoiceRepostiory, invoice.ProvideInvoiceService, invoice.ProvideInvoiceAPI)

	return invoice.InvoiceAPI{}
}
