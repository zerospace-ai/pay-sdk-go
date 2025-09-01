package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/rsa_utils"
	"io"
)

func VerifySignature(publicKey *rsa.PublicKey, data, signature string) error {
	decodedSignature, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("error decoding signature: %v", err)
	}
	hash := sha256.New()

	io.WriteString(hash, data)
	hashInBytes := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashInBytes[:], decodedSignature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %v", err)
	}

	return nil
}

func SignData(privateKey *rsa.PrivateKey, data string) (string, error) {
	hash := sha256.New()

	io.WriteString(hash, data)
	hashInBytes := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashInBytes[:])
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(signature), nil
}

func main() {
	rawStr := "aaaaa"
	RsaPrivateKey := `xxxxxx`

	privateKey, err := rsa_utils.LoadPrivateKeyFromBase64(RsaPrivateKey)
	if err != nil {
		return
	}
	publicKey := &privateKey.PublicKey
	sign := "xxxxxx"

	err = VerifySignature(publicKey, rawStr, sign)
	if err != nil {
		fmt.Println("verify signature failed", err)
	} else {
		fmt.Println("verify signature success")
	}

	goSign, err := SignData(privateKey, rawStr)
	if err != nil {
		fmt.Println("sign data failed", err)
	} else {
		fmt.Println("sign data success", goSign)
	}

	err = VerifySignature(publicKey, rawStr, goSign)
	if err != nil {
		fmt.Println("verify signature failed", err)
	} else {
		fmt.Println("verify signature success")
	}

}
