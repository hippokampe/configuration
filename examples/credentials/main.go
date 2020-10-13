package main

import (
	"fmt"
	"github.com/hippokampe/configuration/credentials"
	"os"
)

func main() {
	cred := credentials.New()

	//	fmt.Println(cred.ReadFromFile("/home/davixcky/.config/hippokampe/credentials.json"))

	fmt.Println(cred.ReadFromFile(os.Getenv("HIPPOKAMPE_CREDENTIALS")))

	cred.ID = "1532"
	cred.Password = "password"
	cred.Username = "David Orozco"
	cred.Email = "1532@holbertonschool.com"

	fmt.Println(cred)
	fmt.Println(cred.Save())
	fmt.Println(cred.IsLogged())
	fmt.Println(cred.SetLogged())
	fmt.Println(cred.IsLogged())
	fmt.Println(cred.Logout())
	fmt.Println(cred.IsLogged())
	fmt.Println(credentials.Get())
}
