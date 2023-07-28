package it

func Main() {
	michaelJackson := Singer {
		Person: Person{Name: "Michael Jackson", Age: 30},
		BestSong: "Billie Jean",
	}
	
	alPacino := Actor {
		Person: Person{Name: "Al Pacino", Age: 60},
		BestMovie: "Scarface",
	}
	
	saySomething(michaelJackson)
	michaelJackson.talk()

	saySomething(alPacino)
	alPacino.talk()
}