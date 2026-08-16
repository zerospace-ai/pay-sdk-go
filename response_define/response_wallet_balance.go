package response_define

type ResponseWalletBalance struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Msg       string `json:"msg"`
	Code      int    `json:"code"`
}
