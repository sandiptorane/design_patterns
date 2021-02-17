package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(msg string) error
}

type PrinterImpl1 struct {

}

func (p *PrinterImpl1)PrintMessage(msg string) error{
	fmt.Printf(msg)
	return nil
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2)PrintMessage(msg string) error{
	if p.Writer == nil{
		return errors.New("you need to pass io.Writer to PrinterImpl2")
	}
	_,err := fmt.Fprintf(p.Writer,"%s",msg) //the fmt.Fprintf method takes an io.Writer interface as the first field and a message formatted as the rest, so we simply forward the contents of the msg argument to the io.Writer provided
	if err != nil{
		return err
	}
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter)Write(s[] byte)(n int,err error){
   n = len(s)
   if n>0 {
   	t.Msg = string(s)
	   return n, nil
   }
   err = errors.New("Content received on writer was empty\n")
   return
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct{
   Msg string
   Printer PrinterAPI
}

func (n *NormalPrinter)Print() error{
	_ = n.Printer.PrintMessage(n.Msg)
	return nil
}

type PacktPrinter struct {
	Msg string
	Printer PrinterAPI
}

func (p *PacktPrinter)Print() error{
	err := p.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", p.Msg))
	return err
}

