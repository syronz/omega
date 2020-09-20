package dict

import (
	"encoding/json"
	"omega/pkg/glog"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

// Init terms and put in the main map
func Init(termsPath string, backendTranslation bool) {
	thisTerms = make(map[string]Term)
	translateInBackend = backendTranslation

	file, err := os.Open(termsPath)
	if err != nil {
		glog.Fatal("can't open terms file: ", err, termsPath)
	}
	defer file.Close()

	fileType := strings.ToUpper(filepath.Ext(termsPath))

	switch fileType {

	case ".JSON":

		decoder := json.NewDecoder(file)
		var lines map[string]interface{}
		err = decoder.Decode(&lines)
		if err != nil {
			glog.Fatal("can't decode terms to JSON: ", err)
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

	case ".TOML":

		if _, err := toml.DecodeFile(termsPath, &thisTerms); err != nil {
			glog.Fatal("failed in decoding the toml file for terms", err)
		}

	default:
		glog.Fatal("JSON or TOML accepted for terms type")
	}

}
