package main

import (
	"fmt"
	"holberton/configuration"
	"log"
)

func main() {
	config := configuration.New("/etc/hbtn/general.json")
	if err := config.ReadGeneralConfig(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.SetUsername("root", true))
	fmt.Println(config.SetUsername("davixcky", false))
	fmt.Println(config.SetPort("123"))
	if err := config.WriteConfig(); err != nil {
		log.Fatal(err)
	}
}
