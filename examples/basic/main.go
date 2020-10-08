package main

import (
	"fmt"
	"github.com/hippokampe/configuration"
	"github.com/hippokampe/configuration/credentials"
	"log"
)

func main() {
	config := configuration.New("general.json")
	if err := config.ReadGeneralConfig(); err != nil {
		log.Fatal(err)
	}

	config.InternalStatus.Credentials = &credentials.Credentials{
		ID: "1234",
		Email: "email",
		CredentialsFile: "credentials.json",
	}

	fmt.Println(config.InternalStatus.Credentials.Save())

	fmt.Println(config.SetLogged())
	if err := config.WriteConfig(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.IsLogged())

	if err := config.WriteConfig(); err != nil {
		log.Fatal(err)
	}
}
