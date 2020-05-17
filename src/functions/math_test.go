//
// EPITECH PROJECT, 2020
// 205IQ_2019 [WSL]
// File description:
// math_test
//

package functions

import "testing"

func TestDataDeviation(t *testing.T) {
	LastTest := []int{100, 24}

	result4 := DataDeviation(LastTest)

	if result4 != 0 {
		t.Errorf("failed for the sucess DataDeviation")
	} else {
		t.Logf("\033[32mSucess\032 !")
	}

}

func TestInfIQ(t *testing.T) {
	LastTest := []int{100, 24, 90}

	result4 := InfIQ(LastTest)

	if result4 != 0 {
		t.Errorf("failed for the sucess InfIQ")
	} else {
		t.Logf("Sucess !")
	}

}

func TestBetweenIQ(t *testing.T) {
	LastTest := []int{100, 24, 90, 95}

	result4 := BetweenIQ(LastTest)

	if result4 != 0 {
		t.Errorf("failed for the sucess Between IQ")
	} else {
		t.Logf("Sucess !")
	}

}
