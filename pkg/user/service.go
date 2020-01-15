package user

import (
	// "omega/utils/password"
	"omega/internal/core"
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
	user.Password = "123456"
	p.engine.Debug("INSIDE THE SERVICE .....................................")

	s4 := p.Repo.Save(user)

	return s4
}

func (p *Service) Delete(user User) {
	p.Repo.Delete(user)
}
