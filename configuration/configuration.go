// Package configuration handles modifying, saving, and
// outputting the configuration file.
package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/eugenetriguba/of/fs"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// Configuration represents the fields in the
// configuration file.
type Configuration struct {
	MailDropEmail string `json:"mailDropEmail"`
	GmailUsername string `json:"gmailUsername"`
	GmailPassword string `json:"gmailPassword"`
}

// Init initializes the configuration file by
// creating the configuration folder at ~/.of/ and
// configuration file at ~/.of/config.json. It only
// creates these if they do not exist.
//
// It is intended to be run when the application
// starts up to make sure that the config directory
// and file is created.
func (config *Configuration) Init() error {
	err := config.createFolder()

	configFilePath, err := config.GetConfigFilePath()
	if err != nil {
		return errors.Wrap(err, "Retrieving the configuration file path failed")
	}

	created, err := config.Create(configFilePath)
	if err != nil || !created {
		return errors.Wrap(err, "Creating the configuration file failed")
	}

	return nil
}

// GetConfigDirPath retrieves the absolute path to
// the configuration directory in an OS independent way.
//
// The configuration folder is at `~/.of`.
// If an error occurs, an empty string is returned as the path.
func (config *Configuration) GetConfigDirPath() (string, error) {
	dirPath, err := homedir.Expand("~/.of/")
	if err != nil {
		return "", errors.Wrap(err, "Expanding the configuration directory path failed")
	}

	return dirPath, nil
}

// GetConfigFilePath retrieves the absolute path to the
// configuration file in an OS independent way.
//
// The configuration file is at `~/.of/config.json`.
// If an error occurs, an empty string is returned as the path.
func (config *Configuration) GetConfigFilePath() (string, error) {
	filePath, err := homedir.Expand("~/.of/config.json")
	if err != nil {
		return "", errors.Wrap(err, "Expanding the configuration file path failed")
	}

	return filePath, nil
}

// Parse parses the configuration file from `~/.of/config.json`
// and reads the fields into this Configuration.
func (config *Configuration) Parse() error {
	filePath, err := config.GetConfigFilePath()
	if err != nil {
		return errors.Wrap(err, "Retreiving the configuration file path failed")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "Opening configuration file failed")
	}
	defer fs.CloseFile(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return errors.Wrap(err, "Decoding the configuration file failed")
	}

	return nil
}

// Save saves the current state of this Configuration to
// `~/.of/config.json`.
func (config *Configuration) Save() error {
	filePath, err := config.GetConfigFilePath()
	if err != nil {
		return errors.Wrap(err, "Retreiving the configuration file path failed")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, "Creating the configuration file failed")
	}
	defer fs.CloseFile(file)

	data, err := json.MarshalIndent(*config, "", "    ")
	if err != nil {
		return errors.Wrap(err, "Marshaling the configuration file failed")
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return errors.Wrap(err, "Writing out the configuration file failed")
	}

	return nil
}

// Output outputs the configuration file to stdout.
func (config *Configuration) Output() error {
	filePath, err := config.GetConfigFilePath()
	if err != nil {
		return errors.Wrap(err, "Retreiving the configuration file path failed")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrap(err, "Opening configuration file failed")
	}
	defer fs.CloseFile(file)

	fileContents, err := ioutil.ReadAll(file)
	fmt.Println(string(fileContents))
	return nil
}

// Create is what creates the configuration file.
func (config *Configuration) Create(filePath string) (bool, error) {
	fileExists, err := fs.FileExists(filePath)
	if err != nil {
		return false, errors.Wrap(err, "error while checking if the given file path exists")
	}

	if !fileExists {
		file, err := os.Create(filePath)
		if err != nil {
			return false, errors.Wrap(err, "Creating the configuration file failed")
		}
		defer fs.CloseFile(file)

		data, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return false, errors.Wrap(err, "Marshaling the configuration file failed")
		}

		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			return false, errors.Wrap(err, "Writing out the configuration file failed")
		}

		return true, nil
	}

	return false, nil
}

// createFolder checks whether the configuration folder exists and
// if not, creates it.
func (config *Configuration) createFolder() error {
	configDirPath, err := config.GetConfigDirPath()
	if err != nil {
		return errors.Wrap(err, "Retrieving the configuration folder failed")
	}

	dirExists, err := fs.DirExists(configDirPath)
	if !dirExists {
		err := os.Mkdir(configDirPath, 0751)
		if err != nil {
			return errors.Wrap(err, "Creating the configuration directory failed")
		}
	}

	return nil
}
