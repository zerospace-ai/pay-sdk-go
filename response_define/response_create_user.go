package response_define

type ResponseCreateUserData struct {
	OpenId string `json:"OpenId"`
}

type ResponseCreateUser struct {
	Code      int                    `json:"code"`
	Timestamp string                 `json:"timestamp"`
	Msg       string                 `json:"msg"`
	Data      ResponseCreateUserData `json:"data"`
	Sign      string                 `json:"sign"`
}
