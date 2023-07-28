package it

type Talker interface {
	talk()
}

func saySomething(p Talker) {
	p.talk()
}