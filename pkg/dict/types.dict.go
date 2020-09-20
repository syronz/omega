package dict

// Lang is used for type of event
type Lang string

// R is used for parameters which we want to translated
type R string

// Lang enums
const (
	En Lang = "en"
	Ku Lang = "ku"
	Ar Lang = "ar"
)

// Langs represents all accepted languages
var Langs = []Lang{
	En,
	Ku,
	Ar,
}

// Term is list of languages
type Term struct {
	En string `toml:"en"`
	Ku string `toml:"ku"`
	Ar string `toml:"ar"`
}

// thisTerms used for holding language identifier as a string and Term Struct as value
var thisTerms map[string]Term
var translateInBackend bool
