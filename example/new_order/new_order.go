package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/request_define"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	_, apiObj := common.Init()

	req := request_define.RequestNewOrder{
		OutOrderNo: "TEST_ORDER_123456",
		TokenId:    6, // Example Token ID
		Quantity:   "0.01",
		NotifyUrl:  "https://your-domain.com/callback",
	}

	reqBody, timestamp, sign, clientSign, err := apiObj.NewOrder(req)
	if err != nil {
		logrus.Warnln("Error building request: ", err)
		return
	}

	var rsp response_define.ResponseNewOrder
	err = common.ExecuteRequest(api.PathNewOrder, reqBody, timestamp, sign, clientSign, &rsp)
	if err != nil {
		logrus.Warnln("Error executing request: ", err)
		return
	}

	fmt.Printf("NewOrder Success: %+v\n", rsp)
	fmt.Printf("Cashier URL: %s\n", rsp.Data.CashierUrl)
}
