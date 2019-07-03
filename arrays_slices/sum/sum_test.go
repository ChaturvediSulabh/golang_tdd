package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of numbers #1", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	})
}
