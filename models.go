package configuration

type Configuration struct {
	internalBrowsers map[string]Browser
	filename         string
	BrowserSelected  string    `json:"browser_selected"`
	Browsers         []Browser `json:"browsers,omitempty"`
	Port             string    `json:"port,omitempty"`
	Owner            struct {
		Home     string `json:"home,omitempty"`
		Username string `json:"username,omitempty"`
	} `json:"owner,omitempty"`
	InternalStatus struct {
		Logged            bool   `json:"logged,omitempty"`
		ConfigurationFile string `json:"configuration_file,omitempty"`
		CredentialsFile   string `json:"credentials_file,omitempty"`
	} `json:"internal_status,omitempty"`
}

type Browser struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Path    string `json:"path,omitempty"`
}
