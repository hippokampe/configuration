package configuration

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hippokampe/configuration/v2/browser"
)

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
		internal.BrowserSelected = nil
		return err
	}

	internal.BrowserSelected = bw
	return nil
}

func (internal InternalSettings) GetPathBrowser() (string, error) {
	if internal.BrowserSelected == nil {
		return "", errors.New("browser is not set")
	}

	return internal.BrowserSelected.Path, nil
}

func (internal *InternalSettings) SetPort(port string) error {
	if _, err := strconv.Atoi(port); err != nil {
		return err
	}

	internal.Port = fmt.Sprintf(":%s", port)
	return nil
}

func (internal *InternalSettings) GetPort() string {
	if internal.Port == "" {
		return ":5678"
	}

	return internal.Port
}
