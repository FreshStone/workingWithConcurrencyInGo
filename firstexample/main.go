package main
import (
	"fmt"
	"sync"
)

func print(s string, wg *sync.WaitGroup){ //wg must be passed by pointer mostly avoid copying
	defer wg.Done()
	fmt.Println(s)
}

func main(){
	var wg sync.WaitGroup
	wg.Add(5) //if number added here is greater than the number of goroutines returning than it will cause deadlock
	a := []string{"1","2","3","4","5"}
	for _, s := range a{
		go print(s, &wg)  //completion time of each go routine may be different 
	}
	//time.Sleep(1*time.Second()) //not the correct way
	wg.Wait()
	wg.Add(1) //because waitgroup counter decreased to zero previously
	print("finished", &wg)
}