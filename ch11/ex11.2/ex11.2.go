package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(i * 2)
		i++
	}
}
