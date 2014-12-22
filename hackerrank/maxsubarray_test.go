package main

import "testing"

func TestMax(t *testing.T) {
	var max int

	max = MaxContiguous([]int{1, 2, 3})
	if max != 6 {
		t.Error("max is not 6: ", max)
	}
}

func TestMaxSplit(t *testing.T) {
	var max int

	max = MaxContiguous([]int{1, 2, 3, -10, 4, 5, 6})
	if max != 15 {
		t.Error("max is not 15: ", max)
	}
}

func TestMaxSplitOverTheTop(t *testing.T) {
	var max int

	max = MaxContiguous([]int{1, 2, 3, -1, 4, 5, 6})
	if max != 20 {
		t.Error("max is not 20: ", max)
	}
}

func TestMaxOfAllNegative(t *testing.T) {
	var max int

	max = MaxContiguous([]int{-3, -2, -1})
	if max != -1 {
		t.Error("max is not 1: ", max)
	}
}

func TestMaxOfLonePositiveAtEnd(t *testing.T) {
	var max int

	max = MaxContiguous([]int{-3, -2, -2, 5})
	if max != 5 {
		t.Error("max is not 5: ", max)
	}
}
