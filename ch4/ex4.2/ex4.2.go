package main

import "fmt"

func main() {
	var minimunWage int = 10 // 변수 선언 및 초기화
	var workingHour int = 20 // 변수 선언 및 초기화

	//변수 income 선언 및 초기화
	var income int = minimunWage * workingHour

	//변수 세개 출력
	fmt.Println(minimunWage, workingHour, income)
}
