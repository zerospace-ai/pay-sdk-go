package request_define

type RequestCreateWallet struct {
	ChainID string `json:"ChainID"`
	OpenID  string `json:"OpenId"`
}
