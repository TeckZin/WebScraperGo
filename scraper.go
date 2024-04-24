package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type item struct {
	Name string   `json:"name"`
	Menu []string `json:"menu"`
	Urls []string `json:"urls"`
}

func main() {

	c := colly.NewCollector(
	//colly.AllowedDomains("amazon.com"),
	)

	var items []item

	c.OnHTML("div.content div[class='ui large vertical fluid tabular menu'] div.item", func(h *colly.HTMLElement) {
		//fmt.Println(h.ChildText("div.item div.header"))

		//fmt.Println(h.ChildText("div.menu"))
		//fmt.Println(len(texts))

		//c.OnHTML("div.menu", func(h2 *colly.HTMLElement) {
		//	texts = append(texts, h2.ChildText("a.item"))
		//	fmt.Println(h2.ChildText("a.item"))
		//
		//})

		//fmt.Println(h.ChildText("div.menu a.item"))

		var texts []string
		var urls []string

		h.ForEach("div.menu a", func(_ int, h *colly.HTMLElement) {
			//fmt.Println("next")

			texts = append(texts, h.Text)
			urls = append(urls, h.Attr("href"))

			//fmt.Println(h.Attr("href"))

			//fmt.Println(h.Text)

		})

		item := item{
			Name: h.ChildText("div.header"),
			Menu: texts,
			Urls: urls,
		}

		//fmt.Println(item.Name)
		//fmt.Println(item.Menu)

		items = append(items, item)

		//fmt.Println(h.Text)

	})

	// interact with button
	//c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
	//	nextPage := h.Request.AbsoluteURL(h.Attr("href"))
	//	c.Visit(nextPage)
	//})

	err := c.Visit("https://go-colly.org/docs/introduction/start/")
	if err != nil {
		fmt.Println("fail")
		return
	}

	fmt.Println(items)

}
