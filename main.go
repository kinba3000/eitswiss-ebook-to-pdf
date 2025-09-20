package main

import "github.com/kinba3000/eitswiss-ebook-to-pdf/cmd"
import (
    "context"
    "log"

    "github.com/playwright-community/playwright-go"
)
func main() {
	// cmd.Execute()
	pw, _ := playwright.Run()
    browser, _ := pw.Chromium.Launch()
    context_, _ := browser.NewContext()
    page, _ := context_.NewPage()
    page.Goto("https://ebook.eitswiss.ch/login")

    // ... Login durchführen ...

    // Storage-Status in Datei speichern
    context_.StorageState(playwright.BrowserContextStorageStateOptions{
        Path: playwright.String("state.json"),
    })

    log.Println("Login-State gespeichert")
}

