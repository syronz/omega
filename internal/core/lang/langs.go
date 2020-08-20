package lang

// Language is used for type of event
type Language string

// Language enums
const (
	En Language = "en"
	Ku Language = "ku"
	Ar Language = "ar"
	Fa Language = "fa"
)

// Languages represents all accepted languages
var Languages = []string{
	string(En),
	string(Ku),
	string(Ar),
	string(Fa),
}
