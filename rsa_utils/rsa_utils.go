package rsa_utils

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

func ComposeParams(params map[string]string) string {
	var keys []string
	for key := range params {
		if key != "sign" {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var builder strings.Builder
	for i, key := range keys {
		if i > 0 {
			builder.WriteString("&")
		}
		builder.WriteString(fmt.Sprintf("%s=%s", key, params[key]))
	}

	return builder.String()
}

func StructToMap(data interface{}) (map[string]string, error) {
	result := make(map[string]string)

	val := reflect.ValueOf(data)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("json")

		switch field.Kind() {
		case reflect.String:
			result[fieldName] = field.String()
		default:
			return nil, fmt.Errorf("unsupported field type: %s", field.Kind())
		}
	}

	return result, nil
}

func LoadPrivateKeyFromBase64(base64String string) (*rsa.PrivateKey, error) {

	keyBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, fmt.Errorf("base64.StdEncoding.DecodeString err %v please check privateKey", err)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("x509.ParsePKCS8PrivateKey err %v please check privateKey", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return rsaPrivateKey, nil
}

func SignData(privateKey *rsa.PrivateKey, data string) (string, error) {
	hash := md5.New()

	io.WriteString(hash, data)
	hashInBytes := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashInBytes[:])
	if err != nil {
		return "", fmt.Errorf("SignData rsa.SignPKCS1v15 err %v please check privateKey", err)
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

func VerifySignature(publicKey *rsa.PublicKey, data, signature string) error {
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("VerifySignature base64.StdEncoding.DecodeString err %v please check signature %s", err, signature)
	}
	hash := md5.New()

	io.WriteString(hash, data)
	hashInBytes := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.MD5, hashInBytes[:], decodedSignature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %v", err)
	}

	return nil
}

func ParsePublicKey(publicKeyPEM string) (*rsa.PublicKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyPEM)
	if err != nil {
		return nil, errors.New("base64.StdEncoding.DecodeString err " + err.Error() + " please check publicKeyPEM")
	}

	publicKey, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return nil, errors.New("x509.ParsePKIXPublicKey err " + err.Error() + " please check publicKeyPEM " + publicKeyPEM)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("parsed public key is not an RSA public key please check publicKeyPEM " + publicKeyPEM)
	}

	return rsaPublicKey, nil
}
