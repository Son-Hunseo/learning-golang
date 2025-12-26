package main

import "fmt"

// constant 즉, 상수는 다른 언어에서와 마찬가지로 상수값 (변할 수 없음)
const HOST string = "localhost"

func main() {
	fmt.Println(HOST)

	// HOST = HOST + ":8080" // 불변하기 때문에 이렇게 재할당 하려고 하면 컴파일 오류남
}
