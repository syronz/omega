package user

import (
	"fmt"
	"omega/utils/password"
)

// BuildSave is used for saving user by following builder design pattern
func (p *Service) BuildSave(model interface{}) (createdUser interface{}, err error) {
	user := *(model.(*User))
	user.Password, err = password.Hash(user.Password, p.Engine.Environments.Setting.PasswordSalt)
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	createdUser, err = p.Repo.Save(user)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))

	return
}
