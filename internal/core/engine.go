package core

import (
	"omega/internal/types"

	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
	goaes "github.com/syronz/goAES"
)

// Engine to keep all database connections and
// logs configuration and environments and etc
type Engine struct {
	DB         *gorm.DB
	ActivityDB *gorm.DB
	APILog     *logrus.Logger
	Envs       types.Envs
	AES        goaes.BuildModel
	Setting    map[types.Setting]types.SettingMap
}

// Clone return an engine just like before
func (e *Engine) Clone() *Engine {
	var DB gorm.DB
	DB = *e.DB
	var clonedEngine Engine
	clonedEngine = *e
	clonedEngine.DB = &DB

	return &clonedEngine
}
