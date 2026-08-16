package api

import (
	"github.com/zerospace-ai/pay-sdk-go/request_define"
)

// GetWalletBalance query wallet balance
// @param address wallet address
// @param contractAddress token contract address or token symbol (e.g. XRP, USDT)
// @param chainId chain id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) GetWalletBalance(address, contractAddress string, chainId int) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestWalletBalance{
			Address:         address,
			ContractAddress: contractAddress,
			ChainId:         chainId,
		},
	)
}

// WalletBalance query wallet balance
// @param address wallet address
// @param contractAddress token contract address or token symbol (e.g. XRP, USDT)
// @param chainId chain id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) WalletBalance(address, contractAddress string, chainId int) ([]byte, string, string, string, error) {
	return s.GetWalletBalance(address, contractAddress, chainId)
}

// GetWalletBalanceByReq query wallet balance with RequestWalletBalance
// @param req request_define.RequestWalletBalance
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) GetWalletBalanceByReq(req request_define.RequestWalletBalance) ([]byte, string, string, string, error) {
	return s.signPack(req)
}
