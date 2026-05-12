# CryptoPay Go SDK

![Go Version](https://img.shields.io/badge/go-1.18+-blue.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Telegram](https://img.shields.io/badge/chat-Telegram-blue?logo=telegram)](https://t.me/ZeroSerivce)

## Bienvenido a CryptoPay Go SDK

CryptoPay Go SDK es un SDK de servicio de criptomonedas profesional implementado en Golang, que proporciona funciones como registro de usuarios, generación de billeteras, notificaciones de devolución de llamadas de depósitos y retiros.
Ha sido ampliamente utilizado y ha demostrado ser seguro, estable y fácilmente ampliable.

## Instalación

```bash
go get github.com/zerospace-ai/pay-sdk-go
```

> **Nota:** La compilación requiere Go 1.18+.

## Inicio Rápido

### 1. Preparar la Configuración

Antes de utilizar el SDK, debe preparar el archivo de configuración `config.yaml`, que contiene la información de autenticación del comerciante y las claves públicas/privadas:

```yaml
ApiKey: "your_api_key"
ApiSecret: "your_api_secret"
PlatformPubKey: "platform_public_key"
PlatformRiskPubKey: "platform_risk_public_key"
RsaPrivateKey: "your_rsa_private_key"
```

> **💡 Consejo:** Para obtener detalles sobre cómo generar el propio par de claves RSA del comerciante (`RsaPrivateKey`) y los mecanismos detallados de autenticación y seguridad, lea [Autenticación y Seguridad (authentication.md)](./authentication.md).

### 2. Inicializar el SDK y Enviar Solicitud

Aquí hay un ejemplo completo que demuestra cómo inicializar la instancia del SDK y llamar a la API "Crear Nuevo Usuario":

```go
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zerospace-ai/pay-sdk-go/api"
)

func main() {
	// 1. Cargar configuración
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 2. Crear instancia del SDK
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	// 3. Llamar a la API: Construir la solicitud para crear un nuevo usuario
	openId := "PT00001" // Identificador único del usuario
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Printf("Error al construir la solicitud: %v\n", err)
		return
	}

	// (La ejecución de la solicitud de red y el análisis de respuesta se omiten aquí, consulte examples.md para obtener el código completo)
	fmt.Printf("¡Solicitud CreateUser construida con éxito!\nBody: %s\n", string(reqBody))
	fmt.Printf("Headers preparados: timestamp=%s, sign=%s, clientSign=%s\n", timestamp, sign, clientSign)
}
```

## Conceptos Clave y Navegación

Para utilizar mejor este SDK, le recomendamos que lea los documentos restantes en el siguiente orden:

1. **[Autenticación y Seguridad (authentication.md)](./authentication.md)**: Aprenda cómo generar pares de claves RSA y el mecanismo de verificación de firmas entre el SDK y la plataforma.
2. **[Referencia de API (api-reference.md)](./api-reference.md)**: Contiene instrucciones detalladas para todos los puntos finales de API compatibles (por ejemplo, creación de billetera, retiro) y formatos de webhook.
3. **[Ejemplos y Herramientas (examples.md)](./examples.md)**: Vea ejemplos de código basados en escenarios más complejos e instrucciones sobre el uso de las herramientas CLI integradas del SDK.
4. **[Apéndice (appendix.md)](./appendix.md)**: Información de diccionario estático, como ChainIDs admitidos, tipos de tokens y direcciones de contratos.

## Contacto

Si tiene alguna pregunta, comuníquese con el proveedor de servicios:  
Telegram: [@ZeroSerivce](https://t.me/ZeroSerivce)