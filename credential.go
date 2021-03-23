package mws

import "os"

var TestCredential = &Credential{
	AccessKey: "AKIAI3ZG4SY4QFVXVYCA",
	SecretKey: "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",
	SellerID:  "A10LKG3OIMP8CS",
	AuthToken: "amzn.mws.dcb8aa31-07a1-6d36-7455-85cd115694a8",
	Region:    GetRegionByCountry("CA"),
}

//Credential 令牌
type Credential struct {
	AccessKey string //开发者凭证: AWSAccessKeyId
	SecretKey string //开发者凭证: SecretKey
	SellerID  string //卖家ID
	AuthToken string //卖家授权TOKEN
	Region    Region //凭证地区信息
}

//GetCredentialFromEnv 从环境变量获取授权令牌
func GetCredentialFromEnv(sellerID, authToken string) *Credential {
	return &Credential{
		SellerID:  sellerID,
		AuthToken: authToken,
		AccessKey: os.Getenv("AccessKey"),
		SecretKey: os.Getenv("SecretKey"),
		Region:    GetRegionByCountry(os.Getenv("Country")),
	}
}

//GetCredentialForTest 从环境变量获取授权令牌（用于测试）
func GetCredentialForTest() *Credential {
	return GetCredentialFromEnv(os.Getenv("TestSellerId"), os.Getenv("TestAuthToken"))
}
