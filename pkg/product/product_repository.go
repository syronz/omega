package product

import (
	"omega/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(c config.CFG) ProductRepository {
	return ProductRepository{DB: c.DB}
}

func (p *ProductRepository) FindAll() []Product {
	var products []Product
	p.DB.Find(&products)

	return products
}

func (p *ProductRepository) FindByID(id uint) Product {
	var product Product
	p.DB.First(&product, id)

	return product
}

func (p *ProductRepository) Save(product Product) Product {
	p.DB.Save(&product)

	return product
}

func (p *ProductRepository) Delete(product Product) {
	p.DB.Delete(&product)
}
