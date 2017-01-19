package alipay

import "net/http"

const (
	K_ALI_PAY_TRADE_STATUS_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" // 交易创建，等待买家付款
	K_ALI_PAY_TRADE_STATUS_TRADE_CLOSED   = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款
	K_ALI_PAY_TRADE_STATUS_TRADE_SUCCESS  = "TRADE_SUCCESS"  // 交易支付成功
	K_ALI_PAY_TRADE_STATUS_TRADE_FINISHED = "TRADE_FINISHED" // 交易结束，不可退款
)

// https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.8AmJwg&treeId=203&articleId=105286&docType=1
type TradeNotification struct {
	AppId             string `json:"app_id"`              // 开发者的app_id
	AuthAppId         string `json:"auth_app_id"`         // App Id
	NotifyId          string `json:"notify_id"`           // 通知校验ID
	NotifyType        string `json:"notify_type"`         // 通知类型
	NotifyTime        string `json:"notify_time"`         // 通知时间
	TradeNo           string `json:"trade_no"`            // 支付宝交易号
	TradeStatus       string `json:"trade_status"`        // 交易状态
	TotalAmount       string `json:"total_amount"`        // 订单金额
	ReceiptAmount     string `json:"receipt_amount"`      // 实收金额
	InvoiceAmount     string `json:"invoice_amount"`      // 开票金额
	BuyerPayAmount    string `json:"buyer_pay_amount"`    // 付款金额
	SellerId          string `json:"seller_id"`           // 卖家支付宝用户号
	SellerEmail       string `json:"seller_email"`        // 卖家支付宝账号
	BuyerId           string `json:"buyer_id"`            // 买家支付宝用户号
	BuyerLogonId      string `json:"buyer_logon_id"`      // 买家支付宝账号
	FundBillList      string `json:"fund_bill_list"`      // 支付金额信息
	Charset           string `json:"charset"`             // 编码格式
	PointAmount       string `json:"point_amount"`        // 集分宝金额
	OutTradeNo        string `json:"out_trade_no"`        // 商户订单号
	OutBizNo          string `json:"out_biz_no"`          // 商户业务号
	GmtCreate         string `json:"gmt_create"`          // 交易创建时间
	GmtPayment        string `json:"gmt_payment"`         // 交易付款时间
	GmtRefund         string `json:"gmt_refund"`          // 交易退款时间
	GmtClose          string `json:"gmt_close"`           // 交易结束时间
	Subject           string `json:"subject"`             // 总退款金额
	Body              string `json:"body"`                // 商品描述
	RefundFee         string `json:"refund_fee"`          // 总退款金额
	Version           string `json:"version"`             // 接口版本
	SignType          string `json:"sign_type"`           // 签名类型
	Sign              string `json:"sign"`                // 签名
	PassbackParams    string `json:"passback_params"`     // 回传参数
	VoucherDetailList string `json:"voucher_detail_list"` // 优惠券信息
}

func GetTradeNotification(req *http.Request) (noti *TradeNotification) {
	if req == nil {
		return nil
	}
	req.ParseForm()

	noti = &TradeNotification{}
	noti.AppId = req.PostFormValue("app_id")
	noti.AuthAppId = req.PostFormValue("auth_app_id")
	noti.NotifyId = req.PostFormValue("notify_id")
	noti.NotifyType = req.PostFormValue("notify_type")
	noti.NotifyTime = req.PostFormValue("notify_time")
	noti.TradeNo = req.PostFormValue("trade_no")
	noti.TradeStatus = req.PostFormValue("trade_status")
	noti.TotalAmount = req.PostFormValue("total_amount")
	noti.ReceiptAmount = req.PostFormValue("receipt_amount")
	noti.InvoiceAmount = req.PostFormValue("invoice_amount")
	noti.BuyerPayAmount = req.PostFormValue("buyer_pay_amount")
	noti.SellerId = req.PostFormValue("seller_id")
	noti.SellerEmail = req.PostFormValue("seller_email")
	noti.BuyerId = req.PostFormValue("buyer_id")
	noti.BuyerLogonId = req.PostFormValue("buyer_logon_id")
	noti.FundBillList = req.PostFormValue("fund_bill_list")
	noti.Charset = req.PostFormValue("charset")
	noti.PointAmount = req.PostFormValue("point_amount")
	noti.OutTradeNo = req.PostFormValue("out_trade_no")
	noti.OutBizNo = req.PostFormValue("out_biz_no")
	noti.GmtCreate = req.PostFormValue("gmt_create")
	noti.GmtPayment = req.PostFormValue("gmt_payment")
	noti.GmtRefund = req.PostFormValue("gmt_refund")
	noti.GmtClose = req.PostFormValue("gmt_close")
	noti.Subject = req.PostFormValue("subject")
	noti.Body = req.PostFormValue("body")
	noti.RefundFee = req.PostFormValue("refund_fee")
	noti.Version = req.PostFormValue("version")
	noti.SignType = req.PostFormValue("sign_type")
	noti.Sign = req.PostFormValue("sign")
	noti.PassbackParams = req.PostFormValue("passback_params")
	noti.VoucherDetailList = req.PostFormValue("voucher_detail_list")

	if len(noti.NotifyId) == 0 {
		return nil
	}
	return noti
}
