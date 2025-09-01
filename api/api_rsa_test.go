package api

import (
	"encoding/json"
	"fmt"
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

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
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
