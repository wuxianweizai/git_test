package main

//haha
import "fmt"
import "sync"
import _ "net/http/pprof"
import "net/http"

var wg sync.WaitGroup

func Count(ch chan int, i int) {
	ch <- i
	fmt.Println("count")
	wg.Done()
}

//hello word  java go

func main() {
	wg.Add(11)
	//chs := make([]chan int,10)
	var chs = [10]chan int{}
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	go func() {
		http.ListenAndServe("0.0.0.0:8888", nil)
	}()

	wg.Wait()
}
