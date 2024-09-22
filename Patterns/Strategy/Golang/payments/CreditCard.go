package payments

import "fmt"

type CreditCard struct {
	name string
	card string
}

func NewCreditCard(name string, card string) *CreditCard {
	return &CreditCard{
		name: name,
		card: card,
	}
}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Pay the amount %f with credit card number %s and name %s", amount, c.card, c.name)
}
