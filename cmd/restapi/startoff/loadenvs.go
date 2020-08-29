package startoff

import (
	"omega/domain/base"
	"omega/internal/core"
	"omega/internal/types"
	"os"
)

func LoadEnvs() *core.Engine {
	var engine core.Engine
	var envs types.Envs

	envs = make(types.Envs, 29)

	envs[core.Port] = os.Getenv("OMEGA_CORE_PORT")
	envs[core.Addr] = os.Getenv("OMEGA_CORE_ADDR")
	envs[core.DatabaseDataDSN] = os.Getenv("OMEGA_CORE_DATABASE_DATA_DSN")
	envs[core.DatabaseDataType] = os.Getenv("OMEGA_CORE_DATABASE_DATA_TYPE")
	envs[core.DatabaseDataLog] = os.Getenv("OMEGA_CORE_DATABASE_DATA_LOG")
	envs[core.DatabaseActivityDSN] = os.Getenv("OMEGA_CORE_DATABASE_ACTIVITY_DSN")
	envs[core.DatabaseActivityType] = os.Getenv("OMEGA_CORE_DATABASE_ACTIVITY_TYPE")
	envs[core.DatabaseActivityLog] = os.Getenv("OMEGA_CORE_DATABASE_ACTIVITY_LOG")
	envs[core.AutoMigrate] = os.Getenv("OMEGA_CORE_AUTO_MIGRATE")
	envs[core.ServerLogFormat] = os.Getenv("OMEGA_CORE_SERVER_LOG_FORMAT")
	envs[core.ServerLogOutput] = os.Getenv("OMEGA_CORE_SERVER_LOG_OUTPUT")
	envs[core.ServerLogLevel] = os.Getenv("OMEGA_CORE_SERVER_LOG_LEVEL")
	envs[core.ServerLogJSONIndent] = os.Getenv("OMEGA_CORE_SERVER_LOG_JSON_INDENT")
	envs[core.APILogFormat] = os.Getenv("OMEGA_CORE_API_LOG_FORMAT")
	envs[core.APILogOutput] = os.Getenv("OMEGA_CORE_API_LOG_OUTPUT")
	envs[core.APILogLevel] = os.Getenv("OMEGA_CORE_API_LOG_LEVEL")
	envs[core.APILogJSONIndent] = os.Getenv("OMEGA_CORE_API_LOG_JSON_INDENT")
	envs[core.TermsPath] = os.Getenv("OMEGA_CORE_TERMS_PATH")
	envs[core.DefaultLang] = os.Getenv("OMEGA_CORE_DEFAULT_LANGUAGE")
	envs[core.TranslateInBackend] = os.Getenv("OMEGA_CORE_TRANSLATE_IN_BACKEND")
	envs[core.ExcelMaxRows] = os.Getenv("OMEGA_CORE_EXCEL_MAX_ROWS")
	envs[core.ErrPanel] = os.Getenv("OMEGA_CORE_ERR_PANEL")
	envs[core.OriginalError] = os.Getenv("OMEGA_CORE_ORIGINAL_ERROR")
	envs[core.GindMode] = os.Getenv("GIN_MODE")

	envs[base.PasswordSalt] = os.Getenv("OMEGA_BASE_PASSWORD_SALT")
	envs[base.JWTSecretKey] = os.Getenv("OMEGA_BASE_JWT_SECRET_KEY")
	envs[base.JWTExpiration] = os.Getenv("OMEGA_BASE_JWT_EXPIRATION")
	envs[base.RecordRead] = os.Getenv("OMEGA_BASE_RECORD_READ")
	envs[base.RecordWrite] = os.Getenv("OMEGA_BASE_RECORD_WRITE")
	envs[base.AdminUsername] = os.Getenv("OMEGA_BASE_ADMIN_USERNAME")
	envs[base.AdminPassword] = os.Getenv("OMEGA_BASE_ADMIN_PASSWORD")

	engine.Envs = envs

	return &engine
}
