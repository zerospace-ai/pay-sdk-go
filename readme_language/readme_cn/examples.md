# 示例 📝

本文档提供了 CryptoPay Go SDK 的使用示例，包括 Demo 运行、密钥生成和回调处理。

## 1 SDK 实例对象 🛠️

### 1.1 所需配置 ⚙️

1. 注册您的业务名称并获取 `ApiKey` 和 `ApiSecret`；

2. 生成您自己的 `RSA` 密钥对；

3. 准备平台的 `RSA` 公钥；

### 1.2 创建签名对象 🔏

1. 添加配置文件 `config.yaml`。

```yaml
# 配置业务信息
ApiKey: ""
ApiSecret: ""
# 平台公钥
PlatformPubKey: ""
# 用于阻塞平台的公钥
PlatformRiskPubKey: ""
# 您自己的私钥
RsaPrivateKey: ""
```

2. 加载配置文件并创建 API 对象。

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

### 1.3 创建并签名请求数据。 ✍️

以用户创建为例。

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

### 1.4 填充并发起请求 🚀

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

### 1.5 验证解析返回数据 ✅

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

## 2. 生成可执行接口命令

* 1. 在 SDK 根目录执行 make 命令，在 bin 目录生成每个功能命令的二进制可执行文件。

* 2. 带有 ".exe" 后缀的文件在 64 位 Windows 机器上运行；没有 ".exe" 后缀的文件在 Linux/Mac 上运行。例如，create_user.exe 和 create_user 可执行文件。

* 3. 将配置好的 config.yaml 文件复制到 bin 目录。

## 3. 调用命令 📞

### 3.1. 注册新用户 🆕


转到 SDK 的 bin 目录，并在其中的 config.yaml 文件中修改 UserOpenId 字段。

运行 create_user 或 create_user.exe 可执行文件，在平台上注册新用户。

如果尝试注册已注册的新 UserOpenId，将返回错误。


### 3.2. 钱包注册 💼

转到 SDK 的 bin 目录，并在 config.yaml 文件中指定 `UserOpenId` 和 `ChainID` 字段。

运行 `create_wallet` 或 `create_wallet.exe` 可执行文件，完成用户在平台上的钱包注册。

### 3.3. 获取充值地址 📍

转到 SDK 的 bin 目录，并在 config.yaml 中指定 `UserOpenId` 和 `ChainIDs` 字段。

运行 `get_wallet_addresses` 或 `get_wallet_addresses.exe` 可执行文件。

### 3.4. 提现 💸

转到 SDK 的 bin 目录，并在 config.yaml 中指定 `UserOpenId`、`TokenId`、`Amount`、`AddressTo`、`SafeCheckCode` 和 `CallbackUrl` 字段。

运行 `user_withdraw_by_open_id` 或 `user_withdraw_by_open_id.exe` 可执行文件。