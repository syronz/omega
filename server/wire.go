//+build wireinject

package server

import (
	"rest-gin-gorm/pkg/invoice"
	"rest-gin-gorm/pkg/product"
	"rest-gin-gorm/pkg/user"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func InitProductAPI(db *gorm.DB) product.ProductAPI {
	wire.Build(product.ProvideProductRepostiory, product.ProvideProductService, product.ProvideProductAPI)

	return product.ProductAPI{}
}

func InitInvoiceAPI(db *gorm.DB) invoice.InvoiceAPI {
	wire.Build(invoice.ProvideInvoiceRepostiory, invoice.ProvideInvoiceService, invoice.ProvideInvoiceAPI)

	return invoice.InvoiceAPI{}
}

func InitUserAPI(db *gorm.DB, log *logrus.Logger) user.UserAPI {
	wire.Build(user.ProvideUserRepostiory, user.ProvideUserService, user.ProvideUserAPI)

	return user.UserAPI{}
}
