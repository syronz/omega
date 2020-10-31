package table

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/glog"
)

// InsertUsers for add required users
func InsertUsers(engine *core.Engine) {
	userRepo := basrepo.ProvideUserRepo(engine)
	userService := service.ProvideBasUserService(userRepo)

	// reset the users table
	userRepo.Engine.DB.Table(basmodel.UserTable).Unscoped().Delete(basmodel.User{})

	users := []basmodel.User{
		{
			FixedCol: types.FixedCol{
				ID:        11,
				CompanyID: 1001,
				NodeID:    101,
			},
			RoleID:   1,
			Name:     engine.Envs[base.AdminUsername],
			Username: engine.Envs[base.AdminUsername],
			Password: engine.Envs[base.AdminPassword],
			Lang:     dict.Ku,
		},
		{
			FixedCol: types.FixedCol{
				ID:        12,
				CompanyID: 1001,
				NodeID:    101,
			},
			RoleID:   2,
			Name:     "cashier",
			Username: "cashier",
			Password: "cashier2020",
			Lang:     dict.En,
		},
		{
			FixedCol: types.FixedCol{
				ID:        13,
				CompanyID: 1001,
				NodeID:    101,
			},
			RoleID:   3,
			Name:     "reader",
			Username: "reader",
			Password: "reader2020",
			Lang:     dict.Ar,
		},
	}

	for _, v := range users {
		if _, err := userService.Save(v); err != nil {
			glog.Fatal(err)
		}
	}

}
