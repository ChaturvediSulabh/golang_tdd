package iteration

import (
	"fmt"
	"testing"
)

func TestIterations(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want s%s'", got, want)
		}
	}
	t.Run("repeat a char per the count provided", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Sum of all numbers in an array", func(t *testing.T) {
		got := Sum([5]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		want := 47
		assertCorrectMessage(t, got, want)
	})
}

func ExampleRepeat() {
	fmt.Println(Repeat("b", 4))
	// Output: bbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
