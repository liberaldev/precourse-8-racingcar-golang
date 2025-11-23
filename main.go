package main

import (
	"errors"
	"fmt"
	"os"
	"precourse-8-racingcar-golang/cars"

	"github.com/spf13/cobra"
)

var times int

var rootCmd = &cobra.Command{
	Use:   "racing [car1] [car2] [car3]...",
	Short: "자동차 경주",
	Long:  "여러 자동차의 이름을 입력받아 경주를 진행합니다.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		carNames := args

		// times가 설정되지 않았으면 입력받기
		if times == 0 {
			fmt.Print("시도 횟수를 입력하세요: ")
			_, err := fmt.Scan(&times)
			if err != nil {
				return errors.New("숫자를 입력하세요")
			}
		}

		if times < 1 {
			return errors.New("1보다 큰 정수를 입력하세요")
		}

		c := cars.Cars{}
		if err := c.Init(carNames); err != nil {
			return err
		}

		for i := 0; i < times; i++ {
			fmt.Println()
			c.MoveCarsByRandomNumber()
			c.CarsStepPrint()
		}

		return nil
	},
}

func init() {
	rootCmd.Flags().IntVarP(&times, "times", "t", 0, "시도 횟수")

	// 한글 도움말 템플릿 설정
	rootCmd.SetUsageTemplate(`사용법:
  {{.UseLine}}

플래그:
{{.LocalFlags.FlagUsages}}`)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
