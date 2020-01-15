//+build wireinject

package server

import (
	"omega/config"
	"omega/pkg/invoice"
	"omega/pkg/product"
	"omega/pkg/sample4"
	"omega/pkg/sample5"
	"omega/pkg/user"

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

func initSample4API(c config.CFG) sample4.Sample4API {
	wire.Build(sample4.ProvideSample4Repostiory, sample4.ProvideSample4Service, sample4.ProvideSample4API)

	return sample4.Sample4API{}
}

func initSample5API(c config.CFG) sample5.Sample5API {
	wire.Build(sample5.ProvideSample5Repostiory,
		sample5.ProvideSample5Service,
		sample5.ProvideSample5Controller,
		sample5.ProvideSample5API)

	return sample5.Sample5API{}
}
