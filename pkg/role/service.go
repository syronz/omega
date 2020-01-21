package role

import (
	"fmt"
	"omega/engine"
	"omega/internal/param"
)

// Service for injecting auth repo
type Service struct {
	Repo   Repo
	Engine engine.Engine
}

// ProvideService for role is used in wire
func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// FindAll roles
func (p *Service) FindAll() (roles []Role, err error) {
	roles, err = p.Repo.FindAll()
	p.Engine.CheckError(err, "all roles")
	return
}

// List of roles
func (p *Service) List(params param.Param) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	data["roles"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "roles list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "roles count")

	return
}

// FindByID for role
func (p *Service) FindByID(id uint64) (role Role, err error) {
	role, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Role with id %v", id))

	return
}

// Save role
func (p *Service) Save(role Role) (createdRole Role, err error) {
	createdRole, err = p.Repo.Save(role)
	// p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving role for %+v", role))

	return
}

// Delete role
func (p *Service) Delete(role Role) error {
	return p.Repo.Delete(role)
}
