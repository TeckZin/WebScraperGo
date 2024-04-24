package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type item struct {
	Name string   `json:"name"`
	Menu []string `json:"menu"`
}

func main() {

	c := colly.NewCollector(
	//colly.AllowedDomains("amazon.com"),
	)

	var items []item

	c.OnHTML("div.content div[class='ui large vertical fluid tabular menu'] div.item", func(h *colly.HTMLElement) {
		//fmt.Println(h.ChildText("div.item div.header"))

		texts := strings.Split(h.ChildText("div.menu"), "\n")
		//fmt.Println(h.ChildText("div.menu"))
		//fmt.Println(len(texts))

		//c.OnHTML("div.menu", func(h2 *colly.HTMLElement) {
		//	texts = append(texts, h2.ChildText("a.item"))
		//	fmt.Println(h2.ChildText("a.item"))
		//
		//})

		//fmt.Println(h.ChildText("div.menu a.item"))

		item := item{
			Name: h.ChildText("div.header"),
			Menu: texts,
		}
		//fmt.Println(item.Name)
		//fmt.Println(item.Menu)

		items = append(items, item)

		//fmt.Println(h.Text)

	})

	err := c.Visit("https://go-colly.org/docs/introduction/start/")
	if err != nil {
		fmt.Println("fail")
		return
	}

	fmt.Println(items)

}
