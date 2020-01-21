package shared

import (
	"errors"
	"omega/engine"
	"strings"

	"github.com/gin-gonic/gin"
)

// Shared is used a builder
type Shared struct {
	NewModel   interface{}
	ExistModel interface{}
	SavedModel interface{}
	Context    *gin.Context
	Engine     engine.Engine
	Service    interface{}
	err        error
	Message    string
	part       string
}

func (s Shared) Error() error {
	return s.err
}

// New initiate the builder
func New(context *gin.Context, engine engine.Engine, service interface{}, part string) Shared {
	return Shared{
		Context: context,
		Engine:  engine,
		Service: service,
		part:    part,
	}
}

// Bind is used for binding models with JSON
func (s Shared) Bind(v interface{}) Shared {
	if s.err != nil {
		return s
	}

	err := s.Context.BindJSON(v)
	if err != nil {
		s.err = err
		s.Message = "Error in binding " + s.part
		return s
	}

	s.NewModel = v
	return s
}

// CheckAccess by using ACL control user's permission
func (s Shared) CheckAccess(resource string) Shared {
	if s.err != nil {
		return s
	}
	if !s.Engine.CheckAccess(s.Context, resource) {
		s.err = errors.New("Access denied")
		s.Message = "You don't have permission " + s.part
		s.Engine.Record(s.Context, s.part+"-forbidden", s.ExistModel, s.NewModel)
	}
	return s
}

// SaveFunc is a template for returning exact model
type SaveFunc func(model interface{}) (interface{}, error)

// Save used for creating and updating
func (s Shared) Save(saveFunc SaveFunc) Shared {
	if s.err != nil {
		return s
	}
	s.SavedModel, s.err = saveFunc(s.NewModel)
	if s.err != nil {
		errMessage := s.err.Error()
		if strings.Contains(strings.ToUpper(errMessage), "DUPLICATE") {
			s.Message = "Duplication happened for " + s.part
		} else {
			s.Message = "Error in saving " + s.part
		}
	}

	return s
}

// Record the process to the database
func (s Shared) Record() Shared {
	if s.err != nil {
		return s
	}

	s.Engine.Record(s.Context, s.part, s.ExistModel, s.NewModel)

	return s

}
