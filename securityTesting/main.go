package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func generateRandomString(length int) string{
	rand.Seed(time.Now().UnixNano())

	// Caracterele posibile pentru string-ul generat
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+"

	// Generăm string-ul aleatoriu
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(randomString)
}

func generateFuzzingStrings(size int) []string {
	fuzzingStrings := make([]string, size)

	for i := 0; i < size; i++ {
		// Generăm un string aleatoriu cu lungimea între 10 și 20
		randomLength := rand.Intn(256) + 10
		fuzzingStrings[i] = generateRandomString(randomLength)
	}

	return fuzzingStrings
}

func main() {
	
	fuzzingStrings := generateFuzzingStrings(500)
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
	
	for i < 500 {
		time.Sleep(1 * time.Second)
		page.Fill(`input[id="search"]`, fuzzingStrings[i])
		i++;
	}

	time.Sleep(60 * time.Second)
	context.Close()
	browser.Close()
}