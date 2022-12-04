package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/db"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/gui"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/models"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/scraper"
	"github.com/Antonio-Costa00/Go-Price-Monitor/pkg/sendmessage"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	from := os.Getenv("FROM")
	token := os.Getenv("X_API_TOKEN")

	DB := db.Init(dbHost, dbUser, dbPass, dbName, dbPort)

	ug := gui.UserGUI()
	tel := ug.Telephone
	prod := ug.Product
	val := ug.Value

	isProdEmpty := prod == ""
	if isProdEmpty {
		fmt.Println("Please enter a product name")
		return
	}

	var products []models.Product
	// Product with the lowest price
	var bestProd models.Product
	const url = "https://www.brasiltronic.com.br"
	// Scrape the website until the lowest price is found
	for {
		products = scraper.Scraper(prod, url)
		fmt.Println("Scraping...")
		fmt.Println("Products found: ", len(products))
		if len(products) == 0 {
			fmt.Println("No products found")
			return
		}
		bestProd = scraper.GetLowestPriceProd(products)
		fmt.Println("Best product: ", bestProd)
		isProdInSale := val > bestProd.FullPrice
		if isProdInSale {
			msg := fmt.Sprintf("🚨 *Price dropped!* The product %s is on sale for R$%.2f. \nAccess the link to buy: %s", bestProd.Name, bestProd.FullPrice, bestProd.Link)
			sendmessage.SendMessage(from, token, tel, msg)
			fmt.Println("Message sent with success!")
			return
		}
		// Insert the products in the database if they don't exist
		for _, p := range products {
			DB.FirstOrCreate(&p, p)
		}
		// Wait for 10 seconds to check again
		fmt.Println("The product is not on sale. Waiting for 10 seconds to check again...")
		time.Sleep(10 * time.Second)
	}
}
