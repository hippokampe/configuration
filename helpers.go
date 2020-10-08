package configuration

import (
	"errors"
	"fmt"
	"github.com/hippokampe/configuration/credentials"
	"os/user"
	"strconv"
	"strings"
)

func (c *Configuration) GetPathBrowser(browser string) (string, error) {
	if v, ok := c.internalBrowsers[browser]; ok {
		return v.Path, nil
	}

	return "", errors.New("browser not found")
}

func (c *Configuration) SetPort(port string) error {
	if _, err := strconv.Atoi(port); err != nil {
		return err
	}

	c.Port = fmt.Sprintf(":%s", port)
	return nil
}

func (c *Configuration) GetPort() string {
	if c.Port == "" {
		return ":5678"
	}

	return c.Port
}

func (c *Configuration) SetUsername(username string, allowRoot bool) (string, error) {
	owner, err := user.Lookup(username)
	if err != nil {
		return "", err
	}

	if !allowRoot && owner.Username == "root" {
		return "", errors.New("user cannot be root")
	}

	c.Owner.Username = owner.Username
	c.Owner.Home = owner.HomeDir
	return c.Owner.Home, nil
}

func (c *Configuration) SetBrowser(browserName string) {
	browserName = strings.ToLower(browserName)
	c.BrowserSelected = browserName
}

func (c *Configuration) SetLogged() bool {
	cred := c.InternalStatus.Credentials

	if cred == nil {
		c.InternalStatus.Logged = false
		return c.InternalStatus.Logged
	}

	c.InternalStatus.Logged = len(cred.ID) == 4
	return c.InternalStatus.Logged
}

func (c *Configuration) SetLogout() (bool, error) {
	c.InternalStatus.Logged = false

	if c.InternalStatus.Credentials == nil {
		return c.InternalStatus.Logged, nil
	}

	c.InternalStatus.Credentials.Remove()
	c.InternalStatus.Credentials = &credentials.Credentials{}

	if err := c.WriteConfig(); err != nil {
		return false, err
	}

	return c.InternalStatus.Logged, nil
}

func (c *Configuration) IsLogged() bool {
	if c.InternalStatus.Credentials == nil {
		return false
	}

	return c.InternalStatus.Logged
}