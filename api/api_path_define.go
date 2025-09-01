package api

const (
	DevNetEndpoint = "https://sandbox-api.privatex.io/sdk"

	PathCreateUser           = "/user/create"
	PathCreateWallet         = "/wallet/create"
	PathUserWithdrawByOpenID = "/partner/UserWithdrawByOpenID"
	PathGetBlockHeader       = "/hash/getBlockHeader"
	PathGetTXDetail          = "/hash/getTxDetail"
	PathQueryTokenList       = "/partner/QueryTokenList"
	PathGetWalletAddresses   = "/wallet/getWalletAddresses"
	PathGetTokenPrice        = "/token/price"
)
