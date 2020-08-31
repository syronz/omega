package core

import (
	"omega/internal/types"

	"github.com/jinzhu/gorm"
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

// Debug print struct with details with logrus ability
// func (e *Engine) Debug2(objs ...interface{}) {
// 	for _, v := range objs {
// 		parts := make(map[string]interface{}, 2)
// 		parts["type"] = fmt.Sprintf("%T", v)
// 		parts["value"] = v
// 		dataInJSON, _ := json.Marshal(parts)

// 		e.ServerLog.Debug(string(dataInJSON))
// 	}
// }

// CheckError print all errors which happened inside the services, mainly they just have
// an error and a message
// func (e *Engine) CheckError2(err error, message string, data ...interface{}) {
// 	if err != nil {
// 		e.ServerLog.WithFields(logrus.Fields{
// 			"err": err.Error(),
// 		}).Error(message)
// 		if data != nil {
// 			e.Debug2(data...)
// 		}
// 	}
// }

// CheckInfo print all errors which happened inside the services, mainly they just have
// an error and a message
// func (e *Engine) CheckInfo(err error, message string, data ...interface{}) {
// 	if err != nil {
// 		e.ServerLog.WithFields(logrus.Fields{
// 			"err": err.Error(),
// 		}).Info(message)
// 		if data != nil {
// 			e.Debug2(data...)
// 		}
// 	}
// }

// T Translating the term
// func (e *Engine) T2(str string, language lang.Lang, params ...interface{}) string {
// 	return e.Dict.Translate(str, language, params...)
// }

// SafeT Translating the term and if the word won't exist return false
// func (e *Engine) SafeT(str string, language lang.Lang, params ...interface{}) (string, bool) {
// 	return e.Dict.SafeTranslate(str, language, params...)
// }
