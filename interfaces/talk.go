package interfaces

type Talker interface {
	talk()
}

func saySomething(p Talker) {
	p.talk()
}