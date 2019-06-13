package iteration

import "testing"

func TestIterations(t *testing.T) {
	got := repeat("a", 5)
	want := "aaaaa"
	if got != want {
		t.Errorf("got '%s' want s%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 5)
	}
}
