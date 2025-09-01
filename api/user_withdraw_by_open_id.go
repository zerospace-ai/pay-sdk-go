package api

import (
	"github.com/zerospace-ai/pay-sdk-go/request_define"
)

// UserWithdrawByOpenID user withdraw by openId
// @param tokenId token
// @param amount Amount of tokens
// @param addressTo address to
// @param callbackUrl callback url
// @param safeCheckCode safe check code
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) UserWithdrawByOpenID(openId string,
	tokenId,
	amount,
	addressTo,
	callbackUrl,
	safeCheckCode string,
) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestWithdrawByOpenID{
			OpenID:        openId,
			TokenID:       tokenId,
			Amount:        amount,
			AddressTo:     addressTo,
			CallBackURL:   callbackUrl,
			SafeCheckCode: safeCheckCode,
		},
	)
}
