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
		chromedp.Navigate("https://www.yolo247.com/"),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Evaluate JavaScript to find and click the button to login
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`document.querySelector('div.sc-hKMtZM.eTgRHo div.sc-gsnTZi.jxpWyQ button').click();`, nil),
		chromedp.WaitVisible(`input[name="username"]`, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println("Error clicking button and waiting for username input:", err)
		return
	}

	// Fill in the username and password fields
	err = chromedp.Run(ctx,
		chromedp.SendKeys(`#username`, "kiranreddy20"), // Use the ID selector for the username field
		chromedp.SendKeys(`input[name="password"]`, "Solutions@123"),
		chromedp.SendKeys(`input[name="password"]`, "\n"), // \n represents "Enter" key
		// chromedp.WaitVisible(`#contentArea`, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println("Error filling in credentials and waiting for content:", err)
		return
	}
	time.Sleep(4 * time.Second)
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://www.yolo247.com/online-casino-games/roulette"), // Replace with the URL you want to navigate to
	)
	if err != nil {
		fmt.Println("Error navigating to another page:", err)
		return
	}
	cssSelector := `#root > div.sc-cFcpLw.deWdXm > div:nth-child(2) > div > div > div > div.sc-eUuhEu.cZwlDa > div > div > div.sc-rJRYz.WqVYM > div > div > div > div.sc-kACOFk.sc-fiMvTd.cARyEF.fytwlt > div:nth-child(2) > div > span > div > div > div > button`

	// Execute JavaScript code to click the element directly using its CSS selector.
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
	// Sleep for 3 seconds
	// elementXPath := `//*[@id="root"]/div[1]/div[2]/div/div[9]/div[2]/div[2]/a[2]/div/div/div/button`
	// err = chromedp.Run(ctx,
	// 	chromedp.ScrollIntoView(elementXPath),
	// )
	// if err != nil {
	// 	fmt.Println("Error scrolling to the element:", err)
	// 	return
	// }

	// // Click the element using JavaScript
	// jsCode := fmt.Sprintf(`document.evaluate('%s', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue.click();`, elementXPath)
	// var jsResult interface{}
	// err = chromedp.Run(ctx,
	// 	chromedp.Evaluate(jsCode, &jsResult),
	// )
	// if err != nil {
	// 	fmt.Println("Error executing JavaScript to click the element:", err)
	// 	return
	// }
	// // err = chromedp.Run(ctx,
	// // 	chromedp.WaitVisible("#contentArea", chromedp.ByQuery),
	// // )
	// // if err != nil {
	// // 	fmt.Println("Error waiting for content area:", err)
	// // 	return
	// // }
	// time.Sleep(20 * time.Second)
	// fmt.Println("fghjk")
	// // Define the target data-role attribute value you want to click
	// Code := `//*[@id="48z5pjps3ntvqc1b"]/article/div[1]/div[2].click();`

	// // Execute the JavaScript code to click the element.
	// if err := chromedp.Run(ctx,
	// 	chromedp.Evaluate(Code, nil),
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// // Wait for a few seconds to observe the result.
	// time.Sleep(5 * time.Second)

	// fmt.Println("Clicked on 'Auto-Roulette' element.")

	// fmt.Println("youadsbnam")
}
