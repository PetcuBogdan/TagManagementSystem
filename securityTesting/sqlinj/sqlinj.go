package main

import (
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	
	fuzzingStrings := [6]string{`" or ""="`, `auto;DROP TABLE Suppliers`,`auto' UNION SELECT 1;--`,`auto', id=24`,`' ,rez=(select group_concat(name || "," || tag || ":") from documenttag_document;)`,` 1' UNION SELECT NULL-- -]`}
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("http://localhost:3000/account/login")
	assertErrorToNilf("could not goto: %w", err)

	time.Sleep(1 * time.Second)
	page.Fill(`input[name="email"]`, "petcu_b@yahoo.ro")

	time.Sleep(1 * time.Second)
	page.Fill(`input[name="password"]`, "Test123@")
	time.Sleep(1 * time.Second)
	page.Click(`button[class="formkit-input"]`)
	time.Sleep(1 * time.Second)

	_, err = page.Goto("http://localhost:3000/user/search")
	assertErrorToNilf("could not goto: %w", err)

	time.Sleep(1 * time.Second)
	
	i := 0;
	
	for i < 6 {
		time.Sleep(1 * time.Second)
		page.Fill(`input[id="search"]`, fuzzingStrings[i])
		i++;
	}

	time.Sleep(60 * time.Second)
	context.Close()
	browser.Close()
}

