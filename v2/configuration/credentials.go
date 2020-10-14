package configuration

import (
	"errors"
	"github.com/hippokampe/configuration/v2/credentials"
)

func (internal *InternalSettings) BindCredentials(cred *credentials.Credentials) error {
	if cred == nil {
		return errors.New("credentials cannot be nil")
	}

	internal.cred = cred
	internal.credentialsPath = cred.GetFilename()
	return nil
}

func (internal *InternalSettings) GetCredentials() (*credentials.Credentials, error) {
	if internal.cred == nil {
		return nil, errors.New("credentials must be bind first")
	}

	return internal.cred, nil
}

func (internal *InternalSettings) IsLogged() (bool, error) {
	var err error

	if internal.cred == nil {
		return false, errors.New("credentials must be bind first")
	}

	return internal.logged, nil
}

func (internal *InternalSettings) SetLogged() (bool, error) {
	var err error

	if internal.cred == nil {
		return false, errors.New("credentials must be bind first")
	}

	internal.logged, err = internal.cred.SetLogged()
	if err != nil {
		return false, err
	}

	if err = internal.cred.Save(); err != nil {
		return false, err
	}

	if err := internal.Save(); err != nil {
		return false, err
	}

	return internal.logged, nil
}

func (internal *InternalSettings) Logout() (bool, error) {
	if internal.cred == nil {
		return false, errors.New("credentials must be bind first")
	}

	internal.logged = false
	if status, _ := internal.cred.Logout(); !status {
		return false, nil
	}

	if err := internal.Save(); err != nil {
		return false, err
	}

	return true, nil
}
