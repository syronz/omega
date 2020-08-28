package base

import "omega/internal/types"

const (
	PasswordSalt  types.Envkey = "PASSWORD_SALT"
	JWTSecretKey  types.Envkey = "JWT_SECRET_KEY"
	JWTExpiration types.Envkey = "JWT_EXPIRATION"
	RecordRead    types.Envkey = "RECORD_READ"
	RecordWrite   types.Envkey = "RECORD_WRITE"
	AdminUsername types.Envkey = "ADMIN_USERNAME"
	AdminPassword types.Envkey = "ADMIN_PASSWORD"
)
