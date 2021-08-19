package speedy

import "testing"

func TestOoklaTest(t *testing.T) {
	t.Log("TestOoklaTest")
}

func BenchmarkStartOoklaTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartOoklaTest()
	}
}
