# 例とツール

このドキュメントは 2 つの部分に分かれています。
1. **シナリオベースのコード例:** 実際のコードで API 呼び出しと検証を処理する方法を示します。
2. **CLI ツールガイド:** クイックテストのために SDK に含まれているコンパイル済み実行可能ファイルの使用方法について説明します。

---

## 1. シナリオベースのコード例

### 1.1 完全な API 呼び出しと応答の検証

次のコードは、SDK を使用して「ユーザー作成」リクエストを構築し、HTTP リクエストを送信し、プラットフォームから返されたデータの署名のセキュリティを検証する方法を示しています。

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. SDK の初期化と Resty クライアントの再利用 (前提条件: config.yaml が構成されていること)
	_, apiObj := common.Init()

	// 2. リクエストパラメーターと署名ヘッダーを生成する
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("リクエストの構築に失敗しました: ", err)
		return
	}

	// 3. リクエストを送信し、応答署名を自動的に検証する
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("リクエストまたは検証に失敗しました: ", err)
		return
	}

	fmt.Println("✅ リクエストが成功し、検証されました! 返された OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. CLI ツールの使用ガイド

SDK には、各 API エンドポイントをすばやくテストするためのコマンドラインインターフェース (CLI) バイナリファイルが用意されています。

### 2.1 実行可能ファイルのコンパイル

SDK のルートディレクトリで `make` コマンドを実行すると、システムは `bin` ディレクトリの各機能のバイナリ実行可能ファイルを生成します。
* **Windows:** `.exe` で終わるファイルを生成します (例: `create_user.exe`)。
* **Mac/Linux:** 拡張子のないファイルを生成します (例: `create_user`)。

### 2.2 設定ファイルの準備

ツールを実行する前に、設定済みの `config.yaml` ファイルが `bin` ディレクトリに配置されていることを確認してください。

### 2.3 エンドポイントコマンドのテスト

#### 新規ユーザーの登録
1. `bin/config.yaml` の `UserOpenId` フィールドを変更します。
2. `./create_user` を実行します (または `create_user.exe` をダブルクリックします)。
3. OpenId がすでに登録されている場合、ツールはエラーを返します。

#### ウォレットの登録
1. `bin/config.yaml` で `UserOpenId` と `ChainID` を指定します。
2. `./create_wallet` を実行します。

#### 入金アドレスの取得
1. `bin/config.yaml` で、`UserOpenId` とクエリする `ChainIDs` (例: "1,56") を指定します。
2. `./get_wallet_addresses` を実行します。

#### 出金の申請
1. `bin/config.yaml` で次を指定します。
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (一意の注文重複防止コード)
   * `CallbackUrl`
2. `./user_withdraw_by_open_id` を実行します。