# 예제 및 도구

이 문서는 두 부분으로 나뉩니다.
1. **시나리오 기반 코드 예제:** 실제 코드에서 API 호출 및 검증을 처리하는 방법을 보여줍니다.
2. **CLI 도구 가이드:** 빠른 테스트를 위해 SDK에 포함된 컴파일된 실행 파일을 사용하는 방법을 설명합니다.

---

## 1. 시나리오 기반 코드 예제

### 1.1 완전한 API 호출 및 응답 검증

다음 코드는 SDK를 사용하여 "사용자 생성" 요청을 빌드하고 HTTP 요청을 보내고 플랫폼에서 반환된 데이터의 서명 보안을 검증하는 방법을 보여줍니다.

```go
package main

import (
	"fmt"
	"github.com/zerospace-ai/pay-sdk-go/api"
	"github.com/zerospace-ai/pay-sdk-go/example/common"
	"github.com/zerospace-ai/pay-sdk-go/response_define"
)

func main() {
	// 1. SDK 초기화 및 Resty 클라이언트 재사용 (전제 조건: config.yaml이 구성됨)
	_, apiObj := common.Init()

	// 2. 요청 매개변수 및 서명 헤더 생성
	openId := "HASH1756194148"
	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		fmt.Println("요청 빌드 실패: ", err)
		return
	}

	// 3. 요청 보내기 및 응답 서명 자동 검증
	var rspCreateUser response_define.ResponseCreateUser
	err = common.ExecuteRequest(api.PathCreateUser, reqBody, timestamp, sign, clientSign, &rspCreateUser)
	if err != nil {
		fmt.Println("요청 또는 검증 실패: ", err)
		return
	}

	fmt.Println("✅ 요청이 성공적으로 검증되었습니다! 반환된 OpenId:", rspCreateUser.Data.OpenId)
}
```


---

## 2. CLI 도구 사용 가이드

SDK는 각 API 엔드포인트를 빠르게 테스트하기 위한 CLI(명령줄 인터페이스) 바이너리 파일을 제공합니다.

### 2.1 실행 파일 컴파일

SDK 루트 디렉터리에서 `make` 명령을 실행하면 시스템이 `bin` 디렉터리에 각 기능에 대한 바이너리 실행 파일을 생성합니다.
* **Windows:** `.exe`로 끝나는 파일을 생성합니다(예: `create_user.exe`).
* **Mac/Linux:** 확장자 없는 파일을 생성합니다(예: `create_user`).

### 2.2 구성 파일 준비

도구를 실행하기 전에 구성된 `config.yaml` 파일이 `bin` 디렉터리에 있는지 확인하십시오.

### 2.3 엔드포인트 명령 테스트

#### 새 사용자 등록
1. `bin/config.yaml`의 `UserOpenId` 필드를 수정합니다.
2. `./create_user`를 실행합니다(또는 `create_user.exe`를 두 번 클릭합니다).
3. OpenId가 이미 등록된 경우 도구에서 오류를 반환합니다.

#### 지갑 등록
1. `bin/config.yaml`에 `UserOpenId` 및 `ChainID`를 지정합니다.
2. `./create_wallet`을 실행합니다.

#### 입금 주소 얻기
1. `bin/config.yaml`에 `UserOpenId`와 조회할 `ChainIDs`(예: "1,56")를 지정합니다.
2. `./get_wallet_addresses`를 실행합니다.

#### 출금 신청
1. `bin/config.yaml`에 다음을 지정합니다.
   * `UserOpenId`
   * `TokenId`
   * `Amount`
   * `AddressTo`
   * `SafeCheckCode` (고유한 주문 중복 방지 코드)
   * `CallbackUrl`
2. `./user_withdraw_by_open_id`를 실행합니다.