package response_define

type ResponseUserWithdrawByOpenID struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      string `json:"data"`
}

type ResponseUserWithdrawByOpenIDData struct {
	UserPayChargeRecords []UserPayChargeRecords `json:"userPayChargeRecords"`
	TotalCount           int                    `json:"TotalCount"`
}

type UserPayChargeRecords struct {
	OpenID        string  `json:"OpenID"`
	PayOrCharge   int     `json:"PayOrCharge"`
	FromAddress   string  `json:"FromAddress"`
	ToAddress     string  `json:"ToAddress"`
	HashCode      string  `json:"HashCode"`
	SafeCode      string  `json:"SafeCode"`
	Amount        float64 `json:"Amount"`
	Status        int     `json:"Status"`
	NoticeTimes   int     `json:"NoticeTimes"`
	NoticeURL     string  `json:"NoticeUrl"`
	NoticeRespone string  `json:"NoticeRespone"`
	CreateTime    string  `json:"CreateTime"`
}

type RequestWithdrawCallback struct {
	Amount       string `json:"amount"`
	Chainid      string `json:"chainid"`
	Confirm      string `json:"confirm"`
	Createdtime  string `json:"createdtime"`
	From         string `json:"from"`
	Hash         string `json:"hash"`
	SafeCode     string `json:"safecode"`
	Sign         string `json:"sign"`
	Status       string `json:"status"`
	Timestamp    string `json:"timestamp"`
	To           string `json:"to"`
	Tokenaddress string `json:"tokenaddress"`
	TokenId      string `json:"tokenid"`
	Type         string `json:"type"`
}
