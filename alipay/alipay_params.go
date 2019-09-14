package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/iGoogle-ink/gopay"
	"hash"
	"net/url"
)

//	AppId      string `json:"app_id"`      //支付宝分配给开发者的应用ID
//	Method     string `json:"method"`      //接口名称
//	Format     string `json:"format"`      //仅支持 JSON
//	ReturnUrl  string `json:"return_url"`  //HTTP/HTTPS开头字符串
//	Charset    string `json:"charset"`     //请求使用的编码格式，如utf-8,gbk,gb2312等，推荐使用 utf-8
//	SignType   string `json:"sign_type"`   //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用 RSA2
//	Sign       string `json:"sign"`        //商户请求参数的签名串
//	Timestamp  string `json:"timestamp"`   //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
//	Version    string `json:"version"`     //调用的接口版本，固定为：1.0
//	NotifyUrl  string `json:"notify_url"`  //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
//	BizContent string `json:"biz_content"` //业务请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档

type OpenApiRoyaltyDetailInfoPojo struct {
	RoyaltyType  string `json:"royalty_type,omitempty"`
	TransOut     string `json:"trans_out,omitempty"`
	TransOutType string `json:"trans_out_type,omitempty"`
	TransInType  string `json:"trans_in_type,omitempty"`
	TransIn      string `json:"trans_in"`
	Amount       string `json:"amount,omitempty"`
	Desc         string `json:"desc,omitempty"`
}

//设置 应用公钥证书SN
//    appCertSN：应用公钥证书SN，通过 gopay.GetCertSN() 获取
func (this *AliPayClient) SetAppCertSN(appCertSN string) (client *AliPayClient) {
	this.AppCertSN = appCertSN
	return this
}

//设置 支付宝根证书SN
//    alipayRootCertSN：支付宝根证书SN，通过 gopay.GetCertSN() 获取
func (this *AliPayClient) SetAliPayRootCertSN(alipayRootCertSN string) (client *AliPayClient) {
	this.AlipayRootCertSN = alipayRootCertSN
	return this
}

//设置支付后的ReturnUrl
func (this *AliPayClient) SetReturnUrl(url string) (client *AliPayClient) {
	this.ReturnUrl = url
	return this
}

//设置支付宝服务器主动通知商户服务器里指定的页面http/https路径。
func (this *AliPayClient) SetNotifyUrl(url string) (client *AliPayClient) {
	this.NotifyUrl = url
	return this
}

//设置编码格式，如utf-8,gbk,gb2312等，默认推荐使用 utf-8
func (this *AliPayClient) SetCharset(charset string) (client *AliPayClient) {
	if charset == null {
		this.Charset = "utf-8"
	} else {
		this.Charset = charset
	}
	return this
}

//设置签名算法类型，目前支持RSA2和RSA，默认推荐使用 RSA2
func (this *AliPayClient) SetSignType(signType string) (client *AliPayClient) {
	if signType == null {
		this.SignType = "RSA2"
	} else {
		this.SignType = signType
	}
	return this
}

//设置应用授权
func (this *AliPayClient) SetAppAuthToken(appAuthToken string) (client *AliPayClient) {
	this.AppAuthToken = appAuthToken
	return this
}

//设置用户信息授权
func (this *AliPayClient) SetAuthToken(authToken string) (client *AliPayClient) {
	this.AuthToken = authToken
	return this
}

//获取参数签名
func getRsaSign(bm gopay.BodyMap, signType, privateKey string) (sign string, err error) {
	var (
		h              hash.Hash
		key            *rsa.PrivateKey
		hashs          crypto.Hash
		signStr        string
		encryptedBytes []byte
	)
	//解析秘钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return null, errors.New("秘钥错误")
	}
	//log.Println(block.Type)
	key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		//log.Println("x509.ParsePKCS1PrivateKey:", err)
		return null, err
	}

	switch signType {
	case "RSA":
		h = sha1.New()
		hashs = crypto.SHA1
	case "RSA2":
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}

	//signStr = sortAliPaySignParams(bm)
	signStr = bm.EncodeAliPaySignParams()
	//fmt.Println("原始字符串：", signStr)
	_, err = h.Write([]byte(signStr))
	if err != nil {
		return null, err
	}
	encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, key, hashs, h.Sum(nil))
	if err != nil {
		//log.Println("rsa.SignPKCS1v15:", err)
		return null, err
	}
	secretData := base64.StdEncoding.EncodeToString(encryptedBytes)
	return secretData, nil
}

//格式化请求URL参数
func FormatAliPayURLParam(body gopay.BodyMap) (urlParam string) {
	v := url.Values{}
	for key, value := range body {
		v.Add(key, value.(string))
	}
	urlParam = v.Encode()
	//fmt.Println("Encode后参数:", urlParam)
	return
}
