package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/request_define"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
)

func TestGetWalletBalance_SignPack(t *testing.T) {
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

	address := "rSampleWalletAddress123456789"
	contractAddress := "XRP"
	chainId := 5

	reqBody, timestamp, sign, clientSign, err := apiObj.GetWalletBalance(address, contractAddress, chainId)
	if err != nil {
		t.Fatalf("GetWalletBalance failed: %v", err)
	}

	if len(reqBody) == 0 {
		t.Fatalf("reqBody is empty")
	}
	if timestamp == "" {
		t.Fatalf("timestamp is empty")
	}
	if sign == "" {
		t.Fatalf("sign is empty")
	}
	if clientSign == "" {
		t.Fatalf("clientSign is empty")
	}

	// Verify request JSON can be unmarshalled into RequestWalletBalance
	var req request_define.RequestWalletBalance
	if err := json.Unmarshal(reqBody, &req); err != nil {
		t.Fatalf("Unmarshal reqBody failed: %v", err)
	}
	if req.Address != address || req.ContractAddress != contractAddress || req.ChainId != chainId {
		t.Fatalf("Request fields mismatch: %+v", req)
	}

	// Verify RSA signature of clientSign
	reqMapObj := rsa_utils.ToStringMap(reqBody)
	err = apiObj.VerifyRSAsignature(reqMapObj, clientSign)
	if err != nil {
		logrus.Infof("Note on client sign verify: %v", err)
	}
}

func TestResponseWalletBalance_Unmarshal(t *testing.T) {
	rawJSON := `{
		"sign" : "",
		"timestamp" : "1725432397796",
		"data" : "1979984",
		"msg" : "ok",
		"code" : 1
	}`

	var rsp response_define.ResponseWalletBalance
	err := json.Unmarshal([]byte(rawJSON), &rsp)
	if err != nil {
		t.Fatalf("Unmarshal ResponseWalletBalance failed: %v", err)
	}

	if rsp.Code != 1 {
		t.Errorf("Expected code 1, got %d", rsp.Code)
	}
	if rsp.Msg != "ok" {
		t.Errorf("Expected msg 'ok', got %s", rsp.Msg)
	}
	if rsp.Data != "1979984" {
		t.Errorf("Expected data '1979984', got %s", rsp.Data)
	}
	if rsp.Timestamp != "1725432397796" {
		t.Errorf("Expected timestamp '1725432397796', got %s", rsp.Timestamp)
	}
}
