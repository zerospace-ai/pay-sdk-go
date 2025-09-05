# CryptoPay Go SDK

![Go Sürümü](https://img.shields.io/badge/go-1.18+-blue.svg)
[![Lisans: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/chat-Telegram-blue?logo=telegram)](https://t.me/ZeroSerivce)

## 🌟 CryptoPay Go SDK'ya Hoş Geldiniz

CryptoPay Go SDK, Golang ile uygulanan profesyonel bir kripto para hizmeti SDK'sıdır ve kullanıcı kaydı, cüzdan oluşturma, para yatırma geri çağrı bildirimleri, para çekme ve diğer işlevleri sağlar.

Uzun süreli kullanım ile güvenli, istikrarlı ve genişletilmesi kolay olduğu kanıtlanmıştır.

## ⚙️ Kurulum

```bash
go get github.com/zerospace-ai/pay-sdk-go
```

Not: Derleme için Go 1.18+ gereklidir 🛠️.
## 🚀 Hızlı Başlangıç
### 1. ⚙️ config.yaml

```yaml
ApiKey: "your_api_key"
ApiSecret: "your_api_secret"
PlatformPubKey: "platform_public_key"
PlatformRiskPubKey: "platform_risk_public_key"
RsaPrivateKey: "your_rsa_private_key"
```

Alan açıklamaları:

• 🔑 ApiKey / ApiSecret:

Tüccar hesabı kaydederken platform tarafından atanır, API isteği kimlik doğrulaması için kullanılır ✅.

• 🛡️ PlatformPubKey / PlatformRiskPubKey:

Platform tarafından sağlanan genel anahtarlar, platform tarafından döndürülen verileri veya geri çağrı imzalarını doğrulamak için kullanılır, bilgi kaynaklarının güvenilir olmasını sağlar. PlatformRiskPubKey esas olarak risk kontrolü ile ilgili olay doğrulaması için kullanılır ⚠️.

• 🗝️ RsaPrivateKey:

Tüccar tarafından oluşturulan RSA özel anahtarı, istekleri imzalamak için kullanılır, istek içeriğinin değiştirilmediğini sağlar. Önemli not: Özel anahtar gizli tutulmalıdır 🔒, ifşa etmeyin 🚫.

### 2. RSA Anahtar Çifti Oluşturma 🔐

İstekleri imzalamak için RSA anahtar çifti kullanarak veri güvenliğini sağlar. Aşağıda farklı işletim sistemlerinde anahtar çifti oluşturma ve anahtar dizelerini çıkarma açıklanmaktadır.

#### 2.1 OpenSSL Kullanarak Anahtar Çifti Oluşturma

```bash
# 2048-bit özel anahtar oluştur
openssl genrsa -out rsa_private_key.pem 2048

# Özel anahtardan genel anahtar oluştur
openssl rsa -in rsa_private_key.pem -out rsa_public_key.pem -pubout
```

> 💡 İpucu: Oluşturulan genel anahtar, başlangıç ve bitiş -----BEGIN PUBLIC KEY----- / -----END PUBLIC KEY----- kaldırılmalı, satır sonları kaldırılmalı, tek satır dizesine dönüştürülmeli ve platforma gönderilmelidir.
> 
> Anahtar dizelerini çıkarın ve genel anahtarı platforma gönderin 📤.
>
>Mac ve Windows'ta RSA anahtar çiftlerini oluşturma komutları Linux ile aynıdır.

#### 2.2 Anahtar Dizelerini Çıkarma 🔑

Mac/Linux veya Git Bash/WSL/Cygwin'da:

```bash
# Özel anahtar dizesini çıkar
grep -v '^-----' rsa_private_key.pem | tr -d '\n'; echo

# Genel anahtar dizesini çıkar
grep -v '^-----' rsa_public_key.pem | tr -d '\n'; echo
```

Windows

PowerShell ile özel ve genel anahtar dizelerini çıkarma:

```powershell
# Özel anahtar
Write-Output ((Get-Content rsa_private_key.pem | Where-Object {$_ -notmatch "^-----"}) -join "")

# Genel anahtar
Write-Output ((Get-Content rsa_public_key.pem | Where-Object {$_ -notmatch "^-----"}) -join "")
```

> ⚠️ Not: Oluşturulan özel anahtar güvenli tutulmalı ve sızdırılmamalıdır.


### 🛠️ 3. SDK Örneği Oluşturma

```golang
import "github.com/zerospace-ai/pay-sdk-go/api"
import "github.com/spf13/viper"

    viper.SetConfigFile("config.yaml")
    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }
    apiObj := api.NewSDK(api.SDKConfig{
        ApiKey:             viper.GetString("ApiKey"),
        ApiSecret:          viper.GetString("ApiSecret"),
        PlatformPubKey:     viper.GetString("PlatformPubKey"),
        PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
        RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
    })
```

## 🔑 Ana Kavramlar

- 🆔 **OpenId**: Kullanıcının benzersiz tanımlayıcısı, örneğin "HASH1756194148".
- 🔐 **RSA Anahtarı**: İstekleri imzalamak ve doğrulamak için kullanılır, veri güvenliğini sağlar.
- ✍️ **API İmzası**: İstekleri imzalamak için MD5 ve RSA algoritmalarını kullanır, değiştirilmediğini sağlar.

Ayrıntılı API açıklamaları için lütfen [🧩 api-reference.md](./api-reference.md) ve [🧩 examples.md](./examples.md) dosyalarına bakın.

Kimlik Doğrulama ve Güvenlik için lütfen [🧩 authentication.md](./authentication.md) dosyasına bakın

## 📎 Ek

Daha ayrıntılı referanslar için lütfen [Ek](./appendix.md) belgesine bakın, aşağıdaki içerikleri içerir:

- [🧩 ChainID Listesi](./appendix.md#-chainid-listesi)
- [🏷️ Token Türleri](./appendix.md#-token-türü)
- [🌐 Genel Bilgiler](./appendix.md#-genel-bilgiler)
- [🔰 Token Temel Bilgileri](./appendix.md#-token-temel-bilgileri)

> 💡 **İpucu**: Ek, desteklenen zincir bilgilerini, token türlerini ve genel token verilerini sağlar, geliştiricilerin SDK'yi entegre etmesini ve kullanmasını kolaylaştırır.

## 📞 İletişim

Herhangi bir sorunuz varsa, lütfen hizmet sağlayıcısıyla iletişime geçin.  
💬 Telegram: [@ZeroSerivce](https://t.me/ZeroSerivce)