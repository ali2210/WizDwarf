package users

import (
	// paypalSdk "github.com/logpacker/PayPal-Go-SDK"
	"github.com/ali2210/wizdwarf/structs/users/model"
)

var (
	persona string = ""
)

type CreditCardInfo interface {
	SetAuthorizeStoreID(id string)
	GetAuthorizeStoreID() string
	VoidStruct() *model.DigialProfile
}

type DigitalPrint struct{}

func NewClient() CreditCardInfo {
	return &DigitalPrint{}
}

func (*DigitalPrint) SetAuthorizeStoreID(id string) {
	persona = id
}

func (*DigitalPrint) GetAuthorizeStoreID() string {
	return persona
}

func (*DigitalPrint) VoidStruct() *model.DigialProfile {
	return &model.DigialProfile{}
}
