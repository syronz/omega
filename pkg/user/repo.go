package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepostiory(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (p *UserRepository) FindAll() []User {
	var users []User
	p.DB.Find(&users)

	return users
}

func (p *UserRepository) FindByID(id uint) User {
	var user User
	p.DB.First(&user, id)

	return user
}

func (p *UserRepository) Save(user User) User {
	p.DB.Save(&user)

	return user
}

func (p *UserRepository) Delete(user User) {
	p.DB.Delete(&user)
}
