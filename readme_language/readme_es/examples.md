# Ejemplos y Herramientas

Este documento se divide en dos partes:
1. **Ejemplos de Código Basados en Escenarios:** Demuestra cómo manejar llamadas API y verificaciones en código práctico.
2. **Guía de Herramientas CLI:** Explica cómo utilizar los archivos ejecutables compilados incluidos con el SDK para pruebas rápidas.

---

## 1. Ejemplos de Código Basados en Escenarios

### 1.1 Llamada Completa a la API y Verificación de Respuesta

El siguiente código demuestra cómo utilizar el SDK para construir una solicitud "Crear Usuario", enviar la solicitud HTTP y verificar la seguridad de la firma de los datos devueltos por la plataforma.

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. Inicializar el SDK y reutilizar el cliente Resty (Requisito previo: config.yaml está configurado)
	_, apiObj := common.Init()

	// 2. Generar parámetros de solicitud y encabezado de firma
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("Error al construir la solicitud: ", err)
		return
	}

	// 3. Enviar solicitud y verificar automáticamente la firma de la respuesta
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("Error en la solicitud o verificación: ", err)
		return
	}

	fmt.Println("✅ ¡Solicitud exitosa y verificada! OpenId devuelto:", rspCreateUser.Data.OpenId)
}
```


---

## 2. Guía de Uso de Herramientas CLI

El SDK proporciona archivos binarios de la Interfaz de Línea de Comandos (CLI) para probar rápidamente cada punto final de la API.

### 2.1 Compilar Archivos Ejecutables

Ejecute el comando `make` en el directorio raíz del SDK, y el sistema generará archivos ejecutables binarios para cada función en el directorio `bin`.
* **Windows:** Genera archivos que terminan en `.exe` (por ejemplo, `create_user.exe`).
* **Mac/Linux:** Genera archivos sin sufijos (por ejemplo, `create_user`).

### 2.2 Preparar Archivo de Configuración

Antes de ejecutar las herramientas, asegúrese de que el archivo configurado `config.yaml` se encuentre en el directorio `bin`.

### 2.3 Prueba de Comandos de Puntos Finales

#### Registrar Nuevo Usuario
1. Modifique el campo `UserOpenId` en `bin/config.yaml`.
2. Ejecute `./create_user` (o haga doble clic en `create_user.exe`).
3. Si el OpenId ya está registrado, la herramienta devolverá un error.

#### Registro de Billetera
1. Especifique `UserOpenId` y `ChainID` en `bin/config.yaml`.
2. Ejecute `./create_wallet`.

#### Obtener Direcciones de Depósito
1. Especifique `UserOpenId` y los `ChainIDs` a consultar (por ejemplo, "1,56") en `bin/config.yaml`.
2. Ejecute `./get_wallet_addresses`.

#### Solicitar Retiro
1. Especifique lo siguiente en `bin/config.yaml`:
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (Código anti-duplicación de pedido único)
   * `CallbackUrl`
2. Ejecute `./user_withdraw_by_open_id`.