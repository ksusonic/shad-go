//go:build !solution

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	var elapsed time.Duration

	for _, url := range os.Args[1:] {
		wg.Add(1)
		go func(url string, t time.Time) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			duration := time.Since(start)
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("%s\t%d\t%s\n", duration, len(body), url)
			elapsed = duration
		}(url, start)
	}
	wg.Wait()
	fmt.Println(elapsed, "elapsed")
}
