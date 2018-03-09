package product

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
	id                  *string
	title               *string
	description         *string
	link                *string
	imageLink           *string
	additionalImageLink *string
	mobileLink          *string
}

type priceAndAvailability struct {
	availability           *string
	availabilityDate       *string
	expirationDate         *string
	price                  *string
	salePrice              *string
	salePriceEffectiveDate *string
	unitPricingMeasure     *string
	unitPricingBaseMeasure *string
	installment            *string
	loyaltyPoints          *string
}

type productCategory struct {
	googleProductCategory *string
	productType           *string
}

type productIdentifiers struct {
	brand            *string
	gtin             *string
	mpn              *string
	identifierExists *string
}

type detailedProductDescription struct {
	condition                *string
	adult                    *string
	multipack                *uint
	isBundle                 *string
	energyEfficiencyClass    *string
	minEnergyEfficiencyClass *string
	maxEnergyEfficiencyClass *string
	ageGroup                 *string
	color                    *string
	gender                   *string
	material                 *string
	pattern                  *string
	size                     *string
	sizeType                 *string
	sizeSystem               *string
	itemGroupID              *string
}

type shoppingCampaignsAndOtherConfigurations struct {
	adwordsRedirect     *string
	excludedDestination *string
	includedDestination *string
	customLabel0        *string
	promotionID         *string
}

type shipping struct {
	shipping        *string
	shippingLabel   *string
	shippingWeight  *string
	shippingLength  *string
	shippingWidth   *string
	shippingHeight  *string
	maxHandlingTime *string
	minHandlingTime *string
}

type tax struct {
	tax         *string
	taxCategory *string
}
