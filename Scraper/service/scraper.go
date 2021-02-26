package service

import (
	"fmt"
	"github.com/SanjanaKansal/scraper/models"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

// ScrapePage scrapes the web page given its URL.
func ScrapePage (pageURL string) (models.SuccessMessage, models.ScrapedData) {
	var scrapingStatus models.SuccessMessage
	var scrapedData models.ScrapedData
	var productData models.ProductDetails
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div#dp", func(e *colly.HTMLElement) {
		// Read Product Name, Product Image URL and Product Price.
		productData.Name =  e.ChildText("span#productTitle")
		productData.ImageURL = e.ChildAttr("img","src")
		productData.Price = e.ChildText("span#priceblock_ourprice")
		if productData.Price == "" {
			productData.Price = "Product out of stock! can not tell the price."
		}
		// Read Product reviews count.
		reviews := e.ChildText("span#acrCustomerReviewText.a-size-base")
		review := strings.Split(reviews, " ")
		productData.TotalReviews, _ = strconv.Atoi(review[0])

		// Read Product Description.
		var des []string
		e.ForEach("div#feature-bullets", func(_ int, e *colly.HTMLElement) {
			des = append(des, e.ChildText("span.a-list-item"))
		})
		for _, description := range des {
			productData.Description = strings.Split(description, "\n\n\n")
		}
	})

	err := c.Visit(pageURL)

	if err != nil {
		scrapingStatus.Success = false
		scrapingStatus.Message = err.Error()
		return scrapingStatus, scrapedData
	}

	scrapedData.URL = pageURL
	scrapedData.Product = productData
	scrapingStatus.Success = true
	scrapingStatus.Message = "Successfully scraped the page."
	return scrapingStatus, scrapedData
}
