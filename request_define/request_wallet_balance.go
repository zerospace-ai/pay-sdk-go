package request_define

type RequestWalletBalance struct {
	Address         string `json:"address"`
	ContractAddress string `json:"contractAddress"`
	ChainId         int    `json:"chainId"`
}
