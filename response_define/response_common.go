package response_define

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCommon struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data *json.RawMessage `json:"data,omitempty"`
}

type ResponseSuccess struct {
	Code      int              `json:"code"`
	Msg       string           `json:"msg"`
	Timestamp string           `json:"timestamp"`
	Data      *json.RawMessage `json:"data,omitempty"`
	Sign      string           `json:"sign"`
}

type ErrCode int

type Response struct {
	Code ErrCode     `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 0
	SUCCESS = 1
)

func Result(code ErrCode, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "operate success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "operate success", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "operate fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code ErrCode, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
