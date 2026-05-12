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
	tokenId := viper.GetString("TokenId")
	amount := viper.GetString("Amount")
	addressTo := viper.GetString("AddressTo")
	callbackUrl := viper.GetString("CallbackUrl")
	safeCheckCode := viper.GetString("SafeCheckCode")

	reqBody, timestamp, sign, clientSign, err := apiObj.UserWithdrawByOpenID(openId,
		tokenId,
		amount,
		addressTo,
		callbackUrl,
		safeCheckCode,
	)

	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	var rspUserWithdraw response_define.ResponseUserWithdrawByOpenID
	err = common.ExecuteRequest(api.PathUserWithdrawByOpenID, reqBody, timestamp, sign, clientSign, &rspUserWithdraw)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	fmt.Printf("UserWithdrawByOpenID Success: %+v\n", rspUserWithdraw)
}
