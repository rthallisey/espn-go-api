package analytics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func MapStringFloatToJson(location string, data map[string]float64) {
	s, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	defaultDir := "results"

	if _, err := os.Stat(defaultDir + "/" + filepath.Dir(location)); os.IsNotExist(err) {
		os.MkdirAll(defaultDir+"/"+filepath.Dir(location), os.ModePerm)
	}

	err = ioutil.WriteFile(defaultDir+"/"+location, s, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func MapStringIntToJson(location string, data map[string]int) {
	s, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	defaultDir := "results"

	if _, err := os.Stat(defaultDir + "/" + filepath.Dir(location)); os.IsNotExist(err) {
		os.MkdirAll(defaultDir+"/"+filepath.Dir(location), os.ModePerm)
	}

	err = ioutil.WriteFile(defaultDir+"/"+location, s, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
