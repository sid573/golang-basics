package pack

import "testing"

// Unit Testing
func TestAdd(t *testing.T) {
	// to run in parallel
	t.Parallel()
	t.Run("subtest1", func(b *testing.T) {
		s := []int{2, 3}
		v := sum(s)
		if v != 5 {
			b.Fail()
		}
	})
	t.Run("subtest2", func(b *testing.T) {

	})
}

// BenchMarking

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{2, 3}
		sum(s)
	}
}
