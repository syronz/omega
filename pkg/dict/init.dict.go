package dict

import (
	"encoding/json"
	"log"
	"os"
)

// Init terms and put in the main map
func Init(termsPath string) {
	// dict := Dict{}
	thisTerms = make(map[string]Term)

	// get current directory
	// _, dir, _, _ := runtime.Caller(0)
	// termPath := filepath.Join(filepath.Dir(dir), "../..", "env", "terms.json")

	file, err := os.Open(termsPath)
	if err != nil {
		log.Fatal("can't open terms file: ", err, termsPath)
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
		term := Term{
			En: lang["en"].(string),
			Ku: lang["ku"].(string),
			Ar: lang["ar"].(string),
		}
		thisTerms[i] = term
	}

}
