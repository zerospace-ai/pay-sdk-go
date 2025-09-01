package api

import (
	"github.com/zerospace-ai/pay-sdk-go/request_define"
)

// GetWalletAddresses Query the wallet addresses of some token IDs specified by a certain openId. If no results are
// found in the query, the system will automatically create them and return a list.
// @param openId user open id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) GetWalletAddresses(openId, chainIDs string) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestGetWalletAddresses{
			OpenID:   openId,
			ChainIDs: chainIDs,
		},
	)
}
