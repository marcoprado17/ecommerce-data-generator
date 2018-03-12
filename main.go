package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marcoprado17/ecommerce-data-generator/product"
	"github.com/marcoprado17/ecommerce-data-generator/utils"
)

const productURL = "https://enextmarketads-marketplace-a.herokuapp.com/seller/%s/product/"
const imageURL = "https://enextmarketads-marketplace-a.herokuapp.com/public/imgs/%s"

func main() {
	data := make(map[string][]product.Product)

	for i := 0; i < 40; i++ {
		sellerID := utils.GetRandID(10, 10)
		nProducts := 0
		switch i / 10 {
		case 0:
			nProducts = 0
			break
		case 1:
			nProducts = 1
			break
		case 2:
			nProducts = 5
			break
		case 3:
			nProducts = 10
			break
		case 4:
			nProducts = 20
			break
		case 5:
			nProducts = 40
			break
		case 6:
			nProducts = 80
			break
		default:
			nProducts = 1
			break
		}
		fmt.Println(i)
		data[sellerID] = make([]product.Product, nProducts)
		for j := 0; j < nProducts; j++ {
			data[sellerID][j] = *product.NewSimpleValid(
				fmt.Sprintf(productURL, sellerID)+"%s",
				imageURL,
			)
		}
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("data.json", b, 0644)
}
