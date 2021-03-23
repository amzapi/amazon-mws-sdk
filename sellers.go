package mws

type SellerService struct {
	*AmazonClient
}

func NewSellerService() *SellerService {
	return &SellerService{
		AmazonClient: newAmazonClient("/Sellers/2011-07-01", "2011-07-01"),
	}
}

//MarketplaceParticipationsResult
type MarketplaceParticipationsResult struct {
	NextToken      string
	Participations []*Participation `xml:"ListParticipations>Participation"`
	Marketplaces   []*Marketplace   `xml:"ListMarketplaces>Marketplace"`
}

//Participation
type Participation struct {
	MarketplaceID              string `xml:"MarketplaceId"`
	SellerID                   string `xml:"SellerId"`
	HasSellerSuspendedListings string `xml:"HasSellerSuspendedListings"`
}

//Marketplace
type Marketplace struct {
	MarketplaceID       string `xml:"MarketplaceId"`
	Name                string
	DefaultCountryCode  string
	DefaultCurrencyCode string
	DefaultLanguageCode string
	DomainName          string
}

//ListMarketplaceParticipations
func (s *SellerService) ListMarketplaceParticipations(c *Credential) (string, *MarketplaceParticipationsResult, error) {
	data := ActionValues("ListMarketplaceParticipations")
	var result struct {
		BaseResponse
		Result *MarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}

//ListMarketplaceParticipationsByNextToken
func (s *SellerService) ListMarketplaceParticipationsByNextToken(c *Credential, nextToken string) (string, *MarketplaceParticipationsResult, error) {
	data := ActionValues("ListMarketplaceParticipationsByNextToken")
	data.Set("NextToken", nextToken)
	var result struct {
		BaseResponse
		Result *MarketplaceParticipationsResult `xml:"ListMarketplaceParticipationsByNextTokenResult"`
	}
	if _, err := s.GetModel("POST", c, data, nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}
