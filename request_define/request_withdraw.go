package request_define

type RequestWithdrawByOpenID struct {
	OpenID        string `json:"OpenId"`
	TokenID       string `json:"TokenId"`
	Amount        string `json:"Amount"`
	AddressTo     string `json:"AddressTo"`
	CallBackURL   string `json:"CallBackUrl"`
	SafeCheckCode string `json:"SafeCheckCode"`
}
