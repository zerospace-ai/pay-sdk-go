package rsa_utils

import (
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRSAUtils(t *testing.T) {
	rawStr := "OpenId=HASH13900000002"
	prv, err := LoadPrivateKeyFromBase64(`MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQChw1SIrKi9ezrZq2ohekcp2m+cEtFereSZA26vW0TwFvg3VO9cbK2AN0EeQzLvT3HGHi6DS1sYct7RDwkxF3a1Dq5MiYQMFoPA0rI1HQCdQqi2ewoXYRkOkPWIHsYeZb7D2Cxt2+myJCPSwY7Q5G24crRMNvB8eYjzUs3V/kGChHim06I70Gs9IyoNBgrSE2SRyUHYATBRxsilqfdhVPNM9snmTOXTSsY4LMeRcdfs9uNEgFVhkFu+Pu9Hhtj8WwqHizQLOmWx+L+vUlUBe0QVq+/slfacUrugKCDVrVJtYpeW3qlwCDt+Af0DcNcTwic0TakmN0oRUzyd8rYw38knAgMBAAECggEAGbuig3mAAGCNnJnlsLysfG54ycm+j8K29lZy8adhwJXO17KFv1y0fwyLKd7DyAkJztv3R0CiFbIUWwp6ylnystvKg3duz6N3QyHhYoyiD+JOR3UsRkrWexg5TvWiS6yg27PUYYeV1hZksn7DEtz7zVdF8kfdGtgUnqyh/PRNBI+APG3dN+HB+kVzBTjNFdpPLNQWaIhOiGRmJE6VZQmvHu2zPwETr5FW/BNs52izIqD1s11N9Zne4aXC96VEHzU0hRLTtYXyCra0NYtEh/N8o0cylY8c/L07B8I0jmpRWVC6AsBfZMRC3uEeM8OzwmxGYcZB6pMgYZXN3/QOnpPOeQKBgQDgDrO56hK5zJQB0YtB9A3n9k5mVMRT9b+zsR0WnjK+WCpy6uCZjS2/LX6hoh3VXQGhq2ubLdnT4WQ9R9KhkcbOOVNSn+QcQHghst4/E1+Hchj5y0wBR0RxW9x1teq1Bjctn0EhRFLyWLTVFvbNUO20g3jFxXkdyleAyLqqkuS9fwKBgQC40xqNBGFn67etphDVcrF+vYHGtiKlgrdNHsR0I7HmXVikzgM4v7wLqQCBuefkeWO1qY6vWdbPFaWarHftTBLP2csuG9kgpz9mha4+JWf6f80cilGeT/SMtp5jGcvgwSy7XkJE37TBIZG8jyHu/iqnXEKYlkKDoIs16jccSWoYWQKBgDrjgldl70AIRgpoUcqm45TJBvgcZXGP0K/g6D1OnRWsRJPpqdiR46kwwuymmGLAzDH6xRCHL45h2h/FJdYzY2ZGaOD5h1Hpm+l3grRfWidWVLwqs7JOUe3dcAc8JhEoLg5+ofalZ4usxvKO05VJJQDh4TdR3LB4wyDlITK94wLvAoGAM6Az8GPi2WZMPiL+3MBWt/IzR9AkSVsD3HTVpyM5VAGK+y3YVIep/Q/N5m6JSZZmtZ2RD4XmrJ19ToESVqRDNO+/AzggJDDTUs3QZ+eG7b+5CVnRhokK5Cs2frP78OHEnumrVRWvpaq8zJCmh91TCMCKkZiXJ7E2cW7kye/vxLECgYB1ApKOp6jF5T/TwHtyX08dCzG2zebmN2yUNWryj/LY7gJqTgeDrNq4/kLq+1Uk5sdCrc3YskPyVb9fQ+oyIK4Z1801FY1VU2Lb9dxG2vIcCtZkb0/iYWwPqV37e29QjZC97JyHf5Ol1U2/L0iXIIaSUnpgXkSs5mAkdzggSXGJOA==`)
	if err != nil {
		logrus.Warnln("err:", err)
		return
	}
	sign, err := SignData(prv, rawStr)
	if err != nil {
		logrus.Warnln("err:", err)
		return
	}
	logrus.Infoln("sign:", sign)
}

func TestJson(t *testing.T) {
	jsonStr := `{"key1": "value1", "key2": {"subKey1": "subValue1"}, "key3": ["arrayValue1", {"arraySubKey": "arraySubValue"}]}`
	var jEle json.RawMessage = []byte(jsonStr)

	resMap := ToStringMap(jEle)
	for key, value := range resMap {
		println(key, ":", value)
	}
}
