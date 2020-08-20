package node

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

// ProvideService for node is used in wire
func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// FindAll nodes
func (p *Service) FindAll() (nodes []Node, err error) {
	nodes, err = p.Repo.FindAll()
	p.Engine.CheckError(err, "all nodes")
	return
}

// List of nodes
func (p *Service) List(params param.Param) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	data["nodes"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "nodes list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "nodes count")

	return
}

// FindByID for node
func (p *Service) FindByID(id uint64) (node Node, err error) {
	node, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Node with id %v", id))

	return
}

// Save node
func (p *Service) Save(node Node) (createdNode Node, err error) {
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", node))

	createdNode, err = p.Repo.Save(node)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving node for %+v", node))

	return
}

// func (p *Service) SaveSimple(model interface{}) (createdNode interface{}, err error) {
// 	node := *(model.(*Node))
// 	node.Password, err = password.Hash(node.Password, p.Engine.Environments.Setting.PasswordSalt)
// 	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", node))

// 	createdNode, err = p.Repo.Save(node)
// 	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving node for %+v", node))

// 	// createdNode.Password = ""

// 	return
// }

// Delete node
func (p *Service) Delete(node Node) error {
	return p.Repo.Delete(node)
}
