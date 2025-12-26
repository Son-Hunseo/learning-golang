package main

import "fmt"

func main() {
	// 자바에서는 String name 이런식으로 자료형이 왼쪽에 오지만
	// golang에서는 자료형이 오른쪽에 온다.

	var name string // cf) string 자료형의 기본 값은 빈 문자열 "" 이다.
	name = "Hello, World!"
	fmt.Println(name)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	// 한 줄에 여러 변수 선언 가능
	var firstName, lastName string
	firstName = "Hunseo"
	lastName = "Son"
	fmt.Println(firstName, lastName)

	// 변수 선언과 초기화 동시에
	email := "test@test.com"
	fmt.Println(email)

	age := 28
	fmt.Println(age)

	// 변수 선언과 초기화를 동시에 하는 경우는 위처럼 초기화 값에 따라 타입이 정해지지만
	// 이를 명시하고 싶을 경우 아래와 같이 사용하는 것도 가능하다.
	var year int = 2025
	fmt.Println(year)
}
