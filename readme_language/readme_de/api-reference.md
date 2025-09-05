# API-Referenz 📚

Dieses Dokument beschreibt alle API-Schnittstellen des CryptoPay Go SDK im Detail, einschließlich Anfrageparametern, Rückgabeparametern und Beispielen.

## Neuen Benutzer registrieren (create_user)🆕🧑‍💻

### Konzept
Erstellen Sie einen neuen Plattformbenutzer, der die eindeutige ID des Benutzers erfordert, d.h. UserOpenId.

### HTTP-Anfrage
- Schnittstellenname: create_user
- URL: https://sandbox-api.privatex.io/sdk/user/create
- Methode: POST

### Anfrageparameter
| Parametername | Erforderlich | Typ   | Beschreibung                                             |
| ------ | ---- | ------ | ------------------------------------------------ |
| OpenId | Ja   | string | Empfohlen, den Plattform-Standardpräfix + eindeutige Benutzer-ID zu verwenden, um OpenId zu bilden |

### Rückgabeparameter
| Parametername      | Typ   | Beschreibung            |
| ----------- | ------ | --------------- |
| code        | int    | Globaler Statuscode      |
| msg         | string | Statusbeschreibung        |
| data.OpenId | string | Eindeutige Benutzer-OpenId |
| sign        | string | Plattformsignatur        |

### Beispiel
Anfragebeispiel:
```bash
curl --location 'https://sandbox-api.privatex.io/sdk/user/create' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{ 
  "OpenId":"PT00001"
}'
```
Rückgabebeispiel:
```json
{
    "code": 1,
    "msg": "ok",
    "data": {
        "OpenId": "PT00001"
    },
    "sign": "..."
}
```

Für Authentifizierung & Sicherheit siehe [🧩 authentication.md](./authentication.md)

## Wallet erstellen (create_wallet) 💰

### Konzept
Erstellen Sie ein Wallet-Konto für den Benutzer im entsprechenden Blockchain-Netzwerk.

### HTTP-Anfrage
- Schnittstellenname: create_wallet
- URL: https://sandbox-api.privatex.io/sdk/wallet/create
- Methode: POST

### Anfrageparameter
| Parametername  | Erforderlich | Typ   | Beschreibung            |
| ------- | ---- | ------ | --------------- |
| ChainID | Ja   | string | Chain-ID           |
| OpenId  | Ja   | string | Eindeutige Benutzer-OpenId |

### Rückgabeparameter
| Parametername       | Typ   | Beschreibung            |
| ------------ | ------ | --------------- |
| code         | int    | Globaler Statuscode      |
| msg          | string | Statusbeschreibung        |
| data.address | string | Wallet-Adresse        |
| data.OpenId  | string | Eindeutige Benutzer-OpenId |
| sign         | string | Plattformsignatur        |

### Beispiel
Anfragebeispiel:
```bash
curl --location 'https://sandbox-api.privatex.io/sdk/wallet/create' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{
  "OpenId":"PT00001",
  "ChainID":"1"
}'
```
Rückgabebeispiel:
```json
{
    "code": 1,
    "msg": "ok",
    "data": {
        "address": "...",
        "OpenId": "PT00001"
    },
    "sign": "..."
}
```

## Einzahlungsadresse abrufen (get_wallet_addresses)💳

### Konzept
Rufen Sie die Blockchain-Wallet-Einzahlungsadresse des Benutzers ab.

### HTTP-Anfrage
- Schnittstellenname: get_wallet_addresses
- URL: https://sandbox-api.privatex.io/sdk/wallet/getWalletAddresses
- Methode: POST

### Anfrageparameter
| Parametername   | Erforderlich | Typ   | Beschreibung                  |
| -------- | ---- | ------ | --------------------- |
| OpenId   | Ja   | string | Eindeutige Benutzer-OpenId       |
| ChainIDs | Ja   | string | Mehrere Chain-IDs, durch Kommas getrennt |

### Rückgabeparameter
| Parametername         | Typ   | Beschreibung       |
| -------------- | ------ | ---------- |
| code           | int    | Globaler Statuscode |
| msg            | string | Statusbeschreibung   |
| data.Addresses | array  | Adressenliste   |
| sign           | string | Plattformsignatur   |

### Beispiel
Anfragebeispiel:
```bash
curl --location 'https://sandbox-api.privatex.io/sdk/wallet/getWalletAddresses' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{
  "OpenId":"PT00001",
  "ChainIDs":"56,2"
}'
```
Rückgabebeispiel:
```json
{
    "code": 1,
    "msg": "ok",
    "data": {
        "Addresses": [
            {
                "chainID": 56,
                "address": "..."
            }
        ]
    },
    "sign": "..."
}
```

## Benutzer-Auszahlung (user_withdraw_by_open_id)💸

### Konzept
Benutzer-Auszahlungsoperation, Überweisung vom Partnerkonto zur benutzerspezifizierten Adresse.

* Funktion: Benutzer-Auszahlungsoperationsschnittstelle. Auszahlungen müssen vom Partnerkonto im entsprechenden Token-Auszahlungspool zum benutzerspezifizierten Auszahlungs-Wallet-Adresse überwiesen werden. Partner können eine sichere Callback-Adresse einrichten, um die Legitimität der Auszahlung zu überprüfen. Wenn als gültig überprüft, kann die Auszahlung direkt aus dem Wallet des Händler-Fonds-Pools abgeschlossen werden.

* Die Auszahlungstransaktionsschnittstelle überprüft, ob das Standard-Auszahlungs-Hot-Wallet ausreichend Auszahlungsvermögen und Gebühren hat.

* Standardmäßig verwendet die Auszahlungsschnittstelle einen Sicherheitsüberprüfungscode als eindeutigen Parameteranforderung für Auszahlungstransaktionen. Es wird generell empfohlen, die eindeutige Auszahlungsauftragsnummer der Geschäftsplattform als Sicherheitskode zu verwenden. Das Einreichen eines doppelten Sicherheitsüberprüfungscodes führt zu einem Fehler.

* Alle Auszahlungstransaktionsanfragen werden mit den auf der Kanalplattform konfigurierten Risikokontrollprüfungsregeln abgeglichen. Wenn die Parameteranfrage gültig ist, wird die Transaktionsanfrage akzeptiert. Auszahlungstransaktionen, die den automatischen Prüfungsregeln entsprechen, werden sofort an die Netzwerktransaktion übermittelt und die Hash-Informationen der übermittelten Transaktion werden zurückgegeben (Rückgabefeld data). Auszahlungstransaktionsanfragen, die eine sekundäre Prüfung auf dem Kanal erfordern, geben (code=2) zurück. Die Auszahlungsanfrage muss nicht erneut eingereicht werden. Der Administrator muss die sekundäre Prüfung auf der Kanalplattform abschließen. Nach Abschluss der sekundären Prüfung wird der Transaktionsauftrag callback-benachrichtigt, um den Status der Auszahlungstransaktion zu ändern.

* Voraussetzung: Der Fonds-Pool der entsprechenden Währung muss ausreichend Mittel für die Auszahlung haben (insbesondere für Token-Auszahlungen im ETH-Netzwerk, die ein bestimmtes ETH-Transaktionsgebühren-Guthaben im Fonds-Pool-Wallet erfordern).

* ⚠️ Hinweis: **Für Blockchain-Auszahlungen stellen Sie bitte sicher, dass der Vorabgenehmigungsprozess abgeschlossen ist, bevor Sie die Schnittstelle aufrufen. Sobald eine Blockchain-Transaktion initiiert wird, kann sie nicht widerrufen oder zurückgegeben werden.**

### HTTP-Anfrage
- Schnittstellenname: user_withdraw_by_open_id
- URL: https://sandbox-api.privatex.io/sdk/partner/UserWithdrawByOpenID
- Methode: POST

### Anfrageparameter
| Parametername        | Erforderlich | Typ   | Beschreibung            |
| ------------- | ---- | ------ | --------------- |
| OpenId        | Ja   | string | Eindeutige Benutzer-OpenId |
| TokenId       | Ja   | string | Token-ID         |
| Amount        | Ja   | float  | Auszahlungsbetrag        |
| AddressTo     | Ja   | string | Zieladresse        |
| CallBackUrl   | Nein   | string | Callback-URL        |
| SafeCheckCode | Nein   | string | Sicherheitsüberprüfungscode      |

### Rückgabeparameter

| Parametername | Typ | Beschreibung |
| :-------- | :----- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| code | int | Statuscode</br>0 Parameterfehler, doppelte Auftragsnummer, falsches Auszahlungsadressenformat oder unzureichende Auszahlungs-Wallet-Gebühren. Detaillierte Informationen finden Sie in msg.</br>1 Die Auszahlungstransaktion wurde erfolgreich übermittelt und an das Blockchain-Netzwerk übermittelt. Der eindeutige Hash der übermittelten Transaktion ist in data enthalten.</br>2 Die Auszahlungstransaktion wurde erfolgreich übermittelt und erfordert eine sekundäre Kanalprüfung, bevor die Transaktion abgeschlossen werden kann. Nach Abschluss der Prüfung werden die Transaktionsinformationen durch einen Callback aktualisiert.</br>-1 Die Auszahlungstransaktion ist fehlgeschlagen. Sie können die Auszahlungsanfrage erneut einreichen. |
| msg | string | Statusbeschreibung |
| data | string | Transaktionshash. Wenn intelligente Auszahlung aktiviert ist, wird dieses Feld als leere Zeichenkette zurückgegeben. |
| sign | string | Plattformsignatur |
| timestamp | string | Aktueller Zeitstempel in Millisekunden, umgewandelt in eine Zeichenkette |


### Beispiel
Anfragebeispiel:
```bash
curl --location 'https://sandbox-api.privatex.io/sdk/partner/UserWithdrawByOpenID' \
--header 'key: vratson2i5hjxgkd' \
--header 'sign: 0592dc64d480fb119d1e07ce06011db8' \
--header 'clientSign: xxxxxxxxxxxxxxxxx' \
--header 'Content-Type: application/json' \
--header 'timestamp: 1725076567682' \
--data '{ 
  "OpenId": "PT00001", 
  "TokenId": "4", 
  "Amount": "0.02", 
  "AddressTo": "TQdL5yttJPTx7hJmBhGfo2LcE7AXLPtHSg", 
  "CallBackUrl": "http://xxxxxx/withdraw_callback", 
  "SafeCheckCode": "1000000000000000"
}'
```

Rückgabebeispiel
```json
{
    "sign": "D+VTPNiwGLzh9eIvkrscwS4UlGKzdnrBgB63RDG4HeobZT6FXqUwYCPgKojynKaxwm5PkmW0xhIASZ4asSCvnYfi0NSFehchZAtUnQIispxKcjsiudWsUznbkEIQ2h2TA/mbUZ1X9+wyh7QhNo6+RkxtgRyRpVb7ARG8pL14cdTAs OTtMLO0W1GO0M83VAv2ybBZNObncX9qy6tdwLQV/KYuNJYyMN0dL0nLKYHnj9Q4d3lEDM45AVJ0153/YIiIgcF BnOWhsQ9rVARcFeXeWd9KJ5OZpmxlxnhcJGcEUY2UDC4zKLZxtUet7CPAyehAMQ5plkpvRrR3Z6lA5zl6GQ==",
    "timestamp": "1725439986754",
    "data": "94f4c29eba73d53dcd3aa1b8cf90a98108d0acf82f38b97a4032dcdf7ff172e7",
    "msg": "ok",
    "code": 1
}
```

## Sekundäre Prüfung des Auszahlungsauftrags 💳

* Funktion: Sekundäre Risikokontrollprüfungsschnittstelle für Händler-Auszahlungsaufträge
* ⚠️ Hinweis: **Die Plattform weist Händlern einen separaten Risikokontroll-RSA-öffentlichen Schlüssel zu (anders als der Ein-/Auszahlungs-Callback-Benachrichtigungs-öffentliche Schlüssel)**
* Auslösungszeit: Nachdem der Administrator die Risikokontroll-Callback-URL-Parameter auf der Händlerseite konfiguriert hat (Systemeinstellungen), fügt der Kanal für jede Auszahlungstransaktionsanfrage eine zusätzliche Risikokontroll-Callback-sekundäre Prüfung hinzu. Nur wenn die Risikokontroll-URL auf der Händlerseite einen korrekten Überprüfungspass-Code zurückgibt, ist die Transaktion gültig übermittelt.
* Technische Anforderungen: Technische Implementierung und Konfiguration der sekundären Prüfungs-Callback-Schnittstelle auf der Händlerseite sind erforderlich.

#### HTTP-Anfrage

Die Plattform sendet eine Auszahlungsprüfungsanfrage an den Händler

> POST: `/withdrawal/order/check`

#### Anfrageparameter

| Parametername | Erforderlich | Typ | Beschreibung |
| :-------- | :--- | :----- | :------------------------------------------------------------------------- |
| safeCode | Nein | string | Eindeutige Transaktions-ID, die vom Händler eingereicht wurde, entspricht im Allgemeinen der Auszahlungsauftrags-ID des Händlers (SafeCheckCode für Auszahlungstransaktionen) |
| openId | Ja | string | Benutzer-ID des Händlers, der die Auszahlungstransaktion einreicht |
| tokenId | Ja | string | Währungs-ID, basierend auf der von der Plattform bereitgestellten Währungs-ID |
| toAddress | Ja | string | Auszahlungsadresse |
| amount | Ja | string | Auszahlungsbetrag |
| timestamp | Ja | int | Aktueller Zeitstempel |
| sign | Ja | string | Signatur: Nur die Parameter im Data-Feld werden signiert; die Korrektheit der Signatur muss mit dem Risikokontroll-RSA-öffentlichen Schlüssel der Plattform überprüft werden.

#### Rückgabeparameterbeschreibung

| Parametername | Typ | Beschreibung |
| :-------- | :----- | :----------------------------------------------------- |
| code | int | Überprüfungsergebnis. 0 bedeutet bestanden; andere Codes sind ungültig.
| timestamp | int | Aktueller Zeitstempel, in Sekunden.
| message | string | Rückgabemeldung.
| sign | string | Signatur: Die RSA-Privatschlüssel-Signatur des Händlers für das Data-Feld im Antwortparameter.

## Ein- und Auszahlungs-Callback-Benachrichtigungen

1. Ein- und Auszahlungstransaktionen lösen mehrere Callback-Benachrichtigungen aus. Die Transaktionsinformationen und der Status der letzten Callback-Benachrichtigung werden verwendet.
2. Die Geschäftseite muss eine gültige Callback-Nachricht zurückgeben. Das Format ist in der Rückgabeparameterbeschreibung beschrieben. Ein Rückgabecode von 0 bedeutet, dass die Callback-Nachricht verarbeitet wurde und keine weiteren Benachrichtigungen erforderlich sind. Andernfalls wird der Callback fortgesetzt (zunächst alle 2 Sekunden für 50 Mal, und dann alle 10 Minuten danach), bis eine Bestätigungsnachricht mit Code 0 zurückgegeben wird.

Kontaktieren Sie den Dienstleister, um die Callback-URL einzustellen.

> POST

* Funktion: Definiert das Callback-Nachrichtenformat, das die Plattform verwendet, um die Anwendungsseite über Token-Transaktionsinformationen (Benutzer-Auszahlung oder Einzahlung) zu benachrichtigen. Diese Nachricht eignet sich für Anwendungsseitige Ereignisbenachrichtigungen bezüglich des Token-Transaktionsstatus (Auszahlung oder Einzahlung). Anwendungen können die Callback-Benachrichtigungsschnittstelle optional basierend auf ihrer Anwendungsfunktionalität unterstützen.

### Anfrageparameter

| Parametername | Erforderlich | Typ | Beschreibung                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| :----------- | :--- | :----- |:----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| openid | ja | string | Eindeutige Kanalbenutzer-ID                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| totalvalue | ja | string | USDT-Wert, der der Ein- oder Auszahlungstransaktion entspricht (basierend auf dem Marktpreis zum Zeitpunkt der Transaktion berechnet)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| status | ja | int | Transaktionsstatus:</br>1 Die Transaktion ist abgeschlossen und wurde erfolgreich an das Blockchain-Netzwerk übermittelt. Transaktionsdetails können on-chain mit dem Hash abgefragt werden.</br>-1 Die Transaktion wurde an das Blockchain-Netzwerk übermittelt, aber die on-chain-Transaktion ist fehlgeschlagen. Sie können sie in Merchant Management --> Transaction Management --> [Submit Order Security Code] erneut prüfen. Die Geschäftsplattform muss den Statuswechsel nicht verarbeiten und kann einfach auf den Kanal warten, um die neue Statusbenachrichtigung callback zu rufen.</br>-2 Die Auszahlungstransaktionsanfrage wurde vom Händler-Backend abgelehnt. Die Auszahlungsanfrage ist abgelaufen. Es wird empfohlen, dass die Geschäftsplattform die Auszahlungsanfrage des Benutzers nach Erhalt der Benachrichtigung zurückgibt.</br>2 Die Auszahlungstransaktion wurde an das Merchant Management übermittelt. Da sie die konfigurierten Währungssicherheitsrisikokontrollanforderungen ausgelöst hat, muss der Administrator sich bei Merchant Management --> Transaction Management --> Withdrawal Review anmelden, um die Auszahlungsanfrage zu bearbeiten.</br>3 Während der Blockchain-Verarbeitung der Auszahlungstransaktion muss die Geschäftsplattform den Statuswechsel nicht aktualisieren und kann einfach auf den Kanal warten, um eine neue Statusbenachrichtigung zu erhalten. </br>⛑️**Spezielle Erinnerung: Für Auszahlungstransaktions-Callbacks, die von der Geschäftsplattform empfangen werden, wenn status = -1, wird der Callback ignoriert. Nachdem der Administrator sich im Management-Backend angemeldet und die Transaktion erneut übermittelt hat, wird gleichzeitig eine neue Statusbenachrichtigung an die Plattform gepusht.** || type | ja | int | 1 für Einzahlungstransaktionen; 2 für Auszahlungstransaktionen |
| hash | ja | string | Transaktions-Hash-Wert                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| confirm | ja | int | Anzahl der on-chain-Bestätigungen für die Transaktion                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| createdtime | ja | string | Erstellungszeit                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| from | ja | string | Adresse des Transaktionsinitiators                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| to | ja | string | Empfangsadresse der Transaktion                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |