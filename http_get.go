package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func httpGet(url string) (result string) {
	resp,err := http.Get(url)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	result = string(body)
	return
}
func main() {
		fmt.Println(httpGet("http://www.baidu.com"))
}
