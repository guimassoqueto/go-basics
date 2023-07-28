package concurrencyV1

import "fmt"


func counterV1(n int, v string) {
	for i := 0; i < n; i ++ {
		fmt.Printf("%s: %d\n", v, i)
	}
}

func CountersV1() {
	go counterV1(5, "C----")
	go counterV1(10, "A++++")
}