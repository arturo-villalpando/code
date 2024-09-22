package cart

import (
	i "strategy/items"
	p "strategy/payments"
)

type ShoppingCart struct {
	items         []i.Item
	paymentMethod p.PaymenthMethod
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{}
}

func (s *ShoppingCart) SetItem(item i.Item) {
	s.items = append(s.items, item)
}

func (s *ShoppingCart) GetItems() []i.Item {
	return s.items
}

func (s *ShoppingCart) SetPaymentMethod(paymentMethod p.PaymenthMethod) {
	s.paymentMethod = paymentMethod
}

func (s *ShoppingCart) GetPaymentMethod() p.PaymenthMethod {
	return s.paymentMethod
}

func (s *ShoppingCart) Checkout() string {
	var total float64
	for _, item := range s.items {
		total += item.GetPrice()
	}

	return s.paymentMethod.Pay(total)
}
