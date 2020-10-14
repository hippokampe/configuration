package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hippokampe/configuration/v2/browser"
	"github.com/hippokampe/configuration/v2/configuration"
	"github.com/hippokampe/configuration/v2/credentials"
)

func main() {
	bw := &browser.Browser{Name: "firefox", Path: "/etc", Version: "1.2.3"}
	browser.AddBrowser(bw)

	cred := credentials.New()
	cred.SetFilename(os.Getenv("HIPPOKAMPE_CREDENTIALS"))
	if err := cred.ReadFromFile(); err != nil {
		log.Println(err)
	}

	config := configuration.New()
	config.SetFilename(os.Getenv("HIPPOKAMPE_CONFIG"))
	if err := config.ReadFromFile(); err != nil {
		log.Println(err)
	}

	fmt.Println(config.BindCredentials(cred))
	fmt.Println(config.GetCredentials())

	fmt.Println(config.IsLogged())

	_ = cred.SetValue("id", "1532")
	_ = cred.SetValue("password", "password")
	_ = cred.SetValue("username", "David Orozco")
	_ = cred.SetValue("email", "1532@holbertonschool.com")

	fmt.Println(config.GetCredentials())

	if err := cred.Save(); err != nil {
		log.Println(err)
	}

	fmt.Println(config.SetLogged())
	fmt.Println(config.SetPort("5600"))
	fmt.Println(config.SetBrowser("firefox"))

	println()

	cred, _ = config.GetCredentials()
	fmt.Println(cred)

	fmt.Println(config.Logout())
}
