package basrepo

import (
	"omega/internal/core"
	"omega/internal/types"
)

// AccessRepo for injecting engine
type AccessRepo struct {
	Engine *core.Engine
}

// ProvideAccessRepo is used in wire
func ProvideAccessRepo(engine *core.Engine) AccessRepo {
	return AccessRepo{Engine: engine}
}

// GetUserResources is used for finding all resources
func (p *AccessRepo) GetUserResources(userID types.RowID) (result string, err error) {
	resources := struct {
		Resources string
	}{}

	err = p.Engine.DB.Table("bas_users").Select("bas_roles.resources").
		Joins("INNER JOIN bas_roles ON bas_users.role_id = bas_roles.id").
		Where("bas_users.id = ?", userID).Scan(&resources).Error

	result = resources.Resources

	return
}
