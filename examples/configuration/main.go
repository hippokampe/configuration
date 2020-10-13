package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hippokampe/configuration/browser"
	"github.com/hippokampe/configuration/configuration"
	"github.com/hippokampe/configuration/credentials"
)

func main() {
	bw := &browser.Browser{Name: "firefox", Path: "/etc", Version: "1.2.3"}
	browser.AddBrowser(bw)

	cred := credentials.New()
	if err := cred.ReadFromFile(os.Getenv("HIPPOKAMPE_CREDENTIALS")); err != nil {
		log.Println(err)
	}

	cred.ID = "1532"
	cred.Password = "password"
	cred.Username = "David Orozco"
	cred.Email = "1532@holbertonschool.com"

	config := configuration.New()
	if err := config.ReadFromFile(os.Getenv("HIPPOKAMPE_CONFIG")); err != nil {
		log.Println(err)
	}

	if err := cred.Save(); err != nil {
		log.Println(err)
	}

	fmt.Println(config.BindCredentials(cred))
	fmt.Println(config.SetLogged())
	fmt.Println(config.SetPort("5600"))
	fmt.Println(config.SetBrowser("firefox"))

	if err := config.Save(); err != nil {
		log.Fatal(err)
	}
}
