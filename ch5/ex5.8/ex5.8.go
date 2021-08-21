package main

import (
	"bufio" //io를 담당하는 패키지
	"fmt"
	"os" //표준 입출력 등을 가지고 있는 패키지
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	// //표준 입력을 읽는 객체, NewReader()함수로 표준입력스트림(os.Stdin)을 이용해서 Reader 객체를 생성
	// NewReader returns a new Reader whose buffer has the default size.

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
