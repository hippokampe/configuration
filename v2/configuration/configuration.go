package configuration

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (internal *InternalSettings) SetFilename(filename string) {
	internal.filename = filename
}

func (internal *InternalSettings) ReadFromFile() error {
	file, err := ioutil.ReadFile(internal.filename)
	if err != nil {
		return err
	}

	tmpInternal := struct {
		BrowserSelected *browser.Browser `json:"browser_selected"`
		Logged          bool             `json:"logged"`
		Port            string           `json:"port"`
		Owner           owner            `json:"owner"`
		CredentialsPath string           `json:"credentials_path,omitempty"`
		GeneralPath     string           `json:"general_path,omitempty"`
	}{}

	if err := json.Unmarshal(file, &tmpInternal); err != nil {
		return err
	}

	internal.browserSelected = tmpInternal.BrowserSelected
	internal.logged = tmpInternal.Logged
	internal.port = tmpInternal.Port
	internal.owner = tmpInternal.Owner
	internal.credentialsPath = tmpInternal.CredentialsPath
	internal.generalPath = tmpInternal.GeneralPath

	if internal.logged && internal.cred != nil {
		_, _ = internal.cred.SetLogged()
	}

	return nil
}

func (internal *InternalSettings) Save() error {
	if internal.filename == "" {
		return errors.New("you must set the credential file")
	}

	internal.port = internal.GetPort()

	tmpInternal := struct {
		BrowserSelected *browser.Browser `json:"browser_selected"`
		Logged          bool             `json:"logged"`
		Port            string           `json:"port"`
		Owner           owner            `json:"owner"`
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
