package cli

import "testing"

func TestCLI(t *testing.T) {
	wanted := "Hello World!"
	result := GetCliText()
	if wanted != result {
		t.Errorf("Wanted %s but got %s", wanted, result)
	}
}
