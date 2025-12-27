package main

import (
	"fmt"
)

// Enum이란 데이터에 이름이라는 옷을 입혀서 의미 있는 집합으로 만드는 과정이다.
// 이에 사실 아래 예시 1에서도 const Sunday = 0, ... const Saturday = 6 이런식으로 일일히 할당해도 되지만
// 수정과 확장 측면, 실수를 방지 하는 측면에서 아래와 같이 const 블록과 iota를 사용하여 관용적으로 enum을 정의하는 방법으로 사용한다.

// Go의 const 블록 안에서 iota는 항상 0에서 시작하며, 새로운 줄로 넘어갈 때 마다 값이 1씩 자동으로 증가한다.

// 예시 1
const (
	Sunday = iota // 0부터 시작이 아닌 1부터 시작을 하고싶다면, Sunday = iota + 1 이렇게 시작하면 된다.
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 예시 2
type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn
	LogInfo
	LogDebug
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Saturday)

	fmt.Println(LogDebug)
}
