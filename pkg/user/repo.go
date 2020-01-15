package user

import (
	"omega/config"

	// "omega/internal/glog"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserRepository struct {
	cfg config.CFG
}

func ProvideUserRepostiory(c config.CFG) UserRepository {
	return UserRepository{cfg: c}
}

func (p *UserRepository) FindAll() []User {
	var users []User
	p.cfg.DB.Find(&users)

	return users
}

func (p *UserRepository) FindByID(id uint) User {
	var user User
	_ = p.cfg.DB.First(&user, id).Error

	user.Extra = struct {
		LastVisit string
		Mark      int
	}{
		"2019",
		-15,
	}

	// glog.Debug(user, id, err)

	return user
}

func (p *UserRepository) Save(user User) (s4 User) {
	p.cfg.DB.Create(&user).Scan(&s4)

	// p.cfg.Log.Debug(s4)
	// glog.Debug(s4)
	// err = i.DB.Create(&i.Item).Scan(&item).Error

	return s4
}

func (p *UserRepository) Delete(user User) {
	p.cfg.DB.Delete(&user)
}
