package product

import (
	"fmt"
	"math/rand"

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

func str(s string) *string {
	return &s
}

func GetImage(imageLink string, fileName string) string {
	return fmt.Sprintf(imageLink, fileName)
}

// NewSimpleValid returns a new instance of a simple valid product
func NewSimpleValid(link, imageLink string) *Product {
	var pID = utils.GetRandWord(10, 10)
	var pTitle = "Panela Tramontina Grande Vermelha"
	var pDescription = "Panela para conzinhar todos os tipos de alimentos"
	var pLink = fmt.Sprintf(link, pID)
	var pImageLink = GetImage(imageLink, "panela.jpg")
	var pPrice = "69.90 BRL"
	var pAvailability = "in stock"
	var pBrand = "Tramontina"
	var pGtin = "1488713532784"
	var pCondition = "new"

	var cID = utils.GetRandWord(10, 10)
	var cTitle = "Camisa Social Ralph Lauren Azul"
	var cDescription = "Camisa Social Ralph Lauren Azul com fios de linho"
	var cLink = fmt.Sprintf(link, cID)
	var cImageLink = GetImage(imageLink, "camisa_azul_claro.jpg")
	var cPrice = "149.90 BRL"
	var cAvailability = "in stock"
	var cBrand = "Ralph Lauren"
	var cGtin = "1498713532120"
	var cCondition = "new"

	var tID = utils.GetRandWord(10, 10)
	var tTitle = "Tênis Mizuno Rx-37"
	var tDescription = "Tênis para corrida"
	var tLink = fmt.Sprintf(link, tID)
	var tImageLink = GetImage(imageLink, "tenis.jpg")
	var tPrice = "199.90 BRL"
	var tAvailability = "in stock"
	var tBrand = "Mizuno"
	var tGtin = "1508713532120"
	var tCondition = "new"

	var vID = utils.GetRandWord(10, 10)
	var vTitle = "Playstation 4"
	var vDescription = "Playstation 4"
	var vLink = fmt.Sprintf(link, vID)
	var vImageLink = GetImage(imageLink, "videogame.jpg")
	var vPrice = "2500.00 BRL"
	var vAvailability = "in stock"
	var vBrand = "Sony"
	var vGtin = "1518713532120"
	var vCondition = "new"

	var sID = utils.GetRandWord(10, 10)
	var sTitle = "Smartphone Sansung J5"
	var sDescription = "Smartphone Sansung J5"
	var sLink = fmt.Sprintf(link, sID)
	var sImageLink = GetImage(imageLink, "celular.jpg")
	var sPrice = "875.50 BRL"
	var sAvailability = "in stock"
	var sBrand = "Sony"
	var sGtin = "1528713532120"
	var sCondition = "new"

	var vestID = utils.GetRandWord(10, 10)
	var vestTitle = "Vestido Azul Estampado Maria Valentina"
	var vestDescription = "Vestido Azul Estampado Maria Valentina"
	var vestLink = fmt.Sprintf(link, vestID)
	var vestImageLink = GetImage(imageLink, "vestido_azul.jpg")
	var vestPrice = "399.99 BRL"
	var vestAvailability = "in stock"
	var vestBrand = "Maria Valentina"
	var vestGtin = "1538713532120"
	var vestCondition = "new"

	var vestBrancoID = utils.GetRandWord(10, 10)
	var vestBrancoTitle = "Vestido Branco Maria Valentina"
	var vestBrancoDescription = "Vestido Branco Maria Valentina"
	var vestBrancoLink = fmt.Sprintf(link, vestBrancoID)
	var vestBrancoImageLink = GetImage(imageLink, "vestido_branco.jpg")
	var vestBrancoPrice = "800.00 BRL"
	var vestBrancoAvailability = "in stock"
	var vestBrancoBrand = "Maria Valentina"
	var vestBrancoGtin = "1548713532120"
	var vestBrancoCondition = "new"

	var ovoID = utils.GetRandWord(10, 10)
	var ovoTitle = "Ovo de Páscoa Bis"
	var ovoDescription = "Ovo de Páscoa Bis"
	var ovoLink = fmt.Sprintf(link, ovoID)
	var ovoImageLink = GetImage(imageLink, "ovo.jpg")
	var ovoPrice = "32.80 BRL"
	var ovoAvailability = "in stock"
	var ovoBrand = "Lacta"
	var ovoGtin = "1558713532120"
	var ovoCondition = "new"

	var camisaID = utils.GetRandWord(10, 10)
	var camisaTitle = "Camisa Azul Vila Romana"
	var camisaDescription = "Camisa Azul Vila Romana"
	var camisaLink = fmt.Sprintf(link, camisaID)
	var camisaImageLink = GetImage(imageLink, "camisa_azul_escuro.jpg")
	var camisaPrice = "259.90 BRL"
	var camisaAvailability = "in stock"
	var camisaBrand = "Vila Romana"
	var camisaGtin = "1568713532120"
	var camisaCondition = "new"

	var noteID = utils.GetRandWord(10, 10)
	var noteTitle = "Notebook Sansung RX89"
	var noteDescription = "Notebook Sansung RX89"
	var noteLink = fmt.Sprintf(link, noteID)
	var noteImageLink = GetImage(imageLink, "note.jpg")
	var notePrice = "1999.90 BRL"
	var noteAvailability = "in stock"
	var noteBrand = "Sansung"
	var noteGtin = "1578713532120"
	var noteCondition = "new"

	var sampleBasicProductDatas = [...]basicProductData{
		basicProductData{
			ID:          &pID,
			Title:       &pTitle,
			Description: &pDescription,
			Link:        &pLink,
			ImageLink:   &pImageLink,
		},
		basicProductData{
			ID:          &cID,
			Title:       &cTitle,
			Description: &cDescription,
			Link:        &cLink,
			ImageLink:   &cImageLink,
		},
		basicProductData{
			ID:          &tID,
			Title:       &tTitle,
			Description: &tDescription,
			Link:        &tLink,
			ImageLink:   &tImageLink,
		},
		basicProductData{
			ID:          &vID,
			Title:       &vTitle,
			Description: &vDescription,
			Link:        &vLink,
			ImageLink:   &vImageLink,
		},
		basicProductData{
			ID:          &sID,
			Title:       &sTitle,
			Description: &sDescription,
			Link:        &sLink,
			ImageLink:   &sImageLink,
		},
		basicProductData{
			ID:          &vestID,
			Title:       &vestTitle,
			Description: &vestDescription,
			Link:        &vestLink,
			ImageLink:   &vestImageLink,
		},
		basicProductData{
			ID:          &vestBrancoID,
			Title:       &vestBrancoTitle,
			Description: &vestBrancoDescription,
			Link:        &vestBrancoLink,
			ImageLink:   &vestBrancoImageLink,
		},
		basicProductData{
			ID:          &ovoID,
			Title:       &ovoTitle,
			Description: &ovoDescription,
			Link:        &ovoLink,
			ImageLink:   &ovoImageLink,
		},
		basicProductData{
			ID:          &camisaID,
			Title:       &camisaTitle,
			Description: &camisaDescription,
			Link:        &camisaLink,
			ImageLink:   &camisaImageLink,
		},
		basicProductData{
			ID:          &noteID,
			Title:       &noteTitle,
			Description: &noteDescription,
			Link:        &noteLink,
			ImageLink:   &noteImageLink,
		},
	}

	var samplePriceAndAvailabilities = [...]priceAndAvailability{
		priceAndAvailability{
			Price:        &pPrice,
			Availability: &pAvailability,
		},
		priceAndAvailability{
			Price:        &cPrice,
			Availability: &cAvailability,
		},
		priceAndAvailability{
			Price:        &tPrice,
			Availability: &tAvailability,
		},
		priceAndAvailability{
			Price:        &vPrice,
			Availability: &vAvailability,
		},
		priceAndAvailability{
			Price:        &sPrice,
			Availability: &sAvailability,
		},
		priceAndAvailability{
			Price:        &vestPrice,
			Availability: &vestAvailability,
		},
		priceAndAvailability{
			Price:        &vestBrancoPrice,
			Availability: &vestBrancoAvailability,
		},
		priceAndAvailability{
			Price:        &ovoPrice,
			Availability: &ovoAvailability,
		},
		priceAndAvailability{
			Price:        &camisaPrice,
			Availability: &camisaAvailability,
		},
		priceAndAvailability{
			Price:        &notePrice,
			Availability: &noteAvailability,
		},
	}

	var sampleProductIdentifiers = [...]productIdentifiers{
		productIdentifiers{
			Gtin:  &pGtin,
			Brand: &pBrand,
		},
		productIdentifiers{
			Gtin:  &cGtin,
			Brand: &cBrand,
		},
		productIdentifiers{
			Gtin:  &tGtin,
			Brand: &tBrand,
		},
		productIdentifiers{
			Gtin:  &vGtin,
			Brand: &vBrand,
		},
		productIdentifiers{
			Gtin:  &sGtin,
			Brand: &sBrand,
		},
		productIdentifiers{
			Gtin:  &vestGtin,
			Brand: &vestBrand,
		},
		productIdentifiers{
			Gtin:  &vestBrancoGtin,
			Brand: &vestBrancoBrand,
		},
		productIdentifiers{
			Gtin:  &ovoGtin,
			Brand: &ovoBrand,
		},
		productIdentifiers{
			Gtin:  &camisaGtin,
			Brand: &camisaBrand,
		},
		productIdentifiers{
			Gtin:  &noteGtin,
			Brand: &noteBrand,
		},
	}

	var sampleDetailedProductDescriptions = [...]detailedProductDescription{
		detailedProductDescription{
			Condition: &pCondition,
		},
		detailedProductDescription{
			Condition: &cCondition,
		},
		detailedProductDescription{
			Condition: &tCondition,
		},
		detailedProductDescription{
			Condition: &vCondition,
		},
		detailedProductDescription{
			Condition: &sCondition,
		},
		detailedProductDescription{
			Condition: &vestCondition,
		},
		detailedProductDescription{
			Condition: &vestBrancoCondition,
		},
		detailedProductDescription{
			Condition: &ovoCondition,
		},
		detailedProductDescription{
			Condition: &camisaCondition,
		},
		detailedProductDescription{
			Condition: &noteCondition,
		},
	}

	var sampleProducts = [...]Product{
		Product{
			basicProductData:           sampleBasicProductDatas[0],
			priceAndAvailability:       samplePriceAndAvailabilities[0],
			productIdentifiers:         sampleProductIdentifiers[0],
			detailedProductDescription: sampleDetailedProductDescriptions[0],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[1],
			priceAndAvailability:       samplePriceAndAvailabilities[1],
			productIdentifiers:         sampleProductIdentifiers[1],
			detailedProductDescription: sampleDetailedProductDescriptions[1],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[2],
			priceAndAvailability:       samplePriceAndAvailabilities[2],
			productIdentifiers:         sampleProductIdentifiers[2],
			detailedProductDescription: sampleDetailedProductDescriptions[2],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[3],
			priceAndAvailability:       samplePriceAndAvailabilities[3],
			productIdentifiers:         sampleProductIdentifiers[3],
			detailedProductDescription: sampleDetailedProductDescriptions[3],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[4],
			priceAndAvailability:       samplePriceAndAvailabilities[4],
			productIdentifiers:         sampleProductIdentifiers[4],
			detailedProductDescription: sampleDetailedProductDescriptions[4],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[5],
			priceAndAvailability:       samplePriceAndAvailabilities[5],
			productIdentifiers:         sampleProductIdentifiers[5],
			detailedProductDescription: sampleDetailedProductDescriptions[5],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[6],
			priceAndAvailability:       samplePriceAndAvailabilities[6],
			productIdentifiers:         sampleProductIdentifiers[6],
			detailedProductDescription: sampleDetailedProductDescriptions[6],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[7],
			priceAndAvailability:       samplePriceAndAvailabilities[7],
			productIdentifiers:         sampleProductIdentifiers[7],
			detailedProductDescription: sampleDetailedProductDescriptions[7],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[8],
			priceAndAvailability:       samplePriceAndAvailabilities[8],
			productIdentifiers:         sampleProductIdentifiers[8],
			detailedProductDescription: sampleDetailedProductDescriptions[8],
		},
		Product{
			basicProductData:           sampleBasicProductDatas[9],
			priceAndAvailability:       samplePriceAndAvailabilities[9],
			productIdentifiers:         sampleProductIdentifiers[9],
			detailedProductDescription: sampleDetailedProductDescriptions[9],
		},
	}

	return &sampleProducts[rand.Intn(10)]

	// fmt.Println(sampleProducts)

	// product := new(Product)

	// ID := utils.GetRandWord(10, 10)
	// product.ID = &ID
	// title := utils.GetRandPhrase(5, 7, 1, 3)
	// product.Title = &title
	// description := utils.GetRandPhrase(3, 7, 10, 20)
	// product.Description = &description
	// productLink := fmt.Sprintf(link, *product.ID)
	// product.Link = &productLink
	// productImageLink := utils.GetImage(imageLink)
	// product.ImageLink = &productImageLink

	// availability, _ := utils.GetValue([]utils.ValueProbPair{
	// 	utils.ValueProbPair{V: "in stock", P: 0.8},
	// 	utils.ValueProbPair{V: "out of stock", P: 0.2},
	// }).(string)
	// product.Availability = &availability
	// price := utils.GetRandPrice(1.0, 150.0, "BRL")
	// product.Price = &price

	// brand := utils.GetRandPhrase(3, 10, 1, 2)
	// product.Brand = &brand
	// gtin := utils.GetRandGtin()
	// product.Gtin = &gtin

	// condition, _ := utils.GetValue([]utils.ValueProbPair{
	// 	utils.ValueProbPair{V: "new", P: 0.33},
	// 	utils.ValueProbPair{V: "refurbished", P: 0.33},
	// 	utils.ValueProbPair{V: "used", P: 0.33},
	// }).(string)
	// product.Condition = &condition

	// return product
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
