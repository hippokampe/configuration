package configuration

import (
	"github.com/hippokampe/configuration/v2/browser"
	"github.com/hippokampe/configuration/v2/credentials"
)

type InternalSettings struct {
	browserSelected *browser.Browser
	port            string
	credentialsPath string
	generalPath     string
	logged          bool
	owner           owner
	cred            *credentials.Credentials
	filename        string
}

type owner struct {
	Username string
	Home     string
}
