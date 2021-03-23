package mws

import (
	"bytes"
)

type FeedService struct {
	*AmazonClient
}

func NewFeedService() *FeedService {
	return &FeedService{
		AmazonClient: newAmazonClient("/Feeds/2009-01-01", "2009-01-01"),
	}
}

func (s *FeedService) SubmitFeed(c *Credential, marketplaceIdList []string, feedType FeedType, feedContent []byte, params ...Values) (string, *SubmitFeedResult, error) {
	data := ActionValues("SubmitFeed")
	data.Set("FeedType", string(feedType))
	data.Sets("MarketplaceIdList.Id", marketplaceIdList...)
	data.Set("ContentMD5Value", FeedContentMd5(feedContent))
	data.SetAll(params...)

	//data.Set("PurgeAndReplace", strconv.FormatBool(purgeAndReplace))

	var result struct {
		BaseResponse
		Result *SubmitFeedResult `xml:"SubmitFeedResult"`
	}
	if _, err := s.GetModel("POST", c, data, bytes.NewBuffer(feedContent), &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

func (s *FeedService) GetFeedSubmissionList(c *Credential, params ...Values) (string, *GetFeedSubmissionListResult, error) {
	data := ActionValues("GetFeedSubmissionList")
	data.SetAll(params...)

	//data.Sets("FeedSubmissionIdList.Id", feedSubmissionIdList...)
	//data.SetInt("MaxCount", maxCount)
	//data.Sets("FeedTypeList.Type", feedTypeList...)
	//data.Sets("FeedProcessingStatusList.Status", feedProcessingStatusList...)
	//data.SetTime("SubmittedFromDate", submittedFromDate)
	//data.SetTime("SubmittedToDate", submittedToDate)

	var result struct {
		BaseResponse
		Result *GetFeedSubmissionListResult `xml:"GetFeedSubmissionListResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

func (s *FeedService) GetFeedSubmissionListByNextToken(c *Credential, nextToken string) (string, *GetFeedSubmissionListResult, error) {
	data := ActionValues("GetFeedSubmissionListByNextToken")
	data.Set("NextToken", nextToken)
	var result struct {
		BaseResponse
		Result *GetFeedSubmissionListResult `xml:"GetFeedSubmissionListByNextTokenResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

func (s *FeedService) GetFeedSubmissionCount(c *Credential, params ...Values) (string, int, error) {

	data := ActionValues("GetFeedSubmissionCount")
	data.SetAll(params...)

	//data := NewValues()
	//data.Sets("FeedTypeList.Type", feedTypeList...)
	//data.Sets("FeedProcessingStatusList.Status", feedProcessingStatusList...)
	//data.SetTime("SubmittedFromDate", submittedFromDate)
	//data.SetTime("SubmittedToDate", submittedToDate)

	var result struct {
		BaseResponse
		Count int `xml:"GetFeedSubmissionCountResult>Count"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", 0, err
	}
	return result.RequestID, result.Count, nil
}

func (s *FeedService) GetFeedSubmissionResult(c *Credential, feedSubmissionId string) (*GetFeedSubmissionResult, error) {
	data := ActionValues("GetFeedSubmissionResult")
	data.Set("FeedSubmissionId", feedSubmissionId)
	var result GetFeedSubmissionResult
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
