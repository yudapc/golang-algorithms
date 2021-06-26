package strategypattern

import paymentmethods "golang-algorithms/cheats/design-patterns/strategy-pattern/payment-methods"

func ImplementationPayment() {
	name := "Budi"
	issuer := "Janji Jiwa"
	phone := "08123123123"
	amount := 10000
	paymentType := "cash"

	var paymentStrategy PaymentStrategy

	// pay with cash
	paymentStrategy = &paymentmethods.CashPayment{Amount: amount}
	paymentStrategy.Pay()

	// pay with card
	paymentStrategy = &paymentmethods.CardPayment{
		Name:   name,
		Type:   paymentType,
		Issuer: issuer,
		Amount: amount,
	}
	paymentStrategy.Pay()

	// pay with ovo
	paymentStrategy = &paymentmethods.OvoPayment{
		Phone:  phone,
		Amount: amount,
	}
	paymentStrategy.Pay()

}
