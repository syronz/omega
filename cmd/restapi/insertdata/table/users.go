package table

import (
	"fmt"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/param"
	"omega/pkg/dict"
	"omega/pkg/glog"
)

// InsertUsers for add required users
func InsertUsers(engine *core.Engine) {
	userRepo := basrepo.ProvideUserRepo(engine)
	userService := service.ProvideBasUserService(userRepo)
	params := param.Param{}
	users := []basmodel.User{
		{
			ID:       1,
			RoleID:   1,
			Username: engine.Envs[base.AdminUsername],
			Password: engine.Envs[base.AdminPassword],
			Lang:     dict.Ku,
		},
		{
			ID:       2,
			RoleID:   2,
			Username: "cashier",
			Password: "cashier2020",
			Lang:     dict.En,
		},
		{
			ID:       3,
			RoleID:   3,
			Username: "reader",
			Password: "reader2020",
			Lang:     dict.Ar,
		},
	}

	for _, v := range users {
		if _, err := userService.Save(v, params); err != nil {
			fmt.Printf(">>>>>>>>>>>>>>>>> $$$ %+v \n", err)
			glog.Fatal(err)
		}
	}

}
