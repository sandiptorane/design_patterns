package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestWriter(t *testing.T){
	//sub := NewWriterSubscriber(0,nil)
}

func TestSubscriber(t *testing.T){
	sub := NewWriterSubscriber(0,nil)

	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)
	stdoutPrinter := sub.(*WriteSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(result string){
			if !strings.Contains(result, msg) {
				t.Fatal(fmt.Errorf("incorrect string: %s", result))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil{
		wg.Done()
		t.Error(err)
	}

	wg.Wait()
	sub.Close()
}

type mockWriter struct {
	testingFunc func(string)
}

func (m *mockWriter) Write(p []byte)(n int ,err error){
	m.testingFunc(string(p))
	return len(p),nil
}
