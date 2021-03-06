package configuration

import (
	"errors"
	"fmt"
	"os/user"
	"strconv"
	"strings"

	"github.com/hippokampe/configuration/v2/browser"
)

func (internal *InternalSettings) SetUsername(username string, allowRoot bool) (string, error) {
	owner, err := user.Lookup(username)
	if err != nil {
		return "", err
	}

	if !allowRoot && owner.Username == "root" {
		return "", errors.New("user cannot be root")
	}

	internal.owner.Username = owner.Username
	internal.owner.Home = owner.HomeDir
	return internal.owner.Home, nil
}

func (internal *InternalSettings) GetBrowser(browserName string) (*browser.Browser, error) {
	browserName = strings.ToLower(browserName)
	bw, err := browser.GetBrowser(browserName)
	if err != nil {
		return nil, err
	}

	return bw, nil
}

func (internal *InternalSettings) SetBrowser(browserName string) error {
	bw, err := internal.GetBrowser(browserName)
	if err != nil {
		internal.browserSelected = nil
		return err
	}

	internal.browserSelected = bw
	return nil
}

func (internal InternalSettings) GetPathBrowser() (string, error) {
	if internal.browserSelected == nil {
		return "", errors.New("browser is not set")
	}

	return internal.browserSelected.Path, nil
}

func (internal *InternalSettings) SetPort(port string) error {
	if _, err := strconv.Atoi(port); err != nil {
		return err
	}

	internal.port = fmt.Sprintf(":%s", port)
	return nil
}

func (internal *InternalSettings) GetPort() string {
	if internal.port == "" {
		return ":5678"
	}

	return internal.port
}

func (internal *InternalSettings) GetOwnerUsername() string {
	return internal.owner.Username
}

func (internal InternalSettings) Filename() string {
	return internal.filename
}