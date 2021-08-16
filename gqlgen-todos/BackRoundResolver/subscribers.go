package BackRoundResolver

import "log"

type ISubscriber interface {
	GetId() string
	HandleMessage()
	GetStop() <-chan struct{}
	GetEvents() chan IEvent
}

type Subscriber struct {
	Id     string
	Stop   <-chan struct{}
	Events chan IEvent
}

type specialSubscriber struct {
	Subscriber
}

func (s Subscriber) GetEvents() chan IEvent {
	return s.Events
}

func (s Subscriber) GetStop() <-chan struct{} {
	return s.Stop
}

func (s Subscriber) GetId() string {
	return s.Id
}

func (s Subscriber) HandleMessage() {
	go func() {
		for {
			select {
			case ev := <-s.Events:
				log.Println("Simple subs", ev.Handle)
			}
		}
	}()
}

/// overrides method
func (s specialSubscriber) HandleMessage() {
	go func() {
		for {
			select {
			case ev := <-s.Events:
				log.Println("Special subs", ev.Handle)
			}
		}
	}()
}
