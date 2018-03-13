package seller

import (
	"github.com/marcoprado17/ecommerce-data-generator/product"
)

// Seller represents a seller
type Seller struct {
	StoreName  *string
	SellerName *string
	Products   []product.Product
}

// AddProduct adds a new product to the seller
func (seller *Seller) AddProduct(prod *product.Product) {
	seller.Products = append(seller.Products, *prod)
}
