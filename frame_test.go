package main

import "testing"

func TestTotalPoints_PerfectGame(t *testing.T) {
	knockDownFunc := func(standingPins int) int {
		switch standingPins {
		case 0:
			return 0
		default:
			return totalPins
		}
	}

	executeTestWithConstantScore(t, knockDownFunc, 300)
}

func TestTotalPoints_NineZero(t *testing.T) {
	knockDownFunc := func(standingPins int) int {
		switch standingPins {
		case 0:
			return 0
		case 10:
			return 9
		default:
			return 0
		}
	}

	executeTestWithConstantScore(t, knockDownFunc, 90)
}

func TestTotalPoints_AlwaysFive(t *testing.T) {
	knockDownFunc := func(standingPins int) int {
		switch standingPins {
		case 0:
			return 0
		default:
			return 5
		}
	}

	executeTestWithConstantScore(t, knockDownFunc, 150)
}

func TestSpareAfterMiss(t *testing.T) {
	originalGetKnockedDownPins := getKnockedDownPins
	defer func() { getKnockedDownPins = originalGetKnockedDownPins }()
	testFrame := frame{}

	getKnockedDownPins = func(standingPins int) int { return 0 }
	testFrame.roll(totalPins)
	getKnockedDownPins = func(standingPins int) int { return totalPins }
	testFrame.roll(totalPins)

	if testFrame.hadStrike() {
		t.Error("Expected frame to not be a strike when first roll was a miss!")
	}

	if !testFrame.hadSpare() {
		t.Error("Expected frame to be a spare when 10 pins were knocked down in the second roll!")
	}
}

func executeTestWithConstantScore(t *testing.T, knockDownFunc func(int) int, expectedTotalPoints int) {
	originalGetKnockedDownPins := getKnockedDownPins
	getKnockedDownPins = knockDownFunc
	defer func() { getKnockedDownPins = originalGetKnockedDownPins }()

	testGame := playNewGame()
	totalPoints := testGame[len(testGame)-1].getTotalPoints()
	if totalPoints != expectedTotalPoints {
		t.Errorf("Expected to have %d points in total, but got %d!", expectedTotalPoints, totalPoints)
	}
}
