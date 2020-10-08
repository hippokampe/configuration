package credentials

type Credentials struct {
	ID       string `json:"id,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	CredentialsFile   string                   `json:"credentials_file,omitempty"`
}
