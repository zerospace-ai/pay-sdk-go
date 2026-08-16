package response_define

type ResponseNewOrderData struct {
	OrderNo    string `json:"orderNo"`
	OutOrderNo string `json:"outOrderNo"`
	OutUserId  string `json:"outUserId"`
	TokenId    int    `json:"tokenId"`
	Quantity   string `json:"quantity"`
	Addresses  string `json:"addresses"`
	CashierUrl string `json:"cashierUrl"`
}

type ResponseNewOrder struct {
	Code      int                  `json:"code"`
	Msg       string               `json:"msg"`
	Timestamp string               `json:"timestamp"`
	Sign      string               `json:"sign"`
	Data      ResponseNewOrderData `json:"data"`
}
