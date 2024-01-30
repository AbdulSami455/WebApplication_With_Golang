package main

import "testing"

func testadd(t *testing.T) {

	result := add(1, 2)
	if result != 3 {
		t.Errorf("error is there", result)
	}

}
