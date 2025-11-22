package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] // 첫 번째 인자(프로그램 이름) 제외

	var carNames []string
	var times int

	// --times 플래그 찾기
	for i := 0; i < len(args); i++ {
		if args[i] == "--times" && i+1 < len(args) {
			if t, err := strconv.Atoi(args[i+1]); err == nil {
				times = t
			}
			break
		}
		carNames = append(carNames, args[i])
	}

	fmt.Printf("자동차 이름: %v\n", carNames)
	fmt.Printf("시도 횟수: %d\n", times)
}
