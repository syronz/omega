package node

import (
	"omega/engine"
	"omega/internal/param"
	"omega/utils/search"
)

// Repo for injecting engine
type Repo struct {
	Engine engine.Engine
}

var pattern = `(nodes.name LIKE '%[1]v%%' OR
		nodes.id = '%[1]v' OR
		nodes.phone LIKE '%[1]v%%')`

// ProvideRepo is used in wire
func ProvideRepo(engine engine.Engine) Repo {
	return Repo{Engine: engine}
}

// FindAll nodes
func (p *Repo) FindAll() (nodes []Node, err error) {
	err = p.Engine.DB.Select("id, name").Find(&nodes).Error
	return
}

// List nodes
func (p *Repo) List(params param.Param) (nodes []Node, err error) {
	err = p.Engine.DB.Select(params.Select).
		Where(search.Parse(params, pattern)).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&nodes).Error

	return
}

// Count nodes
func (p *Repo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("nodes").
		Select(params.Select).
		Where("deleted_at = null").
		Where(search.Parse(params, pattern)).
		Count(&count).Error
	return
}

// FindByID for node
func (p *Repo) FindByID(id uint64) (node Node, err error) {
	err = p.Engine.DB.First(&node, id).Error

	return
}

// FindByNodename for node
func (p *Repo) FindByNodename(nodename string) (node Node, err error) {
	err = p.Engine.DB.Where("nodename = ?", nodename).First(&node).Error
	return
}

// Save node
func (p *Repo) Save(node Node) (u Node, err error) {
	err = p.Engine.DB.Save(&node).Scan(&u).Error
	return
}

// Delete node
func (p *Repo) Delete(node Node) (err error) {
	err = p.Engine.DB.Delete(&node).Error
	return
}
