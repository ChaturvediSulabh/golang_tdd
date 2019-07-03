package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of whole numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	})

	message := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Colllection of Slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		message(t, got, want)
	})

	t.Run("Collection of all tails of slices", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		message(t, got, want)
	})

	t.Run("Tail of an empty slice", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{1, 4, 5})
		want := []int{0, 9}
		message(t, got, want)
	})

	t.Run("Tail of an empty slice", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{1, 4, 5})
		want := []int{0, 9}
		message(t, got, want)
	})
}
