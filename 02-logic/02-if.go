package main

import "fmt"

func main() {

	tmp := 25
	if tmp > 30 {
		fmt.Println("greater than 30")
	} else { // Java랑 다르게 줄바꿈을 하면 안됨
		fmt.Println("lower than 30")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}

	// 조금 더 세련된 방법

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	// hasAccess에 userAccess["jane"] 값이 할당된다.
	// ok는 userAccess 맵에 "jane"이라는 키가 존재하는지 여부를 알려주는 bool 값이다.
	// ok를 사용하는 이유는 go의 경우 키가 없어도 에러가 나지 않고 디폴트값(false 혹은 0)을 반환하기 때문에
	// 이에 해당 맵에 값이 false인지, 아니면 키가 없는 것인지를 구분하기 위해 ok 라는 변수를 관용적으로 사용한다.
	if hasAccess, ok := userAccess["jane"]; ok && hasAccess {
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("access not granted")
	}
}
