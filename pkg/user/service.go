package user

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
	p.UserRepository.Save(user)

	return user
}

func (p *UserService) Delete(user User) {
	p.UserRepository.Delete(user)
}
