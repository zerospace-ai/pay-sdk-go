# Örnekler 📝

Bu belge, CryptoPay Go SDK'sı için kullanım örnekleri sağlar, Demo çalıştırma, anahtar üretimi ve geri arama işleme dahil.

## 1 SDK Örnek Nesnesi 🛠️

### 1.1 Gerekli Yapılandırma ⚙️

1. İşletme adınızı kaydedin ve `ApiKey` ve `ApiSecret` elde edin;

2. Kendi `RSA` anahtar çiftinizi üretin;

3. Platformun `RSA` genel anahtarını hazırlayın;

### 1.2 İmza Nesnesi Oluşturma 🔏

1. Bir yapılandırma dosyası `config.yaml` ekleyin.

```yaml
# İşletme bilgilerini yapılandırın
ApiKey: ""
ApiSecret: ""
# Platform genel anahtarı
PlatformPubKey: ""
# Platformu engellemek için genel anahtar
PlatformRiskPubKey: ""
# Kendi özel anahtarınız
RsaPrivateKey: ""
```

2. Yapılandırma dosyasını yükleyin ve API nesnesini oluşturun.

```golang

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Yapılandırma yüklenemedi: %s", err))
	}
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

```

### 1.3 İstek Verilerini Oluşturma ve İmzalama ✍️

Kullanıcı oluşturmayı örnek olarak kullanalım.

```golang

  // ....
	openId := "HASH1756194148"

	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		logrus.Warnln("Hata: ", err)
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

### 1.4 İsteği Doldurma ve Başlatma 🚀

```golang
  // ....
	
	finalURL, err := url.JoinPath(api.DevNetEndpoint, api.PathCreateWallet)
	if err != nil {
		logrus.Warnln("Hata: ", err)
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

### 1.5 Dönüş Verilerini Doğrulama ve Ayrıştırma ✅

```golang

	rspCommon := response_define.ResponseCommon{}
	err = json.Unmarshal(body, &rspCommon)
	if err != nil {
		logrus.Warnln("Hata: ", err)
		return
	}
	logrus.Infoln("Yanıt: ", rspCommon)

	if rspCommon.Code != response_define.SUCCESS {
		logrus.Warnln("Yanıt başarısız Kodu", rspCommon.Code, "Msg", rspCommon.Msg)
		return
	}

	rspCreateUser := response_define.ResponseCreateUser{}
	err = json.Unmarshal(body, &rspCreateUser)
	if err != nil {
		logrus.Warnln("Hata: ", err)
		return
	}
	logrus.Infoln("ResponseCreateUser: ", rspCreateUser)

	mapObj := rsa_utils.ToStringMap(body)
	err = apiObj.VerifyRSAsignature(mapObj, rspCreateUser.Sign)
	if err != nil {
		logrus.Warnln("Hata: ", err)
		return
	}

```

## 2. Yürütülebilir Arayüz Komutları Oluşturma

* 1. SDK kök dizininde make komutunu çalıştırarak bin dizininde her fonksiyon komutu için ikili yürütülebilir dosyalar üretin.

* 2. ".exe" sonekli dosya 64-bit Windows makinelerinde çalışır; soneksiz dosya Linux/Mac'te çalışır. Örneğin, create_user.exe ve create_user yürütülebilir dosyaları.

* 3. Yapılandırılmış config.yaml dosyasını bin dizinine kopyalayın.

## 3. Komutu Çağırma 📞

### 3.1. Yeni Kullanıcı Kaydetme 🆕


SDK'nın bin dizinine gidin ve oradaki config.yaml dosyasındaki UserOpenId alanını değiştirin.

create_user veya create_user.exe yürütülebilir dosyasını çalıştırarak platformda yeni bir kullanıcı kaydedin.

Zaten kaydedilmiş bir UserOpenId'yi kaydetmeye çalışırsanız, hata döndürülür.


### 3.2. Cüzdan Kaydı 💼

SDK'nın bin dizinine gidin ve `config.yaml` dosyasındaki `UserOpenId` ve `ChainID` alanlarını belirtin.

`create_wallet` veya `create_wallet.exe` yürütülebilir dosyasını çalıştırarak platformda kullanıcının cüzdan kaydını tamamlayın.

### 3.3. Yatırma Adresini Alma 📍

SDK'nın bin dizinine gidin ve `config.yaml`deki `UserOpenId` ve `ChainIDs` alanlarını belirtin.

`get_wallet_addresses` veya `get_wallet_addresses.exe` yürütülebilir dosyasını çalıştırın.

### 3.4. Para Çekme 💸

SDK'nın bin dizinine gidin ve `config.yaml`deki `UserOpenId`, `TokenId`, `Amount`, `AddressTo`, `SafeCheckCode` ve `CallbackUrl` alanlarını belirtin.

`user_withdraw_by_open_id` veya `user_withdraw_by_open_id.exe` yürütülebilir dosyasını çalıştırın.