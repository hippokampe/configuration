package credentials

func (c *Credentials) SetLogged() bool {
	if c.credentialsFile == "" {
		c.isLogged = false
		return c.isLogged
	}

	c.isLogged = true
	return c.isLogged
}

func (c Credentials) IsLogged() bool {
	return c.isLogged
}

func (c *Credentials) Logout() (bool, error) {
	if !c.isLogged {
		return false, nil
	}

	if err := c.Remove(); err != nil {
		return false, err
	}

	c.ID = ""
	c.Password = ""
	c.Username = ""
	c.Email = ""
	c.credentialsFile = ""
	c.isLogged = false

	credentials = nil

	return true, nil
}
