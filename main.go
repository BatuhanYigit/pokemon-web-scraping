package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PokemonProduct struct {
	link      string
	name      string
	price     string
	imagelink string
}

func main() {
	var pokemonProducts []PokemonProduct
	c := colly.NewCollector()

	c.OnHTML("li.page-numbers, a.page-numbers", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))

	})

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.link = h.ChildAttr("a", "href")
		pokemonProduct.name = h.ChildText("h2")
		pokemonProduct.price = h.ChildText("span.price")
		pokemonProduct.imagelink = h.ChildAttr("img", "src")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.Visit("https://scrapeme.live/shop/page/1/")

	file, err := os.Create("./pokemon.csv")
	if err != nil {
		log.Fatalln("Failed create csv ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"link",
		"name",
		"price",
		"imagelink",
	}
	writer.Write(headers)

	for _, pokemonProduct := range pokemonProducts {
		record := []string{
			pokemonProduct.link,
			pokemonProduct.name,
			pokemonProduct.price,
			pokemonProduct.imagelink,
		}

		writer.Write(record)

	}

	defer writer.Flush()

	fmt.Println("Pokemon List CSV Created ! ")

}
