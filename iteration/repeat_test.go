package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("Expected %q but got %q", want, got)
		}
	})

	t.Run("3 times", func(t *testing.T) {
		got := Repeat("b", 3)
		want := "bbb"

		if got != want {
			t.Errorf("Expected %q but got %q", want, got)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
