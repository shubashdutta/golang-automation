package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Navigate to the login page
	err := chromedp.Run(ctx,
		chromedp.Navigate("url"),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Evaluate JavaScript to find and click the button to login
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`jspath of button you wanted to click ').click();`, nil),
		chromedp.WaitVisible(`input[name="username"]`, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println("Error clicking button and waiting for username input:", err)
		return
	}

	// Fill in the username and password fields
	err = chromedp.Run(ctx,
		chromedp.SendKeys(`#username`, ""), // Use the ID selector for the username field
		chromedp.SendKeys(`input[name="password"]`, ""),
		chromedp.SendKeys(`input[name="password"]`, "\n"), // \n represents "Enter" key
		// chromedp.WaitVisible(`#contentArea`, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println("Error filling in credentials and waiting for content:", err)
		return
	}
	time.Sleep(4 * time.Second)
	err = chromedp.Run(ctx,
		chromedp.Navigate(""), // Replace with the URL you want to navigate to
	)
	if err != nil {
		fmt.Println("Error navigating to another page:", err)
		return
	}
	cssSelector := `xpath`

	
	jsCode := `
		var element = document.querySelector("` + cssSelector + `");
		if (element) {
			element.click();
		}
	`

	fmt.Printf("Clicking the element with CSS selector: %s using JavaScript...\n", cssSelector)
	if err := chromedp.Run(ctx, chromedp.Evaluate(jsCode, nil)); err != nil {
		log.Fatal("Failed to simulate a click using JavaScript:", err)
	}

	fmt.Println("Element clicked successfully.")
	fmt.Println("Element clicked successfully.")
	time.Sleep(10 * time.Minute)
}
