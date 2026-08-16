# 示例代码与工具

本文档分为两部分：
1. **场景化代码示例：** 演示在实际代码中如何处理 API 的调用与验证。
2. **命令行工具指南：** 介绍如何使用 SDK 附带的编译后执行文件进行快速测试。

---

## 1. 场景化代码示例

### 1.1 完整的 API 调用与响应验证

以下代码展示了如何利用 SDK 构建一个“创建用户”请求，发送 HTTP 请求，并对平台返回的数据签名进行安全验证。

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. 初始化 SDK 与复用 Resty 客户端 (前提：config.yaml 已配置)
	_, apiObj := common.Init()

	// 2. 生成请求参数和签名 Header
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("构建请求失败: ", err)
		return
	}

	// 3. 发送请求并自动验证响应签名
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("请求或验证失败: ", err)
		return
	}

	fmt.Println("✅ 请求成功并验证通过！返回的 OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. 命令行工具使用指南

SDK 提供了快速测试各接口的命令行工具二进制文件（CLI）。

### 2.1 编译执行文件

在 SDK 根目录下执行 `make` 命令，系统将在 `bin` 目录下生成各功能的二进制可执行文件。
* **Windows:** 生成 `.exe` 结尾的文件（如 `create_user.exe`）。
* **Mac/Linux:** 生成无后缀的文件（如 `create_user`）。

### 2.2 准备配置文件

运行工具前，请确保将配置好的 `config.yaml` 文件放置在 `bin` 目录中。

### 2.3 测试各接口命令

#### 注册新用户
1. 在 `bin/config.yaml` 中修改 `UserOpenId` 字段。
2. 运行 `./create_user` (或双击 `create_user.exe`)。
3. 若该 OpenId 已注册，工具将返回错误。

#### 钱包注册
1. 在 `bin/config.yaml` 中指定 `UserOpenId` 和 `ChainID`。
2. 运行 `./create_wallet`。

#### 获取充值地址
1. 在 `bin/config.yaml` 中指定 `UserOpenId` 和需要查询的 `ChainIDs` (例如 "1,56")。
2. 运行 `./get_wallet_addresses`。

#### 申请提现
1. 在 `bin/config.yaml` 中指定：
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (唯一的订单防重码)
   * `CallbackUrl`
2. 运行 `./user_withdraw_by_open_id`。

#### 创建收银台订单
1. 在 `bin/config.yaml` 中指定 `OutOrderNo`, `TokenId`, `Quantity`, `NotifyUrl`。
2. 运行 `./new_order`。

#### 查询钱包余额
1. 在 `bin/config.yaml` 中指定 `WalletAddress`, `ContractAddress`, `WalletBalanceChainId`。
2. 运行 `./wallet_balance`。