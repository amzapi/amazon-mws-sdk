package mws

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	netUrl "net/url"
	"runtime"
	"strings"
	"time"
)

type AmazonClient struct {
	signatureMethod  string
	signatureVersion string
	userAgent        string
	api              string
	version          string
	client           *http.Client
}

func newAmazonClient(api, version string) *AmazonClient {
	return &AmazonClient{
		signatureVersion: "2",
		signatureMethod:  "HmacSHA256",
		userAgent:        fmt.Sprintf("amazon-mws-sdk/v%s (Language=%s; Platform=%s-%s)", VERSION, strings.Replace(runtime.Version(), "go", "go/", -1), runtime.GOOS, runtime.GOARCH),
		api:              api,
		version:          version,
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:             http.ProxyFromEnvironment,             //使用系统的代理
				DisableKeepAlives: true,                                  //如果为true,禁用HTTP保持活动状态,并仅将与服务器的连接用于单个HTTP请求。
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, //不再对证书进行校验
				//DialContext: (&net.Dialer{
				//	KeepAlive: 30 * time.Second, // KeepAlive指定两次保持活动的间隔探测活动的网络连接。
				//	Timeout:   30 * time.Second,
				//}).DialContext,
				//MaxIdleConns:          100,              //连接池最大数量
				//IdleConnTimeout:       90 * time.Second, //闲置连接在连接池中的保留时间
				//TLSHandshakeTimeout:   10 * time.Second, //TLS握手超时时间
				//ResponseHeaderTimeout: 10 * time.Second, //读取响应报文头超时时间
				//ExpectContinueTimeout: 1 * time.Second,  //等待收到一个go-ahead响应报文所用的时间
			},
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *AmazonClient) doSignature(credential *Credential, method string, data Values) {
	data.Set(keyVersion, c.version)
	data.Set(keyAWSAccessKeyID, credential.AccessKey)
	data.Set(keyMWSAuthToken, credential.AuthToken)
	data.Set(keySellerID, credential.SellerID)
	data.Set(keySignatureMethod, c.signatureMethod)
	data.Set(keySignatureVersion, c.signatureVersion)
	data.Set(keyTimestamp, time.Now().UTC().Format(time.RFC3339))
	data.Delete(keySignature)

	s := method + "\n" + credential.Region.Endpoint + "\n" + c.api + "\n" + data.Encode()

	mac := hmac.New(sha256.New, []byte(credential.SecretKey))
	mac.Write([]byte(s))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	data.Set(keySignature, signature)
}

func (c *AmazonClient) Request(method string, credential *Credential, data Values, body *bytes.Buffer) ([]byte, error) {

	if credential.AccessKey == "" || credential.SecretKey == "" || credential.SellerID == "" || credential.Region.Endpoint == "" {
		return nil, errors.New("incomplete request parameters")
	}

	c.doSignature(credential, method, data)

	url, err := netUrl.Parse(fmt.Sprintf("https://%s%s", credential.Region.Endpoint, c.api))
	if err != nil {
		return nil, err
	}
	url.RawQuery = data.Encode()

	var request *http.Request
	var contentType string
	var xmlContent string

	if body == nil {
		request, err = http.NewRequest(method, url.String(), nil)
		contentType = "application/x-www-form-urlencoded"
		xmlContent = ""
	} else {
		request, err = http.NewRequest(method, url.String(), body)
		contentType = "text/xml"
		xmlContent = body.String()
	}

	//记录请求信息
	var fields = Fields{
		"URL":    url.String(),
		"Params": JsonMarshalIndentToString(data),
		"Body":   xmlContent,
	}

	if err != nil {
		errResp := &ErrorNetwork{Err: err, Fields: fields}
		return nil, errResp
	}

	request.Header.Set("User-Agent", c.userAgent)
	request.Header.Set("Content-Type", contentType)

	resp, err := c.client.Do(request)
	if err != nil {
		errResp := &ErrorNetwork{Err: err, Fields: fields}
		return nil, errResp
	}

	defer resp.Body.Close()
	v, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		errResp := &ErrorNetwork{Err: err, Fields: fields}
		return nil, errResp
	}

	//记录返回信息
	fields["StatusCode"] = resp.StatusCode
	fields["Response"] = string(v)

	if resp.StatusCode != 200 {
		errResp := ErrorResponse{}
		err = xml.Unmarshal(v, &errResp)
		if err != nil {
			return v, err
		}
		errResp.Fields = fields
		return v, &errResp
	}

	return v, nil
}

func (c *AmazonClient) GetModel(method string, credential *Credential, data Values, body *bytes.Buffer, out interface{}) ([]byte, error) {
	v, err := c.Request(method, credential, data, body)
	if err != nil {
		return nil, err
	}
	return v, xml.Unmarshal(v, out)
}

func (c *AmazonClient) GetServiceStatus(credential *Credential) (string, *GetServiceStatusResult, error) {
	var result struct {
		BaseResponse
		Result *GetServiceStatusResult `xml:"GetFeedSubmissionListResult"`
	}
	if _, err := c.GetModel(http.MethodGet, credential, ActionValues("GetServiceStatus"), nil, &result); err != nil {
		return "", nil, err
	}
	return result.RequestID, result.Result, nil
}
