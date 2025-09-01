package token_utils

import "fmt"

type TokenInfo struct {
	Mainnet      string `json:"mainnet"`
	ChainID      int64  `json:"chain_id"`
	TokenID      int64  `json:"token_id"`
	TokenAddress string `json:"token_address"`
	Symbol       string `json:"symbol"`
	Decimal      int16  `json:"decimals"`
}

var (
	mapTokenInfo = make(map[string]TokenInfo)
)

func GetTokenInfo(chainID, tokenId int64) *TokenInfo {
	key := GenerateKey(chainID, tokenId)
	if value, ok := mapTokenInfo[key]; ok {
		return &value
	}
	return nil
}

func GenerateKey(chainID, tokenId int64) string {
	return fmt.Sprintf("%d_%d", chainID, tokenId)
}
