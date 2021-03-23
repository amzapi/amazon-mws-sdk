package mws

import "time"

type OrderService struct {
	*AmazonClient
}

func NewOrderService() *OrderService {
	return &OrderService{
		AmazonClient: newAmazonClient("/Orders/2013-09-01", "2013-09-01"),
	}
}

type ListOrdersResult struct {
	NextToken         string
	CreatedBefore     time.Time
	LastUpdatedBefore time.Time
	Orders            []*Order `xml:"Orders>Order"`
}

type Order struct {
	AmazonOrderID                string                        `xml:"AmazonOrderId"`                //亚马逊所定义的订单编码，格式为 3-7-7。
	SellerOrderID                string                        `xml:"SellerOrderId"`                //卖家所定义的订单编码。
	PurchaseDate                 *time.Time                    `xml:"PurchaseDate"`                 //创建订单的日期。
	LastUpdateDate               *time.Time                    `xml:"LastUpdateDate"`               //订单的最后更新日期
	OrderStatus                  string                        `xml:"OrderStatus"`                  //当前的订单状态。
	FulfillmentChannel           string                        `xml:"FulfillmentChannel"`           //订单配送方式：亚马逊配送 (AFN) 或卖家自行配送 (MFN)。
	SalesChannel                 string                        `xml:"SalesChannel"`                 //订单中第一件商品的销售渠道。
	OrderChannel                 string                        `xml:"OrderChannel"`                 //订单中第一件商品的订单渠道。
	ShipServiceLevel             string                        `xml:"ShipServiceLevel"`             //货件服务水平。
	ShippingAddress              Address                       `xml:"ShippingAddress"`              //订单的配送地址。
	OrderTotal                   MoneyType                     `xml:"OrderTotal"`                   //订单的总费用。
	NumberOfItemsShipped         int                           `xml:"NumberOfItemsShipped"`         //已配送的商品数量。
	NumberOfItemsUnShipped       int                           `xml:"NumberOfItemsUnshipped"`       //未配送的商品数量。
	PaymentExecutionDetail       []*PaymentExecutionDetailItem `xml:"PaymentExecutionDetail"`       //
	PaymentMethod                string                        `xml:"PaymentMethod"`                //订单的主要付款方式。COD - 货到付款。仅适用于中国 (CN) 和日本 (JP)。CVS - 便利店。仅适用于日本 (JP)。Other - COD 和 CVS 之外的付款方式。注： 可使用多种次级付款方式为 PaymentMethod = COD的订单付款。每种次级付款方式均表示为 PaymentExecutionDetailItem 对象。
	IsReplacementOrder           bool                          `xml:"IsReplacementOrder"`           //true if this is a replacement order.
	ReplacedOrderId              string                        `xml:"ReplacedOrderId"`              //The AmazonOrderId value for the order that is being replaced.
	MarketplaceID                string                        `xml:"MarketplaceId"`                //订单生成所在商城的匿名编码。
	BuyerEmail                   string                        `xml:"BuyerEmail"`                   //买家的匿名电子邮件地址。
	BuyerName                    string                        `xml:"BuyerName"`                    //买家姓名。
	BuyerCounty                  string                        `xml:"BuyerCounty"`                  //This element is used only in the Brazil marketplace.
	BuyerTaxInfo                 BuyerTaxInfo                  `xml:"BuyerTaxInfo"`                 //
	ShipmentServiceLevelCategory string                        `xml:"ShipmentServiceLevelCategory"` //订单的配送服务级别分类。 Expedited, FreeEconomy, NextDay, SameDay, SecondDay, Scheduled, Standard
	EasyShipShipmentStatus       string                        `xml:"EasyShipShipmentStatus"`       //
	OrderType                    string                        `xml:"OrderType"`                    //订单类型。StandardOrder - 包含当前有库存商品的订单。Preorder -所含预售商品（发布日期晚于当前日期）的订单。 注： Preorder 仅在日本 (JP) 是可行的OrderType 值。
	EarliestShipDate             *time.Time                    `xml:"EarliestShipDate"`             //您承诺的订单发货时间范围的第一天。日期格式为 ISO 8601。 仅对卖家配送网络 (MFN) 订单返回。
	LatestShipDate               *time.Time                    `xml:"LatestShipDate"`               //您承诺的订单发货时间范围的最后一天。日期格式为 ISO 8601。对卖家配送网络 (MFN)	和亚马逊物流 (AFN) 订单返回。
	EarliestDeliveryDate         *time.Time                    `xml:"EarliestDeliveryDate"`         //您承诺的订单送达时间范围的第一天。日期格式为 ISO 8601。仅对没有 PendingAvailability、Pending 或 Canceled状态的 MFN 订单返回。
	LatestDeliveryDate           *time.Time                    `xml:"LatestDeliveryDate"`           //您承诺的订单送达时间范围的最后一天。日期格式为 ISO 8601。仅对没有 PendingAvailability、Pending 或 Canceled状态的 MFN 订单返回。
	IsBusinessOrder              bool                          `xml:"IsBusinessOrder"`              //true if the order is an Amazon Business order. An Amazon Business order is an order where the buyer is a Verified Business Buyer and the seller is an Amazon Business Seller. For more information about the Amazon Business Seller Program
	IsSoldByAB                   bool                          `xml:"IsSoldByAB"`                   //
	PurchaseOrderNumber          string                        `xml:"PurchaseOrderNumber"`          //	he purchase order (PO) number entered by the buyer at checkout.
	IsPrime                      bool                          `xml:"IsPrime"`                      //true if the order is a seller-fulfilled Amazon Prime order.
	IsPremiumOrder               bool                          `xml:"IsPremiumOrder"`               //true if the order has a Premium Shipping Service Level Agreement. For more information about Premium Shipping orders, see "Premium Shipping Options" in the Seller Central Help for your marketplace.
	IsGlobalExpressEnabled       bool                          `xml:"IsGlobalExpressEnabled"`       //是否启用了全球快递
	PromiseResponseDueDate       *time.Time                    `xml:"PromiseResponseDueDate"`       //
	IsEstimatedShipDateSet       bool                          `xml:"IsEstimatedShipDateSet"`       //
}

//Address  地址信息
type Address struct {
	Name          string `xml:"Name"`          //名称。
	AddressLine1  string `xml:"AddressLine1"`  //街道地址。
	AddressLine2  string `xml:"AddressLine2"`  //其他街道地址信息（如果需要）。
	AddressLine3  string `xml:"AddressLine3"`  //其他街道地址信息（如果需要）。
	City          string `xml:"City"`          //城市。
	Municipality  string `xml:"Municipality"`  //自治市。
	County        string `xml:"County"`        //县。
	District      string `xml:"District"`      //区。
	StateOrRegion string `xml:"StateOrRegion"` //州或地区。
	PostalCode    string `xml:"PostalCode"`    //邮政编码。
	CountryCode   string `xml:"CountryCode"`   //两位数国家/地区代码。格式为 ISO 3166-1-alpha 2 。
	Phone         string `xml:"Phone"`         //电话号码。未退还亚马逊物流订单（FBA）。
	AddressType   string `xml:"AddressType"`   //指示地址是商业地址还是住宅地址。 此元素仅在美国市场中使用。AddressType values: Commercial, Residential
}

//PaymentExecutionDetailItem  付款信息
type PaymentExecutionDetailItem struct {
	Payment       MoneyType `xml:"Payment"`
	PaymentMethod string    `xml:"PaymentMethod"`
}

type PaymentMethodDetail struct {
	PaymentMethodDetail string `xml:"PaymentMethodDetail"`
}

type BuyerTaxInfo struct {
	CompanyLegalName   string `xml:"CompanyLegalName"`
	TaxingRegion       string `xml:"TaxingRegion"`
	TaxClassifications struct {
		TaxClassification TaxClassification `xml:"TaxClassification"`
	} `xml:"TaxClassifications"`
}

type TaxClassification struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

// http://docs.developer.amazonservices.com/en_US/orders-2013-09-01/Orders_ListOrders.html#RequestParameters
func (s *OrderService) ListOrders(c *Credential, marketplaceIdList []string, params ...Values) (string, *ListOrdersResult, error) {
	data := ActionValues("ListOrders")
	data.Sets("MarketplaceId.Id", marketplaceIdList...)
	data.SetAll(params...)

	//data.SetTime("LastUpdatedAfter", lastUpdatedAfter)
	//data.SetTime("LastUpdatedBefore", lastUpdatedBefore)
	//data.SetTime("CreatedAfter", createdAfter)
	//data.SetTime("CreatedBefore", createdBefore)
	//data.Sets("OrderStatus.Status",orderStatus)
	//data.Set("FulfillmentChannel", fulfillmentChannel)  //AFN,MFN,Default: All
	//data.Set("PaymentMethod",paymentMethod)
	//data.Set("BuyerEmail",buyerEmail)
	//data.Set("SellerOrderId",sellerOrderId)
	//data.SetInt("MaxResultsPerPage", maxPerPage)

	var result struct {
		BaseResponse
		Result *ListOrdersResult `xml:"ListOrdersResult"`
	}

	if _, err := s.GetModel("GET", c, data, nil, &result); err != nil {
		return "", nil, err
	}

	return result.RequestID, result.Result, nil
}

func (s *OrderService) ListOrdersByNextToken(c *Credential, nextToken string) (string, *ListOrdersResult, error) {
	data := ActionValues("ListOrdersByNextToken")
	data.Set("NextToken", nextToken)

	var result struct {
		BaseResponse
		Result *ListOrdersResult `xml:"ListOrdersByNextTokenResult"`
	}

	if _, err := s.GetModel("GET", c, data, nil, &result); err != nil {
		return "", nil, err
	}

	return result.RequestID, result.Result, nil
}

type ListOrderItemsResult struct {
	NextToken     string
	AmazonOrderID string       `xml:"AmazonOrderId"`
	OrderItems    []*OrderItem `xml:"OrderItems>OrderItem"`
}

type OrderItem struct {
	ASIN                          string     `xml:"ASIN"`                               //商品的亚马逊标准识别号 (ASIN)。
	OrderItemID                   string     `xml:"OrderItemId"`                        //亚马逊定义的订单商品识别号。
	SellerSKU                     string     `xml:"SellerSKU"`                          //商品的卖家 SKU。
	BuyerCustomizedURL            string     `xml:"BuyerCustomizedInfo->CustomizedURL"` //
	Title                         string     `xml:"Title"`                              //商品名称
	QuantityOrdered               int        `xml:"QuantityOrdered"`                    //下单的商品数量
	QuantityShipped               int        `xml:"QuantityShipped"`                    //已配送的商品数量。
	PointsNumber                  int        `xml:"PointsGranted->PointsNumber"`        //
	PointsMonetaryValue           *MoneyType `xml:"PointsGranted->PointsMonetaryValue"` //
	ProductNumberOfItems          int        `xml:"ProductInfo>NumberOfItems"`          //一件，包含是数量
	ItemPrice                     *MoneyType `xml:"ItemPrice"`                          //订单商品的总价。注意，订单商品指的商品和数量。这意味着，ItemPrice 的价值等于商品售价乘以订购数量。请注意： ItemPrice 不包括ShippingPrice 和GiftWrapPrice
	ShippingPrice                 *MoneyType `xml:"ShippingPrice"`                      //运费
	GiftWrapPrice                 *MoneyType `xml:"GiftWrapPrice"`                      //商品的礼品包装金额。
	TaxCollectionModel            string     `xml:"TaxCollection>Model"`                //TaxCollection Model
	TaxCollectionResponsibleParty string     `xml:"TaxCollection>ResponsibleParty"`     //TaxCollection ResponsibleParty
	ItemTax                       *MoneyType `xml:"ItemTax"`                            //商品价格税。
	ShippingTax                   *MoneyType `xml:"ShippingTax"`                        //运费税。
	GiftWrapTax                   *MoneyType `xml:"GiftWrapTax"`                        //礼品包装价格的税金。
	ShippingDiscount              *MoneyType `xml:"ShippingDiscount"`                   //运费折扣。
	ShippingDiscountTax           *MoneyType `xml:"ShippingDiscountTax"`                //运费折价税。
	PromotionDiscount             *MoneyType `xml:"PromotionDiscount"`                  //优惠中所有促销折扣的总和。
	PromotionDiscountTax          *MoneyType `xml:"PromotionDiscountTax"`               //优惠中所有促销折扣总额的税。
	PromotionIDs                  []string   `xml:"PromotionIds"`                       //PromotionId 元素列表。
	CODFee                        *MoneyType `xml:"CODFee"`                             //COD 服务费用。 注： CODFee 是仅在日本 (JP) 使用的响应元素。
	CODFeeDiscount                *MoneyType `xml:"CODFeeDiscount"`                     //货到付款费用的折扣。注： CODFeeDiscount 是仅在日本 (JP) 使用的响应元素。
	IsGift                        bool       `xml:"IsGift"`                             //买家提供的礼品消息。
	GiftMessageText               string     `xml:"GiftMessageText"`                    //买家提供的礼品消息。
	GiftWrapLevel                 string     `xml:"GiftWrapLevel"`                      //买家指定的礼品包装等级。
	ConditionNote                 string     `xml:"ConditionNote"`                      //卖家描述的商品状况。
	ConditionID                   string     `xml:"ConditionId"`                        //商品的状况。New, Used, Collectible, Refurbished, Preorder, Club
	ConditionSubtypeID            string     `xml:"ConditionSubtypeId"`                 //商品的子状况。New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, Any, Other
	ScheduledDeliveryStartDate    *time.Time `xml:"ScheduledDeliveryStartDate"`         //订单预约送货上门的开始日期（目的地时区）。日期格式为 ISO 8601。
	ScheduledDeliveryEndDate      *time.Time `xml:"ScheduledDeliveryEndDate"`           //订单预约送货上门的终止日期（目的地时区）。日期格式为 ISO 8601 注： 预约送货上门仅适用于日本 (JP)。
	PriceDesignation              string     `xml:"PriceDesignation"`                   //价格指定 BusinessPrice-仅适用于Amazon Business订单的特殊价格。
	IsTransparency                bool       `xml:"IsTransparency"`                     //如果需要透明代码，则为true。
	SerialNumberRequired          bool       `xml:"SerialNumberRequired"`               //如果该产品的产品类型具有序列号，则为true。
}

func (s *OrderService) ListOrderItems(c *Credential, amazonOrderID string) (requestID string, orderItemsResult *ListOrderItemsResult, err error) {
	data := ActionValues("ListOrderItems")
	data.Set("AmazonOrderId", amazonOrderID)

	var result struct {
		BaseResponse
		Result *ListOrderItemsResult `xml:"ListOrderItemsResult"`
	}

	if _, err := s.GetModel("GET", c, data, nil, &result); err != nil {
		return "", nil, err
	}

	return result.RequestID, result.Result, nil
}

func (s *OrderService) ListOrderItemsByNextToken(c *Credential, nextToken string) (string, *ListOrderItemsResult, error) {
	data := ActionValues("ListOrderItemsByNextToken")
	data.Set("NextToken", nextToken)

	var result struct {
		BaseResponse
		Result *ListOrderItemsResult `xml:"ListOrderItemsByNextTokenResult"`
	}

	if _, err := s.GetModel("GET", c, data, nil, &result); err != nil {
		return "", nil, err
	}

	return result.RequestID, result.Result, nil
}

func (s *OrderService) GetOrder(c *Credential, amazonOrderIDs []string) (string, *ListOrdersResult, error) {
	data := ActionValues("GetOrder")
	data.Sets("AmazonOrderId.Id", amazonOrderIDs...)

	var result struct {
		BaseResponse
		Result *ListOrdersResult `xml:"GetOrderResult"`
	}

	if _, err := s.GetModel("GET", c, data, nil, &result); err != nil {
		return "", nil, err
	}

	return result.RequestID, result.Result, nil
}
