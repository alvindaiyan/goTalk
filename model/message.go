package model

type Message struct {
	Content       string
	UserIdSend    int
	UserIdReceive int
}

func Addmsg(a Message, c chan Message) {
	// todo this is a adapter sub
	c <- a
}
