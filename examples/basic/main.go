package main

import (
	"fmt"
	"log"

	general "github.com/hippokampe/configuration"
	"github.com/hippokampe/configuration/browser"
)

func main() {
	config := general.New("/etc/hippokampe/general.json")
	if err := config.ReadGeneralConfig(); err != nil {
		log.Fatal(err)
	}

	for _, bw := range config.Browsers {
		browser.AddBrowser(bw)
		fmt.Println(bw)
	}

	println()
	fmt.Println(browser.GetBrowser("firefox"))
	fmt.Println(browser.GetBrowser("chromium"))
	fmt.Println(browser.GetBrowser("chrome"))
}
