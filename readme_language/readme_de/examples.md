# Beispiele und Tools

Dieses Dokument ist in zwei Teile gegliedert:
1. **Szenariobasierte Codebeispiele:** Demonstriert, wie API-Aufrufe und Verifizierungen im praktischen Code gehandhabt werden.
2. **CLI-Tools-Handbuch:** Erklärt, wie die mit dem SDK kompilierten ausführbaren Dateien für schnelle Tests verwendet werden.

---

## 1. Szenariobasierte Codebeispiele

### 1.1 Vollständiger API-Aufruf und Antwortverifizierung

Der folgende Code demonstriert, wie das SDK verwendet wird, um eine "Benutzer erstellen"-Anfrage (Create User) zu erstellen, die HTTP-Anfrage zu senden und die Sicherheit der Signatur der von der Plattform zurückgegebenen Daten zu überprüfen.

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. SDK initialisieren und Resty-Client wiederverwenden (Voraussetzung: config.yaml ist konfiguriert)
	_, apiObj := common.Init()

	// 2. Anfrageparameter und Signatur-Header generieren
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Anfrage: ", err)
		return
	}

	// 3. Anfrage senden und Antwortsignatur automatisch verifizieren
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("Anfrage oder Verifizierung fehlgeschlagen: ", err)
		return
	}

	fmt.Println("✅ Anfrage erfolgreich und verifiziert! Zurückgegebene OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. Handbuch zur Verwendung von CLI-Tools

Das SDK bietet ausführbare Binärdateien (Command Line Interface, CLI) zum schnellen Testen jedes API-Endpunkts.

### 2.1 Ausführbare Dateien kompilieren

Führen Sie den Befehl `make` im Stammverzeichnis des SDK aus. Das System generiert dann ausführbare Binärdateien für jede Funktion im Verzeichnis `bin`.
* **Windows:** Generiert Dateien mit der Endung `.exe` (z. B. `create_user.exe`).
* **Mac/Linux:** Generiert Dateien ohne Dateiendung (z. B. `create_user`).

### 2.2 Konfigurationsdatei vorbereiten

Stellen Sie vor dem Ausführen der Tools sicher, dass sich die konfigurierte Datei `config.yaml` im Verzeichnis `bin` befindet.

### 2.3 Endpunkt-Befehle testen

#### Neuen Benutzer registrieren
1. Ändern Sie das Feld `UserOpenId` in `bin/config.yaml`.
2. Führen Sie `./create_user` aus (oder doppelklicken Sie auf `create_user.exe`).
3. Wenn die OpenId bereits registriert ist, gibt das Tool einen Fehler zurück.

#### Wallet-Registrierung
1. Geben Sie `UserOpenId` und `ChainID` in `bin/config.yaml` an.
2. Führen Sie `./create_wallet` aus.

#### Einzahlungsadressen abrufen
1. Geben Sie `UserOpenId` und die abgefragten `ChainIDs` (z. B. "1,56") in `bin/config.yaml` an.
2. Führen Sie `./get_wallet_addresses` aus.

#### Auszahlung beantragen
1. Geben Sie Folgendes in `bin/config.yaml` an:
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (Eindeutiger Anti-Duplikations-Code für Bestellungen)
   * `CallbackUrl`
2. Führen Sie `./user_withdraw_by_open_id` aus.