//+build wireinject

package server

import (
	"rest-gin-gorm/config"
	"rest-gin-gorm/pkg/invoice"
	"rest-gin-gorm/pkg/product"
	"rest-gin-gorm/pkg/user"

	"github.com/google/wire"
)

func initProductAPI(c config.CFG) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}

func initInvoiceAPI(c config.CFG) invoice.InvoiceAPI {
	wire.Build(invoice.ProvideInvoiceRepostiory, invoice.ProvideInvoiceService, invoice.ProvideInvoiceAPI)

	return invoice.InvoiceAPI{}
}

func initUserAPI(c config.CFG) user.UserAPI {
	wire.Build(user.ProvideUserRepostiory, user.ProvideUserService, user.ProvideUserAPI)

	return user.UserAPI{}
}
