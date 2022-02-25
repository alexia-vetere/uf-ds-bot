package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	allFacts := make([]string, 0)

	c := colly.NewCollector()

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			ufValor := el.ChildText("td:nth-child(2)")
			allFacts = append(allFacts, ufValor)
		})

	})
	c.Visit("https://valoruf.cl/")
	fmt.Println("1 UF =", allFacts[1], "CLP")
}
