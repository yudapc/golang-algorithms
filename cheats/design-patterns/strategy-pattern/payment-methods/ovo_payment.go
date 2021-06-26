package paymentmethods

import (
	"fmt"
	"strconv"
)

type OvoPayment struct {
	Phone  string
	Amount int
}

func (o *OvoPayment) Pay() string {
	message := "paid with ovo with amount " + strconv.Itoa(o.Amount)
	fmt.Println(message)
	return message
}
