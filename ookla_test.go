package speedy

import "testing"

func TestStartOoklaTest(t *testing.T) {
	r, err := StartOoklaTest()
	if err != nil {
		t.Errorf("StartooklaTest() = %s", err)
	}
	t.Logf("StartooklaTest() = download: %f , upload: %f", r.Download, r.Upload)
}

func BenchmarkStartOoklaTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartOoklaTest()
	}
}
