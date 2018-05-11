package main

import(
	"fmt"
	"time"
	"log"
	"net/http"
	"context"
	"sync"
)

func main() {

	loadCache()

	loadCacheProperly()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// loadCache - wrong or absent goroutine management
func loadCache() error {

	go func() {
		for range time.Tick(10 * time.Second){
			fmt.Printf("%v, Do cache loading!\n", time.Now())
		}
	}()
	return nil
}

// loadCacheProperly - proper goroutine management
func loadCacheProperly() error {

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		cancel()
		wg.Wait()
	}()

	wg.Add(1)
	go func() {

		defer wg.Done()

		select {
		case <- ctx.Done():
			return // avoid leaking of this goroutine when ctx is done.
		default:
			for range time.Tick(10 * time.Second){
				fmt.Printf("%v, Do proper cache loading!\n", time.Now())
			}
		}
	}()
	return nil
}