package mws

type ProductService struct {
	*AmazonClient
}

//Products 创建商品服务
func NewProductService() *ProductService {
	return &ProductService{
		AmazonClient: newAmazonClient("/Products/2011-10-01", "2011-10-01"),
	}
}

//ListMatchingProducts 根据搜索查询返回产品及其属性的列表。
func (s *ProductService) ListMatchingProducts(c *Credential, marketplaceId, query string, params ...Values) (string, *ListMatchingProductsResult, error) {
	data := ActionValues("ListMatchingProducts")
	data.Set("MarketplaceId", marketplaceId)
	data.Set("Query", query)
	//data.Set("QueryContextId",queryContextId)
	data.SetAll(params...)

	var result struct {
		BaseResponse
		Result *ListMatchingProductsResult `xml:"ListMatchingProductsResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetMatchingProduct 根据ASIN值列表返回产品及其属性的列表。
func (s *ProductService) GetMatchingProduct(c *Credential, marketplaceId string, asinList ...string) (string, []*GetMatchingProductResult, error) {
	data := ActionValues("GetMatchingProduct")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("ASINList.ASIN", asinList...)
	var result struct {
		BaseResponse
		Result []*GetMatchingProductResult `xml:"GetMatchingProductResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetMatchingProductForId 根据ASIN，GCID，SellerSKU，UPC，EAN，ISBN和JAN值的列表返回产品及其属性的列表。
func (s *ProductService) GetMatchingProductForId(c *Credential, marketplaceId, idType string, idList ...string) (string, []*GetMatchingProductForIdResult, error) {
	data := ActionValues("GetMatchingProductForId")
	data.Set("MarketplaceId", marketplaceId)
	data.Set("IdType", string(idType))
	data.Sets("IdList.Id", idList...)
	var result struct {
		BaseResponse
		Result []*GetMatchingProductForIdResult `xml:"GetMatchingProductForIdResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetCompetitivePricingForSKU 根据SellerSKU返回产品的当前有竞争力的价格。
func (s *ProductService) GetCompetitivePricingForSKU(c *Credential, marketplaceId string, sellerSKUList ...string) (string, []*GetCompetitivePricingForSKUResult, error) {
	data := ActionValues("GetCompetitivePricingForSKU")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("SellerSKUList.SellerSKU", sellerSKUList...)
	var result struct {
		BaseResponse
		Result []*GetCompetitivePricingForSKUResult `xml:"GetCompetitivePricingForSKUResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetCompetitivePricingForASIN 根据ASIN返回产品的当前竞争价格。
func (s *ProductService) GetCompetitivePricingForASIN(c *Credential, marketplaceId string, asinList ...string) (string, []*GetCompetitivePricingForASINResult, error) {
	data := ActionValues("GetCompetitivePricingForASIN")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("ASINList.ASIN", asinList...)
	var result struct {
		BaseResponse
		Result []*GetCompetitivePricingForASINResult `xml:"GetCompetitivePricingForASINResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetLowestOfferListingsForSKUResult
type GetLowestOfferListingsForSKUResult struct {
	AllOfferListingsConsidered bool        `xml:"AllOfferListingsConsidered"`
	Product                    ProductType `xml:"Product"`
}

//GetLowestOfferListingsForSKU 根据SellerSKU返回最多20种产品的最低价格有效报价清单的定价信息。
func (s *ProductService) GetLowestOfferListingsForSKU(c *Credential, marketplaceId string, sellerSKUList ...string) (string, []*GetLowestOfferListingsForSKUResult, error) {
	data := ActionValues("GetLowestOfferListingsForSKU")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("SellerSKUList.SellerSKU", sellerSKUList...)
	//data.Set("ItemCondition", string(itemCondition))
	var result struct {
		BaseResponse
		Result []*GetLowestOfferListingsForSKUResult `xml:"GetLowestOfferListingsForSKUResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetLowestOfferListingsForASINResult
type GetLowestOfferListingsForASINResult struct {
	AllOfferListingsConsidered bool        `xml:"AllOfferListingsConsidered"`
	Product                    ProductType `xml:"Product"`
}

//GetLowestOfferListingsForASIN 返回基于ASIN的多达20种产品的最低价格有效报价清单的定价信息。
func (s *ProductService) GetLowestOfferListingsForASIN(c *Credential, marketplaceId string, asinList ...string) (string, []*GetLowestOfferListingsForASINResult, error) {
	data := ActionValues("GetLowestOfferListingsForASIN")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("ASINList.ASIN", asinList...)
	//data.Set("ItemCondition", string(itemCondition))
	var result struct {
		BaseResponse
		Result []*GetLowestOfferListingsForASINResult `xml:"GetLowestOfferListingsForASINResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetLowestPricedOffersForSKU TODO:根据SellerSKU返回单个产品的最低报价。 http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestPricedOffersForSKU.html
func (s *ProductService) GetLowestPricedOffersForSKU(c *Credential, marketplaceId, sellerSKU, itemCondition string) (string, *GetLowestPricedOffersForSKUResult, error) {
	data := ActionValues("GetLowestPricedOffersForSKU")
	data.Set("MarketplaceId", marketplaceId)
	data.Set("SellerSKU", sellerSKU)
	data.Set("ItemCondition", itemCondition)
	var result struct {
		BaseResponse
		Result *GetLowestPricedOffersForSKUResult `xml:"GetLowestPricedOffersForSKUResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetLowestPricedOffersForASIN TODO:根据ASIN返回单个产品的最低报价。 http://docs.developer.amazonservices.com/en_US/products/Products_GetLowestPricedOffersForASIN.html
func (s *ProductService) GetLowestPricedOffersForASIN(c *Credential, marketplaceId, asin, itemCondition string) {
	data := ActionValues("GetLowestPricedOffersForASIN")
	data.Set("MarketplaceId", marketplaceId)
	data.Set("ASIN", asin)
	data.Set("ItemCondition", string(itemCondition))
}

//GetMyFeesEstimate TODO:返回产品列表的估计费用。  http://docs.developer.amazonservices.com/en_US/products/Products_GetMyFeesEstimate.html
func (s *ProductService) GetMyFeesEstimate(c *Credential, marketplaceId string) {
	data := ActionValues("GetMyFeesEstimate")
	data.Set("MarketplaceId", marketplaceId)
}

//GetMyPriceForSKU 根据SellerSKU返回您自己的活动商品清单的定价信息。
func (s *ProductService) GetMyPriceForSKU(c *Credential, marketplaceId string, sellerSKUList ...string) (string, []*GetMyPriceForSKUResult, error) {
	data := ActionValues("GetMyPriceForSKU")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("SellerSKUList.SKU", sellerSKUList...)
	//data.Set("ItemCondition", string(itemCondition))
	var result struct {
		BaseResponse
		Result []*GetMyPriceForSKUResult `xml:"GetMyPriceForSKUResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetMyPriceForASIN 根据ASIN返回您自己的活动商品清单的定价信息。
func (s *ProductService) GetMyPriceForASIN(c *Credential, marketplaceId string, asinList ...string) (string, []*GetMyPriceForASINResult, error) {
	data := ActionValues("GetMyPriceForASIN")
	data.Set("MarketplaceId", marketplaceId)
	data.Sets("ASINList.ASIN", asinList...)
	//data.Set("ItemCondition", string(itemCondition))
	var result struct {
		BaseResponse
		Result []*GetMyPriceForASINResult `xml:"GetMyPriceForASINResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetProductCategoriesForSKU 根据SellerSKU返回产品所属的产品类别和父类别。
func (s *ProductService) GetProductCategoriesForSKU(c *Credential, marketplaceId, sellerSKU string) (string, *GetProductCategoriesForSKUResult, error) {
	data := ActionValues("GetProductCategoriesForSKU")
	data.Set("MarketplaceId", marketplaceId)
	data.Set("SellerSKU", sellerSKU)
	var result struct {
		BaseResponse
		Result *GetProductCategoriesForSKUResult `xml:"GetProductCategoriesForSKUResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//GetProductCategoriesForASIN 根据ASIN返回产品所属的产品类别和父类别。
func (s *ProductService) GetProductCategoriesForASIN(c *Credential, marketplaceId, asin string) (string, *GetProductCategoriesForASINResult, error) {
	data := ActionValues("GetProductCategoriesForASIN")
	data.Set("MarketplaceId", marketplaceId)
	data.Set("ASIN", asin)
	var result struct {
		BaseResponse
		Result *GetProductCategoriesForASINResult `xml:"GetProductCategoriesForASINResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}
