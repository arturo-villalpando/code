package main

import (
	"fmt"

	s "strategy/cart"
	i "strategy/items"
	p "strategy/payments"
)

func main() {
	fmt.Println("Initializing the app")
	fmt.Println("--------------------")
	fmt.Println("Creating items")
	shampoo := i.NewItem("Quita caspa", 24.85)
	smartphone := i.NewItem("Smart Phone LK892X", 149.00)
	fmt.Println("--------------------")
	fmt.Println("We need add items to the shopping cart")
	shoppingCart := s.NewShoppingCart()
	shoppingCart.SetItem(*shampoo)
	shoppingCart.SetItem(*smartphone)
	fmt.Println("--------------------")
	fmt.Println("Print items in cart")
	cartItems := shoppingCart.GetItems()
	for _, item := range cartItems {
		itemPrint := fmt.Sprintf("The item '%s' has the price %f", item.GetName(), item.GetPrice())
		fmt.Println(itemPrint)
	}
	fmt.Println("--------------------")
	fmt.Println("We need to create the payment method")
	//creditCard := p.NewCreditCard("Arturo Villalpando", "1234-1234-1232-1111")
	paypal := p.NewPayPal("arturo.villalpando@gmail.com")
	fmt.Println("--------------------")
	fmt.Println("Add paymenth method to the shopping cart")
	shoppingCart.SetPaymentMethod(paypal)
	fmt.Println("--------------------")
	fmt.Println("Print checkout")
	fmt.Println(shoppingCart.Checkout())
}
