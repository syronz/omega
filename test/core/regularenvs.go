package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"omega/config"
	"os"
	"path/filepath"
	"runtime"
)

func getRegularEnvs() config.Environment {

	_, dir, _, _ := runtime.Caller(0)
	dir = filepath.Dir(dir)

	jsonFile, err := os.Open(dir + "/regularenvs.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var envs config.Environment
	json.Unmarshal(byteValue, &envs)
	return envs

}
