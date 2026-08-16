# 示例代碼與工具

本文檔分為兩部分：
1. **場景化代碼示例：** 演示在實際代碼中如何處理 API 的調用與驗證。
2. **命令行工具指南：** 介紹如何使用 SDK 附帶的編譯後執行文件進行快速測試。

---

## 1. 場景化代碼示例

### 1.1 完整的 API 調用與響應驗證

以下代碼展示了如何利用 SDK 構建一個“創建用戶”請求，發送 HTTP 請求，並對平台返回的數據簽名進行安全驗證。

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. 初始化 SDK 與複用 Resty 客戶端 (前提：config.yaml 已配置)
	_, apiObj := common.Init()

	// 2. 生成請求參數和簽名 Header
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("構建請求失敗: ", err)
		return
	}

	// 3. 發送請求並自動驗證響應簽名
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("請求或驗證失敗: ", err)
		return
	}

	fmt.Println("✅ 請求成功並驗證通過！返回的 OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. 命令行工具使用指南

SDK 提供了快速測試各接口的命令行工具二進制文件（CLI）。

### 2.1 編譯執行文件

在 SDK 根目錄下執行 `make` 命令，系統將在 `bin` 目錄下生成各功能的二進制可執行文件。
* **Windows:** 生成 `.exe` 結尾的文件（如 `create_user.exe`）。
* **Mac/Linux:** 生成無後綴的文件（如 `create_user`）。

### 2.2 準備配置文件

運行工具前，請確保將配置好的 `config.yaml` 文件放置在 `bin` 目錄中。

### 2.3 測試各接口命令

#### 註冊新用戶
1. 在 `bin/config.yaml` 中修改 `UserOpenId` 字段。
2. 運行 `./create_user` (或雙擊 `create_user.exe`)。
3. 若該 OpenId 已註冊，工具將返回錯誤。

#### 錢包註冊
1. 在 `bin/config.yaml` 中指定 `UserOpenId` 和 `ChainID`。
2. 運行 `./create_wallet`。

#### 獲取充值地址
1. 在 `bin/config.yaml` 中指定 `UserOpenId` 和需要查詢的 `ChainIDs` (例如 "1,56")。
2. 運行 `./get_wallet_addresses`。

#### 申請提現
1. 在 `bin/config.yaml` 中指定：
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (唯一的訂單防重碼)
   * `CallbackUrl`
2. 運行 `./user_withdraw_by_open_id`。

#### 創建收銀台訂單
1. 在 `bin/config.yaml` 中指定 `OutOrderNo`, `TokenId`, `Quantity`, `NotifyUrl`。
2. 運行 `./new_order`。

#### 查詢錢包餘額
1. 在 `bin/config.yaml` 中指定 `WalletAddress`, `ContractAddress`, `WalletBalanceChainId`。
2. 運行 `./wallet_balance`。