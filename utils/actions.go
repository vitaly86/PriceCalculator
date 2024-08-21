package utils

import (
	"fmt"

	"project.com/price-calculator/app_input"
)

func DisplayProducts() {
	ids, products := FindProducts()
	prices, err, _ := FindPrices()
	if err != nil {
		fmt.Println(err)
	}
	var listProducts []app_input.Product = app_input.ShowNew(ids, products, prices)
	outputProducts := make([]string, len(listProducts))
	for idx, product := range listProducts {
		price := fmt.Sprint(product.ExtractPrice())
		outputProducts[idx] = product.ExtractID() + " : " + product.ExtractProduct() + " : " + price
		fmt.Println(outputProducts[idx])
	}
}

func UpdateProducts() {
	ids, products := FindProducts()
	newListPrices := app_input.UpdateNew(ids, products)
	WritePricesToFile(&newListPrices, Path_prices)
}

func ResetProducts() {
	container := app_input.ReadProduct()
	WriteProductsToFile(&container, Path_products)
	WritePricesToFile(&container, Path_prices)
}
