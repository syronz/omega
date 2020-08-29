package htmapi

import (
	"omega/internal/core"

	"github.com/gin-gonic/gin"
)

// ErrDescAPI for injectiong engine
type ErrDescAPI struct {
	Engine *core.Engine
}

// GenErrDescAPI is used in wire for creating associate
func GenErrDescAPI(engine *core.Engine) ErrDescAPI {
	return ErrDescAPI{
		Engine: engine,
	}
}

func (p *ErrDescAPI) List(c *gin.Context) {

	c.JSON(200, gin.H{"name": "diako"})

}
