package corerr

// List of all errors for record in the app
const (
	InternalServerError                 = "internal server error"
	RouteNotFound                       = "route not found"
	Unauthorized                        = "unauthorized"
	PleaseReportErrorToProgrammer       = "please report error to the programmer"
	RecordVVNotFoundInV                 = "record %v:%v not found in %v"
	RecordNotFound                      = "record not found"
	VisNotValid                         = "%v is not valid"
	BindFailed                          = "bind failed"
	ValidationFailed                    = "validation failed"
	VisRequired                         = "%v is required"
	MinimumAcceptedCharacterForVisV     = "minimum accepted character for %v is %v"
	AcceptedValueForVareV               = "accepted value for %v are: [%v]"
	DuplicateHappened                   = "duplicate happened"
	SomeVRelatedToThisVSoItIsNotDeleted = "some %v related to this %v so it is not deleted"
	SomeVRelatedToThisVSoItIsNotCreated = "some %v related to this %v so it is not created"
	ErrorBecauseOfForeignKey            = "error because of foreign key"
	VWithValueVAlreadyExist             = "%v with value %v already exists"
	VisAlreadyExist                     = "%v is already exist"
	ErrorInBindingV                     = "error in binding %v"
	InvalidID                           = "invalid id"
	InvalidVForV                        = "invalid %v for %v"
	YouDontHavePermission               = "you don't have permission"
	TokenIsRequired                     = "token is required"
	TokenIsNotValid                     = "token is not valid"
)
