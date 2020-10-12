package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var internalConfig *InternalSettings

func New() *InternalSettings {
	if internalConfig != nil {
		return internalConfig
	}

	internalConfig = new(InternalSettings)
	return internalConfig
}

func (internal *InternalSettings) ReadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &internal); err != nil {
		return err
	}

	internal.filename = filename
	return nil
}

func (internal *InternalSettings) Save() error {
	if internal.filename == "" {
		return errors.New("you must set the credential file")
	}

	file, err := json.MarshalIndent(internal, "", " ")
	if err != nil {
		return err
	}

	if internal.cred != nil {
		if err := internal.cred.Save(); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(internal.filename, file, 0644)
}

func (internal *InternalSettings) Remove() error {
	if internal.cred != nil {
		if err := internal.cred.Remove(); err != nil {
			return err
		}
	}

	return os.Remove(internal.filename)
}

func Get() InternalSettings {
	if internalConfig == nil {
		return InternalSettings{}
	}
	return *internalConfig
}
