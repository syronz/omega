package user

import (
	"errors"
	"omega/test/core"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//  go test ./pkg/user/ -run TestServiceCreate -v -count=1 -cover
func checkUserResponse(user, resp User) (err error) {
	if user.Name != resp.Name {
		err = errors.New("Name isn't equal")
		return
	}
	if user.Username != resp.Username {
		err = errors.New("Username isn't equal")
		return
	}
	if user.Phone != resp.Phone {
		err = errors.New("Phone isn't equal")
		return
	}
	if user.Password != "" {
		err = errors.New("Password isn't empty")
		return
	}
	return
}

func TestServiceCreate(t *testing.T) {
	engine := core.StartEngine(&User{})
	engine.DB.Exec("TRUNCATE TABLE users;")

	repo := ProvideRepo(engine)
	service := ProvideService(repo)

	samples := []struct {
		in  User
		out User
		err error
	}{
		{
			in: User{
				Name:     "Diako",
				Username: "diako",
				Password: "123456",
				Phone:    "07505149171",
			},
			out: User{
				Name:     "Diako",
				Username: "diako",
				Password: "",
				Phone:    "07505149171",
			},
			err: nil,
		},
	}

	for _, v := range samples {
		result, err := service.Save(v.in)
		if err != nil {
			t.Error(err)
		}
		if err = checkUserResponse(v.out, result); err != nil {
			t.Error(err)
		}
	}

}
