# CryptoPay Go SDK

![Go Version](https://img.shields.io/badge/go-1.18+-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/chat-Telegram-blue?logo=telegram)](https://t.me/ZeroSerivce)

## Willkommen beim CryptoPay Go SDK

Das CryptoPay Go SDK ist ein professionelles Kryptowährungs-Service-SDK, das in Golang implementiert ist. Es bietet Funktionen wie Benutzerregistrierung, Wallet-Erstellung, Einzahlungs-Rückrufbenachrichtigungen und Auszahlungen.
Es ist weit verbreitet und hat sich als sicher, stabil und leicht erweiterbar erwiesen.

## Installation

```bash
go get github.com/zerospace-ai/pay-sdk-go
```

> **Hinweis:** Die Kompilierung erfordert Go 1.18+.

## Schnellstart

### 1. Konfiguration vorbereiten

Bevor Sie das SDK verwenden, müssen Sie die Konfigurationsdatei `config.yaml` vorbereiten, die die Authentifizierungsinformationen des Händlers und die öffentlichen/privaten Schlüssel enthält:

```yaml
ApiKey: "your_api_key"
ApiSecret: "your_api_secret"
PlatformPubKey: "platform_public_key"
PlatformRiskPubKey: "platform_risk_public_key"
RsaPrivateKey: "your_rsa_private_key"
```

> **💡 Tipp:** Für Details zur Generierung des eigenen RSA-Schlüsselpaars (`RsaPrivateKey`) und zu den detaillierten Authentifizierungs- und Sicherheitsmechanismen lesen Sie bitte [Authentifizierung und Sicherheit (authentication.md)](./authentication.md).

### 2. SDK initialisieren und Anfrage senden

Hier ist ein vollständiges Beispiel, das zeigt, wie die SDK-Instanz initialisiert und die API "Neuen Benutzer erstellen" aufgerufen wird:

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
)

func main() {
	// 1. Konfiguration laden
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 2. SDK-Instanz erstellen
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	// 3. API aufrufen: Anfrage für neuen Benutzer erstellen
	openId := "PT00001" // Eindeutige Kennung des Benutzers
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Printf("Fehler beim Erstellen der Anfrage: %v\n", err)
		return
	}

	// (Netzwerkanfrage und Antwort-Parsing werden hier weggelassen, vollständiger Code in examples.md)
	fmt.Printf("CreateUser-Anfrage erfolgreich erstellt!\nBody: %s\n", string(reqBody))
	fmt.Printf("Header vorbereitet: timestamp=%s, sign=%s, clientSign=%s\n", timestamp, sign, clientSign)
}
```

## Wichtige Konzepte und Navigation

Um dieses SDK optimal nutzen zu können, empfehlen wir, die restlichen Dokumente in dieser Reihenfolge zu lesen:

1. **[Authentifizierung und Sicherheit (authentication.md)](./authentication.md)**: Erfahren Sie, wie Sie RSA-Schlüsselpaare generieren und den Signatur-Verifizierungsmechanismus zwischen SDK und Plattform verstehen.
2. **[API-Referenz (api-reference.md)](./api-reference.md)**: Enthält detaillierte Anweisungen für alle unterstützten API-Endpunkte (z. B. Wallet-Erstellung, Auszahlung) und Webhook-Formate.
3. **[Beispiele und Tools (examples.md)](./examples.md)**: Zeigt komplexere szenariobasierte Codebeispiele und Anweisungen zur Verwendung der integrierten CLI-Tools.
4. **[Anhang (appendix.md)](./appendix.md)**: Statische Wörterbuchinformationen wie unterstützte ChainIDs, Token-Typen und Vertragsadressen.

## Kontakt

Wenn Sie Fragen haben, kontaktieren Sie bitte den Dienstanbieter:  
Telegram: [@ZeroSerivce](https://t.me/ZeroSerivce)