package interation

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Passing zero repeat we expect default time (5)", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("Passing minus one repeat we expect default time (5)", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("10 times", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat_positive() {
	repeated := Repeat("b", 10)
	fmt.Println(repeated)
	// Output: bbbbbbbbbb
}

func ExampleRepeat_zero() {
	repeated := Repeat("z", 0)
	fmt.Println(repeated)
	// Output: zzzzz
}

func ExampleRepeat_negative() {
	repeated := Repeat("h", -1)
	fmt.Println(repeated)
	// Output: hhhhh
}
