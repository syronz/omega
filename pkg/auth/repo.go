package auth

import (
	"omega/engine"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

// Repo for injecting engine
type Repo struct {
	Engine engine.Engine
}

// ProvideRepo is used in wire
func ProvideRepo(Engine engine.Engine) Repo {
	return Repo{Engine: Engine}
}

// Logout enable the force login TODO: not implemented yet
func (p *Repo) Logout(user Auth) error {
	return p.Engine.DB.Find(&user).Error
}

// Login for entering the system
func (p *Repo) Login(auth Auth) (user Auth, err error) {
	err = p.Engine.DB.First(&auth).Error
	return
}
