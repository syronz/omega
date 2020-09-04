package basmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/dict"
)

const (
	// UserTable is used inside the repo layer
	UserTable = "bas_users"
	UserPart  = "user"
	UsersPart = "users"
)

// User model
type User struct {
	ID        types.RowID `gorm:"not null;unique" json:"id"`
	RoleID    types.RowID `gorm:"index:role_id_idx" json:"role_id"`
	Username  string      `gorm:"not null;unique" json:"username,omitempty"`
	Password  string      `gorm:"not null" json:"password,omitempty"`
	Lang      dict.Lang   `gorm:"type:varchar(2);default:'en'" json:"language,omitempty"`
	Email     string      `json:"email,omitempty"`
	Extra     interface{} `sql:"-" json:"user_extra,omitempty"`
	Resources string      `sql:"-" json:"resources,omitempty"`
	Role      string      `sql:"-" json:"role,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p User) Pattern() string {
	return `(
		bas_users.username LIKE '%[1]v%%' OR
		bas_users.id = '%[1]v' OR
		bas_users.email LIKE '%[1]v%%' OR
		bas_roles.name LIKE '%[1]v' 
	)`
}

// Columns return list of total columns according to request, useful for inner joins
func (p User) Columns(variate string, params param.Param) (string, error) {
	full := []string{
		"bas_users.id",
		"bas_users.role_id",
		"bas_users.username",
		"bas_users.language",
		"bas_users.email",
		"bas_roles.name as role",
	}

	return validator.CheckColumns(full, variate, params)
}

// Validate check the type of
func (p *User) Validate(act coract.Action, params param.Param) error {
	//fieldError := corerr.NewSilent("E1052981", params, base.Domain, nil).
	//	FieldError("/users/[:userID]", corerr.Validation_failed_for_V, dict.R("user"))

	//switch act {
	//case coract.Create:
	//	fieldError.SetPath("/users").
	//		SetMsg(corerr.Validation_failed_for_V_V, dict.R("create"), dict.R("user"))

	//	validateUserUsername(fieldError, p.Username)
	//	validateUserPassword(fieldError, p.Password)
	//	validateUserRole(fieldError, p.RoleID)
	//	validateUserLang(fieldError, p.Lang)
	//	validateUserEmail(fieldError, p.Email)

	//case coract.Update:
	//	fieldError.SetPath("/users/:userID").
	//		SetMsg(corerr.Validation_failed_for_V_V, dict.R("update"), dict.R("user"))

	//	validateUserUsername(fieldError, p.Username)

	//	if p.Password != "" {
	//		validateUserPassword(fieldError, p.Password)
	//	}

	//	validateUserRole(fieldError, p.RoleID)
	//	validateUserLang(fieldError, p.Lang)
	//	validateUserEmail(fieldError, p.Email)

	////for default we validate all fields
	//default:
	//	validateUserUsername(fieldError, p.Username)
	//	validateUserPassword(fieldError, p.Password)
	//	validateUserRole(fieldError, p.RoleID)
	//	validateUserLang(fieldError, p.Lang)
	//	validateUserEmail(fieldError, p.Email)
	//}

	//return fieldError.Final()
	return nil
}

// func validateUserPassword(fieldError *corerr.CustomError, password string) {
// 	if len(password) < consts.MinimumPasswordChar {
// 		fieldError.Add("password", corerr.Minimum_accepted_character_for_V_is_V,
// 			dict.R("password"), consts.MinimumPasswordChar)
// 	}
// }

// func validateUserUsername(fieldError *corerr.CustomError, username string) {
// 	if username == "" {
// 		fieldError.Add("username", corerr.V_is_required, dict.R("Username"))
// 	}
// }

// func validateUserRole(fieldError *corerr.CustomError, roleID types.RowID) {
// 	if roleID == 0 {
// 		fieldError.Add("role_id", corerr.V_is_required, dict.R("Role"))
// 	}
// }

// func validateUserLang(fieldError *corerr.CustomError, lang dict.Lang) {
// 	if ok, _ := helper.Includes(dict.Langs, lang); !ok {
// 		var str []string
// 		for _, v := range dict.Langs {
// 			str = append(str, string(v))
// 		}
// 		fieldError.Add("language", corerr.Accepted_value_for_V_are_V, dict.R("Resource"),
// 			strings.Join(str, ", "))
// 	}
// }

// func validateUserEmail(fieldError *corerr.CustomError, email string) {
// 	if email != "" {
// 		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 		if !re.MatchString(email) {
// 			fieldError.Add("email", corerr.V_is_not_valid, dict.R("Email"))
// 		}
// 	}
// }
