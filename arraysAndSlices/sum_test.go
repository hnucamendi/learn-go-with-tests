package sum

import (
	"reflect"
	"slices"
	"testing"
)

func CheckExpected(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d wanted: %d", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("Sum Slice of any size", func(t *testing.T) {
		got := Sum([]int{1, 2, 4, 6, 10, 1, 2, 4, 6, 10})
		want := 46

		CheckExpected(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Test summing x number of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 0})
		want := []int{15, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v wanted: %v", got, want)
		}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v wanted: %v", got, want)
		}
	})
}

func TestSumTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got: %v wanted: %v", got, want)
		}
	}

	t.Run("Sum happy path", func(t *testing.T) {
		got := SumTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{1, 2})
		want := []int{5, 11, 17, 2}

		checkSums(t, got, want)
	})

	t.Run("Safely Sum empty slices", func(t *testing.T) {
		got := SumTails([]int{}, []int{1, 3})
		want := []int{0, 3}

		checkSums(t, got, want)
	})
}
