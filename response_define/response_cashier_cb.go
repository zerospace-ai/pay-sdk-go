package response_define

type RequestCashierCb struct {
	OrderId         string `json:"orderId"`
	OutOrderId      string `json:"outOrderId"`
	OrderStatus     string `json:"orderStatus"`
	OrderType       string `json:"orderType"`
	TokenId         string `json:"tokenId"`
	ChainId         string `json:"chainId"`
	ExchangeAddress string `json:"exchangeAddress"`
	Quantity        string `json:"quantity"`
	QuantityPaid    string `json:"quantityPaid"`
	TotalAmount     string `json:"totalAmount"`
	TotalAmountPaid string `json:"totalAmountPaid"`
	OutUserId       string `json:"outUserId"`
	CreatedAt       string `json:"createdAt"`
	Timestamp       string `json:"timestamp"`
	Sign            string `json:"sign"`
}

type ResponseCashierCb struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}
