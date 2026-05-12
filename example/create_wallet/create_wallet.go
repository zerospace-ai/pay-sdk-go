package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	_, apiObj := common.Init()

	openId := viper.GetString("UserOpenId")
	chainID := viper.GetString("ChainID")

	reqBody, timestamp, sign, clientSign, err := apiObj.CreateWallet(openId, chainID)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	var rspCreateWallet response_define.ResponseCreateWallet
	err = common.ExecuteRequest(api.PathCreateWallet, reqBody, timestamp, sign, clientSign, &rspCreateWallet)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	fmt.Printf("CreateWallet Success: %+v\n", rspCreateWallet)
}
