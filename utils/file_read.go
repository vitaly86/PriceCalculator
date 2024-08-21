package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkPath(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("error opening file")
	}
	return file, nil
}

func FindPrices() ([]float64, error, bool) {
	var prices []float64
	file, err := checkPath(Path_prices)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		price, err := strconv.ParseFloat(line, 64)
		if err != nil || line == "" {
			continue
		}
		prices = append(prices, price)
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("failed to parse stored prices"), false
	}

	if len(prices) == 0 {
		return nil, errors.New("file does not contain data"), false
	}

	return prices, nil, true
}

func FindProducts() ([]string, []string) {
	var ids []string
	var products []string
	file, err := checkPath(Path_products)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Id : Product" || line == "" {
			continue
		} else {
			sep := strings.Index(line, ":")
			id := strings.TrimSpace(line[:sep])
			product := strings.TrimSpace(line[sep+1:])
			ids = append(ids, id)
			products = append(products, product)
		}
	}

	return ids, products
}
