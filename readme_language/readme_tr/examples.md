# Örnekler ve Araçlar

Bu belge iki bölüme ayrılmıştır:
1. **Senaryo Tabanlı Kod Örnekleri:** Pratik kodda API çağrılarının ve doğrulamalarının nasıl işleneceğini gösterir.
2. **CLI Araçları Kılavuzu:** Hızlı test için SDK'da bulunan derlenmiş yürütülebilir dosyaların nasıl kullanılacağını açıklar.

---

## 1. Senaryo Tabanlı Kod Örnekleri

### 1.1 Tam API Çağrısı ve Yanıt Doğrulaması

Aşağıdaki kod, bir "Kullanıcı Oluştur" isteği oluşturmak, HTTP isteğini göndermek ve platform tarafından döndürülen verilerin imzasının güvenliğini doğrulamak için SDK'nın nasıl kullanılacağını gösterir.

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. SDK'yı Başlatın ve Resty İstemcisini Yeniden Kullanın (Önkoşul: config.yaml yapılandırılmış olmalıdır)
	_, apiObj := common.Init()

	// 2. İstek parametrelerini ve imza Header'ını oluşturun
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("İstek oluşturulamadı: ", err)
		return
	}

	// 3. İsteği gönderin ve yanıt imzasını otomatik olarak doğrulayın
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("İstek veya doğrulama başarısız oldu: ", err)
		return
	}

	fmt.Println("✅ İstek başarılı ve doğrulandı! Döndürülen OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. CLI Araçları Kullanım Kılavuzu

SDK, her API uç noktasını hızlı bir şekilde test etmek için Komut Satırı Arayüzü (CLI) ikili dosyaları sağlar.

### 2.1 Yürütülebilir Dosyaları Derleme

SDK kök dizininde `make` komutunu çalıştırın ve sistem `bin` dizinindeki her işlev için ikili yürütülebilir dosyalar üretecektir.
* **Windows:** `.exe` ile biten dosyalar oluşturur (örneğin, `create_user.exe`).
* **Mac/Linux:** Uzantısız dosyalar oluşturur (örneğin, `create_user`).

### 2.2 Yapılandırma Dosyasını Hazırlama

Araçları çalıştırmadan önce, yapılandırılmış `config.yaml` dosyasının `bin` dizinine yerleştirildiğinden emin olun.

### 2.3 Uç Nokta Komutlarını Test Etme

#### Yeni Kullanıcı Kaydet
1. `bin/config.yaml` içindeki `UserOpenId` alanını değiştirin.
2. `./create_user` komutunu çalıştırın (veya `create_user.exe`ye çift tıklayın).
3. OpenId zaten kayıtlıysa, araç bir hata döndürür.

#### Cüzdan Kaydı
1. `bin/config.yaml` içinde `UserOpenId` ve `ChainID` değerlerini belirtin.
2. `./create_wallet` komutunu çalıştırın.

#### Para Yatırma Adreslerini Al
1. `bin/config.yaml` içinde `UserOpenId` ve sorgulanan `ChainIDs` değerlerini (örneğin, "1,56") belirtin.
2. `./get_wallet_addresses` komutunu çalıştırın.

#### Para Çekme Başvurusu
1. `bin/config.yaml` içinde aşağıdakileri belirtin:
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (Benzersiz sipariş kopyalamayı önleme kodu)
   * `CallbackUrl`
2. `./user_withdraw_by_open_id` komutunu çalıştırın.

#### Kasa Siparişi Oluştur
1. `bin/config.yaml` dosyasında `OutOrderNo`, `TokenId`, `Quantity` ve `NotifyUrl` bilgilerini belirtin.
2. `./new_order` komutunu çalıştırın.

#### Cüzdan Bakiyesini Sorgula
1. `bin/config.yaml` dosyasında `WalletAddress`, `ContractAddress` ve `WalletBalanceChainId` bilgilerini belirtin.
2. `./wallet_balance` komutunu çalıştırın.