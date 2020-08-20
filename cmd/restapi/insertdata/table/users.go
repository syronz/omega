package table

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/lang"
)

// InsertBasUsers for add required users
func InsertBasUsers(engine *core.Engine) {
	userRepo := basrepo.ProvideBasUserRepo(engine)
	userService := service.ProvideBasUserService(userRepo)
	users := []basmodel.BasUser{
		{
			ID:       1,
			RoleID:   1,
			Username: engine.Envs[base.AdminUsername],
			Password: engine.Envs[base.AdminPassword],
			Language: string(lang.Ku),
		},
		{
			ID:       2,
			RoleID:   2,
			Username: "cashier",
			Password: "cashier",
			Language: string(lang.En),
		},
		{
			ID:       3,
			RoleID:   3,
			Username: "reader",
			Password: "reader",
			Language: string(lang.Ar),
		},
	}

	for _, v := range users {
		if _, err := userService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}
	}

}
