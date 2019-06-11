package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	sum := Sum(numbers)
	expected := 10

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d', given %v", expected, sum, numbers)
	}

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 2}, []int{0, 9, 9})
		want := []int{4, 18}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		t.Run("make the sums of some slices", func(t *testing.T) {
			got := SumAllTails([]int{}, []int{0, 9, 9})
			want := []int{0, 18}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
	})

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1, 2, 3, 4, 5}
		Sum(numbers)
	}
}
