package observer

import (
	"testing"
	"time"
)

func TestObserver(t *testing.T) {

	subscriber1 := NewSubscriber("subscriber1")
	subscriber2 := NewSubscriber("subscriber2")
	str1 := "msg1"
	msg1 := (Msg)(str1)

	getHandler := func(want string) func(*Msg) {

		return func(got *Msg) {
			if string(*got) != want {
				t.Errorf("Computer.Print() got = %v, want %v", *got, want)
			}
		}

	}

	handler1 := getHandler(str1)

	subscriber1.NewHandler(handler1)
	subscriber2.NewHandler(handler1)

	topic1 := NewTopic("topic1")
	topic1.Register(subscriber1)
	topic1.Register(subscriber2)

	topic2 := NewTopic("topic2")
	topic2.Register(subscriber1)
	topic2.Register(subscriber2)

	topic1.NotifyAll(&msg1)
	topic2.NotifyAll(&msg1)

	time.Sleep(500 * time.Millisecond)
}
