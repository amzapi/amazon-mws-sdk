package mws

import (
	"strconv"
	"testing"
)

func TestFeedService_SubmitFeed(t *testing.T) {
	xmlBody := []byte(`<?xml version="1.0" encoding="iso-8859-1"?>
<AmazonEnvelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:noNamespaceSchemaLocation="amzn-envelope.xsd">
  <Header>
    <DocumentVersion>1.01</DocumentVersion>
    <MerchantIdentifier>M_EXAMPLE_123456</MerchantIdentifier>
  </Header>
  <MessageType>Product</MessageType>
  <PurgeAndReplace>false</PurgeAndReplace>
  <Message>
    <MessageID>1</MessageID>
    <OperationType>Update</OperationType>
    <Product>
      <SKU>56789</SKU>
      <StandardProductID>
        <Type>ASIN</Type>
        <Value>B0EXAMPLEG</Value>
      </StandardProductID>
      <ProductTaxCode>A_GEN_NOTAX</ProductTaxCode>
      <DescriptionData>
        <Title>Example Product Title</Title>
        <Brand>Example Product Brand</Brand>
        <Description>This is an example product description.</Description>
        <BulletPoint>Example Bullet Point 1</BulletPoint>
        <BulletPoint>Example Bullet Point 2</BulletPoint>
        <MSRP currency="USD">25.19</MSRP>
        <Manufacturer>Example Product Manufacturer</Manufacturer>
        <ItemType>example-item-type</ItemType>
      </DescriptionData>
      <ProductData>
        <Health>
          <ProductType>
            <HealthMisc>
              <Ingredients>Example Ingredients</Ingredients>
              <Directions>Example Directions</Directions>
            </HealthMisc>
          </ProductType>
        </Health>
      </ProductData>
    </Product>
  </Message>
</AmazonEnvelope>`)

	data := NewValues()
	data.Set("PurgeAndReplace", strconv.FormatBool(false))
	requestID, result, err := NewFeedService().SubmitFeed(TestCredential, []string{TestCredential.Region.MarketPlaceID}, "_POST_PRODUCT_DATA_", xmlBody, data)
	if err != nil {
		t.Errorf("SubmitFeed() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("SubmitFeed() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("SubmitFeed() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestFeedService_GetFeedSubmissionList(t *testing.T) {
	data := NewValues()
	data.Sets("FeedTypeList.Type", []string{"_POST_PRODUCT_DATA_"}...)
	data.Sets("FeedSubmissionIdList.Id", []string{"85984018455"}...)
	requestID, result, err := NewFeedService().GetFeedSubmissionList(TestCredential, data)
	if err != nil {
		t.Errorf("GetFeedSubmissionList() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetFeedSubmissionList() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetFeedSubmissionList() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestFeedService_GetFeedSubmissionResult(t *testing.T) {
	result, err := NewFeedService().GetFeedSubmissionResult(TestCredential, "85984018455")
	if err != nil {
		t.Errorf("GetFeedSubmissionResult() error = %v", err)
		return
	}
	if result == nil {
		t.Errorf("GetFeedSubmissionResult() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestFeedService_GetFeedSubmissionCount(t *testing.T) {
	requestID, count, err := NewFeedService().GetFeedSubmissionCount(TestCredential)
	if err != nil {
		t.Errorf("GetFeedSubmissionCount() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetFeedSubmissionCount() requestID = %v", requestID)
	}
	t.Logf("count = %v", count)
}
