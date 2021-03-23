package mws

import "time"

//http://docs.developer.amazonservices.com/en_US/products/Products_Datatypes.html
//http://g-ecx.images-amazon.com/images/G/01/mwsportal/doc/en_US/products/default.xsd
//http://g-ecx.images-amazon.com/images/G/01/mwsportal/doc/en_US/products/ProductsAPI_Response.xsd

//BuyBoxPriceType The price of an item that is displayed in the Buy Box.
type BuyBoxPriceType struct {
	Condition          string    `xml:"condition,attr"`
	FulfillmentChannel string    `xml:"fulfillmentChannel,attr"`
	LandedPrice        MoneyType `xml:"LandedPrice"`
	ListingPrice       MoneyType `xml:"ListingPrice"`
	Shipping           MoneyType `xml:"Shipping"`
	Points             *Points   `xml:"Points"`
}

//DetailedShippingTimeType The maximum time within which the item will likely be shipped once an order has been placed.
type DetailedShippingTimeType struct {
	Text             string     `xml:",chardata"`
	MinimumHours     int        `xml:"minimumHours,attr"`
	MaximumHours     int        `xml:"maximumHours,attr"`
	AvailableDate    *time.Time `xml:"availableDate,attr"`
	AvailabilityType string     `xml:"availabilityType,attr"`
}

//FeeDetail The type of fee, fee amount, and other details.
type FeeDetail struct {
	//TODO
}

//FeesEstimate The total estimated fees for a product and a list of details.
type FeesEstimate struct {
	//TODO
}

//FeesEstimateIdentifier A product identifier, marketplace, time of request, and other details that identify an estimate.
type FeesEstimateIdentifier struct {
	//TODO
}

//FeesEstimateRequest A product, marketplace, and proposed price used to request estimated fees.
type FeesEstimateRequest struct {
	//TODO
}

//FeesEstimateResult A product identifier and the estimated fees for that product.
type FeesEstimateResult struct {
	//TODO
}

//LowestPriceType The lowest price of an item.
type LowestPriceType struct {
	Condition          string    `xml:"condition,attr"`
	FulfillmentChannel string    `xml:"fulfillmentChannel,attr"`
	LandedPrice        MoneyType `xml:"LandedPrice"`
	ListingPrice       MoneyType `xml:"ListingPrice"`
	Shipping           MoneyType `xml:"Shipping"`
	Points             Points    `xml:"Points"`
}

//OfferCountType The number of offers in a fulfillment channel that meet a specific condition.
type OfferCountType struct {
	Text               string `xml:",chardata"`
	Condition          string `xml:"condition,attr"`
	FulfillmentChannel string `xml:"fulfillmentChannel,attr"`
}

//Points The number of Amazon Points offered with the purchase of an item. The Amazon Points program is only available in Japan.
type Points struct {
	PointsNumber        string    `xml:"PointsNumber"`
	PointsMonetaryValue MoneyType `xml:"PointsMonetaryValue"`
}

//PriceToEstimateFees Price information for a product, used to estimate fees.
type PriceToEstimateFees struct {
	ListingPrice MoneyType `xml:"ListingPrice"`
	Shipping     MoneyType `xml:"Shipping"`
	Points       Points    `xml:"Points"`
}

//SellerFeedbackRating Information about the seller's feedback, including the percentage of positive feedback, and the total count of feedback received.
type SellerFeedbackRating struct {
	SellerPositiveFeedbackRating float64 `xml:"SellerPositiveFeedbackRating"`
	FeedbackCount                int64   `xml:"FeedbackCount"`
}

//ShipsFrom The state and country from where the item is shipped.
type ShipsFromType struct {
	State   string `xml:"State"`   //The state from where the item is shipped.
	Country string `xml:"Country"` //The country from where the item is shipped.
}

type SKUOfferDetail struct {
	MyOffer              bool                     `xml:"MyOffer"`
	SubCondition         string                   `xml:"SubCondition"`
	SellerFeedbackRating SellerFeedbackType       `xml:"SellerFeedbackRating"`
	ShippingTime         DetailedShippingTimeType `xml:"ShippingTime"`
	ListingPrice         MoneyType                `xml:"ListingPrice"`
	Points               Points                   `xml:"Points"`
	Shipping             MoneyType                `xml:"Shipping"`
	ShipsFrom            ShipsFromType            `xml:"ShipsFrom"`
	IsFulfilledByAmazon  bool                     `xml:"IsFulfilledByAmazon"`
	IsBuyBoxWinner       bool                     `xml:"IsBuyBoxWinner"`
	IsFeaturedMerchant   bool                     `xml:"IsFeaturedMerchant"`
}

type MarketplaceASINType struct {
	MarketplaceID string `xml:"MarketplaceId"`
	ASIN          string `xml:"ASIN"`
}

type SellerSKUIdentifier struct {
	MarketplaceId string `xml:"MarketplaceId"`
	SellerId      string `xml:"SellerId"`
	SellerSKU     string `xml:"SellerSKU"`
}

type IdentifierType struct {
	MarketplaceASIN *MarketplaceASINType `xml:"MarketplaceASIN"`
	SKUIdentifier   *SellerSKUIdentifier `xml:"SKUIdentifier"`
}

type DimensionType struct {
	Height DecimalWithUnits `xml:"Height"`
	Width  DecimalWithUnits `xml:"Width"`
	Length DecimalWithUnits `xml:"Length"`
	Weight DecimalWithUnits `xml:"Weight"`
}

type DecimalWithUnits struct {
	Unit  string  `xml:"Units,attr"`
	Value float64 `xml:",chardata"`
}

type Image struct {
	URL    string           `xml:"URL"`
	Width  DecimalWithUnits `xml:"Width"`
	Height DecimalWithUnits `xml:"Height"`
}

type PriceType struct {
	LandedPrice  *MoneyType `xml:"LandedPrice"`  //商品的当前价格（包括进行促销的商品）。
	ListingPrice *MoneyType `xml:"ListingPrice"` //ListingPrice + Shipping - Points.请注意，如果未返回到岸价格，则上市价格代表具有最低到岸价格的产品。
	Shipping     *MoneyType `xml:"Shipping"`     //商品的运费。
}

type OfferType struct {
	BuyingPrice        PriceType `xml:"BuyingPrice"`        //包含价格信息（包括进行促销的商品）以及运费。
	RegularPrice       MoneyType `xml:"RegularPrice"`       //商品的当前价格（不包括进行促销的商品）。不包括运费。
	FulfillmentChannel string    `xml:"FulfillmentChannel"` //商品的配送渠道。Amazon - 亚马逊物流。Merchant - 卖家自行配送。
	ItemCondition      string    `xml:"ItemCondition"`      //商品的状况。有效值：New、Used、Collectible、Refurbished、Club
	ItemSubCondition   string    `xml:"ItemSubCondition"`   //商品的子状况(成色)。有效值：New、Mint、Very Good、Good、Acceptable、Poor、Club、OEM、Warranty、Refurbished Warranty、Refurbished、Open Box 或 Other。
	SellerID           string    `xml:"SellerId"`           //在操作中提交的 SellerId。
	SellerSKU          string    `xml:"SellerSKU"`          //商品的 SellerSKU。
}

type SellerFeedbackType struct {
	SellerPositiveFeedbackRating float64 `xml:"SellerPositiveFeedbackRating"`
	FeedbackCount                int     `xml:"FeedbackCount"`
}

type ShippingTimeType struct {
	Max string `xml:"Max"`
}

type QualifiersType struct {
	ItemCondition                string           `xml:"ItemCondition"`
	ItemSubcondition             string           `xml:"ItemSubcondition"`
	FulfillmentChannel           string           `xml:"FulfillmentChannel"`
	ShipsDomestically            string           `xml:"ShipsDomestically"`
	ShippingTime                 ShippingTimeType `xml:"ShippingTime"`
	SellerPositiveFeedbackRating string           `xml:"SellerPositiveFeedbackRating"`
}

type LowestOfferListingType struct {
	Qualifiers                      QualifiersType `xml:"Qualifiers"`
	NumberOfOfferListingsConsidered int            `xml:"NumberOfOfferListingsConsidered"`
	SellerFeedbackCount             int            `xml:"SellerFeedbackCount"`
	Price                           PriceType      `xml:"Price"`
	MultipleOffersAtLowestPrice     string         `xml:"MultipleOffersAtLowestPrice"`
}

type SalesRankType struct {
	ProductCategoryID string `xml:"ProductCategoryId"`
	Rank              int    `xml:"Rank"`
}

type OfferListingCountType struct {
	Text      string `xml:",chardata"`
	Condition string `xml:"condition,attr"`
}

type CompetitivePricingType struct {
	CompetitivePrices struct {
		CompetitivePrice []struct {
			BelongsToRequester string    `xml:"belongsToRequester,attr"`
			Condition          string    `xml:"condition,attr"`
			Subcondition       string    `xml:"subcondition,attr"`
			CompetitivePriceId string    `xml:"CompetitivePriceId"`
			Price              PriceType `xml:"Price"`
		} `xml:"CompetitivePrice"`
	} `xml:"CompetitivePrices"`
	NumberOfOfferListings struct {
		OfferListingCount []OfferListingCountType `xml:"OfferListingCount"`
	} `xml:"NumberOfOfferListings"`
	TradeInValue MoneyType `xml:"TradeInValue"`
}

type VariationParentType struct {
	MarketplaceASIN MarketplaceASINType `xml:"Identifiers>MarketplaceASIN"`
	SKUIdentifier   SellerSKUIdentifier `xml:"Identifiers>SKUIdentifier"`
}

type VariationChildType struct {
	Color                  string           `xml:"Color"`
	Edition                string           `xml:"Edition"`
	Flavor                 string           `xml:"Flavor"`
	GemType                string           `xml:"GemType"`
	GolfClubFlex           string           `xml:"GolfClubFlex"`
	GolfClubLoft           DecimalWithUnits `xml:"GolfClubLoft"`
	HandOrientation        string           `xml:"HandOrientation"`
	HardwarePlatform       string           `xml:"HardwarePlatform"`
	ItemDimensions         DimensionType    `xml:"ItemDimensions"`
	MaterialType           string           `xml:"MaterialType"`
	MetalType              string           `xml:"MetalType"`
	Model                  string           `xml:"Model"`
	OperatingSystem        string           `xml:"OperatingSystem"`
	PackageQuantity        string           `xml:"PackageQuantity"`
	ProductTypeSubcategory string           `xml:"ProductTypeSubcategory"`
	RingSize               string           `xml:"RingSize"`
	ShaftMaterial          string           `xml:"ShaftMaterial"`
	Scent                  string           `xml:"Scent"`
	Size                   string           `xml:"Size"`
	SizePerPearl           string           `xml:"SizePerPearl"`
	TotalDiamondWeight     DecimalWithUnits `xml:"TotalDiamondWeight"`
	TotalGemWeight         DecimalWithUnits `xml:"TotalGemWeight"`
}

type CreatorType struct {
	Text string `xml:",chardata"`
	Role string `xml:"Role,attr"`
}

type LanguageType struct {
	Name        string `xml:"Name"`
	Type        string `xml:"Type"`
	AudioFormat string `xml:"AudioFormat"`
}

//Relationships
type Relationships struct {
	VariationParent VariationParentType  `xml:"VariationParent"`
	VariationChild  []VariationChildType `xml:"VariationChild"`
}

type ProductType struct {
	Identifiers         IdentifierType           `xml:"Identifiers"`
	ItemAttributes      *ItemAttributesType      `xml:"AttributeSets>ItemAttributes"`
	Relationships       *Relationships           `xml:"Relationships"`
	CompetitivePricing  *CompetitivePricingType  `xml:"CompetitivePricing"`
	SalesRankings       []SalesRankType          `xml:"SalesRankings>SalesRank"`
	LowestOfferListings []LowestOfferListingType `xml:"LowestOfferListings>LowestOfferListing"`
	Offers              []OfferType              `xml:"Offers>Offer"`
}

type ItemAttributesType struct {
	Actor                                string           `xml:"Actor"`
	Artist                               string           `xml:"Artist"`
	AspectRatio                          string           `xml:"AspectRatio"`
	AudienceRating                       string           `xml:"AudienceRating"`
	Author                               string           `xml:"Author"`
	BackFinding                          string           `xml:"BackFinding"`
	BandMaterialType                     string           `xml:"BandMaterialType"`
	Binding                              string           `xml:"Binding"`
	BlurayRegion                         string           `xml:"BlurayRegion"`
	Brand                                string           `xml:"Brand"`
	CEROAgeRating                        string           `xml:"CEROAgeRating"`
	ChainType                            string           `xml:"ChainType"`
	ClaspType                            string           `xml:"ClaspType"`
	Color                                string           `xml:"Color"`
	CPUManufacturer                      string           `xml:"CPUManufacturer"`
	CPUSpeed                             DecimalWithUnits `xml:"CPUSpeed"`
	CPUType                              string           `xml:"CPUType"`
	Creator                              CreatorType      `xml:"Creator"`
	Department                           string           `xml:"Department"`
	Director                             string           `xml:"Director"`
	DisplaySize                          DecimalWithUnits `xml:"DisplaySize"`
	Edition                              string           `xml:"Edition"`
	EpisodeSequence                      string           `xml:"EpisodeSequence"`
	ESRBAgeRating                        string           `xml:"ESRBAgeRating"`
	Feature                              string           `xml:"Feature"`
	Flavor                               string           `xml:"Flavor"`
	Format                               string           `xml:"Format"`
	GemType                              string           `xml:"GemType"`
	Genre                                string           `xml:"Genre"`
	GolfClubFlex                         string           `xml:"GolfClubFlex"`
	GolfClubLoft                         DecimalWithUnits `xml:"GolfClubLoft"`
	HandOrientation                      string           `xml:"HandOrientation"`
	HardDiskInterface                    string           `xml:"HardDiskInterface"`
	HardDiskSize                         DecimalWithUnits `xml:"HardDiskSize"`
	HardwarePlatform                     string           `xml:"HardwarePlatform"`
	HazardousMaterialType                string           `xml:"HazardousMaterialType"`
	ItemDimensions                       DimensionType    `xml:"ItemDimensions"`
	IsAdultProduct                       bool             `xml:"IsAdultProduct"`
	IsAutographed                        bool             `xml:"IsAutographed"`
	IsEligibleForTradeIn                 bool             `xml:"IsEligibleForTradeIn"`
	IsMemorabilia                        bool             `xml:"IsMemorabilia"`
	IssuesPerYear                        string           `xml:"IssuesPerYear"`
	ItemPartNumber                       string           `xml:"ItemPartNumber"`
	Label                                string           `xml:"Label"`
	Languages                            []LanguageType   `xml:"Languages>Language"`
	LegalDisclaimer                      string           `xml:"LegalDisclaimer"`
	ListPrice                            MoneyType        `xml:"ListPrice"`
	Manufacturer                         string           `xml:"Manufacturer"`
	ManufacturerMaximumAge               DecimalWithUnits `xml:"ManufacturerMaximumAge"`
	ManufacturerMinimumAge               DecimalWithUnits `xml:"ManufacturerMinimumAge"`
	ManufacturerPartsWarrantyDescription string           `xml:"ManufacturerPartsWarrantyDescription"`
	MaterialType                         string           `xml:"MaterialType"`
	MaximumResolution                    DecimalWithUnits `xml:"MaximumResolution"`
	MediaType                            string           `xml:"MediaType"`
	MetalStamp                           string           `xml:"MetalStamp"`
	MetalType                            string           `xml:"MetalType"`
	Model                                string           `xml:"Model"`
	NumberOfDiscs                        int              `xml:"NumberOfDiscs"`
	NumberOfIssues                       int              `xml:"NumberOfIssues"`
	NumberOfItems                        int              `xml:"NumberOfItems"`
	NumberOfPages                        int              `xml:"NumberOfPages"`
	NumberOfTracks                       int              `xml:"NumberOfTracks"`
	OperatingSystem                      string           `xml:"OperatingSystem"`
	OpticalZoom                          DecimalWithUnits `xml:"OpticalZoom"`
	PackageDimensions                    DimensionType    `xml:"PackageDimensions"`
	PackageQuantity                      int              `xml:"PackageQuantity"`
	PartNumber                           string           `xml:"PartNumber"`
	PegiRating                           string           `xml:"PegiRating"`
	Platform                             string           `xml:"Platform"`
	ProcessorCount                       int              `xml:"ProcessorCount"`
	ProductGroup                         string           `xml:"ProductGroup"`
	ProductTypeName                      string           `xml:"ProductTypeName"`
	ProductTypeSubcategory               string           `xml:"ProductTypeSubcategory"`
	PublicationDate                      string           `xml:"PublicationDate"`
	Publisher                            string           `xml:"Publisher"`
	RegionCode                           string           `xml:"RegionCode"`
	ReleaseDate                          string           `xml:"ReleaseDate"`
	RingSize                             string           `xml:"RingSize"`
	RunningTime                          DecimalWithUnits `xml:"RunningTime"`
	ShaftMaterial                        string           `xml:"ShaftMaterial"`
	Scent                                string           `xml:"Scent"`
	SeasonSequence                       string           `xml:"SeasonSequence"`
	SeikodoProductCode                   string           `xml:"SeikodoProductCode"`
	Size                                 string           `xml:"Size"`
	SizePerPearl                         string           `xml:"SizePerPearl"`
	SmallImage                           Image            `xml:"SmallImage"`
	Studio                               string           `xml:"Studio"`
	SubscriptionLength                   DecimalWithUnits `xml:"SubscriptionLength"`
	SystemMemorySize                     DecimalWithUnits `xml:"SystemMemorySize"`
	SystemMemoryType                     string           `xml:"SystemMemoryType"`
	TheatricalReleaseDate                string           `xml:"TheatricalReleaseDate"`
	Title                                string           `xml:"Title"`
	TotalDiamondWeight                   DecimalWithUnits `xml:"TotalDiamondWeight"`
	TotalGemWeight                       DecimalWithUnits `xml:"TotalGemWeight"`
	Warranty                             string           `xml:"Warranty"`
	WEEETaxValue                         MoneyType        `xml:"WEEETaxValue"`
}

//ListMatchingProductsResult
type ListMatchingProductsResult struct {
	Products []ProductType `xml:"Products>Product"`
}

//GetMatchingProductResult
type GetMatchingProductResult struct {
	ASIN    string      `xml:"ASIN,attr"`
	Status  string      `xml:"status,attr"`
	Product ProductType `xml:"Product"`
}

//GetMatchingProductForIdResult
type GetMatchingProductForIdResult struct {
	ID       string         `xml:"Id,attr"`
	IdType   string         `xml:"IdType,attr"`
	Status   string         `xml:"status,attr"`
	Products []*ProductType `xml:"Products>Product"`
}

//GetCompetitivePricingForSKUResult
type GetCompetitivePricingForSKUResult struct {
	SellerSKU string      `xml:"SellerSKU,attr"`
	Status    string      `xml:"status,attr"`
	Product   ProductType `xml:"Product"`
}

//GetCompetitivePricingForASINResult
type GetCompetitivePricingForASINResult struct {
	ASIN    string      `xml:"ASIN,attr"`
	Status  string      `xml:"status,attr"`
	Product ProductType `xml:"Product"`
}

//GetMyPriceForSKUResult
type GetMyPriceForSKUResult struct {
	SellerSKU string      `xml:"SellerSKU,attr"` //SKU
	Status    string      `xml:"status,attr"`    //状态
	Product   ProductType `xml:"Product"`
}

//GetMyPriceForASINResult
type GetMyPriceForASINResult struct {
	ASIN    string      `xml:"ASIN,attr"`   //ASIN
	Status  string      `xml:"status,attr"` //状态
	Product ProductType `xml:"Product"`
}

//ProductCategory 分类信息
type ProductCategory struct {
	ProductCategoryId   string           `xml:"ProductCategoryId"`   //分类ID
	ProductCategoryName string           `xml:"ProductCategoryName"` //分类名字
	Parent              *ProductCategory `xml:"Parent"`              //父分类
}

//GetProductCategoriesForSKUResult 产品分类结果
type GetProductCategoriesForSKUResult struct {
	Category []ProductCategory `xml:"Self"`
}

//GetProductCategoriesForASINResult
type GetProductCategoriesForASINResult struct {
	Category []ProductCategory `xml:"Self"`
}

//GetLowestPricedOffersForSKUResult
type GetLowestPricedOffersForSKUResult struct {
	MarketplaceID string `xml:"MarketplaceID,attr"`
	SKU           string `xml:"SKU,attr"`
	ItemCondition string `xml:"ItemCondition,attr"`
	Status        string `xml:"status,attr"`
	Identifier    struct {
		MarketplaceId     string `xml:"MarketplaceId"`
		SellerSKU         string `xml:"SellerSKU"`
		ItemCondition     string `xml:"ItemCondition"`
		TimeOfOfferChange string `xml:"TimeOfOfferChange"`
	} `xml:"Identifier"`
	Summary struct {
		TotalOfferCount                 string            `xml:"TotalOfferCount"`
		NumberOfOffers                  []OfferCountType  `xml:"NumberOfOffers>OfferCount"`
		LowestPrices                    []LowestPriceType `xml:"LowestPrices>LowestPrice"`
		BuyBoxPrices                    []BuyBoxPriceType `xml:"BuyBoxPrices>BuyBoxPrice"`
		ListPrice                       MoneyType         `xml:"ListPrice"`
		SuggestedLowerPricePlusShipping MoneyType         `xml:"SuggestedLowerPricePlusShipping"`
		BuyBoxEligibleOffers            []OfferCountType  `xml:"BuyBoxEligibleOffers>OfferCount"`
		OffersAvailableTime             string            `xml:"OffersAvailableTime"`
	} `xml:"Summary"`
	Offers []SKUOfferDetail `xml:"Offers>Offer"`
}
