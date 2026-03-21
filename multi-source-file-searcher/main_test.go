package main

import "testing"

func TestOnePlusOne(t *testing.T) {
	if 1+1 != 2 {
		t.Fatalf("Something is wrong with the world")
	}

}
