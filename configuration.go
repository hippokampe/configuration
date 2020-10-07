package configuration

import (
	"encoding/json"
	"io/ioutil"
)

var (
	configuration *Configuration
)

func New(filename string) *Configuration {
	if configuration != nil {
		return configuration
	}

	configuration = new(Configuration)
	configuration.internalBrowsers = make(map[string]Browser)
	configuration.filename = filename

	return configuration
}

func (c *Configuration) ReadGeneralConfig() error {
	file, err := ioutil.ReadFile(c.filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &c); err != nil {
		return err
	}

	for _, browser := range c.Browsers {
		c.internalBrowsers[browser.Name] = browser
	}

	return nil
}

func (c *Configuration) WriteConfig() error {
	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.filename, file, 0644)
}
