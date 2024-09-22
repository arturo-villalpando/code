package payments

import (
	"fmt"
)

type PayPal struct {
	email string
}

func NewPayPal(email string) *PayPal {
	return &PayPal{
		email: email,
	}
}

func (p *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Pay the amount %.2f with paypal email %s", amount, p.email)
}
