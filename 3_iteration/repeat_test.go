package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("correct output", func(t *testing.T) {
		repeated := Repeat("b", 5)
		expected := "bbbbb"
		if repeated != expected {
			t.Errorf("got %q want %q", repeated, expected)
		}
	})

	t.Run("correct length", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"
		if len(repeated) != len(expected) {
			t.Errorf("got %q want %q", repeated, expected)
		}
	})

	t.Run("correct repition", func(t *testing.T) {
		repeated := Repeat("a", 3)
		repeatAmount := strings.Count(repeated, "a")
		expected := 3
		if repeatAmount != expected {
			t.Errorf("got %q want %q", repeated, expected)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	character := "ab"
	repeated := Repeat(character, 3)
	fmt.Println(repeated)
	// Output: ababab
}
