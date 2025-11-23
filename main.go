package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:] // 첫 번째 인자(프로그램 이름) 제외

	var carNames []string
	times := -1

	// --times 플래그 찾기
	for i := 0; i < len(args); i++ {
		if args[i] == "--times" {
			if i+1 < len(args) {
				if t, err := strconv.Atoi(args[i+1]); err == nil && t > 0 {
					times = t
				} else if err != nil {
					fmt.Println("올바른 값을 입력하지 않았습니다")
					os.Exit(1)
				} else if t < 0 {
					fmt.Println("1보다 큰 정수를 입력하세요")
					os.Exit(1)
				}
			}
			break
		}
		carNames = append(carNames, args[i])
	}

	// times가 설정되지 않았으면 입력받기
	if times == -1 {
		fmt.Print("시도 횟수를 입력하세요: ")
		_, err := fmt.Scan(&times)
		if err != nil {
			fmt.Println("숫자를 입력하세요")
			os.Exit(1)
		} else if times < 1 {
			fmt.Println("1보다 큰 정수를 입력하세요")
			os.Exit(1)
		}
	}
}
