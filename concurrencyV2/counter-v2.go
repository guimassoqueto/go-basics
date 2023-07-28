package main
import (
	"fmt"
	"sync"
	"gos/concurrencyV1"
)

var wg sync.WaitGroup

func counterV1(s string, countTo int) {
	for i := 0; i < countTo; i++ {
		fmt.Printf("%s: %d\n", s, i)
	}
	fmt.Print("\n")
}

func counterAsync(s string, countTo int) {
	for i := 0; i < countTo; i++ {
		fmt.Printf("%s: %d\n", s, i)
	}
	concurrencyV1.Runtime()
	fmt.Print("\n")
	wg.Done()
}

func main() {
	// interfaces.Main()
	// generics.Main()
	// pointers.BasicPointer()
	concurrencyV1.Runtime()
	wg.Add(2)
	go counterAsync("B---:", 10)
	go counterAsync("A+++:", 10)
	wg.Wait()
	fmt.Print("Done")
}