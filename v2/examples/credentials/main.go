package main

import (
	"fmt"
	"os"

	"github.com/hippokampe/configuration/v2/credentials"
)

func main() {
	cred := credentials.New()

	//	fmt.Println(cred.ReadFromFile("/home/davixcky/.config/hippokampe/credentials.json"))

	fmt.Println(cred.ReadFromFile(os.Getenv("HIPPOKAMPE_CREDENTIALS")))

	cred.id = "1532"
	cred.password = "password"
	cred.username = "David Orozco"
	cred.email = "1532@holbertonschool.com"

	fmt.Println(cred)
	fmt.Println(cred.Save())
	fmt.Println(cred.IsLogged())
	fmt.Println(cred.SetLogged())
	fmt.Println(cred.IsLogged())
	fmt.Println(cred.Logout())
	fmt.Println(cred.IsLogged())
	fmt.Println(credentials.Get())
}
