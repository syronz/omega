package user

import (
	"omega/internal/core"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repo struct {
	engine core.Engine
}

func ProvideRepo(engine core.Engine) Repo {
	return Repo{engine: engine}
}

func (p *Repo) FindAll() []User {
	var users []User
	p.engine.DB.Find(&users)

	return users
}

func (p *Repo) FindByID(id uint) User {
	var user User
	_ = p.engine.DB.First(&user, id).Error

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

func (p *Repo) Save(user User) (s4 User) {
	p.engine.DB.Save(&user).Scan(&s4)

	// p.engine.Log.Debug(s4)
	// glog.Debug(s4)
	// err = i.DB.Create(&i.Item).Scan(&item).Error

	return s4
}

func (p *Repo) Delete(user User) {
	p.engine.DB.Delete(&user)
}
