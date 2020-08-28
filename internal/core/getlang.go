package core

import (
	"omega/pkg/dict"

	"github.com/gin-gonic/gin"
)

// GetLang return suitable language according to 1.query, 2.JWT, 3.environment
func GetLang(c *gin.Context, engine *Engine) dict.Lang {
	var langLevel dict.Lang

	// priority 4: get from environment
	langLevel = dict.Lang(engine.Envs[DefaultLang])

	// priority 3: get lang from company default language in the database
	// TODO: complete this part

	// priority 2
	langJWT, ok := c.Get("LANGUAGE")
	if ok {
		langLevel = langJWT.(dict.Lang)
	}

	// priority 1
	langQuery := c.Query("lang")
	if langQuery != "" {
		langLevel = dict.Lang(langQuery)
	}

	switch langLevel {
	case dict.En:
		return dict.En
	case dict.Ku:
		return dict.Ku
	case dict.Ar:
		return dict.Ar
	}

	return dict.Ku
}
