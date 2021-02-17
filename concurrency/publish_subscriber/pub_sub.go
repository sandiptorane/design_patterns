package main

//Subscriber contains methods for new event
type Subscriber interface {
	Notify(i interface{}) error
	Close()
}
//Publisher to perform subscriber's operations
type Publisher interface {
	start()
	AddSubscriberCh()  chan<- Subscriber
	RemoveSubscriberCh()  chan<- Subscriber
	PublishingCh() chan<- interface{}
	Stop()
}

