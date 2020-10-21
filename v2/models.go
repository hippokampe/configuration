package general

import (
	"github.com/hippokampe/configuration/v2/browser"
)

type Configuration struct {
	filename               string
	CredentialsFilename    string             `json:"credentials_filename,omitempty"`
	CustomSettingsFilename string             `json:"custom_settings_filename,omitempty"`
	Browsers               []*browser.Browser `json:"browsers,omitempty"`
}
