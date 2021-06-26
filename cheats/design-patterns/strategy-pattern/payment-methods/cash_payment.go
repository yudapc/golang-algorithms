package paymentmethods

type CashPayment struct {
	Amount int
}

func (c *CashPayment) Pay() string {
	return "paid with cash"
}
