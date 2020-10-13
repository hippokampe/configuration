package general

import (
	"github.com/hippokampe/configuration/browser"
)

type Configuration struct {
	filename               string
	CredentialsFilename    string
	CustomSettingsFilename string
	Browsers               []*browser.Browser `json:"browsers,omitempty"`
}
