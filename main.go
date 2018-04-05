package main

// TODO: Não duplicar as imagens, usar as imagens originais

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marcoprado17/ecommerce-data-generator/product"
	"github.com/marcoprado17/ecommerce-data-generator/seller"
	"github.com/marcoprado17/ecommerce-data-generator/utils"
)

const productURL = "https://enextmarketads-marketplace-a.herokuapp.com/seller/%s/product/"
const imageURL = "https://enextmarketads-marketplace-a.herokuapp.com/imgs/%s"

func main() {
	data := make(map[string]*seller.Seller)

	for i := 0; i < 70; i++ {
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
		storeName := utils.GetRandPhrase(3, 7, 2, 4)
		sellerName := utils.GetRandPhrase(3, 7, 3, 6)
		data[sellerID] = &seller.Seller{
			StoreName:  &storeName,
			SellerName: &sellerName,
			Products:   make([]product.Product, 0),
		}
		for j := 0; j < nProducts; j++ {
			data[sellerID].AddProduct(product.NewSimpleValid(
				fmt.Sprintf(productURL, sellerID)+"%s",
				imageURL,
			))
		}
	}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("data.json", b, 0644)
}
