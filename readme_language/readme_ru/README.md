# CryptoPay Go SDK

![Go Version](https://img.shields.io/badge/go-1.18+-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/chat-Telegram-blue?logo=telegram)](https://t.me/ZeroSerivce)

## 🌟 Добро пожаловать в CryptoPay Go SDK

CryptoPay Go SDK - это профессиональный SDK для услуг криптовалюты, реализованный на Golang, предоставляющий регистрацию пользователей, генерацию кошельков, уведомления о депозитах, вывод средств и другие функции.

Он доказал свою безопасность, стабильность и легкость расширения через долгосрочное использование.

## ⚙️ Установка

```bash
go get github.com/zerospace-ai/pay-sdk-go
```

Примечание: Компиляция требует Go 1.18+ 🛠️.
## 🚀 Быстрый старт
### 1. ⚙️ config.yaml

```yaml
ApiKey: "your_api_key"
ApiSecret: "your_api_secret"
PlatformPubKey: "platform_public_key"
PlatformRiskPubKey: "platform_risk_public_key"
RsaPrivateKey: "your_rsa_private_key"
```

Описание полей:

• 🔑 ApiKey / ApiSecret:

Назначается платформой при регистрации учетной записи мерчанта, используется для аутентификации API-запросов ✅.

• 🛡️ PlatformPubKey / PlatformRiskPubKey:

Публичные ключи, предоставляемые платформой, используются для проверки данных или подписей колбэков, возвращаемых платформой, обеспечивая надежность источников информации. PlatformRiskPubKey в основном используется для проверки событий, связанных с контролем рисков ⚠️.

• 🗝️ RsaPrivateKey:

RSA-приватный ключ, сгенерированный мерчантом, используется для подписи запросов, обеспечивая отсутствие подделки содержимого запроса. Важное примечание: Приватный ключ должен храниться в секрете 🔒, не раскрывайте его 🚫.

### 2. Генерация пары ключей RSA 🔐

Использование пары ключей RSA для подписи запросов обеспечивает безопасность данных. Ниже описано, как сгенерировать пару ключей и извлечь строки ключей на разных операционных системах.

#### 2.1 Генерация пары ключей с помощью OpenSSL

```bash
# Генерация 2048-битного приватного ключа
openssl genrsa -out rsa_private_key.pem 2048

# Генерация публичного ключа из приватного ключа
openssl rsa -in rsa_private_key.pem -out rsa_public_key.pem -pubout
```

> 💡 Совет: Сгенерированный публичный ключ нужно удалить начало и конец -----BEGIN PUBLIC KEY----- / -----END PUBLIC KEY-----, удалить разрывы строк, преобразовать в однострочную строку и отправить на платформу.
> 
> Извлечь строки ключей и отправить публичный ключ на платформу 📤.
>
>Команды для генерации пары ключей RSA на Mac и Windows такие же, как на Linux.

#### 2.2 Извлечение строк ключей 🔑

На Mac/Linux или Git Bash/WSL/Cygwin:

```bash
# Извлечение строки приватного ключа
grep -v '^-----' rsa_private_key.pem | tr -d '\n'; echo

# Извлечение строки публичного ключа
grep -v '^-----' rsa_public_key.pem | tr -d '\n'; echo
```

Windows

PowerShell извлечение строк приватного и публичного ключей:

```powershell
# Приватный ключ
Write-Output ((Get-Content rsa_private_key.pem | Where-Object {$_ -notmatch "^-----"}) -join "")

# Публичный ключ
Write-Output ((Get-Content rsa_public_key.pem | Where-Object {$_ -notmatch "^-----"}) -join "")
```

> ⚠️ Примечание: Сгенерированный приватный ключ должен храниться в безопасности и не утечь.


### 🛠️ 3. Создание экземпляра SDK

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

## 🔑 Ключевые концепции

- 🆔 **OpenId**: Уникальный идентификатор пользователя, например "HASH1756194148".
- 🔐 **RSA Key**: Используется для подписи и проверки запросов для обеспечения безопасности данных.
- ✍️ **API Signature**: Используйте алгоритмы MD5 и RSA для подписи запросов, обеспечивая отсутствие подделки.

Для подробных описаний API, пожалуйста, обратитесь к [🧩 api-reference.md](./api-reference.md) и [🧩 examples.md](./examples.md).

Для Аутентификации и Безопасности, пожалуйста, обратитесь к [🧩 authentication.md](./authentication.md)

## 📎 Приложение

Для более подробных ссылок, пожалуйста, проверьте документ [Приложение](./appendix.md), который включает следующее содержание:

- [🧩 Список ChainID](./appendix.md#-список-chainid)
- [🏷️ Типы токенов](./appendix.md#-тип-токена)
- [🌐 Общая информация](./appendix.md#-общая-информация)
- [🔰 Базовая информация о токенах](./appendix.md#-базовая-информация-о-токене)

> 💡 **Совет**: Приложение предоставляет информацию о поддерживаемых цепях, типах токенов и общих данных токенов, облегчая интеграцию и использование SDK разработчиками.

## 📞 Контакты

Если у вас есть вопросы, пожалуйста, свяжитесь с поставщиком услуг.  
💬 Telegram: [@ZeroSerivce](https://t.me/ZeroSerivce)