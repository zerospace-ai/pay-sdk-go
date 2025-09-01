package response_define

type RequestWithdrawCb struct {
	Amount    string `json:"amount"`
	OpenId    string `json:"openId"`
	SafeCode  string `json:"safeCode"`
	Timestamp string `json:"timestamp"`
	ToAddress string `json:"toAddress"`
	TokenId   string `json:"tokenId"`
	Sign      string `json:"sign"`
}

type ResponseWithdraw struct {
	Code      string `json:"code"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Sign      string `json:"sign"`
}

type ResponseWithdrawCb struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Sign      string `json:"sign"`
}
