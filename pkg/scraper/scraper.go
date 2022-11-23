package scraper

import (
	"fmt"
	"strings"

	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/models"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/utils"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// A scraper function that takes a url and a selector and returns a slice of Product structs and
// the lowestPriceProdfind in the sscraped site.
func Scraper(prod, baseUrl string) (Products []models.Product) {

	// Replace blank spaces with '%20' to make the URL work
	var urlReplacer *strings.Replacer = strings.NewReplacer(" ", "%20")
	prod = urlReplacer.Replace(prod)
	searchUrl := fmt.Sprintf("%s%s%s", baseUrl, "/pesquisa?pg=1&t=", prod)
	c := colly.NewCollector()
	s := Selector()
	c.OnHTML(s.main, func(e *colly.HTMLElement) {
		e.ForEach(s.container, func(_ int, h *colly.HTMLElement) {
			price := h.ChildText(s.fullPrice)
			fullPrice, _ := utils.Float64Converter(price)
			link := h.ChildAttr(s.link, "href")
			// Join the base URL with the link to make the complete URL
			link = fmt.Sprintf("%s%s", baseUrl, link)
			Products = append(Products, models.Product{
				Name:      h.ChildText(s.name),
				FullPrice: fullPrice,
				Link:      link,
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.Request, "\nError:", err)
	})

	c.OnHTML(s.nextPage, func(e *colly.HTMLElement) {
		nextPage := urlReplacer.Replace(fmt.Sprintf("%s%s", baseUrl, e.Attr("href")))
		nextPageStr := e.Text
		isNextPage := nextPageStr == s.nextPagetext
		if isNextPage {
			c.Visit(nextPage)
		} else {
			e.Request.Abort()
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Uses a random User-Agent in each request
	extensions.RandomUserAgent(c)

	c.Visit(searchUrl)

	return
}
