package golang

import (
	"fmt"
	"time"
)
//交替打印
func main () {
	num := 0
	chan1 := make(chan int)

	go func() {
		println(num)
		num++
		chan1 <- num
		for {

			printNum := <-chan1
			Print(printNum)
			printNum++
			chan1 <- printNum


		}

	}()
	go func() {

		for {

			printNum := <-chan1
			Print(printNum)
			printNum++
			chan1 <- printNum
		}
	}()
	time.Sleep(time.Second*2)
}

func Print (i int) {
	fmt.Println(i)
}

/*func main () {
	urls := []string{"a","b","c","d","e","f","g"}
	ctx := context.Background()
	timeoutContext, _:= context.WithTimeout(ctx, time.Second * 2)
	resChan := make (chan string, len(urls))
	wg := &sync.WaitGroup{}
	for i:=0;i<len(urls);i++ {
		wg.Add(1)
		go doSomething (timeoutContext, urls[i], resChan, wg)
	}
	wg.Wait()
	for i:=0; i< len(resChan);i++ {
		fmt.Println(fmt.Sprintf("res item=%s",<-resChan))
	}
}

func doSomething(ctx context.Context,url string,res chan <-string, wg *sync.WaitGroup) bool {
	resChan1 :=  make(chan string,1)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5-1)+1) * time.Second)
		resChan1<-url
		//fmt.Println(url)
	}()
	defer wg.Done()
	select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return false
		case v:= <-resChan1:
			fmt.Println("success")
			res <-v
			return true
	}
}*/
