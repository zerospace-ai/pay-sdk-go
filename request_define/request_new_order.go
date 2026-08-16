package request_define

type RequestNewOrder struct {
	OutOrderNo string `json:"outOrderNo"`
	TokenId    int    `json:"tokenId"`
	Quantity   string `json:"quantity"`
	NotifyUrl  string `json:"notifyUrl"`
}
