package u

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
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

func Write(path, content string) {
	err := ioutil.WriteFile(path, []byte(content), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
