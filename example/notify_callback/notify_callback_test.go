package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func TestCallbackVerify(t *testing.T) {
	_, apiObj := common.Init()

	body := []byte(viper.GetString("CBRsp"))
	if len(body) == 0 {
		t.Skip("CBRsp is not configured in config.yaml")
	}

	req := response_define.RequestTokenCb{}
	fmt.Println("Raw JSON:", string(body))
	err := json.Unmarshal(body, &req)
	if err != nil {
		t.Fatalf("json.Unmarshal fail: %v", err)
	}

	mapData := make(map[string]string)

	err = json.Unmarshal(body, &mapData)
	if err != nil {
		logrus.Warnln("json.Unmarshal fail, err", err.Error())
		return
	}

	err = apiObj.VerifyRSAsignature(mapData, req.Sign)
	if err != nil {
		logrus.Warnln("VerifyRSAsignature fail, err", err.Error())
		return
	}
	logrus.Infoln(req)
}
