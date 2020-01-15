package user

import (
	// "omega/utils/password"
	"omega/internal/glog"
)

type UserService struct {
	UserRepository UserRepository
}

func ProvideUserService(p UserRepository) UserService {
	return UserService{UserRepository: p}
}

func (p *UserService) FindAll() []User {
	return p.UserRepository.FindAll()
}

func (p *UserService) FindByID(id uint) User {
	return p.UserRepository.FindByID(id)
}

func (p *UserService) Save(user User) User {
	user.Password = "123456"
	glog.Debug(p.UserRepository.cfg.ENV.Setting.PasswordSalt)
	s4 := p.UserRepository.Save(user)

	return s4
}

func (p *UserService) Delete(user User) {
	p.UserRepository.Delete(user)
}
