package logic

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/order"
	"github.com/smartwalle/alipay/v3"
)

var client *alipay.Client

const (
	kAppId = "9021000141609790"
	//私钥
	kPrivateKey = "MIIEpAIBAAKCAQEAhHqLrDbG2mRpV4kDzIkjHcrDUZxjyyiTiCkAoirPXK8AbN62QT+1u/ITAV0vDP23HvRxY6PWPxQc+6FbdNsC6eXagBrNy1SzC2hl2pLtj16vtPWH8jjlCsNgjWY3+v6avFV/4j5CrzutQrDWmEKu9kmEM0TCSwJp4b9Y9zLbHqH9uiHe8FPwQL6rxDomwEglWuJDiyzPbkGWISHQ4rrxqJvJrPdp8PBm6Ijjd6rsMlir2AQsErmxLTXOH6FTol+LaSwXR8Q+D7+tPvgrowfQuvwcoOW+l30hbAtktEBCskf80IF+kFScqRor3KxVrexQRRGbUSVoRsQylqX4FXtrmwIDAQABAoIBAHHvoyhdg/BX43FxWV8ToqwAsrMBuaIEJ7425WCTSPwBVnBLqpu7W2Vk8It/xPN7UX7B5LM5OQ3PEo4nXCzIH11iXeVHKC7zjXkGNHMjYkSnvp33fRAJdpMnLWbuvr6TMik/r1i5C0kCHiT5SdZxo2AWIiilviQc00tf8ffIOUPB7yjHodot4B0tSVJYhZVzAfisA2xXX3bIlL724pOQNACc4H5Uk6gPSwsP7E34jPdl1JvCzPwXWAX0tW2u3P2iZMRJLrrSkcYuKPuns+SMxPmupfoReugFxEJNQkeJMEKBsd/zdkRi/TxMBwp8aELQQMji+i5spDol/+C1ay4inmECgYEA7TPHJ5Q5NsA4cCkT1N/mRNSqNWFufTw49U8f4Cg0UzL1ksKuhZ2F8uQPQ9A/80QE9KwARwsSkauP123RE+bK7XoYAaNbmumOMlU7yNy4Yvjex1UMd4M1jRsZGdwi7of72WOmfL221ANSR4DfubZybyn4mJsez6us+VVb8ZoQbU0CgYEAjvoyVfkcONyE0pSnzQH4ubg+W+s77tOwY+NgpCQV1hpxnT0kdO/vR+Cr4NAJasUASi1YGSPBs6hS53Dq2jW0w0GWQpIbqIqkcttFdm5NpDIMCgZJkTuOEYbzsEnh9Tj8dBT9pIhcC8MUWCpwuRxx4yv5/AmC57aRduKvODDU6IcCgYEAibEDy3LcX27nlBwUOf8asRvuDfyX4bTCr4uVyk4sSZuMN/wwyLkvF1bRkaDibnkuk0UBCDxiNBMt9XT26CGR7JZyNwk31M380DTv9mJB2pny4Cz7hqiwvk9bm8pQsLxZMtEIxOJvrdSbv4wM0sYY2XS4rQVmtlBGoZmuPKodHQUCgYBH/Lb9THimNHl3oTs/YB/AOoDDPRQm8lScZm5RFo4cB6JunctOSUP7t6YygexL4rZ7oor51WQAtd5QAy1sAqBpswsZp4Dfgja8UaSOtYjIBB7FqVa/FfI45u6pMvnewvtw4uS2Q6W8klMa9PteaHD7BThPauaLDgP9jLLWItf/3wKBgQCNBzyQbRH4aeKywmQVUoQlfJcxTOCQBQoFuhQlagfFD3SLXfE8SuTANLi6z3FwX24bORNZxu3vIXdwKsc+Soz1R6k8SbY3azhK8tWe+9V0M7jH9Q7xUfFA3PLS10s1Qr/kMWbol3CTKPqgcCsuzSDrPykXv8rkj22NInyCYL9JZg=="
	// TODO 设置回调地址域名
	kServerDomain = "http://3c4c4ba0.r17.cpolar.top"
)

func init() {
	//支付宝提供了用于开发时测试的 sandbox 环境，对接的时候需要注意相关的 app id 和密钥是 sandbox 环境还是 production 环境的。如果是 sandbox 环境，本参数应该传 false，否则为 true。
	var err error
	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Fatal("初始化支付宝失败", err)
		return
	}

	// 加载证书
	// 加载应用公钥证书
	if err = client.LoadAppCertPublicKeyFromFile("/home/yjy/projects/cmall/e-mall/crts/appPublicCert.crt"); err != nil {
		log.Fatal("加载证书发生错误", err)
		return
	}
	// 加载支付宝根证书
	if err = client.LoadAliPayRootCertFromFile("/home/yjy/projects/cmall/e-mall/crts/alipayRootCert.crt"); err != nil {
		log.Fatal("加载证书发生错误", err)
		return
	}
	// 加载支付宝公钥证书
	if err = client.LoadAlipayCertPublicKeyFromFile("/home/yjy/projects/cmall/e-mall/crts/alipayPublicCert.crt"); err != nil {
		log.Fatal("加载证书发生错误", err)
		return
	}
}

type Payment struct {
	Data   *data.Data
	Client *alipay.Client
}

func NewPayment(data *data.Data) domain.Payment {
	return &Payment{
		Data:   data,
		Client: client,
	}
}

func (pay *Payment) Init(ctx context.Context, orderNum uint64, userID uint64) (*url.URL, error) {
	orderDB, err := pay.Data.DBClient.Order.Query().Where(order.OrderNum(orderNum)).First(ctx)
	if err != nil {
		return nil, err
	}
	if orderDB.UserID != userID {
		return nil, errors.New("不是你的订单~")
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/api/v1/alipay/notify"
	p.ReturnURL = kServerDomain + "/api/v1/alipay/callback"
	p.Subject = "支付测试:" + fmt.Sprintf("%d", orderNum)
	p.OutTradeNo = fmt.Sprintf("%d", orderNum)
	p.TotalAmount = fmt.Sprintf("%d", orderDB.Price)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	fmt.Println(url)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (pay *Payment) Notify(ctx context.Context, v url.Values) error {
	notification, err := client.DecodeNotification(v)
	if err != nil {
		return err
	}

	// 查询订单状态
	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err := client.Request(ctx, p, &rsp); err != nil {
		return err
	}

	if rsp.IsFailure() {
		return errors.New("支付未成功~请检查")
	}

	fmt.Printf("订单 %s 支付成功", notification.OutTradeNo)
	orderNum, err := strconv.ParseUint(notification.OutTradeNo, 10, 64)

	err = pay.Data.DBClient.Order.Update().Where(order.OrderNum(orderNum)).SetType(2).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
