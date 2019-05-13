package main

import (
    "sync"
    "fmt"
    "strconv"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

/*
func main() {
    var wg sync.WaitGroup
    var urls = []string{
        "http://www.golang.org/",
        "http://www.google.com/",
        "http://www.somestupidname.com/",
    }
    for _, url := range urls {
        // Increment the WaitGroup counter.
        wg.Add(1)
        // Launch a goroutine to fetch the URL.
        go func(url string) {
            // Decrement the counter when the goroutine completes.
            defer wg.Done()
            // Fetch the URL.
            http.Get(url)
        }(url)
    }
    // Wait for all HTTP fetches to complete.
    wg.Wait()
}
*/

func main() {
    var wg sync.WaitGroup
    ch := make(chan int, 1000)
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go doSomething(i, &wg, ch)
    }
    wg.Wait()
    fmt.Println("all done")
    for i := 0; i < 1000; i++ {
        dd := <-ch
        fmt.Println("from ch:"+strconv.Itoa(dd))
    }
}

func doSomething(index int, wg  *sync.WaitGroup, ch chan int) {
    wg.Done()

    fmt.Println("start done:" + strconv.Itoa(index))
    //time.Sleep(20 * time.Millisecond)
    ch <- index
}