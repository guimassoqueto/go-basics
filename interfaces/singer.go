package interfaces

import "fmt"

type Singer struct {
	Person
	BestSong string
}

// interface
func (singer Singer) talk() {
	fmt.Printf("I am a SINGER. My name is %s and my best SONG is %s.\n", singer.Name, singer.BestSong) 
}