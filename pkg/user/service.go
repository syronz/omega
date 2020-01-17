package user

import (
	"omega/engine"
	"omega/utils/password"
)

type Service struct {
	Repo   Repo
	Engine engine.Engine
}

func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

func (p *Service) FindAll() ([]User, error) {
	return p.Repo.FindAll()
}

func (p *Service) FindByID(id uint) (User, error) {
	return p.Repo.FindByID(id)
}

func (p *Service) Save(user User) (createdUser User, err error) {
	user.Password, _ = password.Hash(user.Password, p.Engine.Environments.Setting.PasswordSalt)
	createdUser, err = p.Repo.Save(user)
	createdUser.Password = "***"
	return
}

func (p *Service) Delete(user User) error {
	return p.Repo.Delete(user)
}
