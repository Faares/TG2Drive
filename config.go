package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Telegram Telegram `json:"telegram"`
	Drive    Drive    `json:"drive"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	User     string `json:"user"`
}

type Telegram struct {
	Token      string          `json:"token"`
	Authorized map[string]User `json:"authorized"`
}

type Drive struct {
	Type                        string
	Project_ID                  string
	Private_Key_Id              string
	Private_Key                 string
	Client_Email                string
	Client_Id                   string
	Auth_Uri                    string
	TokenUri                    string
	Auth_Provider_x509_Cert_Url string
	Client_x509_Cert_Url        string
}

func getConfig(file string) Config {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}
