package speedy

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

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

func GetChunkedResult(url string) (int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	i := 0
	chunk := 100 * 1024
	for {
		_, err := io.ReadFull(resp.Body, make([]byte, chunk))
		if err != nil {
			return i * chunk, err
		}
		i++
	}
}

func NetflixDownloadTest() (float64, error) {
	token, err := ParseToken()
	if err != nil {
		return 0, err
	}
	baseurl := "https://api.fast.com/"
	url := baseurl + "netflix/speedtest?https=true&token=" + token + "&urlCount=3"
	r, err := GetHtmlResult(url)
	if err != nil {
		return 0, err
	}
	urls := []UrlResponse{}
	uerr := json.Unmarshal([]byte(r), &urls)
	if uerr != nil {
		return 0, uerr
	}

	size, _ := GetFilesize(urls[1].URL)
	start := time.Now().Unix()
	GetHtmlResult(urls[1].URL)
	end := time.Now().Unix()

	return float64(size / (end - start) * 8 / (1024 * 1024)), nil
}

func NetflixUploadTest() (float64, error) {
	// TODO: Implement
	return 0, nil
}

func GetFilesize(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	// Is our request ok?

	if resp.StatusCode != http.StatusOK {
		return 0, err
	}

	// the Header "Content-Length" will let us know
	// the total file size to download
	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	return int64(size), nil
}

func StartNetflixTest() (*Result, error) {
	d, err := NetflixDownloadTest()
	if err != nil {
		return nil, err
	}
	u, err := NetflixUploadTest()
	if err != nil {
		return nil, err
	}
	return &Result{
		Download: d,
		Upload:   u,
	}, nil
}
