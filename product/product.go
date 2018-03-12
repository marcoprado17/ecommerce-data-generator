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
	id                  string
	title               string
	description         string
	link                string
	imageLink           string
	additionalImageLink string
	mobileLink          string
}

type priceAndAvailability struct {
	availability           string
	availabilityDate       string
	expirationDate         string
	price                  string
	salePrice              string
	salePriceEffectiveDate string
	unitPricingMeasure     string
	unitPricingBaseMeasure string
	installment            string
	loyaltyPoints          string
}

type productCategory struct {
	googleProductCategory string
	productType           string
}

type productIdentifiers struct {
	brand            string
	gtin             string
	mpn              string
	identifierExists string
}

type detailedProductDescription struct {
	condition                string
	adult                    string
	multipack                uint
	isBundle                 string
	energyEfficiencyClass    string
	minEnergyEfficiencyClass string
	maxEnergyEfficiencyClass string
	ageGroup                 string
	color                    string
	gender                   string
	material                 string
	pattern                  string
	size                     string
	sizeType                 string
	sizeSystem               string
	itemGroupID              string
}

type shoppingCampaignsAndOtherConfigurations struct {
	adwordsRedirect     string
	excludedDestination string
	includedDestination string
	customLabel0        string
	promotionID         string
}

type shipping struct {
	shipping        string
	shippingLabel   string
	shippingWeight  string
	shippingLength  string
	shippingWidth   string
	shippingHeight  string
	maxHandlingTime string
	minHandlingTime string
}

type tax struct {
	tax         string
	taxCategory string
}

// NewSimpleValid returns a new instance of a simple valid product
func NewSimpleValid(link, imageLink string) *Product {
	product := new(Product)

	product.id = utils.GetRandWord(10, 10)
	product.title = utils.GetRandPhrase(3, 7, 2, 3)
	product.description = utils.GetRandPhrase(3, 7, 10, 20)
	product.link = fmt.Sprintf(link, product.id)
	product.imageLink = utils.GetImage(imageLink)

	product.availability, _ = utils.GetValue([]utils.ValueProbPair{
		utils.ValueProbPair{V: "in stock", P: 0.8},
		utils.ValueProbPair{V: "out of stock", P: 0.2},
	}).(string)
	product.price = utils.GetRandPrice(1.0, 150.0, "BRL")

	product.brand = utils.GetRandPhrase(3, 10, 1, 2)
	product.gtin = utils.GetRandGtin()

	product.condition, _ = utils.GetValue([]utils.ValueProbPair{
		utils.ValueProbPair{V: "new", P: 0.33},
		utils.ValueProbPair{V: "refurbished", P: 0.33},
		utils.ValueProbPair{V: "used", P: 0.33},
	}).(string)

	return product
}

// NewInvalidByAusentTitle returns a new instance of a product that doesn't have title
func NewInvalidByAusentTitle(link, imageLink string) *Product {
	product := NewSimpleValid(link, imageLink)
	product.title = ""
	return product
}
