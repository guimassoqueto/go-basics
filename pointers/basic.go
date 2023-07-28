package pointers

import "fmt"

func BasicPointer() {
	var a int
	fmt.Printf("a =  %d, location: %v, type: %T.\n", a, &a, &a)

	a = 5
	fmt.Printf("a =  %d, location: %v, type: %T.\n", a, &a, &a)

	var p *int
	p = &a
	*p += 1
	fmt.Printf("&p = %v = &a = %v, p = %v = a = %v\n", p, &a, *p, a)
}