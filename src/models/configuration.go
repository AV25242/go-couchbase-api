package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	Bucket     string `json:bucket`
	Connection string `json:connection`
	Port       int    `json:port`
	Username   string `json:username`
	Password   string `json:password`
}

func GetConfiguration() Configuration {
	config := Configuration{}
	content, err := ioutil.ReadFile("config/connection.json")
	if err != nil {
		fmt.Println("Error" + err.Error())
	}
	if content != nil {
		err2 := json.Unmarshal(content, &config)
		if err2 != nil {
			fmt.Println("Marshalling Error" + err2.Error())
		}
	}

	return config

}
