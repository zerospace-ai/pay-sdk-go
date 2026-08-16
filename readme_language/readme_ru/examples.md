# Примеры и инструменты

Этот документ разделен на две части:
1. **Примеры кода на основе сценариев:** Демонстрирует, как обрабатывать вызовы API и проверки в практическом коде.
2. **Руководство по инструментам CLI:** Объясняет, как использовать скомпилированные исполняемые файлы, поставляемые с SDK, для быстрого тестирования.

---

## 1. Примеры кода на основе сценариев

### 1.1 Полный вызов API и проверка ответа

В следующем коде показано, как использовать SDK для создания запроса "Создать пользователя", отправки HTTP-запроса и проверки безопасности подписи данных, возвращаемых платформой.

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. Инициализация SDK и повторное использование клиента Resty (предварительное условие: config.yaml настроен)
	_, apiObj := common.Init()

	// 2. Создание параметров запроса и заголовка подписи
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("Ошибка при создании запроса: ", err)
		return
	}

	// 3. Отправка запроса и автоматическая проверка подписи ответа
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("Ошибка запроса или проверки: ", err)
		return
	}

	fmt.Println("✅ Запрос успешен и проверен! Возвращенный OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. Руководство по использованию инструментов CLI

SDK предоставляет бинарные файлы интерфейса командной строки (CLI) для быстрого тестирования каждой конечной точки API.

### 2.1 Компиляция исполняемых файлов

Выполните команду `make` в корневом каталоге SDK, и система сгенерирует бинарные исполняемые файлы для каждой функции в каталоге `bin`.
* **Windows:** Генерирует файлы, оканчивающиеся на `.exe` (например, `create_user.exe`).
* **Mac/Linux:** Генерирует файлы без суффиксов (например, `create_user`).

### 2.2 Подготовка файла конфигурации

Перед запуском инструментов убедитесь, что настроенный файл `config.yaml` помещен в каталог `bin`.

### 2.3 Тестирование команд конечной точки

#### Регистрация нового пользователя
1. Измените поле `UserOpenId` в `bin/config.yaml`.
2. Запустите `./create_user` (или дважды щелкните `create_user.exe`).
3. Если OpenId уже зарегистрирован, инструмент вернет ошибку.

#### Регистрация кошелька
1. Укажите `UserOpenId` и `ChainID` в `bin/config.yaml`.
2. Запустите `./create_wallet`.

#### Получить адреса пополнения
1. Укажите `UserOpenId` и запрашиваемые `ChainIDs` (например, "1,56") в `bin/config.yaml`.
2. Запустите `./get_wallet_addresses`.

#### Заявка на вывод средств
1. Укажите следующее в `bin/config.yaml`:
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (Уникальный код для предотвращения дублирования заказа)
   * `CallbackUrl`
2. Запустите `./user_withdraw_by_open_id`.

#### Создание кассового ордера
1. Укажите `OutOrderNo`, `TokenId`, `Quantity` и `NotifyUrl` в `bin/config.yaml`.
2. Запустите `./new_order`.

#### Запрос баланса кошелька
1. Укажите `WalletAddress`, `ContractAddress` и `WalletBalanceChainId` в `bin/config.yaml`.
2. Запустите `./wallet_balance`.