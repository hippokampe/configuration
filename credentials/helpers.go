package credentials

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func (c *Credentials) Set(cred Credentials) {
	c.Username = cred.Username
	c.Password = cred.Username
	c.ID = cred.ID
	c.Email = cred.Email
}

func (c *Credentials) Save() error {
	if c.CredentialsFile == "" {
		return errors.New("you must set the credential file")
	}

	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.CredentialsFile, file, 0644)
}
