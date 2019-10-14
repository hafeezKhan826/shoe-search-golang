package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var shoeColorsScraped []string
var urls = []string{"https://shopcrocs.in/shop-by-style/clogs.html?size=259", "https://shopcrocs.in/shop-by-style/clogs.html?p=2&size=259"}

func makeAPICallForShoes() {
	for ui := 0; ui < len(urls); ui++ {

		response, err := http.Get(urls[ui])
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		// Create a goquery document from the HTTP response
		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body. ", err)
		}
		/* document.Find("a.product-item-link").Each(func(index int, element *goquery.Selection) {
			// text := element.Text()
			// fmt.Println("Title: ", strings.TrimSpace(text))
		}) */

		document.Find("ul.color-options,li,a").Each(func(index int, element *goquery.Selection) {
			title, exists := element.Attr("title")
			if exists {
				shoeColorsScraped = append(shoeColorsScraped, strings.TrimSpace(title))
			}
		})
	}

}
func searchForMyShoe(color string) bool {
	for i := 0; i < len(shoeColorsScraped); i++ {

		if strings.Contains(strings.TrimSpace(shoeColorsScraped[i]), strings.TrimSpace(color)) {

			fmt.Println("==============>", color, " shoe is available in the store <==============")
			return true
		}
	}
	return false

}

func main() {
	makeAPICallForShoes()

	got := true
	tries := 8
	const shoeColorToSearch = "Navy/Blue"
	for got && tries > 0 {
		got := searchForMyShoe(shoeColorToSearch)
		if !got {
			tries--
			fmt.Println("==============> Waiting for 5 seconds", "----> ", tries, " tries left <==============")
			time.Sleep(1 * time.Second)
			if tries == 0 {
				fmt.Println("==============> Could not find your ", shoeColorToSearch, "color shoe after 8 tries so terminating search <==============")
			}
		} else {
			got = false
			return
		}
	}
}
