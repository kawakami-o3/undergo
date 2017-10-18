package u

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"time"
)

func U(err error) {
	if err != nil {
		panic(err)
	}
}

func ReplaceAll(target, regex, replace string) string {
	return regexp.MustCompile(regex).ReplaceAllString(target, replace)
}

func Match(target, regex string) bool {
	return regexp.MustCompile(regex).MatchString(target)
}

func MatchAny(target string, list ...string) bool {
	for _, regex := range list {
		if Match(target, regex) {
			return true
		}
	}
	return false
}

func Include(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func Read(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	result := ""
	reader := bufio.NewReaderSize(file, 4096)
	for {
		line, _, err := reader.ReadLine()
		if err == nil {
			result += string(line) + "\n"
		} else if err == io.EOF {
			return result, nil
		} else {
			return "", err
		}
	}
}

func Write(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), os.ModePerm)
}

type HttpClient struct{ http.Client }

var defaultClient = func() *HttpClient {
	jar, _ := cookiejar.New(nil)

	return &HttpClient{http.Client{
		Jar:     jar,
		Timeout: time.Duration(10) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return errors.New("redirect")
		},
	}}
}()

type HttpRequest struct {
	values url.Values
	files  map[string]string
}

func (req *HttpRequest) Add(key, value string) {
	req.values.Add(key, value)
}

func (req *HttpRequest) AddFile(key, fileName string) {
	req.files[key] = fileName
}

func Post(url string, req *HttpRequest) (string, error) {
	return defaultClient.Post(url, req)
}

func (c *HttpClient) Post(url string, req *HttpRequest) (string, error) {
	var resp *http.Response
	var err error
	if len(req.files) == 0 {
		resp, err = c.PostForm(url, req.values)
	} else {
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		for key, fileName := range req.files {
			fileWriter, err := bodyWriter.CreateFormFile(key, fileName)
			if err != nil {
				return "", err
			}

			fh, err := os.Open(fileName)
			if err != nil {
				return "", err
			}
			defer fh.Close()

			_, err = io.Copy(fileWriter, fh)
			if err != nil {
				return "", err
			}
		}

		for key := range req.values {
			bodyWriter.WriteField(key, req.values.Get(key))
		}

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		resp, err = c.Client.Post(url, contentType, bodyBuf)
	}

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
