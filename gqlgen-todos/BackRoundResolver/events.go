package BackRoundResolver

import "log"

type IEvent interface {
	Handle()
}

type EventSimple struct {
	Id  string
	Msg string
}

func (e EventSimple) Handle() {
	log.Println(e.Id + " " + e.Msg)
}
