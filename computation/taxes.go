package taxes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"project.com/price-calculator/app_input"
	"project.com/price-calculator/utils"
)

type AdjustedProduct struct {
	Tax    int      `json:"tax"`
	Prices []string `json:"prices"`
}

func getRate() []int {
	var nbTaxRate int
	fmt.Println("How many Tax Rates?")
	fmt.Scanln(&nbTaxRate)

	var taxContainer []int = make([]int, nbTaxRate)
	var tax int
	for index := 0; index < nbTaxRate; index++ {
		fmt.Printf("Tax Rate %v: ", index+1)
		fmt.Scanln(&tax)
		taxContainer[index] = tax
	}
	return taxContainer
}

func (finalProduct AdjustedProduct) AdjustPrices(products []app_input.Product, tax int) AdjustedProduct {
	var adjusted []string = make([]string, len(products))
	for index, product := range products {
		modPrice := product.IncludeTaxRate(tax)
		strModPrice := fmt.Sprintf("%.2f", modPrice)
		adjusted[index] = strModPrice
	}
	return AdjustedProduct{
		Tax:    tax,
		Prices: adjusted,
	}
}

func (product AdjustedProduct) GetResults() {
	jsonPath := "output/results.json"
	taxes := getRate()
	finalProduct := make([]AdjustedProduct, len(taxes))
	ids, products := utils.FindProducts()
	prices, err, _ := utils.FindPrices()
	if err != nil {
		fmt.Println(err)
	}
	listProducts := app_input.ShowNew(ids, products, prices)

	// Step 1: Create or open the file for writing (truncate existing content)
	file, err := os.OpenFile(jsonPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open the file: %v", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Failed to get file info: %v\n", err)
	}

	for index, tax := range taxes {
		finalProduct[index] = product.AdjustPrices(listProducts, tax)
	}
	if fileInfo.Size() == 0 {
		newERR := initJSON(finalProduct, file)
		if newERR != nil {
			fmt.Println(newERR)
		}
	}
	fmt.Println("Results are saved on results.json file")
	defer file.Close()
}

func initJSON(product []AdjustedProduct, file *os.File) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err := encoder.Encode(product)
	if err != nil {
		return errors.New("failed to write JSON to file")
	}
	return err
}
