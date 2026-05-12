# CryptoPay Go SDK

![Go Version](https://img.shields.io/badge/go-1.18+-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/chat-Telegram-blue?logo=telegram)](https://t.me/ZeroSerivce)

## Добро пожаловать в CryptoPay Go SDK

CryptoPay Go SDK — это профессиональный SDK для криптовалютных сервисов, реализованный на Golang. Он предоставляет такие функции, как регистрация пользователей, создание кошельков, уведомления об обратном вызове (callback) при пополнении и вывод средств.
Он широко используется и доказал свою безопасность, стабильность и простоту расширения.

## Установка

```bash
go get github.com/zerospace-ai/pay-sdk-go
```

> **Примечание:** Для компиляции требуется Go 1.18+.

## Быстрый старт

### 1. Подготовка конфигурации

Перед использованием SDK необходимо подготовить файл конфигурации `config.yaml`, который содержит информацию для аутентификации продавца и открытые/закрытые ключи:

```yaml
ApiKey: "your_api_key"
ApiSecret: "your_api_secret"
PlatformPubKey: "platform_public_key"
PlatformRiskPubKey: "platform_risk_public_key"
RsaPrivateKey: "your_rsa_private_key"
```

> **💡 Подсказка:** Подробную информацию о том, как сгенерировать собственную пару ключей RSA продавца (`RsaPrivateKey`), а также о подробных механизмах аутентификации и безопасности, читайте в разделе [Аутентификация и безопасность (authentication.md)](./authentication.md).

### 2. Инициализация SDK и отправка запроса

Ниже приведен полный пример, демонстрирующий инициализацию экземпляра SDK и вызов API "Создать нового пользователя":

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
)

func main() {
	// 1. Загрузка конфигурации
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 2. Создание экземпляра SDK
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	// 3. Вызов API: Создание запроса на нового пользователя
	openId := "PT00001" // Уникальный идентификатор пользователя
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Printf("Ошибка при создании запроса: %v\n", err)
		return
	}

	// (Выполнение сетевого запроса и синтаксический анализ ответа здесь опущены, полный код см. в examples.md)
	fmt.Printf("Запрос CreateUser успешно создан!\nBody: %s\n", string(reqBody))
	fmt.Printf("Заголовки подготовлены: timestamp=%s, sign=%s, clientSign=%s\n", timestamp, sign, clientSign)
}
```

## Ключевые концепции и навигация

Для более эффективного использования этого SDK мы рекомендуем прочитать остальные документы в следующем порядке:

1. **[Аутентификация и безопасность (authentication.md)](./authentication.md)**: Узнайте, как генерировать пары ключей RSA и как работает механизм проверки подписи между SDK и платформой.
2. **[Справочник по API (api-reference.md)](./api-reference.md)**: Содержит подробные инструкции для всех поддерживаемых конечных точек API (например, создание кошелька, вывод средств) и форматов вебхуков.
3. **[Примеры и инструменты (examples.md)](./examples.md)**: Посмотрите более сложные примеры кода для различных сценариев и инструкции по использовании встроенных инструментов CLI SDK.
4. **[Приложение (appendix.md)](./appendix.md)**: Статическая словарная информация, такая как поддерживаемые ChainID, типы токенов и адреса контрактов.

## Контакты

Если у вас есть какие-либо вопросы, обратитесь к поставщику услуг:  
Telegram: [@ZeroSerivce](https://t.me/ZeroSerivce)