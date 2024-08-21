package utils

import (
	"errors"
	"fmt"
	"os"

	"project.com/price-calculator/app_input"
)

// type permitted interface {
// 	int | float64 | string
// }

const Path_prices = "output/prices.txt"
const Path_products = "output/products.txt"

func checkValidationFile(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return errors.New("failed to find prices.txt file")
	}
	return nil
}

func configPrices(price float64) string {
	strPrice := fmt.Sprintf("%.2f\n", price)
	return strPrice
}

func WritePricesToFile(products *[]app_input.Product, path string) {
	err := checkValidationFile(path)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------------------")
	}
	var listPrices string = "Prices\n"
	for _, product := range *products {
		listPrices += configPrices(product.ExtractPrice())
	}
	os.WriteFile(path, []byte(listPrices), 0644)
}

func WriteProductsToFile(products *[]app_input.Product, path string) {
	err := checkValidationFile(path)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------------------")
	}
	var listProducts string = "Id : Product\n"
	for _, product := range *products {
		listProducts += product.ExtractID() + " : " + product.ExtractProduct() + "\n"
	}
	os.WriteFile(path, []byte(listProducts), 0644)
}
