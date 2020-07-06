package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type post struct {
	title string `json:"title"`
	link  string `json:"link"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	var posts []post

	// On every a element which has href attribute call callback
	// span.a-size-medium sc-product-title a-text-bold
	c.OnHTML(".title a", func(e *colly.HTMLElement) {

		// fmt.Printf(e.Attr("a"))

		indivTitle := e.Text
		indivLink := e.Attr("href")
		posts = append(posts, post{title: indivTitle, link: indivLink})

		// Print link
		fmt.Printf("%q \n%s\n \n", indivTitle, indivLink)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping 1. cart, 2. all carts (fresh, whole foods, amazon)
	// c.Visit("https://www.amazon.com/gp/cart/view.html?ref_=nav_cart")
	c.Visit("https://news.ycombinator.com/")
	// fmt.Printf("%+v\n", posts)

	// Write json to file
	postsJSON, _ := json.Marshal(posts)
	ioutil.WriteFile("postsFile.json", postsJSON, 0644)
}
