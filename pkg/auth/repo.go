package auth

import (
	"omega/engine"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repo struct {
	Engine engine.Engine
}

func ProvideRepo(Engine engine.Engine) Repo {
	return Repo{Engine: Engine}
}

func (p *Repo) Logout(user Auth) error {
	return p.Engine.DB.Find(&user).Error
}

func (p *Repo) Login(auth Auth) (user Auth, err error) {
	err = p.Engine.DB.First(&auth).Error
	return
}