package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
)

func main() {
	_, apiObj := common.Init()
	r := gin.Default()

	r.POST("/withdrawal/order/check", func(c *gin.Context) {
		req := response_define.RequestWithdrawCb{}

		body, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Failed to read request body",
			})
			return
		}

		fmt.Println("Raw JSON:", string(body))
		err = json.Unmarshal(body, &req)
		if err != nil {
			logrus.Warnln("Unmarshal fail")
			return
		}

		mapData, err := rsa_utils.StructToMap(req)
		if err != nil {
			logrus.Warnln("StructToMap fail, err", err.Error())
			response_define.FailWithMessage("StructToMap fail "+err.Error(), c)
			return
		}

		logrus.Infoln("mapData", mapData)
		logrus.Infoln("Sign", req.Sign)

		err = apiObj.VerifyRiskRSAsignature(mapData, req.Sign)
		if err != nil {
			logrus.Warnln("VerifyRiskRSAsignature fail, err", err.Error(), "end")
			response_define.FailWithMessage("verify RSA signature fail "+err.Error(), c)
			return
		}

		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		rsp := response_define.ResponseWithdraw{
			Code:      "0",
			Timestamp: timestamp,
			Message:   "",
			Sign:      "",
		}

		jStr, err := json.Marshal(&req)
		if err != nil {
			logrus.Warnln("json.Marshal fail, err", err.Error())
			return
		}

		reqMapObj := rsa_utils.ToStringMap(jStr)
		clientSign, err := apiObj.GenerateRSASignature(reqMapObj)
		if err != nil {
			logrus.Warnln("apiObj.GenerateRSASignature fail, err", err.Error())
			return
		}
		rsp.Sign = clientSign

		logrus.Infoln("VerifyRiskRSAsignature success.")
		c.JSON(http.StatusOK, rsp)
	})
	r.Run(":9003")
}
