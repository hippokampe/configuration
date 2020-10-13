package configuration

import (
	"github.com/hippokampe/configuration/v2/browser"
	"github.com/hippokampe/configuration/v2/credentials"
)

type InternalSettings struct {
	BrowserSelected *browser.Browser `json:"browser_selected,omitempty"`
	Port            string           `json:"port"`
	CredentialsPath string           `json:"credentials_path"`
	GeneralPath     string           `json:"general_path"`
	Logged          bool             `json:"logged"`
	Owner 			struct {
		Username string `json:"username"`
		Home string `json:"home_dir"`
	}
	cred            *credentials.Credentials
	filename        string
}
