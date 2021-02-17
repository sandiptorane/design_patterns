package bridge

import (
	"strings"
	"testing"
)

func TestPrinterApi1(t *testing.T){
	api1 := PrinterImpl1{}
	err := api1.PrintMessage("hello world")
	if err!=nil{
		t.Errorf("Error trying to use api implementation : Message : %s\n",err.Error())
	}

}

func TestPrinterApi2(t *testing.T) {
	api2 := PrinterImpl2{}
    err := api2.PrintMessage("Hello")

    if err!=nil{
    	expectedMessage := "you need to pass io.Writer to PrinterImpl2"
    	if !strings.Contains(err.Error(),expectedMessage){
    		t.Errorf("Received message is not correct expected : %s \n actual : %s \n",expectedMessage,err.Error())
		}
	}

	testWriter := TestWriter{}

	api2 = PrinterImpl2{
		Writer: &testWriter,
	}
	expectedMessage := "Hello"

	err = api2.PrintMessage(expectedMessage)
	if err!=nil{
		t.Errorf("Error trying to use the API2 implementation: %s\n", err.Error())
	}
	if testWriter.Msg != expectedMessage {
		t.Fatalf("API2 did not write correctly on the io.Writer. \n Actaul: %s\nExpected: %s\n", testWriter.Msg, expectedMessage)
	}
}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normal := &NormalPrinter{
		Msg : expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normal.Print()

	if err!=nil {
		t.Error(err)
	}

	testWriter := TestWriter{}

	normal = &NormalPrinter{
		Msg : expectedMessage,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = normal.Print()

	if err!=nil{
		t.Error(err.Error())
	}
     actualMsg := testWriter.Msg
	assertPrinterMsg(t,actualMsg,expectedMessage)

}

func TestPacktPrinter_Print(t *testing.T) {
	message := "Hello io.Writer"

	packt := &PacktPrinter{
		Msg : message,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()

	if err!=nil {
		t.Error(err)
	}

	testWriter := TestWriter{}

	packt = &PacktPrinter{
		Msg : message,
		Printer: &PrinterImpl2{
			Writer: &testWriter,
		},
	}

	err = packt.Print()

	if err!=nil{
		t.Error(err.Error())
	}
	actualMsg := testWriter.Msg
	expectedMessage := "Message from Packt: Hello io.Writer"
	assertPrinterMsg(t,actualMsg,expectedMessage)

}

func assertPrinterMsg(t *testing.T,actualMsg,expectedMsg string){
	if actualMsg != expectedMsg{
		t.Errorf("The expected Message on the io.Writer doesn't match to actual message : \n Actual : %s \n expected : %s",actualMsg,expectedMsg)
	}
}
