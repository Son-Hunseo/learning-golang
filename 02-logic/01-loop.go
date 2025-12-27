package main

import "fmt"

func main() {
	// go에서는 여러 언어에 있는 여러 스타일의 루프를 모두 지원한다.
	// for, while, do while, foreach 등 모든 스타일의 루프를 for로 표현 가능하다.

	fmt.Println("---------- for style ----------")
	// for 스타일
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------- while style ----------")
	// while 스타일
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("---------- infinite loop ----------")
	// 무한 루프 스타일
	counter := 0
	for {
		fmt.Println("counter:", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	fmt.Println("---------- array ----------")
	// 배열을 활용
	items := [3]string{"Go", "Python", "Java"}
	for index, value := range items {
		fmt.Println(index, value)
	}

	// 인덱스를 사용하고 싶지 않다면
	for _, value := range items {
		fmt.Println(value)
	}
}
