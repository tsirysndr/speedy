package speedy

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const NETFLIX_URL = "https://fast.com/"

type UrlResponse struct {
	URL string `json:"url"`
}

func GetHtmlResult(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ParseJsPath() (string, error) {
	r, err := GetHtmlResult(NETFLIX_URL)
	if err != nil {
		return "", err
	}
	tokenizer := html.NewTokenizer(strings.NewReader(r))
	jspath := ""

	for {
		tokenizer.Next()
		tag, hasAttr := tokenizer.TagName()
		if string(tag) == "script" && hasAttr {
			_, value, _ := tokenizer.TagAttr()
			jspath = string(value)
			break
		}
	}

	return jspath, nil
}

func ParseToken() (string, error) {
	jspath, err := ParseJsPath()
	if err != nil {
		return "", err
	}
	r, err := GetHtmlResult(NETFLIX_URL + jspath)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(r, ",") {
		if strings.Contains(line, "token:") {
			token := strings.Split(line, `"`)[1]
			return token, nil
		}
	}
	return "", nil
}

func DownloadTest() (string, error) {
	token, err := ParseToken()
	if err != nil {
		return "", err
	}
	baseurl := "https://api.fast.com/"
	url := baseurl + "netflix/speedtest?https=true&token=" + token + "&urlCount=3"
	r, err := GetHtmlResult(url)
	if err != nil {
		return "", err
	}
	urls := []UrlResponse{}
	uerr := json.Unmarshal([]byte(r), &urls)
	if uerr != nil {
		return "", uerr
	}
	return urls[0].URL, nil
}

func UploadTest(token string) (int, error) {
	return 0, nil
}

func StartNetflixTest() (*Result, error) {
	return nil, nil
}
