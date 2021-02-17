package main

import (
	"fmt"
	"strings"
)

type WorkerLauncher interface {
	LaunchWorker(chan Request)
}

type PrefixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (w *PrefixSuffixWorker)LaunchWorker(in chan Request){
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PrefixSuffixWorker) uppercase(in <-chan Request) <-chan Request{
	out := make(chan Request)

	go func(){
		for msg := range in {
			s,ok := msg.Data.(string)
			if !ok{
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) append(in <-chan Request) <-chan Request{
	out := make(chan Request)

	go func(){
		for msg := range in {
			uppercaseString,ok := msg.Data.(string)
			if !ok{
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s %s",uppercaseString,w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PrefixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercasedStringWithSuffix))
		}
	}()
}


