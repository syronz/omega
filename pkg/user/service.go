package user

import (
	"omega/internal/core"
	"omega/utils/password"
)

type Service struct {
	Repo   Repo
	engine core.Engine
}

func ProvideService(p Repo) Service {
	return Service{Repo: p, engine: p.engine}
}

func (p *Service) FindAll() []User {
	return p.Repo.FindAll()
}

func (p *Service) FindByID(id uint) User {
	return p.Repo.FindByID(id)
}

func (p *Service) Save(user User) User {
	user.Password, _ = password.Hash(user.Password,
		p.engine.Environments.Setting.PasswordSalt)

	s4 := p.Repo.Save(user)
	s4.Password = ""

	return s4
}

func (p *Service) Delete(user User) {
	p.Repo.Delete(user)
}
