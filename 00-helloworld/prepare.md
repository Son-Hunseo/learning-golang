# `Linter` 설치

> `Linter` 란?
> - `Linter` 란 코드를 탐색해서 문제가 있거나 문제의 소지가 있는 코드를 찾아서 알려주는 자동화 도구
> - 컴파일러와 다른 점은 문법 자체에는 문제가 없어 실행 가능한 코드라도 버그 발생 가능성이 높거나 안티패턴이라고 인식하면 경고를 발생시킨다.

- `Linter` 중 `golang`을 위해 사용할 `Golangci-linter` 설치
  - https://golangci-lint.run/docs/welcome/install/local/
  - 위 링크에 나와있는 방법 중 OS에 맞는 방법으로 설치
- 설치를 완료하면 터미널에 `golangci-lint` 입력해서 설치 확인

---

# `golang-migrate`

> `golang-migrate` 란?
> - DB에서의 `git`과 같은 역할
> - 협업 환경에서 데이터베이스 구조를 변경할 때, 각 팀원이 수동으로 SQL을 실행하면 실수가 발생하기 쉽다. `golang-migrate`를 사용하면 다음과 같은 장점이 있다.
> - 버전 관리: SQL 파일을 001_init.up.sql, 002_add_users.up.sql처럼 순서대로 관리하여 DB 상태를 추적할 수 있다.
> - 일관성 유지: 모든 개발자와 운영 서버가 동일한 DB 구조를 갖도록 보장한다.
> - 롤백 지원: 문제가 생겼을 때 이전 상태로 되돌리는 down 스크립트를 함께 작성할 수 있습니다.
> - 자동화: CI/CD 파이프라인이나 애플리케이션 시작 시 자동으로 DB를 최신 상태로 업데이트할 수 있습니다.
> - cf) `golang`으로 만들어졌지만, `golang`만을 위한 도구는 아니다.

- https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- 위 링크에서 `With Go toolchain` - `Unversioned` 에서 버전에 맞는 명령어 복사해서 설치
- 설치 후 터미널에 `migrate` 입력해서 설치 확인

---
# `swag` 설치

> `swag` 란?
> - `golang`으로 코드를 짤 때, 특정 주석을 달아놓으면 `swag` 도구가 이 주석을 읽고 자동으로 `Swagger`(OpenAPI 2.0) 문서를 만들어주는 도구

- https://github.com/swaggo/swag
- 위 링크에서 `go install` 로 시작하는 명령어로 설치하기
- 설치 후 터미널에 `swag` 입력해서 설치 확인

---
# `gqlgen` 설치

> `gqlgen` 이란?
> - `GraphQL` 형식의 API를 구축할 때 사용하는 도구
> - 스키마를 먼저 작성하면 이를 바탕으로 `golang` 코드를 자동으로 생성해줌

- https://gqlgen.com/ 이 공식 페이지이긴 한데, `go install` 형식의 명령어가 없어서 `go get ~` 명령어 뒷부분 주소로 `go install` 해보자.
- 즉, `go install github.com/99designs/gqlgen@latest` 명령어로 설치
- 설치 후 터미널에 `gqlgen` 입력해서 설치 확인

---
# `pgAdmin` 설치

> `pgAdmin` 이란?
> - 앞으로 다룰 부분에서 `PostgreSQL`을 사용할 것이다.
> - 이를 위한 `PostgreSQL` 관리자 툴이다.
> - `MySQL`의 `MySQL Workbench`와 비슷하다고 생각하면 됨

- https://www.pgadmin.org/
- 위 링크에서 OS에 맞는 명령어로 설치하기
