package users

import (
	// paypalSdk "github.com/logpacker/PayPal-Go-SDK"
	 "github.com/ali2210/wizdwarf/structs/users/model"
)

var(
	persona string = ""
	// clientPaypal paypalSdk.UserInfo = paypalSdk.UserInfo{}
)




type CreditCardInfo interface{
	SetAuthorizeStoreID(id string) 
	GetAuthorizeStoreID()string
	// LinkCard(c *paypalSdk.CreditCard,  info *model.Vistors, publicAddress *string)(DigialProfile)
	VoidStruct() *model.DigialProfile
}

type DigitalPrint struct{}


func NewClient() CreditCardInfo{
	return &DigitalPrint{}
}

func (*DigitalPrint) SetAuthorizeStoreID(id string)  {
	persona = id
}

func (*DigitalPrint) GetAuthorizeStoreID()string {
	return persona
}





// func (*DigitalPrint) LinkCard(c *paypalSdk.CreditCard, info *model.Vistors, publicAddress *string)(DigialProfile)  {
// 	avatarMe := &DigialProfile{}
// 	avatarMe.CreditCard = c
// 	avatarMe.Visitors = info
// 	avatarMe.Public = publicAddress
// 	return *avatarMe
// }

func (*DigitalPrint) VoidStruct() *model.DigialProfile{
	return &model.DigialProfile{}
}

// type Person struct{
// 	Id string
// 	FirstName string
// 	SureName string
// 	Address string
// 	StreetAddress string
// 	Postal string
// 	Email string
// 	PhoneNo string
// 	Gender string
// }

