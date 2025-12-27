package main

// 백준 문제 풀 때 입출력 예시
// boj 3733
import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader // *를 붙여서 포인터를 사용하여 어느 함수에서나 전역으로 접근 가능하게 선언
var writer *bufio.Writer

func solve() {
	var n int
	var s int
	var flag any = nil

	for flag == nil {
		// &도 포인터임, *는 실제 값을 가져올 떄, &는 해당 변수의 주소 값을 나타낼 때 쓰는 포인터
		// Fscan의 반환 첫번째 인자는 성공 갯수, 두번째 인자는 error 이다.
		// 성공 갯수는 필요없는데 선언하면 사용하지 않아서 컴파일 오류 뜨기 때문에 _로 선언
		_, flag = fmt.Fscanln(reader, &n, &s)
		if flag == nil {
			_, _ = fmt.Fprintln(writer, s/(n+1))
		}
	}
}

func main() {
	reader = bufio.NewReader(os.Stdin)  // 표준 입력을 위한 버퍼 생성
	writer = bufio.NewWriter(os.Stdout) // 표준 출력을 위한 버퍼 생성
	// defer는 main 함수가 끝나기 직전에 호출될 로직을 의미
	defer writer.Flush() // 출력 버퍼에 있는 것들을 출력한다.

	solve()
}
