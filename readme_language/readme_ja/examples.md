# 例 📝

このドキュメントは、CryptoPay Go SDKの使用例を提供します。Demoの実行、キーの生成、およびコールバックの処理を含みます。

## 1 SDKインスタンスオブジェクト 🛠️

### 1.1 必要な設定 ⚙️

1. ビジネス名を登録し、`ApiKey` と `ApiSecret` を取得します；

2. 独自の `RSA` キーペアを生成します；

3. プラットフォームの `RSA` 公開鍵を準備します；

### 1.2 署名オブジェクトの作成 🔏

1. 設定ファイル `config.yaml` を追加します。

```yaml
# ビジネス情報を設定
ApiKey: ""
ApiSecret: ""
# プラットフォーム公開鍵
PlatformPubKey: ""
# プラットフォームのブロック公開鍵
PlatformRiskPubKey: ""
# 独自の秘密鍵
RsaPrivateKey: ""
```

2. 設定ファイルをロードし、APIオブジェクトを作成します。

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

### 1.3 リクエストデータの作成と署名 ✍️

ユーザー作成を例にします。

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

### 1.4 リクエストの入力と開始 🚀

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

### 1.5 返却データの検証と解析 ✅

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

## 2. 実行可能なインターフェースコマンドの生成

* 1. SDKルートディレクトリでmakeコマンドを実行し、binディレクトリに各機能コマンドのバイナリ実行可能ファイルを生成します。

* 2. ".exe" サフィックスのファイルは64ビットWindowsマシンで実行されます；サフィックスなしのファイルはLinux/Macで実行されます。例えば、create_user.exe と create_user 実行可能ファイル。

* 3. 設定されたconfig.yamlファイルをbinディレクトリにコピーします。

## 3. コマンドの呼び出し 📞

### 3.1. 新規ユーザーの登録 🆕


SDKのbinディレクトリに移動し、そこにあるconfig.yamlファイルのUserOpenIdフィールドを変更します。

create_user または create_user.exe 実行可能ファイルを実行して、プラットフォームに新規ユーザーを登録します。

すでに登録されている新規UserOpenIdを登録しようとすると、エラーが返されます。


### 3.2. ウォレット登録 💼

SDKのbinディレクトリに移動し、`config.yaml` ファイルで `UserOpenId` と `ChainID` フィールドを指定します。

`create_wallet` または `create_wallet.exe` 実行可能ファイルを実行して、プラットフォームでのユーザーのウォレット登録を完了します。

### 3.3. 入金アドレスの取得 📍

SDKのbinディレクトリに移動し、`config.yaml` で `UserOpenId` と `ChainIDs` フィールドを指定します。

`get_wallet_addresses` または `get_wallet_addresses.exe` 実行可能ファイルを実行します。

### 3.4. 出金 💸

SDKのbinディレクトリに移動し、`config.yaml` で `UserOpenId`、`TokenId`、`Amount`、`AddressTo`、`SafeCheckCode`、および `CallbackUrl` フィールドを指定します。

`user_withdraw_by_open_id` または `user_withdraw_by_open_id.exe` 実行可能ファイルを実行します。