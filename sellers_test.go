package mws

import (
	"testing"
)

func TestSellerService_ListMarketplaceParticipations(t *testing.T) {
	requestID, result, err := NewSellerService().ListMarketplaceParticipations(TestCredential)
	if err != nil {
		t.Errorf("ListMarketplaceParticipations() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("ListMarketplaceParticipations() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("ListMarketplaceParticipations() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}
