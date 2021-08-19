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

func TestDownloadTest(t *testing.T) {
	r, err := DownloadTest()
	if err != nil {
		t.Errorf("DownloadTest() = %s", err)
	}
	t.Logf("DownloadTest() = %s", r)
}

func TestUploadTest(t *testing.T) {
	t.Log("TestUploadTest")
}

func TestNetflixTest(t *testing.T) {
	t.Log("TestNetflixTest")
}
