package cars

import (
	"testing"
)

func TestMoveCarsByRandomNumber(t *testing.T) {
	// Given
	cars := &Cars{}
	_ = cars.Init([]string{"pobi", "woni", "soso", "jisu"})

	// When - 여러 번 실행하여 최소 한 번은 움직이는지 확인
	for i := 0; i < 100; i++ {
		cars.MoveCarsByRandomNumber()
	}

	// Then - 100번 실행했으므로 적어도 일부 차는 움직였을 것
	allCars := cars.GetCars()
	moved := false
	for _, car := range allCars {
		if car.Steps > 0 {
			moved = true
			break
		}
	}

	if !moved {
		t.Error("Expected at least one car to move after 100 iterations")
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		names   []string
		wantErr bool
	}{
		{
			name:    "정상적인 이름들",
			names:   []string{"pobi", "woni", "jun"},
			wantErr: false,
		},
		{
			name:    "빈 이름",
			names:   []string{"pobi", "", "jun"},
			wantErr: true,
		},
		{
			name:    "5자 이상 이름",
			names:   []string{"pobi", "verylongname"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cars := &Cars{}
			err := cars.Init(tt.names)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetWinners(t *testing.T) {
	// Given
	cars := &Cars{}
	if err := cars.Init([]string{"pobi", "woni", "jun"}); err != nil {
		t.Fatal(err)
	}

	// 직접 Steps 값을 설정
	cars.cars[0].Steps = 5
	cars.cars[1].Steps = 3
	cars.cars[2].Steps = 5

	// When
	winners := cars.GetWinners()

	// Then
	if len(winners) != 2 {
		t.Errorf("Expected 2 winners, got %d", len(winners))
	}

	expectedWinners := map[string]bool{"pobi": true, "jun": true}
	for _, winner := range winners {
		if !expectedWinners[winner] {
			t.Errorf("Unexpected winner: %s", winner)
		}
	}
}

func TestGetCars(t *testing.T) {
	// Given
	cars := &Cars{}
	names := []string{"pobi", "woni"}
	if err := cars.Init(names); err != nil {
		t.Fatal(err)
	}

	// When
	result := cars.GetCars()

	// Then
	if len(result) != 2 {
		t.Errorf("Expected 2 cars, got %d", len(result))
	}

	if result[0].Name != "pobi" || result[1].Name != "woni" {
		t.Errorf("Car names don't match expected values")
	}
}
