package migrate

import (
	"omega/engine"
	"omega/internal/models"
	"omega/pkg/user"
)

func migrateUsers(e engine.Engine) {
	userRepo := user.ProvideRepo(e)
	userService := user.ProvideService(userRepo)
	users := []user.User{
		{
			FixedCol: models.FixedCol{
				ID: 1,
			},
			RoleID:   1,
			Name:     "Admin",
			Username: "admin",
			Password: "omega",
		},
	}

	for _, v := range users {
		if _, err := userService.Save(v); err != nil {
			e.ServerLog.Fatal(err)
		}

	}

}
