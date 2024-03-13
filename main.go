package main

import (
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("ul a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnHTML(".testSource .rules:first-of-type", func(e *colly.HTMLElement) {
		println(e.Text)
	})
	c.Visit("https://www.w3.org/Style/CSS/Test/CSS3/Selectors/current/html/full/flat/index.html")

}
