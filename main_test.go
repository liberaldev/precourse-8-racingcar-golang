package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		times       int
		wantErr     bool
		errContains string
	}{
		{
			name:        "인자가 없는 경우",
			args:        []string{},
			times:       0,
			wantErr:     true,
			errContains: "최소 1개 이상의 자동차 이름이 필요합니다",
		},
		{
			name:    "자동차 1대, times 플래그 지정",
			args:    []string{"car1", "--times", "3"},
			times:   3,
			wantErr: false,
		},
		{
			name:    "자동차 여러 대, times 플래그 지정",
			args:    []string{"car1", "car2", "car3", "-t", "5"},
			times:   5,
			wantErr: false,
		},
		{
			name:        "times가 음수인 경우",
			args:        []string{"car1", "--times", "-1"},
			times:       -1,
			wantErr:     true,
			errContains: "1보다 큰 정수",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 전역 변수 초기화
			times = 0

			// 출력 캡처를 위한 버퍼
			buf := new(bytes.Buffer)
			rootCmd.SetOut(buf)
			rootCmd.SetErr(buf)

			// 테스트 인자 설정
			rootCmd.SetArgs(tt.args)

			// 커맨드 실행
			err := rootCmd.Execute()

			// 에러 검증
			if tt.wantErr {
				if err == nil {
					t.Errorf("error를 예상했지만 nil을 받음")
				} else if tt.errContains != "" && !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("error message = %v, want contains %v", err.Error(), tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

			// times 값 검증 (에러가 없는 경우)
			if !tt.wantErr && times != tt.times {
				t.Errorf("times = %v, want %v", times, tt.times)
			}
		})
	}
}

// 표준 입력을 모킹하는 테스트
func TestRootCommandWithStdin(t *testing.T) {
	// 표준 입력 모킹
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// 가짜 입력 생성
	r, w, _ := os.Pipe()
	os.Stdin = r

	// "3\n"을 표준 입력으로 제공
	go func() {
		defer func(w *os.File) {
			err := w.Close()
			if err != nil {
				t.Errorf(err.Error())
			}
		}(w)
		_, _ = io.WriteString(w, "3\n")
	}()

	// 전역 변수 초기화
	times = 0

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"car1", "car2"})

	err := rootCmd.Execute()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if times != 3 {
		t.Errorf("times = %v, want 3", times)
	}

	output := buf.String()
	if !strings.Contains(output, "최종 우승자") {
		t.Error("우승자 출력이 없습니다")
	}
}

// racing 함수 단위 테스트
func TestRacing(t *testing.T) {
	tests := []struct {
		name     string
		carNames []string
		timesVal int
		wantErr  bool
	}{
		{
			name:     "정상 실행",
			carNames: []string{"car1", "car2", "car3"},
			timesVal: 3,
			wantErr:  false,
		},
		{
			name:     "자동차 이름이 긴 경우",
			carNames: []string{"verylongcarname"},
			timesVal: 1,
			wantErr:  true, // cars 패키지에서 에러 발생 예상
		},
		{
			name:     "자동차 1대",
			carNames: []string{"car1"},
			timesVal: 5,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			times = tt.timesVal
			err := racing(tt.carNames)

			if err = w.Close(); err != nil {
				t.Error(err.Error())
			}

			os.Stdout = oldStdout

			if tt.wantErr && err == nil {
				t.Error("error를 예상했지만 nil을 받음")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			var buf bytes.Buffer
			if _, err := io.Copy(os.Stdout, r); err != nil {
				t.Error(err.Error())
			}
			output := buf.String()

			if !strings.Contains(output, "최종 우승자") {
				t.Error("우승자 출력이 없습니다")
			}
		})
	}
}

// scanTimes 함수 테스트
func TestScanTimes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "정상 입력",
			input:   "5\n",
			wantErr: false,
		},
		{
			name:    "문자 입력",
			input:   "abc\n",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 표준 입력 모킹
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			r, w, _ := os.Pipe()
			os.Stdin = r

			go func() {
				defer func(w *os.File) {
					err := w.Close()
					if err != nil {
						t.Error(err.Error())
					}
				}(w)
				_, _ = io.WriteString(w, tt.input)
			}()

			// 전역 변수 초기화
			times = 0

			err := scanTimes()

			if tt.wantErr && err == nil {
				t.Error("error를 예상했지만 nil을 받음")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
