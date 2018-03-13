package product

import (
	"fmt"

	"github.com/marcoprado17/ecommerce-data-generator/utils"
)

// Product represents a product
type Product struct {
	basicProductData
	priceAndAvailability
	productCategory
	productIdentifiers
	detailedProductDescription
	shoppingCampaignsAndOtherConfigurations
	shipping
	tax
}

type basicProductData struct {
	ID                  *string
	Title               *string
	Description         *string
	Link                *string
	ImageLink           *string
	AdditionalImageLink *string
	MobileLink          *string
}

type priceAndAvailability struct {
	Availability           *string
	AvailabilityDate       *string
	ExpirationDate         *string
	Price                  *string
	SalePrice              *string
	SalePriceEffectiveDate *string
	UnitPricingMeasure     *string
	UnitPricingBaseMeasure *string
	Installment            *string
	LoyaltyPoints          *string
}

type productCategory struct {
	GoogleProductCategory *string
	ProductType           *string
}

type productIdentifiers struct {
	Brand            *string
	Gtin             *string
	Mpn              *string
	IdentifierExists *string
}

type detailedProductDescription struct {
	Condition                *string
	Adult                    *string
	Multipack                *uint
	IsBundle                 *string
	EnergyEfficiencyClass    *string
	MinEnergyEfficiencyClass *string
	MaxEnergyEfficiencyClass *string
	AgeGroup                 *string
	Color                    *string
	Gender                   *string
	Material                 *string
	Pattern                  *string
	Size                     *string
	SizeType                 *string
	SizeSystem               *string
	ItemGroupID              *string
}

type shoppingCampaignsAndOtherConfigurations struct {
	AdwordsRedirect     *string
	ExcludedDestination *string
	IncludedDestination *string
	CustomLabel0        *string
	PromotionID         *string
}

type shipping struct {
	Shipping        *string
	ShippingLabel   *string
	ShippingWeight  *string
	ShippingLength  *string
	ShippingWidth   *string
	ShippingHeight  *string
	MaxHandlingTime *string
	MinHandlingTime *string
}

type tax struct {
	Tax         *string
	TaxCategory *string
}

// NewSimpleValid returns a new instance of a simple valid product
func NewSimpleValid(link, imageLink string) *Product {
	product := new(Product)

	ID := utils.GetRandWord(10, 10)
	product.ID = &ID
	title := utils.GetRandPhrase(5, 7, 1, 3)
	product.Title = &title
	description := utils.GetRandPhrase(3, 7, 10, 20)
	product.Description = &description
	productLink := fmt.Sprintf(link, *product.ID)
	product.Link = &productLink
	productImageLink := utils.GetImage(imageLink)
	product.ImageLink = &productImageLink

	availability, _ := utils.GetValue([]utils.ValueProbPair{
		utils.ValueProbPair{V: "in stock", P: 0.8},
		utils.ValueProbPair{V: "out of stock", P: 0.2},
	}).(string)
	product.Availability = &availability
	price := utils.GetRandPrice(1.0, 150.0, "BRL")
	product.Price = &price

	brand := utils.GetRandPhrase(3, 10, 1, 2)
	product.Brand = &brand
	gtin := utils.GetRandGtin()
	product.Gtin = &gtin

	condition, _ := utils.GetValue([]utils.ValueProbPair{
		utils.ValueProbPair{V: "new", P: 0.33},
		utils.ValueProbPair{V: "refurbished", P: 0.33},
		utils.ValueProbPair{V: "used", P: 0.33},
	}).(string)
	product.Condition = &condition

	return product
}

// NewInvalidByAusentTitle returns a new instance of a product that doesn't have title
func NewInvalidByAusentTitle(link, imageLink string) *Product {
	product := NewSimpleValid(link, imageLink)
	product.Title = nil
	return product
}

// NewInvalidByAusentImage returns a new instance of a product that doesn't have the correct image link
func NewInvalidByAusentImage(link, imageLink string) *Product {
	product := NewSimpleValid(link, imageLink)
	*product.ImageLink = fmt.Sprintf(imageLink, utils.GetRandID(10, 10)+".jpg")
	return product
}
