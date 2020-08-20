package initiate

import (
	"encoding/json"
	"log"
	"omega/internal/core"
	"omega/internal/term"
	"os"
)

// LoadTerms terms and put in the main map
func LoadTerms(engine *core.Engine) {
	dict := term.Dict{}
	dict.Terms = make(map[string]term.Term)

	// get current directory
	// _, dir, _, _ := runtime.Caller(0)
	// termPath := filepath.Join(filepath.Dir(dir), "../..", "env", "terms.json")

	file, err := os.Open(engine.Envs[core.TermsPath])
	if err != nil {
		log.Fatal("can't open terms file: ", err, engine.Envs[core.TermsPath])
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var lines map[string]interface{}
	err = decoder.Decode(&lines)
	if err != nil {
		log.Fatal("can't decode terms to JSON: ", err)
	}

	for i, v := range lines {
		lang := v.(map[string]interface{})
		term := term.Term{
			En: lang["en"].(string),
			Ku: lang["ku"].(string),
			Ar: lang["ar"].(string),
		}
		dict.Terms[i] = term
	}

	engine.Dict = dict

}
