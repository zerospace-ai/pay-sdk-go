package api

import (
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
)

func (s *Sdk) GenerateRSASignature(mapData map[string]string) (string, error) {

	rawStr := rsa_utils.ComposeParams(mapData)

	privateKey, err := rsa_utils.LoadPrivateKeyFromBase64(s.config.RsaPrivateKey)
	if err != nil {
		return "", err
	}

	return rsa_utils.SignData(privateKey, rawStr)
}

func (s *Sdk) VerifyRSAsignature(mapData map[string]string, sign string) error {

	rawStr := rsa_utils.ComposeParams(mapData)

	publicKey, err := rsa_utils.ParsePublicKey(s.config.PlatformPubKey)
	if err != nil {
		return err
	}

	err = rsa_utils.VerifySignature(publicKey, rawStr, sign)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sdk) VerifyRiskRSAsignature(mapData map[string]string, sign string) error {

	rawStr := rsa_utils.ComposeParams(mapData)

	publicKey, err := rsa_utils.ParsePublicKey(s.config.PlatformRiskPubKey)
	if err != nil {
		return err
	}

	err = rsa_utils.VerifySignature(publicKey, rawStr, sign)
	if err != nil {
		return err
	}

	return nil
}
