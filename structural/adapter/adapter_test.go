package adapter

import "testing"

func TestAdapter(t *testing.T) {
	msg := "Hello World"

	adapter := &PrinterAdapter{
		OldPrinter: &MyLegacyPrinter{}, Msg : msg,
	}

	returnMsg := adapter.PrintStored()

	if returnMsg != "Legacy Printer : Adapter : Hello World"{
		t.Errorf("msg didn't match: %s\n",returnMsg)
	}

	adapter = &PrinterAdapter{OldPrinter: nil,Msg: msg}
	returnMsg = adapter.PrintStored()

	if returnMsg != "Hello World"{
		t.Errorf("msg didn't match: %s\n",returnMsg)
	}
}
