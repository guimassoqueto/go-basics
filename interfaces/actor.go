package interfaces

import "fmt"


type Actor struct {
	Person
	BestMovie string
}

func (actor Actor) talk() {
	fmt.Printf("I am an ACTOR. My name is %s and my best MOVIE is %s.\n", actor.Name, actor.BestMovie) 
}