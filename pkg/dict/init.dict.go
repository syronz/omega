package dict

import (
	"encoding/json"
	"log"
	"os"
)

// Init terms and put in the main map
func Init(termsPath string, backendTranslation bool) {
	thisTerms = make(map[string]Term)
	translateInBackend = backendTranslation

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
