package main

import (
	"encoding/json"
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
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
	chainIDs := viper.GetString("ChainIDs")

	reqBody, timestamp, sign, clientSign, err := apiObj.GetWalletAddresses(openId, chainIDs)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	fmt.Println("reqBody: ", string(reqBody))

	finalURL, err := url.JoinPath(api.DevNetEndpoint, api.PathGetWalletAddresses)
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
		logrus.Warnln("Error: ", err, "finalURL", finalURL)
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

	rspGetWalletAddresses := response_define.ResponseGetWalletAddresses{}
	err = json.Unmarshal(body, &rspGetWalletAddresses)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}
	logrus.Infoln("ResponseGetWalletAddresses: ", rspGetWalletAddresses)

	mapObj := rsa_utils.ToStringMap(body)
	err = apiObj.VerifyRSAsignature(mapObj, rspGetWalletAddresses.Sign)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	logrus.Infoln("VerifyRSAsignature success")

}
