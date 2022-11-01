package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func errorPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func request(method string, url string, header map[string][]string, cookies []http.Cookie, reqBody io.Reader) (result bool) {
	req, err := http.NewRequest(method, url, reqBody)
	errorPrint(err)
	if err != nil {
		return
	}
	for _, cookie := range cookies {
		req.AddCookie(&cookie)
	}
	req.Header = header
	resp, err := client.Do(req)
	errorPrint(err)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// errorPrint(err)
	// if err!=nil{
	// 	return
	// }
	// result = string(body)
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func goRequest(i int) {
	for {
		data := url.Values{}
		data.Set("fz", "")
		data.Set("ud", "cee")
		data.Set("toip", "")
		data.Set("m3", "+86")
		data.Set("t5", "187"+RandomString(8, []rune("0123456789")))
		data.Set("t3", RandomString(12))
		fmt.Println(i, request(http.MethodPost, "http://"+host+"/user/vx/2.asp", header, cookies, strings.NewReader(data.Encode())))
	}
}

var succesCount = 0
var client = http.Client{
	Timeout: time.Second,
}
var cookies = []http.Cookie{}
var host = "sj.951.vkjjjjzz.xyz"
var header = map[string][]string{
	"User-Agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"},
	"Content-Type":    {"application/x-www-form-urlencoded"},
	"Accept-Encoding": {"gzip, deflate"},
	"Host":            {host},
	"Origin":          {"http://" + host},
	"Referer":         {"http://" + host + "/user/vx/"},
}
var waitGroup sync.WaitGroup

func main() {
	for i := 0; i < 2000; i++ {
		waitGroup.Add(1)
		go goRequest(i)
	}
	waitGroup.Wait()
}
