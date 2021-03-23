package mws

import (
	"fmt"
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

//errCode 亚马逊错误代码翻译中文详细信息
var errCode = map[string]string{
	"InputStreamDisconnected":                 "读取输入流时出错。",
	"InvalidParameterValue":                   "使用了无效的参数值，或者请求大小超过了可接受的最大大小，或者请求过期。",
	"AccessDenied":                            "访问被拒绝。",
	"InvalidAccessKeyId":                      "使用了无效的AWSAccessKeyId值。",
	"SignatureDoesNotMatch":                   "使用的签名与服务器的计算出的签名值不匹配。",
	"InvalidAddress":                          "使用了无效的API部分或操作值，或者使用了无效的路径。",
	"InternalError":                           "内部服务出现故障。",
	"QuotaExceeded":                           "超过一个小时的请求总数。",
	"RequestThrottled":                        "请求的频率大于允许的频率。",
	"AccessToFeedProcessingResultDenied":      "没有足够权限访问上传数据处理结果。",
	"AccessToReportDenied":                    "没有足够权限访问所请求的报告。",
	"ContentMD5Missing":                       "缺少 Content-MD5 标头值。",
	"ContentMD5DoesNotMatch":                  "计算出的 MD5 哈希值与所提供的 Content-MD5 值不一致。",
	"FeedCanceled":                            "当请求已取消的上传数据的处理报告时返回。",
	"FeedProcessingResultNoLongerAvailable":   "无法下载上传数据处理结果。",
	"FeedProcessingResultNotReady":            "处理报告尚未生成。",
	"InputDataError":                          "上传数据内容包含错误。",
	"InvalidMarketplace":                      "提供的 Marketplace Id 请求参数无效或过期",
	"InvalidFeedSubmissionId":                 "提供的上传数据 Submission Id 无效。",
	"InvalidFeedType":                         "所提交的 Feed Type 无效。",
	"InvalidQueryParameter":                   "提交了多余的参数。",
	"InvalidReportId":                         "提供的 Report Id 无效。",
	"InvalidReportType":                       "所提交的 Report Type 无效。",
	"InvalidRequest":                          "请求中由于缺少参数或参数无效，导致请求无法解析。",
	"InvalidScheduleFrequency":                "所提交的计划频率无效。",
	"MissingClientTokenId":                    "缺少 MerchantModel Id 参数或为空。",
	"MissingParameter":                        "查询中缺少必需的参数。",
	"ReportNoLongerAvailable":                 "无法下载指定的报告。",
	"ReportNotReady":                          "报告尚未生成。",
	"UserAgentHeaderLanguageAttributeMissing": "缺少 User-Agent 标头的 Language 属性。",
	"UserAgentHeaderMalformed":                "User-Agent 值不符合所需格式。",
	"UserAgentHeaderMaximumLengthExceeded":    "User-Agent 值超过 500 个字符。",
	"UserAgentHeaderMissing":                  "缺少 User-Agent 标头值。",
}

//ErrorResponse 亚马逊API错误响应
type ErrorResponse struct {
	RequestID string `xml:"RequestID"`
	Errors    []struct {
		Type    string `xml:"Type"`
		Code    string `xml:"Code"`
		Message string `xml:"Message"`
		Detail  string `xml:"Detail"`
	} `xml:"Error"`
	//记录一些请求的参数和返回信息
	Fields Fields `json:"-" xml:"-"`
}

//Error 亚马逊错误信息
func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("ErrorResponse,ErrorMessage = %s", JsonMarshalIndentToString(e))
}

//IsErrorResponse 是否亚马逊API错误响应
func IsErrorResponse(err error) bool {
	if err != nil {
		_, ok := err.(*ErrorResponse)
		return ok
	}
	return false
}

//IsErrorResponseWithCode 是否亚马逊API错误响应
func IsErrorResponseWithCode(err error, code ...string) bool {
	if err != nil {
		if resp, ok := err.(*ErrorResponse); ok {
			for _, e := range resp.Errors {
				if len(code) > 0 {
					for _, c := range code {
						if e.Code == c {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

//IsServiceUnavailable 亚马逊 MWS 服务是否不可用, http响应代码应为 503, 错误表明亚马逊 MWS 服务不可用。。应当使用“指数退避”方法重试请求。
func IsServiceUnavailable(err error) bool {
	return IsErrorResponseWithCode(err, "ServiceUnavailable")
}

//IsRequestThrottled 判断请求是否被限制, http响应代码应为 503, 错误表明您的请求已被限制。请查看您所提交请求类型的相关限制。请设定重试逻辑，以便在适当时间后重新发送请求，以免触发限制。
func IsRequestThrottled(err error) bool {
	return IsErrorResponseWithCode(err, "RequestThrottled")
}

//IsUserAgentHeaderMalformed 判断是否是错误的UserAgent标头, 错误表明请求中所含的 User-Agent 标头为无效格式。请使用亚马逊MWS 客户端库中的代码创建 User-Agent 标头，或参阅相关文档，以了解可接受的 User-Agent 标头格式。
func IsUserAgentHeaderMalformed(err error) bool {
	return IsErrorResponseWithCode(
		err,
		"UserAgentHeaderMissing",
		"UserAgentHeaderMalformed",
		"UserAgentHeaderLanguageAttributeMissing",
		"UserAgentHeaderMaximumLengthExceeded",
	)
}

//IsSignatureDoesNotMatch 所提供的请求签名与服务器计算的签名值不一致。
func IsSignatureDoesNotMatch(err error) bool {
	return IsErrorResponseWithCode(err, "SignatureDoesNotMatch")
}

//IsReportNotReady 报告尚未生成
func IsReportNotReady(err error) bool {
	return IsErrorResponseWithCode(err, "ReportNotReady")
}

//IsInputDataError 上传数据内容包含错误
func IsInputDataError(err error) bool {
	return IsErrorResponseWithCode(err, "InputDataError")
}

//IsInternalError 发生了未知的服务器错误
func IsInternalError(err error) bool {
	return IsErrorResponseWithCode(err, "InternalError")
}

//IsAccessToFeedProcessingResultDenied 没有足够权限访问上传数据处理结果
func IsAccessToFeedProcessingResultDenied(err error) bool {
	return IsErrorResponseWithCode(err, "AccessToFeedProcessingResultDenied")
}

//IsInvalidAccessKeyId 提供的 AWSAccessKeyId 请求参数无效或过期。
func IsInvalidAccessKeyId(err error) bool {
	return IsErrorResponseWithCode(err, "InvalidAccessKeyId")
}

//IsAccessDenied 拒绝访问
func IsAccessDenied(err error) bool {
	return IsErrorResponseWithCode(err, "AccessDenied")
}

//IsInvalidError 是否参数有错误。
func IsInvalidError(err error) bool {
	return IsErrorResponseWithCode(
		err,
		"InvalidAccessKeyId",
		"InvalidFeedSubmissionId",
		"InvalidFeedType",
		"InvalidParameterValue",
		"InvalidQueryParameter",
		"InvalidReportId",
		"InvalidReportType",
		"InvalidRequest",
		"MissingClientTokenId",
		"MissingParameter",
	)

}

//ErrorNetwork 网络错误
type ErrorNetwork struct {
	Err    error
	Fields Fields //一些相关的信息
}

//Error 网络错误信息
func (e *ErrorNetwork) Error() string {
	return fmt.Sprintf("ErrorNetwork,ErrorMessage = %s", e.Err.Error())
}

//IsErrorNetwork 是否网络错误响应
func IsErrorNetwork(err error) bool {
	if err != nil {
		_, ok := err.(*ErrorNetwork)
		return ok
	}
	return false
}
