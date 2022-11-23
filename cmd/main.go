package main

import (
	"fmt"
	"time"

	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/db"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/gui"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/models"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/scraper"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/sendmessage"
)

func main() {

	DB := db.Init()

	ug := gui.UserGUI()
	tel := ug.Telephone
	prod := ug.Product
	val := ug.Value

	var products []models.Product
	// Product with the lowest price
	var bestProd models.Product
	const url = "https://www.brasiltronic.com.br"
	// Scrape the website until the lowest price is found
	for {
		products = scraper.Scraper(prod, url)
		bestProd = scraper.GetLowestPriceProd(products)
		isProdInSale := val > bestProd.FullPrice
		if isProdInSale {
			msg := fmt.Sprintf("🚨 *Abaixou o preço!* O produto %s está em promoção no valor de R$%.2f. \nAcesse o link para compra: %s", bestProd.Name, bestProd.FullPrice, bestProd.Link)
			w := sendmessage.WhatsApp{
				To:       tel,
				Contents: []sendmessage.Contents{{Type: "text", Text: msg}},
			}
			sendmessage.SendMessage(w)
			fmt.Println("Mensagem enviada com sucesso!")
			return
		}
		// Insert the products in the database if they don't exist
		for _, p := range products {
			DB.FirstOrCreate(&p, p)
		}
		// Wait for 10 seconds to check again
		fmt.Println("The product is not in sale. Waiting for 10 seconds to check again...")
		time.Sleep(10 * time.Second)
	}
}
