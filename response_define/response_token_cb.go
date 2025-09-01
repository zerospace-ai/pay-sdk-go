package response_define

const (
	RequestTokenCb_Type_Deposit  = "1"
	RequestTokenCb_Type_Withdraw = "2"
)

type RequestTokenCb struct {
	OpenId       string `json:"openid"`
	TotalValue   string `json:"totalvalue"`
	Amount       string `json:"amount"`
	Chainid      string `json:"chainid"`
	Confirm      string `json:"confirm"`
	Createdtime  string `json:"createdtime"`
	From         string `json:"from"`
	Hash         string `json:"hash,omitempty"`
	Safecode     string `json:"safecode"`
	Sign         string `json:"sign"`
	Status       string `json:"status"`
	Timestamp    string `json:"timestamp"`
	To           string `json:"to"`
	Tokenaddress string `json:"tokenaddress"`
	Tokenid      string `json:"tokenid"`
	Type         string `json:"type"`
	Fee          string `json:"fee"`
}
