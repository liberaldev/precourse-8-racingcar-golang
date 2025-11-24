package cars

import (
	"precourse-8-racingcar-golang/random_number"
	"testing"
)

func TestMoveCarsByRandomNumber(t *testing.T) {
	// Given
	cars := &Cars{}

	if err := cars.Init([]string{"pobi", "woni", "soso", "jisu"}); err != nil {
		t.Fatal(err)
	}

	// Mock 랜덤 값 설정 (4, 3, 3, 4)
	mockRandom := random_number.InitMockRandomGenerator([]int{4, 3, 3, 4})
	cars.randomGenerator = mockRandom

	// When
	if err := cars.MoveCarsByRandomNumber(); err != nil {
		t.Fatal(err)
	}

	// Then
	expected := []Car{
		{Name: "pobi", Steps: 1},
		{Name: "woni", Steps: 0},
		{Name: "soso", Steps: 0},
		{Name: "jisu", Steps: 1},
	}

	actual := cars.GetCars()
	for i, car := range actual {
		if car.Name != expected[i].Name || car.Steps != expected[i].Steps {
			t.Errorf("Index %d: Expected %+v, got %+v", i, expected[i], car)
		}
	}
}

func TestMoveCarsByRandomNumberAllMove(t *testing.T) {
	// Given
	cars := &Cars{}
	if err := cars.Init([]string{"pobi", "woni"}); err != nil {
		t.Fatal(err)
	}

	// Mock 랜덤 값 설정 (모두 4 이상)
	mockRandom := random_number.InitMockRandomGenerator([]int{4, 5})
	cars.randomGenerator = mockRandom

	// When
	if err := cars.MoveCarsByRandomNumber(); err != nil {
		t.Fatal(err)
	}
	// Then
	actual := cars.GetCars()
	if actual[0].Steps != 1 || actual[1].Steps != 1 {
		t.Errorf("Expected all cars to move, got %+v", actual)
	}
}

func TestMoveCarsByRandomNumberNoneMove(t *testing.T) {
	// Given
	cars := &Cars{}
	if err := cars.Init([]string{"pobi", "woni"}); err != nil {
		t.Fatal(err)
	}

	// Mock 랜덤 값 설정 (모두 4 미만)
	mockRandom := random_number.InitMockRandomGenerator([]int{0, 3})
	cars.randomGenerator = mockRandom

	// When
	if err := cars.MoveCarsByRandomNumber(); err != nil {
		t.Fatal(err)
	}

	// Then
	actual := cars.GetCars()
	if actual[0].Steps != 0 || actual[1].Steps != 0 {
		t.Errorf("Expected no cars to move, got %+v", actual)
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
