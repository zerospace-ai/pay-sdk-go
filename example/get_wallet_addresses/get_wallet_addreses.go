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
	chainIDs := viper.GetString("ChainIDs")

	reqBody, timestamp, sign, clientSign, err := apiObj.GetWalletAddresses(openId, chainIDs)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	var rspGetWalletAddresses response_define.ResponseGetWalletAddresses
	err = common.ExecuteRequest(api.PathGetWalletAddresses, reqBody, timestamp, sign, clientSign, &rspGetWalletAddresses)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	fmt.Printf("GetWalletAddresses Success: %+v\n", rspGetWalletAddresses)
}
