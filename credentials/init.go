package credentials

var credentials *Credentials

func New() *Credentials {
	if credentials != nil {
		return credentials
	}

	credentials = new(Credentials)
	return credentials
}

func GetCredentials() *Credentials {
	return credentials
}