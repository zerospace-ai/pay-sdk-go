package response_define

type ResponseGetWalletAddresses struct {
	Msg       string                 `json:"msg"`
	Code      int64                  `json:"code"`
	Data      GetWalletAddressesData `json:"data"`
	Sign      string                 `json:"sign"`
	Timestamp string                 `json:"timestamp"`
}

type GetWalletAddressesData struct {
	Addresses []GetWalletAddressesAddress `json:"Addresses"`
	PartnerID int64                       `json:"PartnerId"`
	OpenID    string                      `json:"OpenId"`
}

type GetWalletAddressesAddress struct {
	Address string `json:"address"`
	ChainID int64  `json:"chainID"`
}
