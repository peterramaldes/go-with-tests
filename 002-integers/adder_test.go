package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertResult := func(t testing.TB, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("2 + 2 = 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertResult(t, sum, expected)
	})

	t.Run("2 + 7 = 9", func(t *testing.T) {
		sum := Add(2, 7)
		expected := 9
		assertResult(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
