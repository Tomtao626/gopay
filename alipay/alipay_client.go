package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iGoogle-ink/gopay"
	"log"
	"strings"
	"time"
)

type AliPayClient struct {
	AppId            string
	PrivateKey       string
	AlipayRootCertSN string
	AppCertSN        string
	ReturnUrl        string
	NotifyUrl        string
	Charset          string
	SignType         string
	AppAuthToken     string
	AuthToken        string
	IsProd           bool
}

//alipay.trade.fastpay.refund.query(统一收单交易退款查询)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (this *AliPayClient) AliPayTradeFastPayRefundQuery(body gopay.BodyMap) (aliRsp *AliPayTradeFastpayRefundQueryResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.fastpay.refund.query")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeFastPayRefundQuery::::", string(bytes))
	aliRsp = new(AliPayTradeFastpayRefundQueryResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal：%v", err.Error())
	}
	if aliRsp.AliPayTradeFastpayRefundQueryResponse.Code != "10000" {
		info := aliRsp.AliPayTradeFastpayRefundQueryResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.order.settle(统一收单交易结算接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (this *AliPayClient) AliPayTradeOrderSettle(body gopay.BodyMap) (aliRsp *AliPayTradeOrderSettleResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_request_no")
	trade2 := body.Get("trade_no")
	if trade1 == null || trade2 == null {
		return nil, errors.New("out_request_no or trade_no are not allowed to be null")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.order.settle")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeOrderSettle::::", string(bytes))
	aliRsp = new(AliPayTradeOrderSettleResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradeOrderSettleResponse.Code != "10000" {
		info := aliRsp.AliPayTradeOrderSettleResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.create(统一收单交易创建接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.create
func (this *AliPayClient) AliPayTradeCreate(body gopay.BodyMap) (aliRsp *AliPayTradeCreateResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("buyer_id")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and buyer_id are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.create")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeCreateResponse::::", string(bytes))
	aliRsp = new(AliPayTradeCreateResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradeCreateResponse.Code != "10000" {
		info := aliRsp.AliPayTradeCreateResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.close(统一收单交易关闭接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.close
func (this *AliPayClient) AliPayTradeClose(body gopay.BodyMap) (aliRsp *AliPayTradeCloseResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.close")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeCloseResponse::::", string(bytes))
	aliRsp = new(AliPayTradeCloseResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradeCloseResponse.Code != "10000" {
		info := aliRsp.AliPayTradeCloseResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.cancel(统一收单交易撤销接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.cancel
func (this *AliPayClient) AliPayTradeCancel(body gopay.BodyMap) (aliRsp *AliPayTradeCancelResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.cancel")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeCancel::::", string(bytes))
	aliRsp = new(AliPayTradeCancelResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradeCancelResponse.Code != "10000" {
		info := aliRsp.AliPayTradeCancelResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.refund(统一收单交易退款接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.refund
func (this *AliPayClient) AliPayTradeRefund(body gopay.BodyMap) (aliRsp *AliPayTradeRefundResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.refund")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeRefundResponse::::", string(bytes))
	aliRsp = new(AliPayTradeRefundResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradeRefundResponse.Code != "10000" {
		info := aliRsp.AlipayTradeRefundResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.refund(统一收单退款页面接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.page.refund
func (this *AliPayClient) AliPayTradePageRefund(body gopay.BodyMap) (aliRsp *AliPayTradePageRefundResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "	alipay.trade.page.refund")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradePageRefundResponse::::", string(bytes))
	aliRsp = new(AliPayTradePageRefundResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradePageRefundResponse.Code != "10000" {
		info := aliRsp.AliPayTradePageRefundResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.precreate(统一收单线下交易预创建)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.precreate
func (this *AliPayClient) AliPayTradePrecreate(body gopay.BodyMap) (aliRsp *AlipayTradePrecreateResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	if trade1 == null {
		return nil, errors.New("out_trade_no is not allowed to be null")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.precreate")
	if err != nil {
		return nil, err
	}
	//log.Println("AlipayTradePrecreateResponse::::", string(bytes))
	aliRsp = new(AlipayTradePrecreateResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayTradePrecreateResponse.Code != "10000" {
		info := aliRsp.AlipayTradePrecreateResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.pay(统一收单交易支付接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.pay
func (this *AliPayClient) AliPayTradePay(body gopay.BodyMap) (aliRsp *AliPayTradePayResponse, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return nil, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	//body.Set("product_code", "FACE_TO_FACE_PAYMENT")
	bytes, err = this.doAliPay(body, "alipay.trade.pay")
	if err != nil {
		return nil, err
	}

	//log.Println("AliPayTradePayResponse::::", string(bytes))
	aliRsp = new(AliPayTradePayResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradePayResponse.Code != "10000" {
		info := aliRsp.AliPayTradePayResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.query(统一收单线下交易查询)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.query
func (this *AliPayClient) AliPayTradeQuery(body gopay.BodyMap) (aliRsp *AliPayTradeQueryResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_trade_no")
	trade2 := body.Get("trade_no")
	if trade1 == null && trade2 == null {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	bytes, err = this.doAliPay(body, "alipay.trade.query")
	if err != nil {
		return nil, err
	}
	//log.Println("AliPayTradeQueryResponse::::", string(bytes))
	aliRsp = new(AliPayTradeQueryResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPayTradeQueryResponse.Code != "10000" {
		info := aliRsp.AliPayTradeQueryResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.app.pay(app支付接口2.0)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (this *AliPayClient) AliPayTradeAppPay(body gopay.BodyMap) (payParam string, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return null, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	//body.Set("product_code", "QUICK_MSECURITY_PAY")
	bytes, err = this.doAliPay(body, "alipay.trade.app.pay")
	if err != nil {
		return null, err
	}
	payParam = string(bytes)
	return payParam, nil
}

//alipay.trade.wap.pay(手机网站支付接口2.0)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.wap.pay
func (this *AliPayClient) AliPayTradeWapPay(body gopay.BodyMap) (payUrl string, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return null, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	body.Set("product_code", "QUICK_WAP_WAY")
	bytes, err = this.doAliPay(body, "alipay.trade.wap.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	return payUrl, nil
}

//alipay.trade.page.pay(统一收单下单并支付页面接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (this *AliPayClient) AliPayTradePagePay(body gopay.BodyMap) (payUrl string, err error) {
	var bytes []byte
	trade := body.Get("out_trade_no")
	if trade == null {
		return null, errors.New("out_trade_no is not allowed to be null")
	}
	//===============product_code值===================
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")
	bytes, err = this.doAliPay(body, "alipay.trade.page.pay")
	if err != nil {
		//log.Println("err::", err.Error())
		return null, err
	}
	payUrl = string(bytes)
	return payUrl, nil
}

//alipay.fund.trans.toaccount.transfer(单笔转账到支付宝账户接口)
//    文档地址：https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
func (this *AliPayClient) AlipayFundTransToaccountTransfer(body gopay.BodyMap) (aliRsp *AlipayFundTransToaccountTransferResponse, err error) {
	var bytes []byte
	trade1 := body.Get("out_biz_no")
	if trade1 == null {
		return nil, errors.New("out_biz_no is not allowed to be null")
	}
	bytes, err = this.doAliPay(body, "alipay.fund.trans.toaccount.transfer")
	if err != nil {
		return nil, err
	}
	//log.Println("AlipayFundTransToaccountTransferResponse::::", string(bytes))
	aliRsp = new(AlipayFundTransToaccountTransferResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayFundTransToaccountTransferResponse.Code != "10000" {
		info := aliRsp.AlipayFundTransToaccountTransferResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//alipay.trade.orderinfo.sync(支付宝订单信息同步接口)
//    文档地址：https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync
func (this *AliPayClient) AliPayTradeOrderinfoSync(body gopay.BodyMap) {

}

//alipay.system.oauth.token(换取授权访问令牌)
//    文档地址：https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (this *AliPayClient) AliPaySystemOauthToken(body gopay.BodyMap) (aliRsp *AliPaySystemOauthTokenResponse, err error) {
	var bs []byte
	grantType := body.Get("grant_type")
	if grantType == null {
		return nil, errors.New("grant_type is not allowed to be null")
	}
	code := body.Get("code")
	refreshToken := body.Get("refresh_token")
	if code == null && refreshToken == null {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}

	bs, err = aliPaySystemOauthToken(this.AppId, this.PrivateKey, body, "alipay.system.oauth.token", this.IsProd)
	if err != nil {
		return nil, err
	}
	//log.Println("AliPaySystemOauthToken::::", string(bytes))
	aliRsp = new(AliPaySystemOauthTokenResponse)
	err = json.Unmarshal(bs, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AliPaySystemOauthTokenResponse.AccessToken == null {
		info := aliRsp.ErrorResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

//alipay.user.info.share(支付宝会员授权信息查询接口)
//    body：此接口无需body参数
//    文档地址：https://docs.open.alipay.com/api_2/alipay.user.info.share
func (this *AliPayClient) AlipayUserInfoShare() (aliRsp *AlipayUserInfoShareResponse, err error) {
	var bs []byte
	bs, err = this.doAliPay(nil, "alipay.user.info.share")
	if err != nil {
		return nil, err
	}
	//log.Println("AlipayUserInfoShare::::", string(bs))
	aliRsp = new(AlipayUserInfoShareResponse)
	err = json.Unmarshal(bs, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayUserInfoShareResponse.Code != "10000" {
		info := aliRsp.AlipayUserInfoShareResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

//alipay.open.auth.token.app(换取应用授权令牌)
//    文档地址：https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
func (this *AliPayClient) AlipayOpenAuthTokenApp(body gopay.BodyMap) (aliRsp *AlipayOpenAuthTokenAppResponse, err error) {
	var bs []byte
	grantType := body.Get("grant_type")
	if grantType == null {
		return nil, errors.New("grant_type is not allowed to be null")
	}
	code := body.Get("code")
	refreshToken := body.Get("refresh_token")
	if code == null && refreshToken == null {
		return nil, errors.New("code and refresh_token are not allowed to be null at the same time")
	}

	bs, err = this.doAliPay(body, "alipay.open.auth.token.app")
	if err != nil {
		return nil, err
	}
	//log.Println("AlipayOpenAuthTokenApp::::", string(bs))
	aliRsp = new(AlipayOpenAuthTokenAppResponse)
	err = json.Unmarshal(bs, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.AlipayOpenAuthTokenAppResponse.Code != "10000" {
		info := aliRsp.AlipayOpenAuthTokenAppResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bs)
	return aliRsp, nil
}

//zhima.credit.score.get(芝麻分)
//    文档地址：https://docs.open.alipay.com/api_8/zhima.credit.score.get
func (this *AliPayClient) ZhimaCreditScoreGet(body gopay.BodyMap) (aliRsp *ZhimaCreditScoreGetResponse, err error) {
	var bytes []byte

	trade1 := body.Get("product_code")
	if trade1 == null {
		body.Set("product_code", "w1010100100000000001")
	}
	trade2 := body.Get("transaction_id")
	if trade2 == null {
		return nil, errors.New("transaction_id is not allowed to be null")
	}
	bytes, err = this.doAliPay(body, "zhima.credit.score.get")
	if err != nil {
		return nil, err
	}
	//log.Println("ZhimaCreditScoreGet::::", string(bytes))
	aliRsp = new(ZhimaCreditScoreGetResponse)
	err = json.Unmarshal(bytes, aliRsp)
	if err != nil {
		return nil, err
	}
	if aliRsp.ZhimaCreditScoreGetResponse.Code != "10000" {
		info := aliRsp.ZhimaCreditScoreGetResponse
		return nil, fmt.Errorf(`{"code":"%v","msg":"%v","sub_code":"%v","sub_msg":"%v"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	aliRsp.SignData = getSignData(bytes)
	return aliRsp, nil
}

//向支付宝发送请求
func (this *AliPayClient) doAliPay(body gopay.BodyMap, method string) (bytes []byte, err error) {
	//===============转换body参数===================
	var bodyStr string
	if body != nil {
		bodyBs, err := json.Marshal(body)
		if err != nil {
			log.Println("json.Marshal:", err)
			return nil, err
		}
		bodyStr = string(bodyBs)
	}
	//fmt.Println(string(bodyStr))
	//===============生成参数===================
	pubBody := make(gopay.BodyMap)
	pubBody.Set("app_id", this.AppId)
	pubBody.Set("method", method)
	pubBody.Set("format", "JSON")
	if this.AppCertSN != null {
		pubBody.Set("app_cert_sn", this.AppCertSN)
	}
	if this.AlipayRootCertSN != null {
		pubBody.Set("alipay_root_cert_sn", this.AlipayRootCertSN)
	}
	if this.ReturnUrl != null {
		pubBody.Set("return_url", this.ReturnUrl)
	}
	if this.Charset == null {
		pubBody.Set("charset", "utf-8")
	} else {
		pubBody.Set("charset", this.Charset)
	}
	if this.SignType == null {
		pubBody.Set("sign_type", "RSA2")
	} else {
		pubBody.Set("sign_type", this.SignType)
	}
	pubBody.Set("timestamp", time.Now().Format(gopay.TimeLayout))
	pubBody.Set("version", "1.0")
	if this.NotifyUrl != null {
		pubBody.Set("notify_url", this.NotifyUrl)
	}
	if this.AppAuthToken != null {
		pubBody.Set("app_auth_token", this.AppAuthToken)
	}
	if this.AuthToken != null {
		pubBody.Set("auth_token", this.AuthToken)
	}
	//fmt.Println("biz_content", string(bodyStr))
	if bodyStr != null {
		pubBody.Set("biz_content", bodyStr)
	}
	//===============获取签名===================
	pKey := FormatPrivateKey(this.PrivateKey)
	sign, err := getRsaSign(pubBody, pubBody.Get("sign_type"), pKey)
	if err != nil {
		return nil, err
	}
	pubBody.Set("sign", sign)
	//fmt.Println("rsaSign:", sign)
	//===============发起请求===================
	urlParam := FormatAliPayURLParam(pubBody)
	//fmt.Println("urlParam:", urlParam)
	if method == "alipay.trade.app.pay" {
		return []byte(urlParam), nil
	}
	if method == "alipay.trade.page.pay" {
		if !this.IsProd {
			//沙箱环境
			return []byte(zfb_sanbox_base_url + "?" + urlParam), nil
		} else {
			//正式环境
			return []byte(zfb_base_url + "?" + urlParam), nil
		}
	}
	var url string
	agent := gopay.HttpAgent()
	if !this.IsProd {
		//沙箱环境
		url = zfb_sanbox_base_url_utf8
		//fmt.Println(url)
		agent.Post(url)
	} else {
		//正式环境
		url = zfb_base_url_utf8
		//fmt.Println(url)
		agent.Post(url)
	}
	res, bs, errs := agent.
		Type("form-data").
		SendString(urlParam).
		EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	//fmt.Println("res:", res.StatusCode)
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %v", res.StatusCode)
	}
	if method == "alipay.trade.wap.pay" {
		//fmt.Println("rsp:::", rsp.Body)
		if res.Request.URL.String() == zfb_sanbox_base_url || res.Request.URL.String() == zfb_base_url {
			return nil, errors.New("请求手机网站支付出错，请检查各个参数或秘钥是否正确")
		}
		return []byte(res.Request.URL.String()), nil
	}
	return bs, nil
}

func getSignData(bs []byte) (signData string) {
	str := string(bs)
	indexStart := strings.Index(str, `":`)
	indexEnd := strings.Index(str, `,"sign"`)
	signData = str[indexStart+2 : indexEnd]
	return
}
