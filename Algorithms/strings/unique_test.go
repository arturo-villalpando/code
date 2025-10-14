package strings

import "testing"

func TestIsUnique(t *testing.T) {
	firstTest := IsUnique("hola")
	if !firstTest {
		t.Errorf("Result was incorrect, got: %t, want: %t", firstTest, true)
	}
	secondTest := IsUnique("holafa")
	if secondTest {
		t.Errorf("Result was incorrect, got: %t, want: %t", secondTest, false)
	}
}
