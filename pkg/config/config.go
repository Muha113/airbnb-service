package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func NewConfig(cfg string, obj interface{}) error {
	file, err := os.Open(cfg)
	if err != nil {
		return err
	}

	buff, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buff, obj)
	if err != nil {
		return err
	}

	fmt.Println("Ну кстати нихуева выглядит")

	return nil
}
