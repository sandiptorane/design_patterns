package main

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

type mockSubscriber struct{
	notifyTestingFunc func(m interface{})
	closeTestingFunc  func()
}

func (m *mockSubscriber)Close(){
	m.closeTestingFunc()
}

func (m *mockSubscriber)Notify(msg interface{}) error{
	m.notifyTestingFunc(msg)
	return nil
}

func TestPublisher(t *testing.T){
     msg := "Hello"

     p := NewPublisher()
     go p.start()

     var wg sync.WaitGroup
    sub := &mockSubscriber{
     	notifyTestingFunc: func(msg interface{}){
     		defer wg.Done()

     		s, ok := msg.(string)
     		if !ok{
     			t.Fatal(errors.New("could not assert result"))
			}
            fmt.Println(msg)
			if s!=msg{
				t.Fail()
			}
		},
		closeTestingFunc: func(){
			wg.Done()
		},
	 }
	 p.AddSubscriberCh() <-sub

	 wg.Add(1)

    //publish message
	 p.PublishingCh() <-msg

	 pubCon := p.(*publisher)
	if len(pubCon.subscribers) != 1 {
		t.Error("Unexpected number of subscribers")
	}
	 wg.Add(1)
	 p.RemoveSubscriberCh() <-sub
	 wg.Wait()

	 //number of subscribers restored to zero
	 if len(pubCon.subscribers) != 0{
	 	t.Error("expected no subscribers")
	 }
	 p.Stop()
}


