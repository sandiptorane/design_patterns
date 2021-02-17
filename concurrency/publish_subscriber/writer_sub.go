package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type WriteSubscriber struct {
	in chan interface{}
	id int
	Writer io.Writer
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber{
	if out == nil{
		out = os.Stdout
	}
	s := &WriteSubscriber{
		in : make(chan interface{}),
		id : id,
		Writer: out,
	}
	go func(){
		for msg := range s.in{
			fmt.Fprintf(s.Writer,"(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}

func (w *WriteSubscriber)Notify(msg interface{}) (err error){
	defer func(){
		if rec := recover();rec != nil{
			err = fmt.Errorf("%v",rec)
		}
	}()

	select {
		case w.in <-msg :
		case <- time.After(time.Second):
			err = fmt.Errorf("timeout")
	}

	return
}

func (w *WriteSubscriber)Close(){
	close(w.in)
}
