package core

import (
	"omega/pkg/dict"

	"github.com/gin-gonic/gin"
)

// GetLang return suitable language according to 1.query, 2.JWT, 3.environment
func GetLang(c *gin.Context, engine *Engine) dict.Language {
	var langLevel string

	// priority 4: get from environment
	langLevel = engine.Envs[DefaultLanguage]

	// priority 3: get lang from company default language in the database
	// TODO: complete this part

	// priority 2
	langJWT, ok := c.Get("LANGUAGE")
	if ok {
		langLevel = langJWT.(string)
	}

	// priority 1
	langQuery := c.Query("lang")
	if langQuery != "" {
		langLevel = langQuery
	}

	switch langLevel {
	case "en":
		return dict.En
	case "ku":
		return dict.Ku
	case "ar":
		return dict.Ar
	}

	return dict.Ku
}
