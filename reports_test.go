package mws

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReportService_RequestReport(t *testing.T) {
	data := NewValues()
	data.Set("ReportOptions", "MarketplaceId="+TestCredential.Region.MarketPlaceID)
	requestID, result, err := NewReportService().RequestReport(TestCredential, "_GET_XML_BROWSE_TREE_DATA_", data)
	if err != nil {
		t.Errorf("RequestReport() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("RequestReport() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("RequestReport() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestReportService_GetReportRequestList(t *testing.T) {
	data := NewValues()
	// data.Sets("ReportRequestIdList.Id", reportRequestIdList...)
	data.Sets("ReportTypeList.Type", []string{"_GET_XML_BROWSE_TREE_DATA_"}...)
	// data.Sets("ReportProcessingStatusList.Status", reportProcessingStatusList...)
	// data.Set("MaxCount", maxCount)
	// data.SetTime("RequestedFromDate", startDate)
	// data.SetTime("RequestedToDate", endDate)
	requestID, result, err := NewReportService().GetReportRequestList(TestCredential, data)
	if err != nil {
		t.Errorf("GetReportRequestList() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetReportRequestList() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetReportRequestList() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestReportService_GetReportList(t *testing.T) {
	data := NewValues()
	// data.Sets("ReportRequestIdList.Id", reportRequestIdList...)
	data.Sets("ReportTypeList.Type", []string{"_GET_XML_BROWSE_TREE_DATA_"}...)
	requestID, result, err := NewReportService().GetReportList(TestCredential, data)
	if err != nil {
		t.Errorf("GetReportList() error = %v", err)
		return
	}
	if requestID == "" {
		t.Errorf("GetReportList() requestID = %v", requestID)
	}
	if result == nil {
		t.Errorf("GetReportList() result = %v", result)
	}
	t.Logf("result = %v", JsonMarshalIndentToString(result))
}

func TestReportService_GetReport(t *testing.T) {

	fileBytes, err := NewReportService().GetReport(TestCredential, "21778784883018455")

	if err != nil {
		t.Errorf("GetReport() error = %v", err)
		return
	}

	if fileBytes == nil {
		t.Errorf("GetReport() result = %v", fileBytes)
	}

	ioutil.WriteFile("BROWSE_TREE_DATA.xml", fileBytes, os.ModePerm)

	t.Logf("data length = %v", len(fileBytes))

	fileBytes1, err := NewReportService().GetReport(TestCredential, "23038325575018514")

	if err != nil {
		t.Errorf("GetReport() error = %v", err)
		return
	}

	if fileBytes1 == nil {
		t.Errorf("GetReport() result = %v", fileBytes1)
	}

	ioutil.WriteFile("_GET_MERCHANT_LISTINGS_DATA_.xml", fileBytes1, os.ModePerm)

	t.Logf("data length = %v", len(fileBytes1))
}
