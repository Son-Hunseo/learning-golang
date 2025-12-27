package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

// go 에서는 리시버를 사용해서 특정 타입에 해당 함수를 귀속시킴
// 이를 통해서 객체지향 언어의 메서드 처럼 사용할 수 잇게하여 객체지향 프로그래밍을 가능하게함
// 나중에 다룰 것이므로 일단 넘어가기
func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError { // 정의하지 않은 레벨의 로그 예외처리
		return "Unknown"
	}

	return levelNames[l] // leveNames 배열의 인덱스와 로그레벨이 같다.
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10)
}
