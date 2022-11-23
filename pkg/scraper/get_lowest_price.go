package scraper

import "github.com/Antonio-Costa00/Go-Price-Monitor/pkg/models"

// Returns the product with the lowest price in the slice of products
func GetLowestPriceProd(products []models.Product) (lowestPriceProd models.Product) {
	lowestPrice := products[0].FullPrice
	for _, product := range products {
		if product.FullPrice < lowestPrice {
			lowestPrice = product.FullPrice
			lowestPriceProd = product
		}
	}
	return
}
