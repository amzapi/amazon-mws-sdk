package mws

type BaseResponse struct {
	RequestID string `xml:"ResponseMetadata>RequestId"`
}

//MoneyType An amount of money in a specified currency.
type MoneyType struct {
	CurrencyCode string  `xml:"CurrencyCode"` //The total value.
	Amount       float64 `xml:"Amount"`       //The currency code in ISO 4217 format .One of the following:USD,EUR,GBP,RMB,INR,JPY,CAD,MXN
}

type GetServiceStatusResult struct {
	Status    string `xml:"Status"`
	Timestamp string `xml:"Timestamp"`
}
