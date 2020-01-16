package auth

import (
	"omega/internal/core"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repo struct {
	Engine core.Engine
}

func ProvideRepo(Engine core.Engine) Repo {
	return Repo{Engine: Engine}
}

func (p *Repo) Logout() []Auth {
	var users []Auth
	p.Engine.DB.Find(&users)

	return users
}

func (p *Repo) Login(id uint) Auth {
	var auth Auth
	_ = p.Engine.DB.First(&auth, id).Error

	// glog.Debug(auth, id, err)

	return auth
}
