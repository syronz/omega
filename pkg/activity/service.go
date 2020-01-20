package activity

import (
	"fmt"
	"omega/engine"
	"omega/internal/models"
	"omega/internal/param"
)

// Service for injecting auth repo
type Service struct {
	Repo   Repo
	Engine engine.Engine
}

// ProvideService for activity is used in wire
func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// List of activities
func (p *Service) List(params param.Param) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	data["activities"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "activities list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "activities count")

	return
}

// FindByID for activity
func (p *Service) FindByID(id uint64) (activity models.Activity, err error) {
	activity, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Activity with id %v", id))

	return
}
