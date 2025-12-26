// 여느 다른 언어와 마찬가지로
// main에서 프로그램이 시작된다.

// 따라서 당연하게도 하나의 프로젝트에는 main이 1개만 존재해야하지만
// 지금은 학습용 프로젝트이기 때문에 하나의 프로젝트에서 폴더별로 나누고
// 각각의 파일에 package main, main()이 존재하는 것임
package main

// 라이브러리 import 하는 영역
import "fmt"

// 컴파일러는
// 1. package main을 찾음
// 2. main() 을 실행함
func main() {
	fmt.Println("Hello, world!")
}

// 이 파일을 실행하는 방법은
// 터미널로 이 파일의 디렉토리로 가서
// go run hello.go
// go run = 빌드 -> 파일(빌드 결과) 실행 -> 파일 삭제

// 이 파일을 빌드하는 방법은
// go build hello.go -> 빌드 결과물 exe 파일 나옴

// 코드 포매팅
// go fmt hello.go

// 외부 패키지 설치
// go get/install
// get과 install의 차이는 나중에 자세하게 다룸
