package main

import "testing"

func TestOnePlusOne(t *testing.T) {
	if 1+1 != 2 {
		t.Fatalf("Something went wrong with the world")
	}
}
