package mws

import (
	"testing"
)

func TestProductService_ListMatchingProducts(t *testing.T) {
	requestID, result, err := NewProductService().ListMatchingProducts(TestCredential, TestCredential.Region.MarketPlaceID, "B084WZNV33")
	if err != nil {
		t.Errorf("ListMatchingProducts() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("ListMatchingProducts() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("ListMatchingProducts() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetMatchingProduct(t *testing.T) {
	requestID, result, err := NewProductService().GetMatchingProduct(TestCredential, TestCredential.Region.MarketPlaceID, "B084WZNV33", "B0EXAMPLEG")
	if err != nil {
		t.Errorf("GetMatchingProduct() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetMatchingProduct() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetMatchingProduct() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetMatchingProductForId(t *testing.T) {
	requestID, result, err := NewProductService().GetMatchingProductForId(TestCredential, TestCredential.Region.MarketPlaceID, "EAN", "6003551978396", "4003994155486")
	if err != nil {
		t.Errorf("GetMatchingProductForId() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetMatchingProductForId() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetMatchingProductForId() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetCompetitivePricingForSKU(t *testing.T) {
	requestID, result, err := NewProductService().GetCompetitivePricingForSKU(TestCredential, TestCredential.Region.MarketPlaceID, "MG-356J-1QOR", "56789")
	if err != nil {
		t.Errorf("GetCompetitivePricingForSKU() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetCompetitivePricingForSKU() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetCompetitivePricingForSKU() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetCompetitivePricingForASIN(t *testing.T) {
	requestID, result, err := NewProductService().GetCompetitivePricingForASIN(TestCredential, TestCredential.Region.MarketPlaceID, "B084WZNV33", "B0EXAMPLEG")
	if err != nil {
		t.Errorf("GetCompetitivePricingForASIN() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetCompetitivePricingForASIN() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetCompetitivePricingForASIN() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetMyPriceForSKU(t *testing.T) {
	requestID, result, err := NewProductService().GetMyPriceForSKU(TestCredential, TestCredential.Region.MarketPlaceID, "MG-356J-1QOR", "56789")
	if err != nil {
		t.Errorf("GetMyPriceForSKU() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetMyPriceForSKU() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetMyPriceForSKU() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetMyPriceForASIN(t *testing.T) {
	requestID, result, err := NewProductService().GetMyPriceForASIN(TestCredential, TestCredential.Region.MarketPlaceID, "B084WZNV33", "B0EXAMPLEG")
	if err != nil {
		t.Errorf("GetMyPriceForASIN() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetMyPriceForASIN() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetMyPriceForASIN() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetProductCategoriesForSKU(t *testing.T) {
	requestID, result, err := NewProductService().GetProductCategoriesForSKU(TestCredential, TestCredential.Region.MarketPlaceID, "MG-356J-1QOR")
	if err != nil {
		t.Errorf("GetProductCategoriesForSKU() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetProductCategoriesForSKU() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetProductCategoriesForSKU() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetProductCategoriesForASIN(t *testing.T) {
	requestID, result, err := NewProductService().GetProductCategoriesForASIN(TestCredential, TestCredential.Region.MarketPlaceID, "B084WZNV33")
	if err != nil {
		t.Errorf("GetProductCategoriesForASIN() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetProductCategoriesForASIN() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetProductCategoriesForASIN() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetLowestOfferListingsForSKU(t *testing.T) {
	requestID, result, err := NewProductService().GetLowestOfferListingsForSKU(TestCredential, TestCredential.Region.MarketPlaceID, "MG-356J-1QOR")
	if err != nil {
		t.Errorf("GetLowestOfferListingsForSKU() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetLowestOfferListingsForSKU() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetLowestOfferListingsForSKU() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetLowestOfferListingsForASIN(t *testing.T) {
	requestID, result, err := NewProductService().GetLowestOfferListingsForASIN(TestCredential, TestCredential.Region.MarketPlaceID, "B084WZNV33")
	if err != nil {
		t.Errorf("GetLowestOfferListingsForASIN() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetLowestOfferListingsForASIN() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetLowestOfferListingsForASIN() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestProductService_GetLowestPricedOffersForSKU(t *testing.T) {
	requestID, result, err := NewProductService().GetLowestPricedOffersForSKU(TestCredential, TestCredential.Region.MarketPlaceID, "MG-356J-1QOR", "New")
	if err != nil {
		t.Errorf("GetLowestPricedOffersForSKU() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetLowestPricedOffersForSKU() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetLowestPricedOffersForSKU() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}
