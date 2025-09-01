package request_define

type RequestGetWalletAddresses struct {
	OpenID   string `json:"OpenId"`
	ChainIDs string `json:"ChainIDs"`
}
