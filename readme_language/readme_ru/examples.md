# Примеры 📝

Этот документ предоставляет примеры использования CryptoPay Go SDK, включая запуск Demo, генерацию ключей и обработку обратных вызовов.

## 1 Объект экземпляра SDK 🛠️

### 1.1 Необходимая конфигурация ⚙️

1. Зарегистрируйте название вашего бизнеса и получите `ApiKey` и `ApiSecret`;

2. Сгенерируйте свою собственную пару ключей `RSA`;

3. Подготовьте публичный ключ `RSA` платформы;

### 1.2 Создание объекта подписи 🔏

1. Добавьте файл конфигурации `config.yaml`.

```yaml
# Настройка информации о бизнесе
ApiKey: ""
ApiSecret: ""
# Публичный ключ платформы
PlatformPubKey: ""
# Публичный ключ для блокировки платформы
PlatformRiskPubKey: ""
# Ваш собственный приватный ключ
RsaPrivateKey: ""
```

2. Загрузите файл конфигурации и создайте объект API.

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

### 1.3 Создание и подпись данных запроса. ✍️

Возьмем создание пользователя в качестве примера.

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

### 1.4 Заполнение и инициирование запроса 🚀

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

### 1.5 Проверка и разбор возвращаемых данных ✅

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

## 2. Генерация исполняемых команд интерфейса

* 1. Выполните команду make в корневом каталоге SDK, чтобы сгенерировать бинарные исполняемые файлы для каждой функциональной команды в каталоге bin.

* 2. Файл с суффиксом ".exe" запускается на 64-битных машинах Windows; файл без суффикса ".exe" запускается на Linux/Mac. Например, create_user.exe и create_user исполняемые файлы.

* 3. Скопируйте настроенный файл config.yaml в каталог bin.

## 3. Вызов команды 📞

### 3.1. Регистрация нового пользователя 🆕


Перейдите в каталог bin SDK и измените поле UserOpenId в файле config.yaml там.

Запустите исполняемый файл create_user или create_user.exe, чтобы зарегистрировать нового пользователя на платформе.

Если вы попытаетесь зарегистрировать новый UserOpenId, который уже зарегистрирован, будет возвращена ошибка.


### 3.2. Регистрация кошелька 💼

Перейдите в каталог bin SDK и укажите поля `UserOpenId` и `ChainID` в файле `config.yaml`.

Запустите исполняемый файл `create_wallet` или `create_wallet.exe`, чтобы завершить регистрацию кошелька пользователя на платформе.

### 3.3. Получение адреса пополнения 📍

Перейдите в каталог bin SDK и укажите поля `UserOpenId` и `ChainIDs` в `config.yaml`.

Запустите исполняемый файл `get_wallet_addresses` или `get_wallet_addresses.exe`.

### 3.4. Вывод средств 💸

Перейдите в каталог bin SDK и укажите поля `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode` и `CallbackUrl` в `config.yaml`.

Запустите исполняемый файл `user_withdraw_by_open_id` или `user_withdraw_by_open_id.exe`.