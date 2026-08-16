package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	_, apiObj := common.Init()

	address := viper.GetString("WalletAddress")
	contractAddress := viper.GetString("ContractAddress")
	chainId := viper.GetInt("WalletBalanceChainId")

	reqBody, timestamp, sign, clientSign, err := apiObj.GetWalletBalance(address, contractAddress, chainId)
	if err != nil {
		logrus.Warnln("Error building request: ", err)
		return
	}

	var rspWalletBalance response_define.ResponseWalletBalance
	err = common.ExecuteRequest(api.PathWalletBalance, reqBody, timestamp, sign, clientSign, &rspWalletBalance)
	if err != nil {
		logrus.Warnln("Error executing request: ", err)
		return
	}

	fmt.Printf("GetWalletBalance Success: %+v\n", rspWalletBalance)
	fmt.Printf("Wallet Address: %s\n", address)
	fmt.Printf("Contract / Token: %s\n", contractAddress)
	fmt.Printf("Chain ID: %d\n", chainId)
	fmt.Printf("Balance: %s\n", rspWalletBalance.Data)
}
