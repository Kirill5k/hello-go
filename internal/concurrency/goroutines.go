package concurrency

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func GoRoutines() {
	urls := []string{
		"https://golang.com",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go getContentTypeOf(url, &wg)
	}
	wg.Wait()
}

func getContentTypeOf(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer func(Body io.ReadCloser) {
		fmt.Printf("closing resp.Body")
		if err := Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}
