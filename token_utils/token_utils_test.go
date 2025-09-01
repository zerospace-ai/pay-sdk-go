package token_utils

import (
	"fmt"
	"testing"
)

func TestLoadAllToken(t *testing.T) {
	fmt.Println(GetTokenInfo(ChainID_BNB, TokenID_BNB))
}
