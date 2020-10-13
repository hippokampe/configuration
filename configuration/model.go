package configuration

import (
	"github.com/hippokampe/configuration/browser"
	"github.com/hippokampe/configuration/credentials"
)

type InternalSettings struct {
	BrowserSelected *browser.Browser `json:"browser_selected,omitempty"`
	Port            string           `json:"port"`
	CredentialsPath string           `json:"credentials_path"`
	GeneralPath     string           `json:"general_path"`
	Logged          bool             `json:"logged"`
	cred            *credentials.Credentials
	filename        string
}
