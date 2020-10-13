package browser

import (
	"errors"
	"strings"
)

func AddBrowser(browser *Browser) {
	internalBrowsers = append(internalBrowsers, browser)
}

func SetBrowsers(browsers []Browser) {
	for _, browser := range browsers {
		internalBrowsers = append(internalBrowsers, &browser)
	}
}

func GetBrowser(browserName string) (*Browser, error) {
	browserName = strings.ToLower(browserName)
	for _, browser := range internalBrowsers {
		if strings.ToLower(browser.Name) == browserName {
			return browser, nil
		}
	}

	return nil, errors.New("browser not found")
}
