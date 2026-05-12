# CryptoPay Go SDK

![Go Version](https://img.shields.io/badge/go-1.18+-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/chat-Telegram-blue?logo=telegram)](https://t.me/ZeroSerivce)

## CryptoPay Go SDK'ya Hoş Geldiniz

CryptoPay Go SDK, Golang ile uygulanan ve kullanıcı kaydı, cüzdan oluşturma, para yatırma geri çağırma (callback) bildirimleri ve para çekme gibi işlevler sağlayan profesyonel bir kripto para hizmeti SDK'sıdır.
Yaygın olarak kullanılmış olup güvenli, kararlı ve kolayca genişletilebilir olduğu kanıtlanmıştır.

## Kurulum

```bash
go get github.com/zerospace-ai/pay-sdk-go
```

> **Not:** Derleme için Go 1.18+ gerekir.

## Hızlı Başlangıç

### 1. Yapılandırmayı Hazırlama

SDK'yı kullanmadan önce, tüccarın kimlik doğrulama bilgilerini ve genel/özel anahtarlarını içeren `config.yaml` yapılandırma dosyasını hazırlamanız gerekir:

```yaml
ApiKey: "your_api_key"
ApiSecret: "your_api_secret"
PlatformPubKey: "platform_public_key"
PlatformRiskPubKey: "platform_risk_public_key"
RsaPrivateKey: "your_rsa_private_key"
```

> **💡 İpucu:** Tüccarın kendi RSA anahtar çiftini (`RsaPrivateKey`) nasıl oluşturacağı ve ayrıntılı kimlik doğrulama ile güvenlik mekanizmaları hakkında ayrıntılar için lütfen [Kimlik Doğrulama ve Güvenlik (authentication.md)](./authentication.md) bölümünü okuyun.

### 2. SDK'yı Başlatma ve İstek Gönderme

Aşağıda, SDK örneğini nasıl başlatacağınızı ve "Yeni Kullanıcı Oluştur" API'sini nasıl çağıracağınızı gösteren eksiksiz bir örnek bulunmaktadır:

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
)

func main() {
	// 1. Yapılandırmayı yükle
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 2. SDK örneği oluştur
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	// 3. API Çağrısı: Yeni kullanıcı oluşturma isteği derle
	openId := "PT00001" // Kullanıcının benzersiz kimliği
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Printf("İstek derleme başarısız oldu: %v\n", err)
		return
	}

	// (Ağ isteği yürütme ve yanıt ayrıştırma burada atlanmıştır, tam kod için lütfen examples.md'ye bakın)
	fmt.Printf("CreateUser isteği başarıyla oluşturuldu!\nBody: %s\n", string(reqBody))
	fmt.Printf("Header'lar hazırlandı: timestamp=%s, sign=%s, clientSign=%s\n", timestamp, sign, clientSign)
}
```

## Temel Kavramlar ve Gezinme

Bu SDK'yı daha iyi kullanabilmek için geri kalan belgeleri aşağıdaki sırayla okumanızı öneririz:

1. **[Kimlik Doğrulama ve Güvenlik (authentication.md)](./authentication.md)**: RSA anahtar çiftlerinin nasıl oluşturulacağını ve SDK ile platform arasındaki imza doğrulama mekanizmasını öğrenin.
2. **[API Referansı (api-reference.md)](./api-reference.md)**: Desteklenen tüm API uç noktaları (örneğin, cüzdan oluşturma, para çekme) ve Webhook formatları için ayrıntılı talimatlar içerir.
3. **[Örnekler ve Araçlar (examples.md)](./examples.md)**: Daha karmaşık senaryo tabanlı kod örneklerini ve SDK'nın yerleşik CLI araçlarının kullanım talimatlarını görün.
4. **[Ek (appendix.md)](./appendix.md)**: Desteklenen ChainID'ler, token türleri ve sözleşme adresleri gibi statik sözlük bilgileri.

## İletişim

Herhangi bir sorunuz varsa, lütfen servis sağlayıcısıyla iletişime geçin:  
Telegram: [@ZeroSerivce](https://t.me/ZeroSerivce)