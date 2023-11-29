package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Item struct {
	Link      string `json:"link"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	ImageLink string `json:"imagelink"`
}

func main() {
	c := colly.NewCollector()

	items := []Item{}

	c.OnHTML("li.page-numbers, a.page-numbers", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))

	})

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		i := Item{
			Link:      h.ChildAttr("a", "href"),
			Name:      h.ChildText("h2"),
			Price:     h.ChildText("span.price"),
			ImageLink: h.ChildAttr("img", "src"),
		}
		items = append(items, i)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.Visit("https://scrapeme.live/shop/page/1/")

	data, err := json.MarshalIndent(items, " ", "")
	if err != nil {
		log.Fatal()
	}
	fmt.Println(string(data))

}
