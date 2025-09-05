# 範例 📝

本文件提供 CryptoPay Go SDK 的使用範例，包括 Demo 運行、金鑰生成和回調處理。

## 1 SDK 實例物件 🛠️

### 1.1 所需配置 ⚙️

1. 註冊您的業務名稱並獲取 `ApiKey` 和 `ApiSecret`；

2. 生成您自己的 `RSA` 金鑰對；

3. 準備平台的 `RSA` 公鑰；

### 1.2 創建簽名物件 🔏

1. 添加配置文件 `config.yaml`。

```yaml
# 配置業務信息
ApiKey: ""
ApiSecret: ""
# 平台公鑰
PlatformPubKey: ""
# 用於封鎖平台的公鑰
PlatformRiskPubKey: ""
# 您自己的私鑰
RsaPrivateKey: ""
```

2. 加載配置文件並創建 API 物件。

```golang

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

```

### 1.3 創建並簽名請求數據。 ✍️

讓我們以用戶創建為例。

```golang

  // ....
	openId := "HASH1756194148"

	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

```

```golang
    dataStr := rsa_utils.ComposeParams(mapData)

	timestamp = strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign = s.GenerateMD5Sign(dataStr, timestamp)

	jStr, err := json.Marshal(&req)
	if err != nil {
		return nil, timestamp, sign, clientSign, err
	}

	reqMapObj := rsa_utils.ToStringMap(jStr)
	clientSign, err = s.GenerateRSASignature(reqMapObj)
```

### 1.4 填充並發起請求 🚀

```golang
  // ....
	
	finalURL, err := url.JoinPath(api.DevNetEndpoint, api.PathCreateWallet)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		Post(finalURL)

```

### 1.5 驗證解析返回數據 ✅

```golang

	rspCommon := response_define.ResponseCommon{}
	err = json.Unmarshal(body, &rspCommon)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}
	logrus.Infoln("Response: ", rspCommon)

	if rspCommon.Code != response_define.SUCCESS {
		logrus.Warnln("Response fail Code", rspCommon.Code, "Msg", rspCommon.Msg)
		return
	}

	rspCreateUser := response_define.ResponseCreateUser{}
	err = json.Unmarshal(body, &rspCreateUser)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}
	logrus.Infoln("ResponseCreateUser: ", rspCreateUser)

	mapObj := rsa_utils.ToStringMap(body)
	err = apiObj.VerifyRSAsignature(mapObj, rspCreateUser.Sign)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

```

## 2. 生成可執行接口命令

* 1. 在 SDK 根目錄中執行 make 命令，在 bin 目錄中生成每個功能命令的二進制可執行文件。

* 2. 帶有 ".exe" 後綴的文件在 64 位 Windows 機器上運行；不帶 ".exe" 後綴的文件在 Linux/Mac 上運行。例如，create_user.exe 和 create_user 可執行文件。

* 3. 將配置好的 config.yaml 文件複製到 bin 目錄。

## 3. 調用命令 📞

### 3.1. 註冊新用戶 🆕


前往 SDK 的 bin 目錄，並在其中的 config.yaml 文件中修改 UserOpenId 字段。

運行 create_user 或 create_user.exe 可執行文件，在平台上註冊新用戶。

如果您嘗試註冊已註冊的 UserOpenId，將返回錯誤。


### 3.2. 錢包註冊 💼

前往 SDK 的 bin 目錄，並在 config.yaml 文件中指定 `UserOpenId` 和 `ChainID` 字段。

運行 `create_wallet` 或 `create_wallet.exe` 可執行文件，在平台上完成用戶的錢包註冊。

### 3.3. 獲取充值地址 📍

前往 SDK 的 bin 目錄，並在 config.yaml 中指定 `UserOpenId` 和 `ChainIDs` 字段。

運行 `get_wallet_addresses` 或 `get_wallet_addresses.exe` 可執行文件。

### 3.4. 提現 💸

前往 SDK 的 bin 目錄，並在 config.yaml 中指定 `UserOpenId`、`TokenId`、`Amount`、`AddressTo`、`SafeCheckCode` 和 `CallbackUrl` 字段。

運行 `user_withdraw_by_open_id` 或 `user_withdraw_by_open_id.exe` 可執行文件。