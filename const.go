package mws

import "strings"

var VERSION = "0.0.1"

const (
	keySellerID         = "SellerId"
	keyMWSAuthToken     = "MWSAuthToken"
	keyAWSAccessKeyID   = "AWSAccessKeyId"
	keyVersion          = "Version"
	keyAction           = "Action"
	keySignature        = "Signature"
	keySignatureMethod  = "SignatureMethod"
	keySignatureVersion = "SignatureVersion"
	keyTimestamp        = "Timestamp"
)

type Region struct {
	Region        string //区域
	Country       string //国家代码
	CountryName   string //国家名称
	Endpoint      string //API域名
	MarketPlaceID string //市场ID
}

// Regions 亚马逊店铺地区,用来区分开发者账号
var Regions = []Region{
	//北美 🇺🇸
	{"NA", "US", "美国", "mws.amazonservices.com", "ATVPDKIKX0DER"},
	{"NA", "CA", "加拿大", "mws.amazonservices.ca", "A2EUQ1WTGCTBG2"},
	{"NA", "MX", "墨西哥", "mws.amazonservices.com.mx", "A1AM78C64UM0Y8"},
	//巴西 🇧🇷
	{"BR", "BR", "巴西", "mws.amazonservices.com", "A2Q3Y263D00KWC"},
	//欧洲
	{"EU", "UK", "英国", "mws-eu.amazonservices.com", "A1F83G8C2ARO7P"},
	{"EU", "IT", "意大利", "mws-eu.amazonservices.com", "APJ6JRA9NG5V4"},
	{"EU", "DE", "德国", "mws-eu.amazonservices.com", "A1PA6795UKMFR9"},
	{"EU", "FR", "法国", "mws-eu.amazonservices.com", "A13V1IB3VIYZZH"},
	{"EU", "ES", "西班牙", "mws-eu.amazonservices.com", "A1RKKUPIHCS9HS"},
	//阿拉伯联合酋长国 🇦🇪
	{"AE", "AE", "阿联酋", "mws.amazonservices.ae", "A2VIGQ35RCS4UG"},
	//印度 🇮🇳
	{"IN", "IN", "印度", "mws.amazonservices.in", "A21TJRUUN4KGV"},
	//澳大利亚 🇦🇺
	{"AU", "AU", "澳大利亚", "mws.amazonservices.com.au", "A39IBJ37TRP1C6"},
	//日本,新加坡 🇯🇵 🇸🇬
	{"FE", "JP", "日本", "mws.amazonservices.jp", "A1VC38T7YXB528"},
	{"FE", "SG", "新加坡", "mws-fe.amazonservices.com", "A19VAU5U5O7RUS"},
}

func GetRegionByCountry(country string) Region {
	for _, region := range Regions {
		if strings.EqualFold(region.Country, country) {
			return region
		}
	}
	panic("Invalid region, check your data")
}

func GetRegionByRegionCode(regionCode string) Region {
	for _, region := range Regions {
		if strings.EqualFold(region.Region, regionCode) {
			return region
		}
	}
	panic("Invalid region, check your data")
}

func GetRegionByMarketPlaceID(marketPlaceID string) Region {
	for _, region := range Regions {
		if strings.EqualFold(region.MarketPlaceID, marketPlaceID) {
			return region
		}
	}
	panic("Invalid marketPlaceID, check your data")
}

func GetCountryCodeByMarketPlaceID(marketPlaceID string) string {
	for _, region := range Regions {
		if strings.EqualFold(region.MarketPlaceID, marketPlaceID) {
			return region.Country
		}
	}
	panic("Invalid marketPlaceID, check your data")
}

func GetRegionCodeByMarketPlaceID(marketPlaceID string) string {
	for _, region := range Regions {
		if strings.EqualFold(region.MarketPlaceID, marketPlaceID) {
			return region.Region
		}
	}
	panic("Invalid marketPlaceID, check your data")
}

func GetCountryName(countryCode string) string {
	for _, region := range Regions {
		if strings.EqualFold(region.Country, countryCode) {
			return region.CountryName
		}
	}
	panic("Invalid countryCode, check your data")
}

var FeedProcessingStatusDesc = map[string]string{
	"_AWAITING_ASYNCHRONOUS_REPLY_": "请求正在处理中,但是正在等待外部信息才能完成",
	"_CANCELLED_":                   "由于发生致命错误,该请求已被中止",
	"_DONE_":                        "请求已被处理",
	"_IN_PROGRESS_":                 "请求正在处理中",
	"_IN_SAFETY_NET_":               "请求正在处理中,但是系统已确定Feed可能存在错误（例如，该请求将从卖家账户中删除所有库存。）Amazon卖家支持人员将联系卖家以确认Feed是否应该进行处理",
	"_SUBMITTED_":                   "请求已收到,但尚未开始处理",
	"_UNCONFIRMED_":                 "请求待处理",
}

func GetFeedStatusDesc(status string) string {
	if v, ok := FeedProcessingStatusDesc[status]; ok {
		return v
	}
	return status
}
