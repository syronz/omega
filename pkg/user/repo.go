package user

import (
	"omega/engine"
)

type Repo struct {
	Engine engine.Engine
}

func ProvideRepo(engine engine.Engine) Repo {
	return Repo{Engine: engine}
}

func (p *Repo) FindAll() (users []User, err error) {
	err = p.Engine.DB.Find(&users).Error
	return
}

func (p *Repo) FindByID(id uint) (user User, err error) {
	err = p.Engine.DB.First(&user, id).Error

	user.Extra = struct {
		LastVisit string
		Mark      int
	}{
		"2019",
		-15,
	}

	return
}

func (p *Repo) FindByUsername(username string) (user User, err error) {
	err = p.Engine.DB.Where("username = ?", username).First(&user).Error
	return
}

func (p *Repo) Save(user User) (u User, err error) {
	err = p.Engine.DB.Save(&user).Scan(&u).Error
	return
}

func (p *Repo) Delete(user User) (err error) {
	err = p.Engine.DB.Delete(&user).Error
	return
}
