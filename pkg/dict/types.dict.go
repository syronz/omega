package dict

// Lang is used for type of event
type Lang string

// Lang enums
const (
	En Lang = "en"
	Ku Lang = "ku"
	Ar Lang = "ar"
	Fa Lang = "fa"
)

// Langs represents all accepted languages
var Langs = []string{
	string(En),
	string(Ku),
	string(Ar),
	string(Fa),
}

// Term is list of languages
type Term struct {
	En string
	Ku string
	Ar string
}

// thisTerms used for holding language identifier as a string and Term Struct as value
var thisTerms map[string]Term
