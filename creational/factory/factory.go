package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64)string
}

const(
	Cash = 1
	DebitCard = 2
)

func GetPaymentMethod(m int)(PaymentMethod,error){
	switch m {
		case Cash:
			return new(CashPM),nil
		case DebitCard:
			return new(DebitCardPM),nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct {}

type DebitCardPM struct{}

func (c *CashPM)Pay(amount float64) string{
	return fmt.Sprintf("%0.2f payment done using cash \n",amount)
}

func (d *DebitCardPM)Pay(amount float64) string{
	return fmt.Sprintf("%0.2f payment done using debit card \n",amount)
}

