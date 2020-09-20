package kernel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"omega/domain/base"
	"omega/internal/core"
	"omega/internal/types"
	"os"
	"path/filepath"
	"runtime"
)

type TestEnvs struct {
	Core struct {
		Port                 string `json:"port"`
		Addr                 string `json:"addr"`
		DatabaseDataDSN      string `json:"database_data_dsn"`
		DatabaseDataType     string `json:"database_data_type"`
		DatabaseDataLog      string `json:"database_data_log"`
		DatabaseActivityDSN  string `json:"database_activity_dsn"`
		DatabaseActivityType string `json:"database_activity_type"`
		DatabaseActivityLog  string `json:"database_activity_log"`
		AutoMigrate          string `json:"auto_migrate"`
		ServerLogFormat      string `json:"server_log_format"`
		ServerLogOutput      string `json:"server_log_output"`
		ServerLogLevel       string `json:"server_log_level"`
		ServerLogJSONIndent  string `json:"server_log_json_indent"`
		APILogFormat         string `json:"api_log_format"`
		APILogOutput         string `json:"api_log_output"`
		APILogLevel          string `json:"api_log_level"`
		APILogJSONIndent     string `json:"api_log_json_indent"`
		TermsPath            string `json:"terms_path"`
		DefaultLang          string `json:"default_language"`
		TranslateInBackend   string `json:"translate_in_backend"`
		ExcelMaxRows         string `json:"excel_max_rows"`
	} `json:"core"`
	Base struct {
		PasswordSalt  string `json:"password_salt"`
		JWTSecretKey  string `json:"jwt_secret_key"`
		JWTExpiration string `json:"jwt_expiration"`
	} `json:"base"`
}

// LoadTestEnv is used for testing environment
func LoadTestEnv() *core.Engine {
	var engine core.Engine
	// engine := new(core.Engine)
	var testEnvs TestEnvs

	_, dir, _, _ := runtime.Caller(0)
	configJSON := filepath.Join(filepath.Dir(dir), "../../test/", "tdd.json")

	jsonFile, err := os.Open(configJSON)

	if err != nil {
		log.Fatalln(err, "can't open the config file", configJSON)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// err = json.Unmarshal(byteValue, &engine.Env)
	err = json.Unmarshal(byteValue, &testEnvs)
	if err != nil {
		log.Fatalln(err, "error in unmarshal JSON")
	}

	var envs types.Envs
	envs = make(types.Envs, 28)
	envs[core.Port] = testEnvs.Core.Port
	envs[core.Addr] = testEnvs.Core.Addr
	envs[core.DatabaseDataDSN] = testEnvs.Core.DatabaseDataDSN
	envs[core.DatabaseDataType] = testEnvs.Core.DatabaseDataType
	envs[core.DatabaseDataLog] = testEnvs.Core.DatabaseDataLog
	envs[core.DatabaseActivityDSN] = testEnvs.Core.DatabaseActivityDSN
	envs[core.DatabaseActivityType] = testEnvs.Core.DatabaseActivityType
	envs[core.DatabaseActivityLog] = testEnvs.Core.DatabaseActivityLog
	envs[core.AutoMigrate] = testEnvs.Core.AutoMigrate
	envs[core.ServerLogFormat] = testEnvs.Core.ServerLogFormat
	envs[core.ServerLogOutput] = testEnvs.Core.ServerLogOutput
	envs[core.ServerLogLevel] = testEnvs.Core.ServerLogLevel
	envs[core.ServerLogJSONIndent] = testEnvs.Core.ServerLogJSONIndent
	envs[core.APILogFormat] = testEnvs.Core.APILogFormat
	envs[core.APILogOutput] = testEnvs.Core.APILogOutput
	envs[core.APILogLevel] = testEnvs.Core.APILogLevel
	envs[core.APILogJSONIndent] = testEnvs.Core.APILogJSONIndent
	envs[core.TermsPath] = testEnvs.Core.TermsPath
	envs[core.DefaultLang] = testEnvs.Core.DefaultLang
	envs[core.TranslateInBackend] = testEnvs.Core.TranslateInBackend
	envs[core.ExcelMaxRows] = testEnvs.Core.ExcelMaxRows

	envs[base.PasswordSalt] = testEnvs.Base.PasswordSalt
	envs[base.JWTSecretKey] = testEnvs.Base.JWTSecretKey
	envs[base.JWTExpiration] = testEnvs.Base.JWTExpiration

	engine.Envs = envs

	return &engine
}
