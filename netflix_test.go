package speedy

import "testing"

func TestGetHtmlResult(t *testing.T) {
	html, err := GetHtmlResult(NETFLIX_URL)
	if err != nil {
		t.Errorf("GetHtmlResult() = %s", err)
	}
	t.Logf("GetHtmlResult() = %s", html)
}

func TestParseJsPath(t *testing.T) {
	jsPath, err := ParseJsPath()
	if err != nil {
		t.Errorf("ParseJsPath() = %s", err)
	}
	t.Logf("ParseJsPath() = %s", jsPath)
}

func TestParseToken(t *testing.T) {
	token, err := ParseToken()
	if err != nil {
		t.Errorf("ParseToken() = %s", err)
	}
	t.Logf("ParseToken() = %s", token)
}

func TestNetflixDownloadTest(t *testing.T) {
	r, err := NetflixDownloadTest()
	if err != nil {
		t.Errorf("NetflixDownloadTest() = %s", err)
	}
	t.Logf("NetflixDownloadTest() = %f", r)
}

func TestNetflixUploadTest(t *testing.T) {
	t.Log("TestNetflixUploadTest")
}

func TestStartNetflixTest(t *testing.T) {
	t.Log("StartTestNetflixTest")
}

func BenchmarkStartNetflixTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartNetflixTest()
	}
}
