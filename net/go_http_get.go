package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)
var succesCount  = 0
func httpGet(url string) (result string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	if resp!=nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		if body != nil {
			succesCount+=1
			fmt.Printf("succesCount:%d\n",succesCount)
		}
		result = string(body)
		return
	}
	return "false"
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			httpGet("https://www.baidu.com/")
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
