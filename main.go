package main

import (
	"fmt"
	"strings"

	"project.com/price-calculator/app_input"
	taxes "project.com/price-calculator/computation"
	"project.com/price-calculator/utils"
)

func main() {
	for {
		_, err, valid := utils.FindPrices()
		if err != nil {
			fmt.Println(err)
		}
		finalProduct := taxes.AdjustedProduct{}
		if valid {
			var option string
			fmt.Println("-------------------------------------")
			fmt.Println("Show / Update / Reset / Adjust / Quit")
			fmt.Println("-------------------------------------")
			fmt.Scanln(&option)
			option = strings.ToLower(option)
			switch option {
			case "show", "s":
				utils.DisplayProducts()
			case "update", "u":
				utils.UpdateProducts()
			case "reset", "r":
				utils.ResetProducts()
			case "adjust", "a":
				finalProduct.GetResults()
			case "quit", "q":
				fmt.Println("Have a nice day")
				return
			default:
				fmt.Println("Mismatch! Please consider an option")
				continue
			}
		} else {
			container := app_input.ReadProduct()
			utils.WriteProductsToFile(&container, utils.Path_products)
			utils.WritePricesToFile(&container, utils.Path_prices)
		}
	}
}
