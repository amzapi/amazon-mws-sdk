package mws

import (
	"time"
)

type ReportService struct {
	*AmazonClient
}

func NewReportService() *ReportService {
	return &ReportService{
		AmazonClient: newAmazonClient("/Reports/2009-01-01", "2009-01-01"),
	}
}

//ReportRequestListResult 获取报告请求列表结果
type ReportRequestListResult struct {
	NextToken         string
	HasNext           bool
	ReportRequestInfo []*ReportRequestInfo
}

//ReportRequestInfo 获取报告请求列表结果
type ReportRequestInfo struct {
	ReportRequestID        string `xml:"ReportRequestId"`
	ReportType             string
	StartDate              time.Time
	EndDate                time.Time
	Scheduled              bool
	SubmittedDate          time.Time
	ReportProcessingStatus string
	GeneratedReportID      string `xml:"GeneratedReportId"`
	StartedProcessingDate  time.Time
	CompletedDate          time.Time
}

//ReportListResult 获取报告列表的结果模型
type ReportListResult struct {
	HasNext   bool
	NextToken string
	Reports   []*ReportInfo `xml:"ReportInfo"`
}

//ReportInfo 报告信息在结果列表中的模型
type ReportInfo struct {
	ReportID        string `xml:"ReportId"`
	ReportRequestID string `xml:"ReportRequestId"`
	ReportType      string
	Acknowledged    bool
	AvailableDate   *time.Time
}

func (s *ReportService) RequestReport(c *Credential, reportType string, params ...Values) (string, *ReportRequestInfo, error) {
	data := ActionValues("RequestReport")
	data.Set("ReportType", reportType)
	data.SetAll(params...)

	// data.Set("ReportOptions", reportOptions)
	// data.SetTime("StartDate", startDate)
	// data.SetTime("EndDate", endDate)
	// data.Sets("MarketplaceIdList.Id", marketplaceIDList...)

	var response struct {
		BaseResponse
		ReportRequestInfo *ReportRequestInfo `xml:"RequestReportResult>ReportRequestInfo"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestInfo, nil
}

func (s *ReportService) GetReportRequestList(c *Credential, params ...Values) (string, *ReportRequestListResult, error) {
	data := ActionValues("GetReportRequestList")
	data.SetAll(params...)

	// data.Sets("ReportRequestIdList.Id", reportRequestIdList...)
	// data.Sets("ReportTypeList.Type", reportTypeList...)
	// data.Sets("ReportProcessingStatusList.Status", reportProcessingStatusList...)
	// data.Set("MaxCount", maxCount)

	// data.SetTime("RequestedFromDate", startDate)
	// data.SetTime("RequestedToDate", endDate)

	var response struct {
		BaseResponse
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestListResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestList, nil
}

func (s *ReportService) GetReportRequestListByNextToken(c *Credential, nextToken string) (string, *ReportRequestListResult, error) {
	data := ActionValues("GetReportRequestListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		ReportRequestList *ReportRequestListResult `xml:"GetReportRequestLisByNextTokentResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportRequestList, nil
}

func (s *ReportService) GetReport(c *Credential, ReportID string) ([]byte, error) {
	data := ActionValues("GetReport")
	data.Set("ReportId", ReportID)
	return s.Request("POST", c, data, nil)
}

func (s *ReportService) GetReportList(c *Credential, params ...Values) (string, *ReportListResult, error) {
	data := ActionValues("GetReportList")
	data.SetAll(params...)

	// data.SetInt("MaxCount", maxCount)
	// data.SetTime("AvailableFromDate", availableFromDate)
	// data.SetTime("AvailableToDate", availableToDate)
	// data.SetBool("Acknowledged", acknowledged)
	// data.Sets("ReportTypeList.Type", reportTypeList...)
	// data.Sets("ReportRequestIdList.Id", reportRequestIDList...)

	var response struct {
		BaseResponse
		ReportList *ReportListResult `xml:"GetReportListResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportList, nil
}

func (s *ReportService) GetReportListByNextToken(c *Credential, nextToken string) (string, *ReportListResult, error) {
	data := ActionValues("GetReportListByNextToken")
	data.Set("NextToken", nextToken)
	var response struct {
		BaseResponse
		ReportList *ReportListResult `xml:"GetReportListByNextTokenResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &response); err != nil {
		return "", nil, err
	}
	return response.RequestID, response.ReportList, nil
}
