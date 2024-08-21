package app_input

import "fmt"

func ReadProduct() (container []Product) {
	container = make([]Product, 3)
	fmt.Println("Please insert the store items")
	var index int = 0
	var option string
	for {
		confirmation := true
		refusal := true
		fmt.Printf("Item %v\n", index+1)
		fmt.Println("----------------")
		product, err := New()
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			if index < 3 {
				container[index] = *product
			} else {
				container = append(container, *product)
			}
			fmt.Print("New Items? [yes/no] ")
			fmt.Scanln(&option)
			if option == "yes" || option == "y" {
				index++
				continue
			} else if option == "no" || option == "n" {
				fmt.Println("You will get the final results in a JSON format")
				break
			} else {
				var inner string
				confirmation = false
				refusal = false
				for !confirmation && !refusal {
					fmt.Print("You have to chose an option [yes / no]: ")
					fmt.Scanln(&inner)
					switch inner {
					case "yes", "y":
						index++
						confirmation = true
					case "no", "n":
						fmt.Println("You will get the final results in a JSON format")
						refusal = true
					}
				}
			}
			if confirmation {
				continue
			}
			if refusal {
				break
			}
		}
	}
	return container
}
