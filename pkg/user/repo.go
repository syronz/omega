package user

import (
	"omega/engine"
)

// Repo for injecting engine
type Repo struct {
	Engine engine.Engine
}

// ProvideRepo is used in wire
func ProvideRepo(engine engine.Engine) Repo {
	return Repo{Engine: engine}
}

// FindAll users
func (p *Repo) FindAll() (users []User, err error) {
	err = p.Engine.DB.Select("id, name").Find(&users).Error
	return
}

// List users
func (p *Repo) List() (users []User, err error) {
	err = p.Engine.DB.Select("id, name").Find(&users).Error
	return
}

// FindByID for user
func (p *Repo) FindByID(id uint64) (user User, err error) {
	err = p.Engine.DB.First(&user, id).Error

	return
}

// FindByUsername for user
func (p *Repo) FindByUsername(username string) (user User, err error) {
	err = p.Engine.DB.Where("username = ?", username).First(&user).Error
	return
}

// Save user
func (p *Repo) Save(user User) (u User, err error) {
	err = p.Engine.DB.Save(&user).Scan(&u).Error
	return
}

// Delete user
func (p *Repo) Delete(user User) (err error) {
	err = p.Engine.DB.Delete(&user).Error
	return
}
