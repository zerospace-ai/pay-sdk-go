# Ejemplos 📝

Este documento proporciona ejemplos de uso para el CryptoPay Go SDK, incluyendo la ejecución de Demo, generación de claves y manejo de callbacks.

## 1 Objeto de Instancia SDK 🛠️

### 1.1 Configuración Requerida ⚙️

1. Registre su nombre de negocio y obtenga el `ApiKey` y `ApiSecret`;

2. Genere su propio par de claves `RSA`;

3. Prepare la clave pública `RSA` de la plataforma;

### 1.2 Creando un Objeto de Firma 🔏

1. Agregue un archivo de configuración `config.yaml`.

```yaml
# Configurar información de negocio
ApiKey: ""
ApiSecret: ""
# Clave pública de la plataforma
PlatformPubKey: ""
# Clave pública para bloquear la plataforma
PlatformRiskPubKey: ""
# Su propia clave privada
RsaPrivateKey: ""
```

2. Cargue el archivo de configuración y cree el objeto API.

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

### 1.3 Crear y firmar los datos de solicitud. ✍️

Usemos la creación de usuario como ejemplo.

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

### 1.4 Rellenar e Iniciar la Solicitud 🚀

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

### 1.5 Verificar el análisis de los datos de retorno ✅

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

## 2. Generar Comandos de Interface Ejecutables

* 1. Ejecute el comando make en el directorio raíz del SDK para generar archivos binarios ejecutables para cada comando de función en el directorio bin.

* 2. El archivo con el sufijo ".exe" se ejecuta en máquinas Windows de 64 bits; el archivo sin el sufijo ".exe" se ejecuta en Linux/Mac. Por ejemplo, create_user.exe y create_user archivos ejecutables.

* 3. Copie el archivo config.yaml configurado al directorio bin.

## 3. Llamando al Comando 📞

### 3.1. Registrando un Nuevo Usuario 🆕


Vaya al directorio bin del SDK y modifique el campo UserOpenId en el archivo config.yaml allí.

Ejecute el archivo ejecutable create_user o create_user.exe para registrar un nuevo usuario en la plataforma.

Si intenta registrar un nuevo UserOpenId que ya ha sido registrado, se devolverá un error.


### 3.2. Registro de Billetera 💼

Vaya al directorio bin del SDK y especifique los campos `UserOpenId` y `ChainID` en el archivo `config.yaml`.

Ejecute el archivo ejecutable `create_wallet` o `create_wallet.exe` para completar el registro de la billetera del usuario en la plataforma.

### 3.3. Obtener Dirección de Depósito 📍

Vaya al directorio bin del SDK y especifique los campos `UserOpenId` y `ChainIDs` en `config.yaml`.

Ejecute el archivo ejecutable `get_wallet_addresses` o `get_wallet_addresses.exe`.

### 3.4. Retiros 💸

Vaya al directorio bin del SDK y especifique los campos `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode` y `CallbackUrl` en `config.yaml`.

Ejecute el archivo ejecutable `user_withdraw_by_open_id` o `user_withdraw_by_open_id.exe`.