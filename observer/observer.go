package observer

import (
	"sync"
)

type ISubscriber interface {
	Pub(*Msg)
}

type Msg string

type Subscriber struct {
	id      string
	Channel chan *Msg
}

func (s *Subscriber) Pub(m *Msg) {
	s.Channel <- m
}

func (s *Subscriber) SetHandler(f func(m *Msg)) {
	go func() {
		for m := range s.Channel {
			f(m)
		}
	}()
}

type Topic struct {
	Subscribers []ISubscriber
	name        string
	mux         sync.Mutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		name: name,
	}
}

func NewSubscriber(id string) *Subscriber {
	return &Subscriber{
		id:      id,
		Channel: make(chan *Msg, 2),
	}
}

func (t *Topic) Register(subscriber ISubscriber) {
	t.mux.Lock()
	t.Subscribers = append(t.Subscribers, subscriber)
	t.mux.Unlock()
}

func (t *Topic) NotifyAll(m *Msg) {
	for _, subscriber := range t.Subscribers {
		subscriber.Pub(m)
	}
}
