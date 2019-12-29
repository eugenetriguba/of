package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	MailDropEmail string `json:"mailDropEmail"`
	GmailEmail string `json:"gmailEmail"`
	GmailPassword string `json:"gmailPassword"`
}

func (config *Configuration) Parse() error {
	file, err := os.Open("configuration/conf.json")
	if err != nil {
		return err
	}
	defer closeFile(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return err
}

func (config *Configuration) Save() error {
	file, err := os.Create("configuration/conf.json")
	if err != nil {
		return err
	}
	defer closeFile(file)

	data, err := json.MarshalIndent(*config, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("configuration/conf.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func closeFile(file *os.File) {
	err := file.Close()

	if err != nil {
		fmt.Println("Error occurred while trying to close the file: ", err)
		os.Exit(1)
	}
}