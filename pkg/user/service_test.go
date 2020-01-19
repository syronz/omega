package user

import (
	"errors"
	"omega/test/core"
	// "omega/utils/glog"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//  go test ./pkg/user/ -run TestServiceCreate -v -count=1 -cover
func checkUserResponse(user, resp User) (err error) {
	if user.Name != resp.Name {
		err = errors.New("Name isn't equal")
		return
	}
	return
}

func TestServiceCreate(t *testing.T) {
	engine := core.StartEngine(&User{})

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
				Password: "123456",
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
