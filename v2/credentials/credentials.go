package credentials

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
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

func (c *Credentials) SetFilename(filename string) {
	c.credentialsFile = filename
}

func (c *Credentials) GetFilename() string {
	return c.credentialsFile
}

func (c *Credentials) ReadFromFile() error {
	file, err := ioutil.ReadFile(c.credentialsFile)
	if err != nil {
		return err
	}

	tmpCredentials := struct {
		Id       string `json:"id"`
		Password string `json:"password"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}{}

	if err := json.Unmarshal(file, &tmpCredentials); err != nil {
		return err
	}

	c.id = tmpCredentials.Id
	c.password = tmpCredentials.Password
	c.username = tmpCredentials.Username
	c.email = tmpCredentials.Email

	return nil
}

func (c *Credentials) Save() error {
	if c.credentialsFile == "" {
		return errors.New("you must set the credential file")
	}

	if credentials == nil {
		return errors.New("credentials cannot be blank")
	}

	tmpCredentials := struct {
		Id       string `json:"id"`
		Password string `json:"password"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}{
		c.id,
		c.password,
		c.username,
		c.email,
	}

	file, err := json.MarshalIndent(tmpCredentials, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.credentialsFile, file, 0644)
}

func (c *Credentials) Remove() error {
	if credentials == nil {
		return errors.New("credentials not set")
	}

	return os.Remove(c.credentialsFile)
}

func Get() Credentials {
	if credentials == nil {
		return Credentials{}
	}
	return *credentials
}

func (c *Credentials) GetValue(key string) (string, error) {
	key = strings.ToLower(key)
	switch key {
	case "id":
		return c.id, nil
	case "username":
		return c.username, nil
	case "email":
		return c.email, nil
	case "password":
		return c.password, nil
	default:
		return "", errors.New("key does not exists")
	}
}

func (c *Credentials) SetValue(key, value string) error {
	key = strings.ToLower(key)
	switch key {
	case "id":
		c.id = value
	case "username":
		c.username = value
	case "email":
		c.email = value
	case "password":
		c.password = value
	default:
		return errors.New("key does not exists")
	}

	return nil
}

func (c *Credentials) UnsetValue(key string) error {
	return c.SetValue(key, "")
}