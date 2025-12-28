package main

import (
	"fmt"
	"time"
)

func main() {

	day := "Tuesday"
	fmt.Println("Today is", day)

	// if와 else if 로도 같은 로직을 작성할 수 있지만
	// switch로 작성함으로써 코드가 더욱 간결하고 조건에 집중할 수 있다.
	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	// 아래와 같이 할당과 동시에 switch문을 작성할 수도 있다.
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) { // 타입 스위치문 v는 타입 확인이 완료된 값을 담을 새로운 변수
		case int: // i의 타입 v가 int라면
			fmt.Printf("Integer: %d\n", v)
		case string: // i의 타입 v가 string이라면
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.32)
}
