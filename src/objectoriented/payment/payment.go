package payment

type CreditCard struct{
	ownername string
	cardnumber string
	expirationmonth int
	expirationyear int
	securitycode int 
	availablecredit float32
}
func CreateCreditAccount(ownername string,cardnumber string,expirationmonth int,expirationyear int,securitycode int) *CreditCard{
	return &CreditCard{
		ownername: ownername,
		cardnumber: cardnumber
		expirationmonth: expirationmonth,
		expirationyear: expirationyear,
		securitycode: securitycode,
	}
}
func(c CreditCard) getownername() string{
	return c.ownername
}
func(c CreditCard) getownername() string{
	return c.ownername
}
func(c CreditCard) getownername() string{
	return c.ownername
}