package main

import "fmt"

func main() {
	var a int = 10                    // a 변수 선언      - ❶
	var msg string = "Hello Variable" // msg 변수 선언    - ❷

	a = 20               // a 변수 값 변경   - ❸
	msg = "Good Morning" // msg 변수 값 변경 - ❹
	fmt.Println(msg, a)  // msg 와 a 값 출력 - ❺
}
