# 예시 📝

이 문서는 CryptoPay Go SDK의 사용 예시를 제공합니다. 데모 실행, 키 생성 및 콜백 처리 등을 포함합니다.

## 1 SDK 인스턴스 객체 🛠️

### 1.1 필수 구성 ⚙️

1. 비즈니스 이름을 등록하고 `ApiKey`와 `ApiSecret`을 얻습니다;

2. 자신의 `RSA` 키 쌍을 생성합니다;

3. 플랫폼의 `RSA` 공개 키를 준비합니다;

### 1.2 서명 객체 생성 🔏

1. 구성 파일 `config.yaml`을 추가합니다.

```yaml
# 비즈니스 정보 구성
ApiKey: ""
ApiSecret: ""
# 플랫폼 공개 키
PlatformPubKey: ""
# 플랫폼 차단용 공개 키
PlatformRiskPubKey: ""
# 자신의 개인 키
RsaPrivateKey: ""
```

2. 구성 파일을 로드하고 API 객체를 생성합니다.

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

### 1.3 요청 데이터 생성 및 서명 ✍️

사용자 생성을 예로 들어 보겠습니다.

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

### 1.4 요청 채우기 및 시작 🚀

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

### 1.5 반환 데이터 검증 및 파싱 ✅

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

## 2. 실행 가능한 인터페이스 명령 생성

* 1. SDK 루트 디렉토리에서 make 명령을 실행하여 bin 디렉토리에 각 기능 명령의 바이너리 실행 파일을 생성합니다.

* 2. ".exe" 접미사가 있는 파일은 64비트 Windows 머신에서 실행됩니다; 접미사가 없는 파일은 Linux/Mac에서 실행됩니다. 예: create_user.exe 및 create_user 실행 파일.

* 3. 구성된 config.yaml 파일을 bin 디렉토리로 복사합니다.

## 3. 명령 호출 📞

### 3.1. 새 사용자 등록 🆕


SDK의 bin 디렉토리로 이동하여 그곳의 config.yaml 파일에서 UserOpenId 필드를 수정합니다.

create_user 또는 create_user.exe 실행 파일을 실행하여 플랫폼에 새 사용자를 등록합니다.

이미 등록된 새 UserOpenId를 등록하려고 하면 오류가 반환됩니다.


### 3.2. 지갑 등록 💼

SDK의 bin 디렉토리로 이동하여 `config.yaml` 파일에서 `UserOpenId` 및 `ChainID` 필드를 지정합니다.

`create_wallet` 또는 `create_wallet.exe` 실행 파일을 실행하여 플랫폼에서 사용자 지갑 등록을 완료합니다.

### 3.3. 입금 주소 가져오기 📍

SDK의 bin 디렉토리로 이동하여 `config.yaml`에서 `UserOpenId` 및 `ChainIDs` 필드를 지정합니다.

`get_wallet_addresses` 또는 `get_wallet_addresses.exe` 실행 파일을 실행합니다.

### 3.4. 출금 💸

SDK의 bin 디렉토리로 이동하여 `config.yaml`에서 `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode` 및 `CallbackUrl` 필드를 지정합니다.

`user_withdraw_by_open_id` 또는 `user_withdraw_by_open_id.exe` 실행 파일을 실행합니다.