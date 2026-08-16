package common

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
)

var (
	once   sync.Once
	client *resty.Client
	apiObj *api.Sdk
)

func Init() (*resty.Client, *api.Sdk) {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../")    // Support running from example subdirectories
		viper.AddConfigPath("../../") // Support running tests from nested subdirectories
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintf("Failed to load config: %s", err))
		}

		apiObj = api.NewSDK(api.SDKConfig{
			ApiKey:             viper.GetString("ApiKey"),
			ApiSecret:          viper.GetString("ApiSecret"),
			PlatformPubKey:     viper.GetString("PlatformPubKey"),
			PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
			RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
		})

		client = resty.New().
			SetTimeout(15*time.Second).
			SetRetryCount(3).
			SetRetryWaitTime(2*time.Second).
			SetHeader("Content-Type", "application/json").
			SetHeader("key", apiObj.GetApiKey())
	})
	return client, apiObj
}

func ExecuteRequest(path string, reqBody []byte, timestamp, sign, clientSign string, responseObj any) error {
	client, apiObj := Init()

	endpoint := viper.GetString("Endpoint")
	if endpoint == "" {
		endpoint = api.DevNetEndpoint
	}

	finalURL, err := url.JoinPath(endpoint, path)
	if err != nil {
		return fmt.Errorf("url join path error: %w", err)
	}

	resp, err := client.R().
		SetBody(reqBody).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		Post(finalURL)

	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	body := resp.Body()
	logrus.Infof("Response Body: %s", string(body))

	// Verify common response
	var rspCommon response_define.ResponseCommon
	if err := json.Unmarshal(body, &rspCommon); err != nil {
		return fmt.Errorf("unmarshal common response error: %w", err)
	}

	if rspCommon.Code != response_define.SUCCESS {
		return fmt.Errorf("api error: code=%d, msg=%s", rspCommon.Code, rspCommon.Msg)
	}

	// Unmarshal to specific response object
	if responseObj != nil {
		if err := json.Unmarshal(body, responseObj); err != nil {
			return fmt.Errorf("unmarshal specific response error: %w", err)
		}

		// Verify signature if present
		mapObj := rsa_utils.ToStringMap(body)
		// We need to find the sign field in the response object via reflection or just use the map
		// But usually the sign is in the specific response struct.
		// For simplicity, we can pass the sign if we have it, or try to get it from the map.
		if signVal, ok := mapObj["sign"]; ok && signVal != "" {
			if err := apiObj.VerifyRSAsignature(mapObj, signVal); err != nil {
				return fmt.Errorf("verify signature error: %w", err)
			}
			logrus.Infoln("VerifyRSAsignature success")
		}
	}

	return nil
}
