package internal

import "testing"

func TestExample(t *testing.T) {
	if 2+2 != 4 {
		t.Fatal("math broken")
	}
}
