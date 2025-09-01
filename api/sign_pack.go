package api

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
)

func (s *Sdk) signPack(req interface{}) ([]byte, string, string, string, error) {

	timestamp := ""
	sign := ""
	clientSign := ""

	mapData, err := rsa_utils.StructToMap(req)
	if err != nil {
		return nil, timestamp, sign, clientSign, err
	}

	dataStr := rsa_utils.ComposeParams(mapData)

	timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign = s.GenerateMD5Sign(dataStr, timestamp)

	jStr, err := json.Marshal(&req)
	if err != nil {
		return nil, timestamp, sign, clientSign, err
	}

	reqMapObj := rsa_utils.ToStringMap(jStr)
	clientSign, err = s.GenerateRSASignature(reqMapObj)
	if err != nil {
		return nil, timestamp, sign, clientSign, err
	}

	return jStr, timestamp, sign, clientSign, nil
}
