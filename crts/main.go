package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	alipay "github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
)

var client *alipay.Client

const (
	kAppId = "9021000141609790"
	//私钥
	kPrivateKey = "MIIEpAIBAAKCAQEAhHqLrDbG2mRpV4kDzIkjHcrDUZxjyyiTiCkAoirPXK8AbN62QT+1u/ITAV0vDP23HvRxY6PWPxQc+6FbdNsC6eXagBrNy1SzC2hl2pLtj16vtPWH8jjlCsNgjWY3+v6avFV/4j5CrzutQrDWmEKu9kmEM0TCSwJp4b9Y9zLbHqH9uiHe8FPwQL6rxDomwEglWuJDiyzPbkGWISHQ4rrxqJvJrPdp8PBm6Ijjd6rsMlir2AQsErmxLTXOH6FTol+LaSwXR8Q+D7+tPvgrowfQuvwcoOW+l30hbAtktEBCskf80IF+kFScqRor3KxVrexQRRGbUSVoRsQylqX4FXtrmwIDAQABAoIBAHHvoyhdg/BX43FxWV8ToqwAsrMBuaIEJ7425WCTSPwBVnBLqpu7W2Vk8It/xPN7UX7B5LM5OQ3PEo4nXCzIH11iXeVHKC7zjXkGNHMjYkSnvp33fRAJdpMnLWbuvr6TMik/r1i5C0kCHiT5SdZxo2AWIiilviQc00tf8ffIOUPB7yjHodot4B0tSVJYhZVzAfisA2xXX3bIlL724pOQNACc4H5Uk6gPSwsP7E34jPdl1JvCzPwXWAX0tW2u3P2iZMRJLrrSkcYuKPuns+SMxPmupfoReugFxEJNQkeJMEKBsd/zdkRi/TxMBwp8aELQQMji+i5spDol/+C1ay4inmECgYEA7TPHJ5Q5NsA4cCkT1N/mRNSqNWFufTw49U8f4Cg0UzL1ksKuhZ2F8uQPQ9A/80QE9KwARwsSkauP123RE+bK7XoYAaNbmumOMlU7yNy4Yvjex1UMd4M1jRsZGdwi7of72WOmfL221ANSR4DfubZybyn4mJsez6us+VVb8ZoQbU0CgYEAjvoyVfkcONyE0pSnzQH4ubg+W+s77tOwY+NgpCQV1hpxnT0kdO/vR+Cr4NAJasUASi1YGSPBs6hS53Dq2jW0w0GWQpIbqIqkcttFdm5NpDIMCgZJkTuOEYbzsEnh9Tj8dBT9pIhcC8MUWCpwuRxx4yv5/AmC57aRduKvODDU6IcCgYEAibEDy3LcX27nlBwUOf8asRvuDfyX4bTCr4uVyk4sSZuMN/wwyLkvF1bRkaDibnkuk0UBCDxiNBMt9XT26CGR7JZyNwk31M380DTv9mJB2pny4Cz7hqiwvk9bm8pQsLxZMtEIxOJvrdSbv4wM0sYY2XS4rQVmtlBGoZmuPKodHQUCgYBH/Lb9THimNHl3oTs/YB/AOoDDPRQm8lScZm5RFo4cB6JunctOSUP7t6YygexL4rZ7oor51WQAtd5QAy1sAqBpswsZp4Dfgja8UaSOtYjIBB7FqVa/FfI45u6pMvnewvtw4uS2Q6W8klMa9PteaHD7BThPauaLDgP9jLLWItf/3wKBgQCNBzyQbRH4aeKywmQVUoQlfJcxTOCQBQoFuhQlagfFD3SLXfE8SuTANLi6z3FwX24bORNZxu3vIXdwKsc+Soz1R6k8SbY3azhK8tWe+9V0M7jH9Q7xUfFA3PLS10s1Qr/kMWbol3CTKPqgcCsuzSDrPykXv8rkj22NInyCYL9JZg=="
	// TODO 设置回调地址域名
	kServerDomain = "http://3c4c4ba0.r17.cpolar.top"
)

func main() {
	var err error
	//支付宝提供了用于开发时测试的 sandbox 环境，对接的时候需要注意相关的 app id 和密钥是 sandbox 环境还是 production 环境的。如果是 sandbox 环境，本参数应该传 false，否则为 true。
	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}

	// 加载证书
	// 加载应用公钥证书
	if err = client.LoadAppCertPublicKeyFromFile("appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	// 加载支付宝根证书
	if err = client.LoadAliPayRootCertFromFile("alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	// 加载支付宝公钥证书
	if err = client.LoadAlipayCertPublicKeyFromFile("alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	//接口内容加密
	//	if err = client.SetEncryptKey("iotxR/d99T9Awom/UaSqiQ=="); err != nil {
	// log.Println("加载内容加密密钥发生错误", err)
	// return

	//路由函数
	r := gin.Default()
	r.GET("/alipay/pay", pay)
	r.GET("/alipay/callback", callback)
	r.POST("/alipay/notify", notify)

	//http.HandleFunc("/alipay/pay", pay)
	//http.HandleFunc("/alipay/callback", callback)
	//http.HandleFunc("/alipay/notify", notify)

	//http.ListenAndServe(":"+kServerPort, nil)
	r.Run()
}

func pay(c *gin.Context) {
	var tradeNo = fmt.Sprintf("%d", xid.Next())
	fmt.Println(tradeNo)
	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "20.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusTemporaryRedirect, url.String())
	//http.Redirect(writer, request, url.String(), http.StatusTemporaryRedirect)
}

func callback(c *gin.Context) {
	// 解析请求参数
	if err := c.Request.ParseForm(); err != nil {
		log.Println("解析请求参数发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求参数发生错误"})
		return
	}

	// 获取通知参数
	outTradeNo := c.Request.Form.Get("out_trade_no")
	fmt.Println(c.Request.Form)
	// 验证签名
	if err := client.VerifySign(c.Request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "回调验证签名发生错误"})
		return
	}

	log.Println("回调验证签名通过")

	// 查询订单状态
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(c, p)
	fmt.Println("rsp:", rsp)
	if err != nil {
		log.Printf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		log.Printf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("订单 %s 支付成功", outTradeNo)})
}
func notify(c *gin.Context) {
	// 解析请求参数
	if err := c.Request.ParseForm(); err != nil {
		log.Println("解析请求参数发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求参数发生错误"})
		return
	}
	fmt.Println(c.Request.Form)
	// 解析异步通知
	notification, err := client.DecodeNotification(c.Request.Form)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析异步通知发生错误"})
		return
	}

	log.Println("解析异步通知成功:", notification.NotifyId)

	// 查询订单状态
	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err := client.Request(c, p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s", notification.OutTradeNo, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s", notification.OutTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s-%s", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}

	log.Printf("订单 %s 支付成功", notification.OutTradeNo)

	client.ACKNotification(c.Writer)
}
