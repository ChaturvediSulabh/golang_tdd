package iteration

import "testing"

func TestIterations(t *testing.T) {
	got := repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got '%s' want s%s'", got, want)
	}
}
