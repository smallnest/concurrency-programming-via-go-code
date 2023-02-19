package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Wait()

	var urls = []string{"http://baidu.com", "http://bing.com", "http://google.com"}
	var result = make([]bool, len(urls))
	http.DefaultClient.Timeout = time.Second

	wg.Add(3)
	for i := 0; i < 3; i++ {
		i := i
		go func(url string) {
			defer wg.Done()

			log.Println("fetching", url)
			resp, err := http.Get(url)
			if err != nil {
				result[i] = false
				return
			}

			result[i] = resp.StatusCode == http.StatusOK
			resp.Body.Close()

		}(urls[i])
	}

	wg.Wait()
	wg.Wait()
	wg.Wait()
	log.Println("done")
	for i := 0; i < 3; i++ {
		log.Println(urls[i], ":", result[i])
	}
}
