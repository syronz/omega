package core

import (
	"omega/internal/core/lang"

	"github.com/gin-gonic/gin"
)

// GetLang return suitable language according to 1.query, 2.JWT, 3.environment
func GetLang(c *gin.Context, engine *Engine) lang.Language {
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
		return lang.En
	case "ku":
		return lang.Ku
	case "ar":
		return lang.Ar
	}

	return lang.Ku
}
