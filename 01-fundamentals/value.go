package main

import "fmt"

func main() {

	fmt.Printf("Hello" + "World\n")

	fmt.Println("Hello" + "World")

	fmt.Println(1 + 1)

	fmt.Println(3.14)

	fmt.Println(true, false)

	// %v : "Value"의 약자로, 데이터의 기본 형태를 출력한다.
	// + : 구조체(Struct)를 출력할 때 필드 이름까지 함께 보여달라는 뜻 (어떠한 객체의 필드까지 모두 출력)
	// \n : 줄바꿈
	fmt.Printf("%+v\n", []int{1, 2, 3})

	var t any // golang에서의 기본 타입인 any(인터페이스) 타입이다.

	fmt.Printf("%+v\n", t) // 자바에서 int 기본값이 0이듯이 any 타입의 기본값은 <nil> 이다.

}
