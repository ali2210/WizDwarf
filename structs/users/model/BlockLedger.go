package model

type DigialProfile struct{
	
	// ledger my data
	Public string
	Private string
	
	// visitor profile
	Name string
	FName string
	Email string
	Address string
	LAddress string
	City string
	Zip string
	Country string

	// credit card details
	Number string
	ExpireMonth string
	ExpireYear string
	Type string

}