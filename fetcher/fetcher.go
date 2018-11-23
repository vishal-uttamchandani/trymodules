package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Fetcher will fetch urls
type Fetcher struct {
	results chan Result
	wg      sync.WaitGroup
}

// New returns a new fetcher
func New() *Fetcher {
	return &Fetcher{
		results: make(chan Result),
	}
}

// Fetch will fetch the urls
func (f *Fetcher) Fetch(urls ...string) <-chan Result {
	for _, url := range urls {
		f.wg.Add(1)
		go f.fetch(url)
	}

	go func() {
		defer close(f.results)
		f.wg.Wait()
	}()

	return f.results
}

func (f *Fetcher) fetch(url string) {
	defer f.wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		f.results <- Result{url, 0, false, err}
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		f.results <- Result{url, 0, false, err}
	}

	f.results <- Result{url, len(bytes), true, nil}
}

// Result to indicate fetch result
type Result struct {
	URL     string
	Size    int
	Success bool
	Err     error
}

func (r Result) String() string {
	if r.Success {
		return fmt.Sprintf("%s = %d", r.URL, r.Size)
	}

	return fmt.Sprintf("%s = %s", r.URL, r.Err)
}
