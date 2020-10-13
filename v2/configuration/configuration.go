package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/hippokampe/configuration/v2/browser"
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
	internal.filename = filename

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &internal); err != nil {
		return err
	}

	return nil
}

func (internal *InternalSettings) Save() error {
	if internal.filename == "" {
		return errors.New("you must set the credential file")
	}

	internal.port = internal.GetPort()

	tmpInternal := struct {
		BrowserSelected *browser.Browser `json:"browser_selected,omitempty"`
		Logged          bool             `json:"logged,omitempty"`
		Port            string           `json:"port,omitempty"`
		Owner           owner            `json:"owner,omitempty"`
		CredentialsPath string           `json:"credentials_path,omitempty"`
		GeneralPath     string           `json:"general_path,omitempty"`
	}{
		BrowserSelected: internal.browserSelected,
		Logged:          internal.logged,
		Port:            internal.port,
		Owner:           internal.owner,
		CredentialsPath: internal.credentialsPath,
		GeneralPath:     internal.generalPath,
	}

	file, err := json.MarshalIndent(tmpInternal, "", " ")
	if err != nil {
		return err
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
