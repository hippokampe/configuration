package configuration

import (
	"errors"
	"github.com/hippokampe/configuration/credentials"
	"os"
)

func (internal *InternalSettings) BindCredentials(cred *credentials.Credentials) error {
	if cred == nil {
		return errors.New("credentials cannot be nil")
	}

	internal.cred = cred
	internal.CredentialsPath = os.Getenv("HIPPOKAMPE_CREDENTIALS")
	return nil
}

func (internal *InternalSettings) IsLogged() bool {
	internal.Logged = internal.cred.IsLogged()
	return internal.Logged
}

func (internal *InternalSettings) SetLogged() bool {
	internal.Logged = internal.cred.SetLogged()
	return internal.Logged
}
