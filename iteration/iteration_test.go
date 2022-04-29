package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeats the input string the number of times we specify", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkMark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
