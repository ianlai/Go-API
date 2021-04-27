package main

import "testing"

func TestSum(t *testing.T) {
	if sum(2, 5) != 7 {
		t.Fail()
	}
	if sum(2, 100) != 102 {
		t.Fail()
	}
	if sum(222, 100) != 322 {
		t.Fail()
	}
}
