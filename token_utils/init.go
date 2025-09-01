package token_utils

import "encoding/json"

func initToken(token TokenInfo) {
	mapTokenInfo[GenerateKey(token.ChainID, token.TokenID)] = token
}

func init() {
	var tokens []TokenInfo
	json.Unmarshal([]byte(tokenJson), &tokens)
	for _, token := range tokens {
		initToken(token)
	}
}
