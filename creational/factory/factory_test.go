package factory

import (
	"strings"
	"testing"
)

func  TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil{
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(103.0)
	if !strings.Contains(msg,"payment done using cash"){
		t.Fatal("the cash payment method message wasn't correct")
	}
	t.Log("LOG",msg)
}

func  TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil{
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(103.0)
	if !strings.Contains(msg,"payment done using debit card"){
		t.Fatal("the debit card payment method message wasn't correct")
	}
	t.Log("LOG",msg)
}

func  TestGetPaymentMethodNotExists(t *testing.T) {
    _, err := GetPaymentMethod(20)

	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}
	t.Log("LOG:",err)
}