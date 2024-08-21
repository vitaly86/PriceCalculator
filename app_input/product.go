package app_input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

type Product struct {
	id    string
	name  string
	price float64
}

func (product Product) IncludeTaxRate(tax int) float64 {
	return product.price * (1 + float64(tax)/100)
}

func New() (*Product, error) {
	id := getId()
	name := getName()
	price := getPrice()

	if name == "" || price == 0 {
		return nil, errors.New("all fields are required")
	}
	return &Product{
		id:    id,
		name:  name,
		price: price,
	}, nil
}

func ShowNew(ids, products []string, prices []float64) []Product {
	var showProducts []Product = make([]Product, len(ids))

	for index := 0; index < len(ids); index++ {
		id := ids[index]
		name := products[index]
		price := prices[index]

		showProducts[index] = Product{
			id:    id,
			name:  name,
			price: price,
		}
	}
	return showProducts
}

func UpdateNew(ids, products []string) []Product {
	var updatedPrices []Product = make([]Product, len(ids))

	for index := 0; index < len(ids); index++ {
		fmt.Printf("Item %v\n", index+1)
		fmt.Println("----------------")
		fmt.Printf("Product_ID: %v\n", ids[index])
		fmt.Printf("Product_Name: %v\n", products[index])
		id := ids[index]
		name := products[index]
		price := getPrice()
		fmt.Println("----------------")

		updatedPrices[index] = Product{
			id:    id,
			name:  name,
			price: price,
		}
	}
	return updatedPrices
}

func (product *Product) ExtractID() string {
	return product.id
}

func (product *Product) ExtractProduct() string {
	return product.name
}

func (product *Product) ExtractPrice() float64 {
	return product.price
}

func getId() string {
	return randomdata.StringNumber(2, "-")
}
func getName() string {
	output("Product_Name: ")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
func getPrice() float64 {
	var prodPrice float64
	output("Product_Price: ")
	fmt.Scanln(&prodPrice)
	return prodPrice
}

func output(text string) {
	fmt.Print(text)
}
