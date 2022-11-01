package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var client http.Client

func errorPrint(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func request(method string, url string, header map[string][]string, cookies []http.Cookie) (result string) {
	fmt.Println(method, url, header, cookies)
	client = http.Client{}
	req, err := http.NewRequest(method, url, nil)
	for _, cookie := range cookies {
		req.AddCookie(&cookie)
	}
	req.Header = header
	errorPrint(err)
	resp, err := client.Do(req)
	errorPrint(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorPrint(err)
	result = string(body)
	return
}

func main() {
	cookies := []http.Cookie{}
	header := map[string][]string{
		"user-agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"},
	}
	fmt.Println(request(http.MethodGet, "https://useragent.buyaocha.com", header, cookies))
}
