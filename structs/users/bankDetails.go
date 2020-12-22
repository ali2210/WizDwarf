package users

import (
	paypalSdk "github.com/logpacker/PayPal-Go-SDK"
	"github.com/ali2210/wizdwarf/structs/users/model"
)

var(
	persona string = ""
	clientPaypal paypalSdk.UserInfo = paypalSdk.UserInfo{}
)


type DigialProfile struct{
	
	*paypalSdk.CreditCard
	*model.Vistors
	 Public string
}

type CreditCardInfo interface{
	SetAuthorizeNum(id string) 
	GetAuthorizeNum()string
	LinkCard(c *paypalSdk.CreditCard,  info *model.Vistors, eth string)(*DigialProfile)
	VoidStruct() *DigialProfile
}

type DigitalPrint struct{}


func NewClient() CreditCardInfo{
	return &DigitalPrint{}
}

func (*DigitalPrint) SetAuthorizeNum(id string)  {
	persona = id
}

func (*DigitalPrint) GetAuthorizeNum()string {
	return persona
}





func (*DigitalPrint) LinkCard(c *paypalSdk.CreditCard, info *model.Vistors, eth string)(*DigialProfile)  {
	avatarMe := DigialProfile{}
	avatarMe.CreditCard = c
	avatarMe.Vistors = info
	avatarMe.Public = eth
	return &avatarMe
}

func (*DigitalPrint) VoidStruct() *DigialProfile{
	return &DigialProfile{}
}

type Person struct{
	Id string
	FirstName string
	SureName string
	Address string
	StreetAddress string
	Postal string
	Email string
	PhoneNo string
	Gender string
}

// type CardHolder interface{
// 	AddUserInfo(p Person)(paypalSdk.UserInfo)
// 	GetUserInfo()(paypalSdk.UserInfo)
// }

// type Avatar struct{}

// func NewrPersona() CardHolder{
// 	return &Avatar{}
// }

// func (*Avatar) AddUserInfo(p Person)(paypalSdk.UserInfo){
	
// 	clientPaypal = paypalSdk.UserInfo{
// 		ID : p.Id,
// 		Name : p.FirstName,
// 		FamilyName : p.SureName,
// 		Email : p.Email,
// 		Phone : p.PhoneNo,
// 		Gender : p.Gender,
// 		Address : &paypalSdk.Address{
// 			Line1 : p.Address,
// 			Line2 : p.StreetAddress,
// 			PostalCode : p.Postal,
// 		},
// 	}
// 	return clientPaypal
// }

// func (*Avatar) GetUserInfo()(paypalSdk.UserInfo){
// 	 return clientPaypal
// }


