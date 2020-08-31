package corerr

type ErrTerm string

// List of all errors for record in the app
const (
	Record__NotFoundIn_                   = "record %v:%v not found in %v"
	RecordNotFound                        = "record not found"
	InternalServerError                   = "internal server error"
	Internal_Server_Error_Happened___     = "internal server error happened, please aware administration and gave him error code"
	Bind_failed                           = "bind failed"
	V_is_not_valid                        = "%v is not valid"
	Validation_failed                     = "validation failed"
	Validation_failed_for_V               = "validation failed for %v"
	Validation_failed_for_V_V             = "validation failed for %v %v"
	Minimum_accepted_character_for_V_is_V = "minimum accepted character for %v is %v"
	V_is_required                         = "%v is required"
	Accepted_value_for_V_are_V            = "accepted value for %v are: [%v]"
	V_not_exist                           = "%v not exist"
	Duplication_happened                  = "duplication happened"
	This_V_already_exist                  = "this %v already exist"
)
