package BackRoundResolver

import "log"

type Backround_resolver struct {
	MyEvent     chan IEvent
	MySubcriber chan ISubscriber
}

func NewBackroundResolver() *Backround_resolver {
	r := &Backround_resolver{
		MyEvent:     make(chan IEvent),
		MySubcriber: make(chan ISubscriber),
	}

	return r
}

func (r *Backround_resolver) BroadcastEvent() {
	log.Println("Events listening started")
	subscribers := map[string]ISubscriber{}
	unsubscribe := make(chan string)
	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.MySubcriber:
			subscribers[s.GetId()] = s
		case e := <-r.MyEvent:
			for id, s := range subscribers {
				go func(id string, s ISubscriber) {
					select {
					case <-s.GetStop():
						unsubscribe <- id
						return
					default:
					}
					select {
					case <-s.GetStop():
						unsubscribe <- id
					case s.GetEvents() <- e:
						log.Println("Event pushed to", s.GetId())
					}
				}(id, s)
			}
		}
	}
}
