# Examples and Tools

This document is divided into two parts:
1. **Scenario-Based Code Examples:** Demonstrates how to handle API calls and verification in practical code.
2. **CLI Tools Guide:** Explains how to use the compiled executable files included with the SDK for quick testing.

---

## 1. Scenario-Based Code Examples

### 1.1 Complete API Call and Response Verification

The following code demonstrates how to use the SDK to build a "Create User" request, send the HTTP request, and verify the security of the signature of the data returned by the platform.

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. Initialize SDK and Reuse Resty Client (Prerequisite: config.yaml is configured)
	_, apiObj := common.Init()

	// 2. Generate request parameters and signature Header
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("Failed to build request: ", err)
		return
	}

	// 3. Send request and automatically verify response signature
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("Request or verification failed: ", err)
		return
	}

	fmt.Println("✅ Request successful and verified! Returned OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. CLI Tools Usage Guide

The SDK provides Command Line Interface (CLI) binary files for quickly testing each API endpoint.

### 2.1 Compile Executable Files

Run the `make` command in the SDK root directory, and the system will generate binary executable files for each function in the `bin` directory.
* **Windows:** Generates files ending in `.exe` (e.g., `create_user.exe`).
* **Mac/Linux:** Generates files without suffixes (e.g., `create_user`).

### 2.2 Prepare Configuration File

Before running the tools, ensure that the configured `config.yaml` file is placed in the `bin` directory.

### 2.3 Testing Endpoint Commands

#### Register New User
1. Modify the `UserOpenId` field in `bin/config.yaml`.
2. Run `./create_user` (or double-click `create_user.exe`).
3. If the OpenId is already registered, the tool will return an error.

#### Wallet Registration
1. Specify `UserOpenId` and `ChainID` in `bin/config.yaml`.
2. Run `./create_wallet`.

#### Get Deposit Addresses
1. Specify `UserOpenId` and the queried `ChainIDs` (e.g., "1,56") in `bin/config.yaml`.
2. Run `./get_wallet_addresses`.

#### Apply for Withdrawal
1. Specify the following in `bin/config.yaml`:
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (Unique order anti-duplication code)
   * `CallbackUrl`
2. Run `./user_withdraw_by_open_id`.
