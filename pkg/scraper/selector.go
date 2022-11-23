package scraper

// Export ProductSelector to be used like a constant, carrying the selectors of the website
func Selector() ProdSelector {
	return ProdSelector{
		main:         "ul.row",
		container:    "div.inner",
		name:         "h3.name.no-medium.no-tablet",
		fullPrice:    "strong.sale-price > span:nth-child(1)",
		link:         "a",
		nextPage:     "a.page-next",
		nextPagetext: "Próximo",
	}
}
