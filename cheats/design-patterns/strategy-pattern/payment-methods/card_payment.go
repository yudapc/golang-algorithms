package paymentmethods

type CardPayment struct {
	Name   string
	Type   string
	Issuer string
	Amount int
}

func (c *CardPayment) Pay() string {
	return "paid with card"
}
