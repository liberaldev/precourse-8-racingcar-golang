package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"precourse-8-racingcar-golang/cars"
	"strings"

	"github.com/spf13/cobra"
)

var times int

func scanTimes() error {
	fmt.Print("시도 횟수를 입력하세요: ")
	_, err := fmt.Scan(&times)
	if err != nil {
		return errors.New("숫자를 입력하세요")
	}
	return nil
}

func racing(carNames []string) error {
	c := cars.Cars{}
	if err := c.Init(carNames); err != nil {
		return err
	}

	for i := 0; i < times; i++ {
		if err := c.MoveCarsByRandomNumber(); err != nil {
			return err
		}
		c.CarsStepPrint()
		fmt.Println()
	}

	fmt.Println("최종 우승자:", strings.Join(c.GetWinners(), ", "))
	return nil
}

var executableName = filepath.Base(os.Args[0])

var rootCmd = &cobra.Command{
	Use:   executableName + " [car1] [car2] [car3]...",
	Short: "자동차 경주",
	Long:  "여러 자동차의 이름을 입력받아 경주를 진행합니다.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("최소 1개 이상의 자동차 이름이 필요합니다")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		carNames := args
		// times가 설정되지 않았으면 입력받기
		if times == 0 {
			if err := scanTimes(); err != nil {
				return err
			}
		}
		if times < 1 {
			return errors.New("1보다 큰 정수를 입력하세요")
		}
		fmt.Println()
		if err := racing(carNames); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.Flags().IntVarP(&times, "times", "t", 0, "시도 횟수")

	// help 플래그 설명 한글화
	rootCmd.Flags().BoolP("help", "h", false, "도움말 표시")

	// 한글 도움말 템플릿 설정
	rootCmd.SetUsageTemplate(`사용법:
  {{.UseLine}}

플래그:
{{.LocalFlags.FlagUsages}}`)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
