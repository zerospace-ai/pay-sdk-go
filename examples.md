# Examples 📝

This document provides usage examples for the CryptoPay Go SDK, including Demo running, key generation, and callback handling.

## 1 SDK Instance Object 🛠️

### 1.1 Required Configuration ⚙️

1. Register your business name and obtain the `ApiKey` and `ApiSecret`;

2. Generate your own `RSA` key pair;

3. Prepare the platform's `RSA` public key;

### 1.2 Creating a Signature Object 🔏

1. Add a configuration file `config.yaml`.

```yaml
# Configure business information
ApiKey: ""
ApiSecret: ""
# Platform public key
PlatformPubKey: ""
# Public key for blocking the platform
PlatformRiskPubKey: ""
# Your own private key
RsaPrivateKey: ""
```

2. Load the configuration file and create the API object.

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

### 1.3 Create and sign the request data. ✍️

Let's use user creation as an example.

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

### 1.4 Filling in and Initiating the Request 🚀

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

### 1.5 Verify parsing return data ✅

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

## 2. Generate Executable Interface Commands

* 1. Execute the make command in the SDK root directory to generate binary executable files for each function command in the bin directory.

* 2. The file with the ".exe" suffix runs on 64-bit Windows machines; the file without the ".exe" suffix runs on Linux/Mac. For example, create_user.exe and create_user executable files.

* 3. Copy the configured config.yaml file to the bin directory.

## 3. Calling the Command 📞

### 3.1. Registering a New User 🆕


Go to the SDK's bin directory and modify the UserOpenId field in the config.yaml file there.

Run the create_user or create_user.exe executable file to register a new user on the platform.

If you attempt to register a new UserOpenId that has already been registered, an error will be returned.


### 3.2. Wallet Registration 💼

Go to the SDK's bin directory and specify the `UserOpenId` and `ChainID` fields in the `config.yaml` file.

Run the `create_wallet` or `create_wallet.exe` executable file to complete the user's wallet registration on the platform.

### 3.3. Get Deposit Address 📍

Go to the SDK's bin directory and specify the `UserOpenId` and `ChainIDs` fields in `config.yaml`.

Run the `get_wallet_addresses` or `get_wallet_addresses.exe` executable file.

### 3.4. Withdrawals 💸

Go to the SDK's bin directory and specify the `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode`, and `CallbackUrl` fields in `config.yaml`.

Run the `user_withdraw_by_open_id` or `user_withdraw_by_open_id.exe` executable file.
