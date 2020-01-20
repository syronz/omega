package core

import (
	"encoding/json"
	"io/ioutil"
	"omega/config"
	"omega/utils/glog"
	"os"
	"path/filepath"
	"runtime"
)

func getRegularEnvs() config.Environment {

	_, dir, _, _ := runtime.Caller(0)
	dir = filepath.Dir(dir)

	jsonFile, err := os.Open(dir + "/regularenvs.json")
	glog.CheckError(err, "can't open the config file")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var envs config.Environment
	err = json.Unmarshal(byteValue, &envs)
	glog.CheckError(err, "error in unmarshal JSON")

	return envs
}
