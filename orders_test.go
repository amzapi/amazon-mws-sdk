package mws

import (
	"testing"
	"time"
)

func TestOrderService_ListOrders(t *testing.T) {
	data := NewValues()
	data.SetTime("CreatedAfter", time.Now().AddDate(0, -1, 0))
	requestID, result, err := NewOrderService().ListOrders(TestCredential, []string{TestCredential.Region.MarketPlaceID}, data)
	if err != nil {
		t.Errorf("ListOrders() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("ListOrders() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("ListOrders() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestOrderService_ListOrderItems(t *testing.T) {
	requestID, result, err := NewOrderService().ListOrderItems(TestCredential, "701-6952015-5478663")
	if err != nil {
		t.Errorf("ListOrderItems() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("ListOrderItems() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("ListOrderItems() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestOrderService_GetOrder(t *testing.T) {
	requestID, result, err := NewOrderService().GetOrder(TestCredential, []string{"701-6952015-5478663"})
	if err != nil {
		t.Errorf("GetOrder() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetOrder() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetOrder() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}
