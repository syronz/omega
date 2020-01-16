package auth

import (
	"omega/internal/core"
	"omega/pkg/user"
	// "omega/utils/password"
)

type Service struct {
	Repo   Repo
	Engine core.Engine
}

func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

func (p *Service) Logout() []Auth {
	return p.Repo.Logout()
}

func (p *Service) Login(auth Auth) (result Auth, err error) {
	userRepo := user.Repo{Engine: p.Engine}
	user, err := userRepo.FindByUsername(auth.Username)
	p.Engine.Debug(user, err)

	return
}
