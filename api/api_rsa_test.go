package api

import (
	"encoding/json"
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/request_define"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
)

func TestRSAVerify(t *testing.T) {
	jStr := `{
		"sign": "Iogo1jpWdLkP+q1Uc4e/lL6/SCC7Vc5cErNGwcg3um/KhlVdPrvaNZ6herVhI6XxwfHZYNSCda/wmC2UJYczlGF4x2ubsi0Xodj8SSosztwB/OvfoLtQG4UOnAtciRmFaUqQ71kLi+BAQdOj6WzBkOVP1c6lhia8C5CIgL96C786BRCsVhXjVxnkmWX3T7Qnu6S/pSOQ/dCPmQkxtzu7Y3TvQXLAEPMBu0KEmHriA5qWjzy1JwSU9BUtrqzwlVu4V0LtfN/gsZ3N6m6CTh0+vjUma4ltVjYa6ZIeduKbVSUMu1NU/uZ+f7WIlujMjhWSZAbu5QlboFTb0mKRm0lxJQ==",
		"timestamp": "1722478998646",
		"data": {
			"OpenId": "HASH13900000001"
		},
		"msg": "ok",
		"code": 1
	}`

	rsp := response_define.ResponseSuccess{}
	err := json.Unmarshal([]byte(jStr), &rsp)
	if err != nil {
		logrus.Warnln("TestRSAVerify", err)
		return
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := NewSDK(SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	mapData := rsa_utils.ToStringMap([]byte(jStr))

	err = apiObj.VerifyRSAsignature(mapData, rsp.Sign)

	logrus.Infoln("TestRSAVerify", err)
}

func TestRSASignature(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := NewSDK(SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	req := request_define.RequestCreateUser{
		OpenID: "123456",
	}

	jStr, err := json.Marshal(&req)
	if err != nil {
		panic(fmt.Sprintf("Failed to Marshal: %s", err))
	}

	reqMapObj := rsa_utils.ToStringMap(jStr)
	clientSign, err := apiObj.GenerateRSASignature(reqMapObj)

	mapData := rsa_utils.ToStringMap([]byte(jStr))

	err = apiObj.VerifyRSAsignature(mapData, clientSign)

	logrus.Infoln("TestRSAVerify", err)
}

func TestJsonToMap(t *testing.T) {
	jStr := `{
  "sign" : "UDXzq9PIyS/LlaJ+O+/gnJn17N/fdKWD/4+u7i7i2CzF1LSiUEKw7i0KkMii5uVLXb7FaqVzWgG2M5Gzo5d14/nqrQJ0xD1aAg44wb3SjR5BdenOdvBxgcAm4IlA+i4ifJzIR5l1Rgxbjrxqgz+455XuMcccXkinwEgB+c+qG+/lGIcnBzRugqy1SayUFAvcZO1HH67g42MciAEgZjU9qaR+rjjQpXP3YlALQBDWaQg423LlgL+Il2N97CIRVgOtcLlpB1/eq6nHJx5haH3jSHAdeSj9hKRgEsuOneR7BRHSFh5JyLP5GSn2kCcuEES23f9PoQNlnjZ1UGFpC3Z76w==",
  "timestamp" : "1757147834383",
  "data" : {
    "Addresses" : [ {
      "address" : "0x9034670cee1564887e16b73e53228298d8bab302",
      "chainID" : 56
    }, {
      "address" : "TVPePaxhekW4s8TcunbwfK74MGJXfdQLW2",
      "chainID" : 2
    } ],
    "PartnerId" : 133,
    "OpenId" : "HASH202509061500"
  },
  "msg" : "ok",
  "code" : 1
}`

	mapData := rsa_utils.ToStringMap([]byte(jStr))
	logrus.Infoln(mapData)
}
