package numbers

import "testing"

func TestNumbers(t *testing.T) {
	got := Sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
