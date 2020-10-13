package configuration

import (
	"errors"
	"os"

	"github.com/hippokampe/configuration/v2/credentials"
)

func (internal *InternalSettings) BindCredentials(cred *credentials.Credentials) error {
	if cred == nil {
		return errors.New("credentials cannot be nil")
	}

	internal.cred = cred
	internal.CredentialsPath = os.Getenv("HIPPOKAMPE_CREDENTIALS")
	return nil
}

func (internal *InternalSettings) GetCredentials() (*credentials.Credentials, error) {
	if internal.cred == nil {
		return nil, errors.New("credentials must be bind first")
	}

	return internal.cred, nil
}

func (internal *InternalSettings) IsLogged() (bool, error) {
	if internal.cred == nil {
		return false, errors.New("credentials must be bind first")
	}

	internal.Logged = internal.cred.IsLogged()
	return internal.Logged, nil
}

func (internal *InternalSettings) SetLogged() (bool, error) {
	if internal.cred == nil {
		return false, errors.New("credentials must be bind first")
	}

	internal.Logged = internal.cred.SetLogged()
	return internal.Logged, nil
}
