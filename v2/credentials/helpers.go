package credentials

import "errors"

func (c *Credentials) SetLogged() (bool, error) {
	if credentials == nil {
		return false, errors.New("credentials not set")
	}

	if c.credentialsFile == "" {
		c.isLogged = false
		return c.isLogged, errors.New("credentials file not set")
	}

	c.isLogged = true
	return c.isLogged, nil
}

func (c Credentials) IsLogged() (bool, error) {
	if credentials == nil {
		return false, errors.New("credentials not set")
	}

	return c.isLogged, nil
}

func (c *Credentials) Logout() (bool, error) {
	if credentials == nil {
		return false, errors.New("credentials not set")
	}

	if !c.isLogged {
		return false, nil
	}

	if err := c.Remove(); err != nil {
		return false, err
	}

	c.id = ""
	c.password = ""
	c.username = ""
	c.email = ""
	c.credentialsFile = ""
	c.isLogged = false

	credentials = nil

	return true, nil
}
