# CryptoPay Go SDK

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Welcome to the CryptoPay Go SDK

The CryptoPay Go SDK is a professional cryptocurrency service SDK implemented in Golang, providing user registration, wallet generation, deposit callback notifications, withdrawals, and other functions.

It has been proven to be secure, stable, and easily extensible through long-term use.

Installation

```bash
go get github.com/zerospace-ai/pay-sdk-go
```
Note: Compiling requires Go 1.18+.

## 1 Demo Running Information

### 1.1 The config.yaml configuration used by the Demo in the example folder:

```yaml
ApiKey: "dkhl346iwonfw436"

ApiSecret: "11e15f2d36f4e61c8a46fd426ae2189bc0406c1e37c29d8b136f75268a1d4216"

PlatformPubKey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlxPMHGzg5ci2qZU9SxkTvTlsKKq7svS84QuDFYRNxkDTeLmQOmuPpxR a5v8Ujeg1M8QooeQlc3h33BQJVo1jqtxsoh5qo0Cj1J9gFn8TUF7xBfCX8VZ+Hqy+O2xSpFJHi9Uv8jCT4bnjzGw/53qCivJ7R0B Y6mEcv5twrSv+IPqWYw4R5DClmPGVZWf+Bn+s+nsuRAWe3fNODlmWGMsg1nTFGd9JMdLiniud/wa9i2xlIHticCCF0WSap4y3 Kp/QmsK1tl66NgpPi8BUfalQHI74snP/BtKSvcvt8+2OkksuI56x1g9PRrmPIFgUKzcgiom6Avd8/4EzoAYaKp12DwIDAQAB"

PlatformRiskPubKey: "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAq4mKl8NbMD1+oYG5XQHrn4j1RsENAzQmBfjaA40ema6fBSmuHZmrrXw jPIAWFLwRxG6aBcizZfU1EjtxDP5L1oWzQdYn+iG3hXr6T1OFMC7NFFsnEA1GAudsEWGEfq/xXHo8LztcD2eagrHHsR0uP91lIBj LGMcspT6JECApvYWe4M6HQ8CyV+CMt7fDkMhL9rilVt29sqb5istcvDrBhdjK8Ss46WrcVRDDgOoyo6BETI1dRuZgwgZnKxa5 zxG6DEQBh7/Aoa05MjLIquIPfnKzF7WkY5eE/HE+juRU3T0KYVMveVizzYED1hdJ/YKqICsRg+tn6C3Q79ZbLg/5pQIDAQAB"

RsaPrivateKey: "MIIEvQIBADANBgkqhkiG9w0BAQEFAAASCBKcwggSjAgEAAoIBAQDcJ5eadEaYPF/FhyjVdrN/s/Nw5UmNY781NcV9/7GZkLwNQrO7 9Uju2fyB5juZcpVZrW8eR/D8fUep3soGocOnYMO4561Xo/j7b/kdc0TLLl09H+GFisAQMky+IHcEp9D81MTEabnJ8IyXA+YYclVR6C WxILR2ZA4qxndos4uzv0vz1N0hY5tzN3E1J6DiaUM+e2QYL6GvCsSQeriQ87Xo8jNliqVDK/+eZU5GEyuHifhZ5dRC59LbhTTsgo7b udGTBW1mCnwZ2X0mdYfU/vQvfnAEGJYXvz/lJjW2Y/ABKhXq0DXUiZBUVRLHRmXlFaCeOpeejDxMiEjTXQfIjfbxAgMBAAECggEANM RYlfCgqwRtgA6xPtjAlS8wfjK2umjZ/4rv1w9eJgyGxjbilX5pCLa+yvO//Tt1iJUFOSDNJfdxIcoAai6Dkq6iusLszUDEAKAJ+YET UA/A0VG+4B8tgbRMqJVncXo3oSHuN5WPrlM1n0yT7fAZexRFVHseRfIdYytGm5XNOjuBzzaqghodKGw/IvdADw8eNYBMHUBFjgvjJS VKd99rsiKRRkzIVtBcTs7RxFXVdhB/PhvNxWugb3r3ihX52ho5uAoIHUGE7fzdIPO4iCYv9MRzKuSVwrAPIiJvAP7duwoC/INMFsqd CTuX1NdPS+58Ubgkdkmbf+8BK12dzK8EQQKBgQDtvJL2+j6Nd7rEf0+DK8ENjhWf3ktOoWX6dmepxDrJI1nHulpjZSWy5qmztMGiHF 5vU0e7ARZZeFFA4aGXcScTT+9ffuWJP6JA04OJryrWTTEF4qT2aTVHJDOYuOtr6pD7541QdpIVNq6TQu7zITaTBoTbHvNYFVFTdZg i8/Lb9QKBgQDtET57KgDBiedVCwIdODPTTsSGqWzZ5J2qE63CYN4nvkHbtzcPKADnOxq0yaYfrr1olYdqIzi0VWR4bPbvuP7D9jtGb xp7kDI6/ZkdcOIAG/0aFqJaqGGnnzgWQJXv8jn08Z90nrikDrYxBX8U1s/9fMdur//csZmjQPmdBOXtjQKBgQCvhHa4cv61sTypkBi jDi2klU7vzc2pis1gggR8uQxxrXC+XZ4YHfgcQeHudDg1OF6cME8YCHB4s7TBgxOrXHXt8ykWRviuQNXIqKBHiZTFzQ2xe6gw6HHWS SryySu+a9qIsGaLjk7B7LIstND3nYDOQZTatdoRIQP+6yXcQGD/9QKBgEyTIlyEP8REOC33JVKs4ciii8Z3mYp0Vx0lyB2eToQF554 B+03w/QGzzLeS3w8i0Vmj2x7Ei79sSczAXa8nUVuZAKKKpsI83IzDd57T5JxmbgXsQ7sG4qxTOLmvWP8tfd0J4xi3YCrV+bGx9c+UZ 5CYqo6tWPc/gsIB7d7zQxXNAoGAcxv32TAh+eRrVgIC0LMDXyKQ7pKt58RTjL8/SsSwavCKznvAp8S1pEde1/OjUfTiL42muJj1DghytwPIaam57X7/Ikgyz5PxgPzABCWv1BY0P4m37Cv8MYeqKv6e/OtjJs2O+r3GP12SI9RMP1trj7DLt5Z2TUmD5xeDEpdbpbw="
```

Where RsaPrivateKey is the merchant's RSA-generated private key.

Test API Endpoint: https://sandbox-api.privatex.io

Note: Contact the service provider to obtain the official API endpoint and parameters.

### 1.2 OpenId in the Demo:

```bash
openId := "HASH1756194148" // HASH + timestamp
```

This unique identifier verifies the partner's identity. Different partners have different configuration parameters. New partners can contact the platform.

## 2 Generating an RSA Key Pair for Merchants

### 2.1 Command to Generate an RSA Key Pair in Ubuntu:

2.1.1. Create a private key using the command line and keep it secure.

```bash
openssl genrsa -out rsa_private_key.pem 2048
```

2.1.2. Creating a Public Key from a Private Key

```bash
openssl rsa -in rsa_private_key.pem -out rsa_public_key.pem -pubout
```

Remove leading and trailing comments from the generated public key, remove newlines, and convert it into a single string. Submit this to the platform.

### 2.2 Generating an RSA Key Pair on Mac and Windows:

The steps are the same on Mac and Windows, as OpenSSL is a cross-platform tool.

The only difference is:

Mac comes with OpenSSL by default (but some versions use LibreSSL; you can install a newer version using brew install openssl ).

Windows does not have OpenSSL by default; you need to install OpenSSL for Windows or use it in Git Bash, WSL, or PowerShell.

### 2.3 Extract the public and private key strings from the key pair file:

```bash
grep -v '^-----' rsa_private_key.pem | tr -d '\n'; echo
```
Execute the above command to extract the private key string.

```bash
grep -v '^-----' rsa_public_key.pem | tr -d '\n'; echo
```
Execute the above command to extract the public key string.

The above command runs in Git Bash/WSL/Cygwin on Mac/Linux or Windows.

If running on Windows PowerShell:

```powershell
Write-Output ((Get-Content rsa_private_key.pem | Where-Object {$_ -notmatch "^-----"}) -join "")
```

```powershell
Write-Output ((Get-Content rsa_public_key.pem | Where-Object {$_ -notmatch "^-----"}) -join "")
```

## 3 Generate Executable Interface Commands

* 1. Execute the make command in the SDK root directory to generate binary executable files for each function command in the bin directory.

* 2. The file with the ".exe" suffix runs on 64-bit Windows machines; the file without the ".exe" suffix runs on Linux/Mac. For example, create_user.exe and create_user executable files.

* 3. Copy the configured config.yaml file to the bin directory.

## 4 SDK Instance Object

### 4.1 Required Configuration

1. Register your business name and obtain the `ApiKey` and `ApiSecret`;

2. Generate your own `RSA` key pair;

3. Prepare the platform's `RSA` public key;

### 4.2 Creating a Signature Object

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

### 4.3 Create and sign the request data.

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

### 4.4 Filling in and Initiating the Request

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

### 4.5 Verify parsing return data

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

## 5 API Description

### 5.1 Registering a New User

* Function: Creates a new user on the platform. Requires the user's unique ID, i.e., UserOpenId.

#### HTTP Request

Interface Name: create_user

Interface URL: 'https://sandbox-api.privatex.io/sdk/user/create'

Request Method: POST

#### Request Parameters

| Parameter Name | Required | Type | Description |
| :----- | :--- | :----- | :-------------------------------------------------------------------------- |
| OpenId | Yes | string | It is recommended to use a platform-standard prefix (e.g., HASH for partners) + the user's unique ID to form the user's unique OpenId. |

Request Example:

```bash
curl --location 'https://sandbox-api.privatex.io/sdk/user/create' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{ 
  "OpenId":"PT00001"
}'
```

```golang
    resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		Post(finalURL)
```

Request header field description:

key: Partner's key

sign: Generated using `md5(Partner's Secret + "Send data key=value, concatenated with &" + current timestamp in milliseconds, converted to a string)`

Example: md5("mysecret" + "OpenId=PT00001" + "1725076567682")

clientSign: Generated using `rsaWithMd5("Send data key=value, concatenated with &")`

timestamp: Current timestamp in milliseconds, converted to a string

#### Return Parameter Description

| Parameter Name | Type | Description |
| :---------- | :----- | :------------------ |
| code | int | Global Status Code |
| msg | string | Status Description |
| data.OpenId | string | Returns the user's unique OpenId |
| sign | string | Platform signature |

Return data example:

```json
{
    "sign": "HCL5RTmVsBJBQd8caxUjlATQJbnL/P4rtCdqta2g3ISFR/G4J/MRT3755UmmEbn08s4rlcx/j+MBjo6EesQL9akLG0rAeVfJNpg3InFb1UJOYgEZ+cTYpfnCIIOhOohaCQO9NTABRp89kt+Qnsfq4WuoobOsvLzwSOHq7vQvODrihiRItF+EhIM+TFRbONOyv9nk4Ft85BFF9wwANa0g8m2rAdL8msIpS/ywnvyEx3VSJpASlQo23hhQhxYvnTaHeHzL7cHAz8Zasmm7MTR6Ww+boRwO8X2hLPukaejM46Fx6+uXLWWeN8tQrwiwnLY5XHyjiG0QWFAzhNoasRkjlg==",
    "timestamp": "1725431229743",
    "data": {
        "OpenId": "PT00001"
    },
    "msg": "ok",
    "code": 1
}
```

#### Calling the Command

Go to the SDK's bin directory and modify the UserOpenId field in the config.yaml file there.

Run the create_user or create_user.exe executable file to register a new user on the platform.

If you attempt to register a new UserOpenId that has already been registered, an error will be returned.

### 5.2 Wallet Registration

* Function: Create a wallet account for the user corresponding to this blockchain network
* Prerequisite: The user with the specified OpenId has been successfully created

#### HTTP Request

Interface Name: create_wallet

Interface URL: 'https://sandbox-api.privatex.io/sdk/wallet/create'

Request Method: POST

#### Request Parameters

| Parameter Name | Required | Type | Description |
| :------ | :--- | :----- | :---------------- |
| ChainID | Yes | string | Chain ID |
| OpenId | Yes | string | User's unique OpenId |

ChainID List

| Coin Name | Full Name | Blockchain Browser Address | Chain ID |
| :------------ | :------------------ | :------------------------------ | :----------- |
| eth | eth | https://etherscan.io | 1 |
| trx | Tron | https://tronscan.io | 2 |
| btc | btc | https://blockchair.com/bitcoin | 3 |
| sol | solana | https://explorer.solana.com | 4 |
| xrp | xrp | https://xrpscan.com | 5 |
| eth_optimism | optimism | https://optimistic.etherscan.io | 10 |
| bnb | bnb | https://bscscan.com | 56 |
| matic_polygon | MATIC polygon chain | https://polygonscan.com | 137 |
| TON | Toncoin | https://tonscan.org/ | 15186 |

Example request:

```bash
curl --location 'https://sandbox-api.privatex.io/sdk/wallet/create' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{
  "OpenId":"PT00001",
  "ChainID":"1"
}'
```

#### Return Parameter Description

| Parameter Name | Type | Description |
| :----------- | :----- | :---------------- |
| code | int | Global Status Code |
| msg | string | Status Description |
| data.address | string | Wallet Address |
| data.UserId | string | User ID |
| data.ChainID | string | Public Chain ID |
| data.OpenId | string | User's unique OpenId |
| sign | string | Platform Signature |

Return Instance

```json
{
    "sign": "i24t857ix3027CPiuQ+getyC7u3pJHcL/m5NiPUQwmv5XkOEdrDnckoblGXIbdO2hgjpJDg47Lbq/YoKu+NiJHGJTwu10CAYDRzyiimBfLsP9yNdnFxJLTUEfOKPSXupJdceMZL8WXF4XkMpwHCrUqhekyM+aVLDHsfROKf3uP+zdjJ++9Z//3Xukg57OBvspYGPqpgIY5fOmALiXs3DgZTdXRYYN6MBRUR3NEd1lb4dSO1AjAGkahhIjGqwaeqSO6YAcfwoj9Be48QS9CurfVxZ9xM8FvbPzPsa2W8kHG7q+Cji4NTk243LJyrQ9QFRpTDUTo5JNrJ1vne/2js8kg==",
    "timestamp": "1725432397796",
    "data": {
        "address": "TUUYqqUsXA2iwfxhiYNfRTKTW3zXFwK3Xx",
        "UserId": 26178,
        "PartnerId": 87,
        "ChainID": 2,
        "OpenId": "PT00001"
    },
    "msg": "ok",
    "code": 1
}
```

#### Calling the Command

Go to the SDK's bin directory and specify the `UserOpenId` and `ChainID` fields in the `config.yaml` file.

Run the `create_wallet` or `create_wallet.exe` executable file to complete the user's wallet registration on the platform.

### 5.3 Get Deposit Address

* Function: Get the user's blockchain wallet deposit address

#### HTTP Request

Interface Name: get_wallet_addresses

Interface URL: 'https://sandbox-api.privatex.io/sdk/wallet/getWalletAddresses'

Request Method: POST

#### Request Parameters

| Parameter Name | Required | Type | Description |
|:--------| :--- | :----- |:---------------------------|
| OpenId | Yes | string | User's unique OpenId |
| ChainIDs | Yes | string | Multiple chain IDs, separated by commas. For example: "56,2" |

Request example:

```bash
curl --location 'https://sandbox-api.privatex.io/sdk/wallet/getWalletAddresses' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{
  "OpenId":"PT00001",
  "ChainIDs":"56,2"
}'
```

#### Return parameter description

| Parameter name | Type | Description |
|:-----------------------| :----- |:-------------|
| code | int | Global status code |
| msg | string | Status description |
| sign | string | Platform signature |
| timestamp | string | Millisecond timestamp |
| data.OpenId | string | User's unique OpenId |
| data.PartnerId | string | Partner ID |
| data.Addresses.chainID | string | Blockchain ID |
| data.Addresses.address | string | Blockchain wallet address |

Return Example

```json
{
  "sign" : "",
  "timestamp" : "1756375769709",
  "data" : {
    "Addresses" : [ {
      "address" : "0xb4e5e33edc071e8b3a83578477ef1e6134f52afe",
      "chainID" : 56
    }, {
      "address" : "TX6cVuF2yRyztceFHPcp6n2mhXSPknv265",
      "chainID" : 2
    } ],
    "PartnerId" : 133,
    "OpenId" : "HASH1756352640"
  },
  "msg" : "ok",
  "code" : 1
}
```

#### Calling the Command

Go to the SDK's bin directory and specify the `UserOpenId` and `ChainIDs` fields in `config.yaml`.

Run the `get_wallet_addresses` or `get_wallet_addresses.exe` executable file.

### 5.4 Withdrawals

#### Partner User Withdrawals

* Function: User withdrawal operation interface. Withdrawals must be transferred from the partner's account in the corresponding token withdrawal pool to the user's specified withdrawal wallet address. Partners can set a secure callback address to verify the legitimacy of the withdrawal. If verified as valid, the withdrawal can be completed directly from the merchant's fund pool wallet.

* The withdrawal transaction interface checks whether the default withdrawal hot wallet has sufficient withdrawal assets and handling fees.

* By default, the withdrawal interface uses a security verification code as the unique parameter requirement for withdrawal transactions. It is generally recommended to use the unique withdrawal order number of the business platform as the security code. Submitting a duplicate security verification code will result in an error.

* All withdrawal transaction requests will be matched against the risk control review rules configured on the channel platform. If the parameter request is valid, the transaction request will be accepted. Withdrawal transactions that meet the automatic review rules will be immediately submitted to the network transaction and the hash information of the submitted transaction will be returned (return field data). Withdrawal transaction requests that require secondary review on the channel will return (code=2). The withdrawal request does not need to be submitted again. The administrator must complete the secondary review on the channel platform. After the secondary review is completed, the transaction order will callback to notify the withdrawal transaction status change.

* Prerequisite: The corresponding currency's fund pool must have a sufficient amount of funds to withdraw (especially for ETH network token withdrawals, which require a certain amount of ETH transaction fee balance in the fund pool wallet).

* ⚠️ Note: **For blockchain withdrawals, please ensure that the pre-approval process is complete before calling the interface. Once a blockchain transaction is initiated, it cannot be revoked or returned.**

#### HTTP Request

Interface Name: user_withdraw_by_open_id

Interface URL: 'https://sandbox-api.privatex.io/sdk/partner/UserWithdrawByOpenID'

Request Method: POST

#### Request Parameters

| Parameter Name | Required | Type | Description |
| :------------ | :--- | :----- | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| OpenId | Yes | string | User's unique OpenId |
| TokenId | Yes | string | TokenId |
| Amount | Yes | float | The amount of currency the user is withdrawing, accurate to 2 decimal places |
| AddressTo | Yes | string | Destination wallet for withdrawal |
| CallBackUrl | No | string | Callback to notify the user of withdrawal progress. Optional; the partner's default callback URL will be used. |
| SafeCheckCode | No | string | The security verification code for user withdrawal transactions, typically the platform's unique withdrawal order number. This order number must be globally unique. Multiple withdrawal requests from a user require different order number parameters. The withdrawal transaction callback will return this information via the 'safecode' field. The platform can uniquely associate withdrawal requests based on the order number. |

Token Type

| TokenID | Value | Description |
| :------ | :------------ | :------------------------------- |
| 1 | ETH-ETH | ETH Network ETH |
| 2 | ETH-USDT | ETH Network USDT |
| 3 | TRON-TRX | TRON Network TRX |
| 4 | TRON-USDT | TRON Network Token: USDT |
| 5 | BNB-BNB | BNB Smart Chain Network BNB |
| 6 | BNB-USDT | BNB Smart Chain Network Token: USDT |
| 11 | Polygon-MATIC | Polygon Network Matic |
| 12 | Polygon-USDT | Polygon Network Token: USDT |
| 13 | Polygon-USDC | Polygon Network Token: USDC |
| 22 | BNB-USDC | BNB Smart Chain Network Token: USDC |
| 23 | BNB-DAI | BNB Smart Chain Network Token: DAI |
| 24 | ETH-USDC | ETH Network USDC |
| 25 | ETH-DAI | ETH Network DAI |
| 130 | Optimism-ETH | Optimism Network ETH |
| 131 | Optimism-WLD | Optimism Network Token: WLD |
| 132 | Optimism-USDT | Optimism Network Token: USDT |
| 100 | BTC-BTC | BTC Network BTC Main Chain Token |
| 200 | TON-TON | TON Network TON Main Chain Token |

Request Example:

```bash
curl --location 'https://sandbox-api.privatex.io/sdk/partner/UserWithdrawByOpenID' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{ 
  "OpenId": "PT00001", 
  "TokenId": "4", 
  "Amount": "0.02", 
  "AddressTo": "TQdL5yttJPTx7hJmBhGfo2LcE7AXLPtHSg", 
  "CallBackUrl": "http://xxxxxx/withdraw_callback", 
  "SafeCheckCode": "1000000000000000"
}'
```

#### Return Parameter Description

| Parameter Name | Type | Description |
| :-------- | :----- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| code | int | Status Code</br>0 Parameter error, duplicate order number, incorrect withdrawal address format, or insufficient withdrawal wallet fees. Detailed information can be found in msg.</br>1 The withdrawal transaction was successfully submitted and has been submitted to the blockchain network. The unique hash of the submitted transaction is contained in data.</br>2 The withdrawal transaction was successfully submitted and requires secondary channel review before the transaction can be completed. After the review is completed, the transaction information will be updated through a callback.</br>-1 The withdrawal transaction failed. You can resubmit the withdrawal request. |
| msg | string | Status Description |
| data | string | Transaction hash. If smart withdrawal is enabled, this field will be returned as an empty string. |
| sign | string | Platform signature |
| timestamp | string | Current timestamp in milliseconds converted to a string |

Return example

```json
{
    "sign": "D+VTPNiwGLzh9eIvkrscwS4UlGKzdnrBgB63RDG4HeobZT6FXqUwYCPgKojynKaxwm5PkmW0xhIASZ4asSCvnYfi0NSFehchZAtUnQIispxKcjsiudWsUznbkEIQ2h2TA/mbUZ1X9+wyh7QhNo6+RkxtgRyRpVb7ARG8pL14cdTAs OTtMLO0W1GO0M83VAv2ybBZNObncX9qy6tdwLQV/KYuNJYyMN0dL0nLKYHnj9Q4d3lEDM45AVJ0153/YIiIgcF BnOWhsQ9rVARcFeXeWd9KJ5OZpmxlxnhcJGcEUY2UDC4zKLZxtUet7CPAyehAMQ5plkpvRrR3Z6lA5zl6GQ==",
    "timestamp": "1725439986754",
    "data": "94f4c29eba73d53dcd3aa1b8cf90a98108d0acf82f38b97a4032dcdf7ff172e7",
    "msg": "ok",
    "code": 1
}
```

#### Calling the Command

Go to the SDK's bin directory and specify the `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode`, and `CallbackUrl` fields in `config.yaml`.

Run the `user_withdraw_by_open_id` or `user_withdraw_by_open_id.exe` executable file.

### 5.5 Withdrawal Order Secondary Review

* Function: Merchant withdrawal order risk control secondary review interface
* ⚠️ Note: **The platform assigns merchants a separate risk control RSA public key (different from the deposit/withdrawal callback notification public key)**
* Triggering Time: After the administrator configures the risk control callback URL parameters on the merchant side (system settings), the channel will add an additional risk control callback secondary review for each withdrawal transaction request. Only when the merchant-side risk control URL returns a correct verification pass code will the transaction be validly submitted.
* Technical Requirements: Merchant-side technical implementation and configuration of the secondary review callback interface are required.

#### HTTP Request

The platform sends a withdrawal review request to the merchant

> POST: `/withdrawal/order/check`

#### Request Parameters

| Parameter Name | Required | Type | Description |
| :-------- | :--- | :----- | :------------------------------------------------------------------------- |
| safeCode | No | string | Unique transaction ID submitted by the merchant, generally corresponding to the merchant's withdrawal order ID (SafeCheckCode for withdrawal transactions) |
| openId | Yes | string | User ID of the merchant submitting the withdrawal transaction |
| tokenId | Yes | string | Currency ID, based on the currency ID provided by the platform |
| toAddress | Yes | string | Withdrawal address |
| amount | Yes | string | Withdrawal amount |
| timestamp | Yes | int | Current timestamp |
| sign | Yes | string | Signature: Only the parameters in the data field are signed; the correctness of the signature must be verified using the platform's risk control RSA public key.

#### Return Parameter Description

| Parameter Name | Type | Description |
| :-------- | :----- | :----------------------------------------------------- |
| code | int | Verification result. 0 indicates a pass; other codes are invalid.
| timestamp | int | Current timestamp, in seconds.
| message | string | Return message.
| sign | string | Signature: The merchant's RSA private key signature for the data field in the response parameter.

### 5.6 Deposit and Withdrawal Callback Notifications

1. Deposit and withdrawal transactions will trigger multiple callback notifications. The transaction information and status of the last callback notification will be used.
2. The business side is required to return a valid callback message. The format is described in the return parameter description. A return code of 0 indicates that the callback message has been processed and no further notifications are required. Otherwise, the callback will continue to notify (initially every 2 seconds for 50 times, and then every 10 minutes thereafter) until a confirmation message with a code of 0 is returned.

Contact the service provider to set the callback URL.

> POST

* Function: Defines the callback message format that the platform uses to notify the application side of token transaction information (user withdrawal or deposit). This message is suitable for application-side event notifications regarding token transaction status (withdrawal or deposit). Applications can optionally support the callback notification interface based on their application functionality.

#### Request Parameters

| Parameter Name | Required | Type | Description |
| :----------- | :--- | :----- |:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| openid | yes | string | Channel user unique ID |
| totalvalue | yes | string | USDT value corresponding to the deposit or withdrawal transaction (calculated based on the market price at the time of the transaction) |
| status | yes | int | Transaction status:</br>1 The transaction is complete and has been successfully submitted to the blockchain network. Transaction details can be queried on-chain using the hash.</br>-1 The transaction has been submitted to the blockchain network, but the on-chain transaction failed. You can re-review it in Merchant Management --> Transaction Management --> [Submit Order Security Code]. The business platform does not need to process the status change and can simply wait for the channel to callback the new status notification.</br>-2 The withdrawal transaction application was rejected by the merchant backend. The withdrawal application has expired. It is recommended that the business platform return the user's withdrawal application after receiving the notification.</br>2 The withdrawal transaction has been submitted to the merchant management. Because it has triggered the configured currency security risk control requirements, the administrator needs to log in to Merchant Management --> Transaction Management --> Withdrawal Review to complete the withdrawal application processing.</br>3 During the withdrawal transaction blockchain processing, the business platform does not need to update the status change and can simply wait for the channel to receive a new status notification. </br>⛑️**Special Reminder: For withdrawal transaction callbacks received by the business platform, if status = -1, the callback will be ignored. After the administrator logs in to the management backend and resubmits the transaction, a new status notification will be pushed to the platform simultaneously. ** || type | yes | int | 1 for deposit transactions; 2 for withdrawal transactions |
| hash | yes | string | Transaction hash value |
| confirm | yes | int | Number of on-chain confirmations for the transaction |
| createdtime | yes | string | Creation time |
| from | yes | string | Transaction initiator's address |
| to | yes | string | Transaction receiving address |
| amount | yes | string | Transaction amount |
| chainid | yes | string | Transaction chain ID |
| tokenid | yes | string | Transaction tokenid |
| tokenaddress | yes | string | Transaction token contract address |
| safecode | yes | string | Valid for withdrawal orders, typically a unique withdrawal order ID | orderid |
| timestamp | yes | string | Transaction timestamp |
| tag | no | string | Optional, for XRP and EOS |
| sign | yes | string | Platform signature data **The recipient can use the platform public key to verify the authenticity of the data returned by the platform. It is strongly recommended that the recipient verify the validity of the signature** |

Callback Notification Example

```json
{
    "amount": "23.0000000000000000000",
    "chainid": "2",
    "confirm": "1",
    "createdtime": "1732105978000",
    "from": "TPQmWeYVUmW4ZP",
    "hash": "b180f4184be91e12124b01089",
    "safecode": "safecode00001",
    "safecode": "",
    "sign":"",
    "status": "1",
    "timestamp": "1732105988040",
    "to": "TWLd7av6Lumoz9XAUkS8mPG7R51UstVLux",
    "tokenaddress": "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
    "tokenid": "4",
    "totalvalue": "23.00000000",
    "type": "1"
}
```

## 6 Common Information

The following information is common to all interfaces. Each subsequent interface requires rules to generate verification information.

### 6.1 Request Verification Header

Request Header Definition

| Parameter Name | Constraints | Example | Description |
| :--------- | :-------- | :--------------------------------- | :----------------------------- |
| key | Length: 64 | ithujj3onrzbgw5t | Partner key |
| timestamp | Length: 32 | 1722586649000 | Timestamp of request initiation (unit: milliseconds) |
| sign | Length: 32 | 9e0ccfe3915e94bcc5bf7dd51ad4e8d9 | Partner secret signature |
| clientSign | Length: 512 | 9e0ccfe3915e94bcc5bfbBsC5EUxV6 ... | Partner RSA signature |

* `sign` Field Rules

1. Register the partner and obtain the key and secret.
2. Parse the request. Sort the JSON body by ASCII ascending order of the keys in the JSON, and concatenate the strings dataStr=key1=value1&key2=value2&key3=value3&...
3. Generate a timestamp (unit: milliseconds)
4. Encrypt and generate a sign:

Plaintext before encryption: strToHash = Secret+dataStr+timestamp

Perform MD5 encryption on the plaintext strToHash to generate the sign.

The specific code is the GenerateMD5Sign function.

5. Place the key, timestamp, and sign in the HTTP header.

* Detailed Explanation of the clientSign Signature Algorithm

1. Obtain the request parameters and format them to obtain a new formatted string.

2. Sign the data from step 1 with the RSA private key and save the signature result to a variable.

Generate a signature string for the following parameter array: `user_id = 1 coin = eth address = 0x038B8E7406dED2Be112B6c7E4681Df5316957cad amount = 10.001 trade_id = 20220131012030274786`

Sort each key in the array from a to z. If the first letter is the same, look at the second letter, and so on. After sorting, concatenate all array values ​​with the ampersand (&) character, for example, in $dataString:

`address=0x038B8E7406dED2Be112B6c7E4681Df5316957cad&amount=10.001&coin=eth&trade_id=20220131012030274786&user_id=1`

This string is the concatenated string.

Use the private key to sign the data using RSA-MD5.

The specific code is in the GenerateRSASignature function.

### 6.2 Public Information

| Name | Type | Example | Description |
| :--------- | :-------- | :--------------------------------- | :--------------------------------- |
| Global Status Code | integer | 1 | 1 indicates success. See Global Status Code for details. |
| Message | string | ok | Returns text information. |
| Data | json | {"OpenID":"HEX..."} | Returns specific data content. |
| Time | timeStamp | 1722587274000 | UTC time (without time zone, in milliseconds). |
| Signature | sign | 9e0ccfe3915e94bcc5bfbBsC5EUxV6 ... | The platform uses RSA to sign all data. |

## 7 Token Basic Information

| Mainchain Network | chain_id [mainchain ID] | token_id [unique ID] | token_address [contract address] | symbol[Token abbreviation] | decimals[Decimal places] |
| --------------- | ------------------ | ------------------ | ---------------------------------------------------------------- | ----------------- | ---------------- |
| Ethereum | 1 | 2 | 0xdac17f958d2ee523a2206206994597c13d831ec7 | USDT | 6 |
| | 1 | 140 | 0x6982508145454Ce325dDbE47a25d4ec3d2311933 | PEPE | 18 |
| | 1 | 141 | 0xb131f4A55907B10d1F0A50d8ab8FA09EC342cd74 | MEME | 18 |
| | 1 | 64 | 0xEd04915c23f00A313a544955524EB7DBD823143d | ACH | 8 |
| | 1 | 25 | 0x6B175474E89094C44Da98b954EedeAC495271d0F | DAI | 18 |
| | 1 | 24 | 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48 | USDC | 6 |
| | 1 | 142 | 0x163f8C2467924be0ae7B5347228CABF260318753 | WLD | 18 |
| | | 1 | 1 | ETH | 18 |
| Tron | 2 | 40 | THb4CqiFdwNHsWsQCs4JhzwjMWys4aqCbF | ETH | 18 |
| | 2 | 90 | TPYmHEhy5n8TCEfYGqW2rPxsghSfzghPDn | USDD | 18 |
| | 2 | 26 | TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8 | USDC | 6 |
| | 2 | 33 | TSkW873XMKiDCxGZrA4YH8KGeipLdC6Gyu | CVNT | 18 |
| | 2 | 3 | TRX | TRX | 6 |
| | 2 | 4 | TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t | USDT | 6 |
| Bitcoin | 3 | 100 | BTC | BTC | 8 |
| | 3 | 102 | SATS | SATS | 18 |
| | 3 | 103 | RATS | RATS | 18 |
| | 3 | 101 | ORDI | ORDI | 18 |
| Solana | 4 | 400 | Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB | USDT | 6 |
| | 4 | 401 | EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v | USDC | 6 |
| | 4 | 19 | SOL | SOL | 9 |
| | 4 | 410 | nQ1qgSpXWi71twnWPFjyfCtcbUXbVyQb64RfHKwRpKE | DAOT | 9 |
| XRP | 5 | 200 | XRP | XRP | 6 |
| DogeCoin | 9 | 300 | DOGE | DOGE | 8 |
| Optimistic | 10 | 131 | 0xdC6fF44d5d932Cbd77B52E5612Ba0529DC6226F1 | WLD | 18 |
| | 10 | 130 | ETH | ETH | 18 |
| | 10 | 133 | 0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85 | USDC | 6 |
| | 10 | 132 | 0x94b008aA00579c1307B0EF2c499aD98a8ce58e58 | USDT | 6 |
| Bnb Smart Chain | 56 | 62 | 0xc0be866ecc026957fc7160c1a45f2bee9870fd46 | ARK | 18 |
| | 56 | 561 | 0xbA2aE424d960c26247Dd6c32edC70B295c744C43 | DOGE | 8 |
| | 56 | 68 | 0x6FDcdfef7c496407cCb0cEC90f9C5Aaa1Cc8D888 | VET | 18 |
| | 56 | 63 | 0x8540f3D726Aed340Bc57Fd07a61b0ae2a9d5ECa9 | PUC | 18 |
| | 56 | 65 | 0xbc7d6b50616989655afd682fb42743507003056d | ACH | 8 |
| | 56 | 66 | 0xFE8bF5B8F5e4eb5f9BC2be16303f7dAB8CF56aA8 | BIBI | 18 |
| | 56 | 29 | 0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56 | BUSD | 18 |
| | 56 | 31 | 0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c | BTCB | 18 |
| | 56 | 30 | 0x2170Ed0880ac9A755fd29B2688956BD959F933F8 | ETH | 18 |
| | 56 | 6 | 0x55d398326f99059ff775485246999027b3197955 | USDT | 18 |
| | 56 | 23 | 0x1AF3F329e8BE154074D8769D1FFa4eE058B1DBc3 | DAI | 18 |
| | 56 | 22 | 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d | USDC | 18 |
| | 56 | 5 | BNB | BNB | 18 |
| Polygon | 137 | 12 | 0xc2132D05D31c914a87C6611C10748AEb04B58e8F | USDT | 6 |
| | 137 | 13 | 0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174 | USDC | 6 |
| | 137 | 11 | POL | MATIC | 18 |
| | 137 | 110 | 0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359 | USDC | 6 |
| CVN Chain | 2032 | 7 | CVN | CVN | 18 |
| | 2032 | 35 | 0x109B57A29eE6E9A93f33687F6CE553fB18D8EE78 | USDT | 6 |
| | 2032 | 51 | 0x6b94b0a2878c68811c1bd6cecc2b7cc44a9ed7ab | HPT | 8 |
| Merlin | 4200 | 500 | BTC | BTC | 18 |
| | 4200 | 501 | 0x5c46bFF4B38dc1EAE09C5BAc65872a1D8bc87378 | MERL | 18 |
| Base | 8453 | 801 | 0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913 | USDC | 6 |
| | 8453 | 802 | ETH | ETH | 18 |
| TON | 201 | 0: | 105e5589bc66db15f13c177a12f2cf3b94881da2f4b8e7922c58569176625eb5 | JETTON | 9 |
| 15186 | 202 | 0: | b113a994b5024a16719f69139328eb759596c38a25f59028b146fecdc3621dfe | USDT | 6 |
| | 15186 | 200 | TON | TON | 9 |
| Arbitrum One | 42161 | 122 | 0xaf88d065e77c8cC2239327C5EDb3A432268e5831 | USDC | 6 |
| | 42161 | 121 | 0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9 | USDT | 6 |
| | 42161 | 120 | ETH | ETH | 18 |
| | 42161 | 123 | 0x9fE175843Df9deCd99C78E72b2424C47D61Ad2bF | ATM | 18 |
| | 42161 | 124 | 0x58BDf739aE17d1C60C6FD3433E288E38B81C2853 | SAM | 18 |
| Avax Chain C | 43114 | 18 | 0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E | USDC | 6 |
| | 43114 | 17 | 0xc7198437980c041c805A1EDcbA50c1Ce5db95118 | USDT | 6 |
| | 43114 | 16 | AVAX | AVAX | 18 |
| NA Chain | 65143 | 600 | NAC | NAC | 9 |
| | 65143 | 601 | GAT | GAT | 9 |
| ODIN | 666666 | 80 | ODIN | ODIN | 18 |
| THOR | 868868 | 81 | THOR | THOR | 18 |

## 8 Contact Us

Telegram Account: [@ZeroSerivce](https://t.me/ZeroSerivce)

