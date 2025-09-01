package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	client := resty.New()

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

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

	fmt.Println("reqBody: ", string(reqBody))
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	finalURL, err := url.JoinPath(api.DevNetEndpoint, api.PathUserWithdrawByOpenID)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		Post(finalURL)

	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	body := resp.Body()
	fmt.Println(string(body))

	rspCommon := response_define.ResponseCommon{}
	err = json.Unmarshal(body, &rspCommon)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}
	logrus.Infoln("Response: ", rspCommon)

	if rspCommon.Code != response_define.SUCCESS {
		logrus.Warnln("Response fail Code", rspCommon.Code, "Msg", rspCommon.Msg)
		return
	}

	rspCreateUser := response_define.ResponseUserWithdrawByOpenID{}
	err = json.Unmarshal(body, &rspCreateUser)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}
	logrus.Infoln("ResponseUserWithdrawByOpenID: ", rspCreateUser)

	mapObj := make(map[string]string)
	err = json.Unmarshal(body, &mapObj)
	if err != nil {
		logrus.Warnln("StructToMap fail, err", err.Error())
		return
	}

	err = apiObj.VerifyRSAsignature(mapObj, rspCreateUser.Sign)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	logrus.Infoln("VerifyRSAsignature success")

}
