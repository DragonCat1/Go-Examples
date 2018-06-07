package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func httpGet(url string) (result string) {
	resp, err := http.Get(url)
	if err != nil {
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
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			httpGet("http://www.baidu.com")
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
