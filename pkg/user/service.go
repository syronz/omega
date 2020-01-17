package user

import (
	"omega/engine"
	"omega/utils/password"
)

// Service for injecting repo
type Service struct {
	Repo   Repo
	Engine engine.Engine
}

// ProvideService is used in wire
func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// FindAll users
func (p *Service) FindAll() ([]User, error) {
	return p.Repo.FindAll()
}

// FindByID for user
func (p *Service) FindByID(id uint64) (User, error) {
	return p.Repo.FindByID(id)
}

// Save user
func (p *Service) Save(user User) (createdUser User, err error) {
	user.Password, err = password.Hash(user.Password, p.Engine.Environments.Setting.PasswordSalt)
	if err != nil {
		p.Engine.ServerLog.Error(err)
	}
	createdUser, err = p.Repo.Save(user)
	if err != nil {
		p.Engine.ServerLog.Info(err.Error())
	}
	return
}

// Delete user
func (p *Service) Delete(user User) error {
	return p.Repo.Delete(user)
}
