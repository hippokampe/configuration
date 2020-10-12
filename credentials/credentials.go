package credentials

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var credentials *Credentials

func New() *Credentials {
	if credentials != nil {
		return credentials
	}

	credentials = new(Credentials)
	credentials.isLogged = false

	return credentials
}

func (c *Credentials) ReadFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &c); err != nil {
		return err
	}

	c.CredentialsFile = filename
	return nil
}

func (c *Credentials) Save() error {
	if c.credentialsFile == "" {
		return errors.New("you must set the credential file")
	}

	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.credentialsFile, file, 0644)
}

func (c *Credentials) Remove() error {
	return os.Remove(c.credentialsFile)
}

func Get() Credentials {
	if credentials == nil {
		return Credentials{}
	}
	return *credentials
}
