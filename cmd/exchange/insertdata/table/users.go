package table

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/domain/sync"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/glog"
)

// InsertUsers for add required users
func InsertUsers(engine *core.Engine) {
	userRepo := basrepo.ProvideUserRepo(engine)
	userService := service.ProvideBasUserService(userRepo)
	users := []basmodel.User{
		{
			FixedCol: types.FixedCol{
				ID:        11,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
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
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
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
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
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
			glog.Fatal("error in saving users", err)
		}
	}

}
