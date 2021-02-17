package main

import (
	"fmt"
	"log"
	"sync"
)

type Request struct {
	Data interface{}
	Handler RequestHandler
}

type RequestHandler func(interface{})

func NewStringRequest(st string,id int,wg *sync.WaitGroup) Request{
	myRequest := Request{
		Data : fmt.Sprintf(st,id),
		Handler: func( i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			 if !ok{
			 	log.Fatal("Invalid Casting to string")
			 }
			 fmt.Println(s)
		},
	}
	return myRequest
}

func main(){
	buffersize := 100
	var dispatcher = NewDispatcherCreator(buffersize)

	workers := 3
	for i:=0; i< workers; i++{
		var w WorkerLauncher = &PrefixSuffixWorker{
			id : i,
			prefixS:fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: "World",
		}
		dispatcher.LaunchWorker(w)
	}
	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i:=0; i<requests;i++{
		req := NewStringRequest("Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}