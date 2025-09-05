# Beispiele 📝

Dieses Dokument bietet Nutzungsbeispiele für das CryptoPay Go SDK, einschließlich Demo-Ausführung, Schlüsselerzeugung und Callback-Handhabung.

## 1 SDK-Instanzobjekt 🛠️

### 1.1 Erforderliche Konfiguration ⚙️

1. Registrieren Sie Ihren Geschäftsnamen und erhalten Sie den `ApiKey` und `ApiSecret`;

2. Generieren Sie Ihr eigenes `RSA`-Schlüsselpaar;

3. Bereiten Sie den `RSA`-öffentlichen Schlüssel der Plattform vor;

### 1.2 Erstellen eines Signaturobjekts 🔏

1. Fügen Sie eine Konfigurationsdatei `config.yaml` hinzu.

```yaml
# Konfigurieren Sie Geschäftsinformationen
ApiKey: ""
ApiSecret: ""
# Plattform-öffentlicher Schlüssel
PlatformPubKey: ""
# Öffentlicher Schlüssel zum Blockieren der Plattform
PlatformRiskPubKey: ""
# Ihr eigener privater Schlüssel
RsaPrivateKey: ""
```

2. Laden Sie die Konfigurationsdatei und erstellen Sie das API-Objekt.

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

### 1.3 Erstellen und Signieren der Anfragedaten. ✍️

Nehmen wir die Benutzererstellung als Beispiel.

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

### 1.4 Ausfüllen und Initiieren der Anfrage 🚀

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

### 1.5 Überprüfen und Parsen der Rückgabedaten ✅

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

## 2. Generieren ausführbarer Schnittstellenbefehle

* 1. Führen Sie den make-Befehl im SDK-Stammverzeichnis aus, um binäre ausführbare Dateien für jeden Funktionsbefehl im bin-Verzeichnis zu generieren.

* 2. Die Datei mit dem ".exe"-Suffix läuft auf 64-Bit-Windows-Maschinen; die Datei ohne das ".exe"-Suffix läuft auf Linux/Mac. Zum Beispiel create_user.exe und create_user ausführbare Dateien.

* 3. Kopieren Sie die konfigurierte config.yaml-Datei in das bin-Verzeichnis.

## 3. Aufrufen des Befehls 📞

### 3.1. Registrieren eines neuen Benutzers 🆕


Gehen Sie zum bin-Verzeichnis des SDK und ändern Sie das UserOpenId-Feld in der config.yaml-Datei dort.

Führen Sie die create_user oder create_user.exe ausführbare Datei aus, um einen neuen Benutzer auf der Plattform zu registrieren.

Wenn Sie versuchen, eine neue UserOpenId zu registrieren, die bereits registriert wurde, wird ein Fehler zurückgegeben.


### 3.2. Wallet-Registrierung 💼

Gehen Sie zum bin-Verzeichnis des SDK und geben Sie die `UserOpenId` und `ChainID`-Felder in der `config.yaml`-Datei an.

Führen Sie die `create_wallet` oder `create_wallet.exe` ausführbare Datei aus, um die Wallet-Registrierung des Benutzers auf der Plattform abzuschließen.

### 3.3. Einzahlungsadresse abrufen 📍

Gehen Sie zum bin-Verzeichnis des SDK und geben Sie die `UserOpenId` und `ChainIDs`-Felder in `config.yaml` an.

Führen Sie die `get_wallet_addresses` oder `get_wallet_addresses.exe` ausführbare Datei aus.

### 3.4. Auszahlungen 💸

Gehen Sie zum bin-Verzeichnis des SDK und geben Sie die `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode` und `CallbackUrl`-Felder in `config.yaml` an.

Führen Sie die `user_withdraw_by_open_id` oder `user_withdraw_by_open_id.exe` ausführbare Datei aus.