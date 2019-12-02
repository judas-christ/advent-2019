package main

import "testing"

func Test_getFuelRequired_100756_eq_50346(t *testing.T) {
	got := getFuelRequired(100756)
	if got != 50346 {
		t.Errorf("getFuelRequired was incorrect, got: %d, want: %d.", got, 50346)
	}
}
