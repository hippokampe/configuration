package configuration

import (
	"errors"
	"fmt"
	"os/user"
	"strconv"
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
