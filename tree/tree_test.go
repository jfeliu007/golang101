package main

import "testing"

func getYearCycleTest(t *testing.T) {
	if getYearSize(0) != 1 {
		t.Fatal("Size 0 != 1")
	}
	if getYearSize(1) != 2 {
		t.Fatal("Syce 1 != 2")
	}
	if getYearSize(4) != 7 {
		t.Fatal("Syce 4 != 7")
	}
}
