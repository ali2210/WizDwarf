package users

import (
	// paypalSdk "github.com/logpacker/PayPal-Go-SDK"
	"encoding/hex"
	"encoding/json"
	"strconv"

	paypal "github.com/logpacker/PayPal-Go-SDK"
)

var (
	persona string = ""
)

type (
	// Create_User struct {
	// 	Name     string
	// 	Fname    string
	// 	Madam    bool
	// 	Address  string 
	// 	Address2 string 
	// 	Zip      string
	// 	City     string
	// 	Country  string
	// 	Email    string
	// 	Password string
	// 	Secure   bool
	// }
	DigialProfile struct {

		// ledger my data
		Public  string
		Private string

		// visitor profile
		Name     string
		FName    string
		Email    string
		Address  string
		LAddress string
		City     string
		Zip      string
		Country  string
		Phone    string
		Twitter  string

		// credit card details
		Number      string
		ExpireMonth string
		ExpireYear  string
		Type        string
	}

	Visitors struct {
		Id       string `json:"id", omitempty`
		Name     string `json:"name", omitempty`
		Email    string `json:"email", omitempty`
		Password string `json:"password", omitempty`
		LastName    string `json:"lastname", omitempty`
		City     string `json:"city",omitempty `
		Zip      string `json:"zip", omitempty`
		Address  string `json:"address", omitempty`
		Apparment string `json:"apparment", omitempty`
		Country  string `json:"country", omitempty`
		Eve      bool   `json:"eve",omitempty`
		PhoneNo  string `json:"phone", omitempty`
		Twitter string 	`json:"twitter", omitempty`
		// Remember bool `json:"remember", omitempty`
	}

	CreditCardInfo interface {
		SetAuthorizeStoreID(id string)
		GetAuthorizeStoreID() string
		VoidStruct() *DigialProfile
	}

	DigitalPrint struct{}

	// UpdateProfile struct {
	// 	Id           string
	// 	FirstName    string
	// 	LastName     string
	// 	Phone        string
	// 	HouseAddress string
	// 	SubAddress   string
	// 	Country      string
	// 	Zip          string
	// 	Male         bool
	// 	Email        string
	// 	Twitter      string
	// 	City         string
	// }
)

func NewClient() CreditCardInfo {
	return &DigitalPrint{}
}

func (*DigitalPrint) SetAuthorizeStoreID(id string) {
	persona = id
}

func (*DigitalPrint) GetAuthorizeStoreID() string {
	return persona
}

func (*DigitalPrint) VoidStruct() *DigialProfile {
	return &DigialProfile{
		Public:      "",
		Private:     "",
		Name:        "",
		FName:       "",
		Email:       "",
		Address:     "",
		LAddress:    "",
		City:        "",
		Zip:         "",
		Country:     "",
		Phone:       "",
		Twitter:     "",
		Number:      "",
		ExpireMonth: "",
		ExpireYear:  "",
		Type:        "",
	}
}

type CalculationInterface interface {
	CalculateTotalBalance(st1, str2 float64) float64
	CalculateNum(str string) (float64, error)
	MarshalJson(pay paypal.PayoutResponse) ([]byte, error)
	MarshalJsonFees(pay *paypal.PaymentResponse) ([]byte, error)
	Encode(encode []byte) string
}

func (*Analysis) MarshalJSONAmount(pay *paypal.PayoutResponse) ([]byte, error) {

	return json.Marshal(pay.BatchHeader.Amount)
}

func (*Analysis) MarshalJSONFees(pay *paypal.PayoutResponse) ([]byte, error) {
	return json.Marshal(pay.BatchHeader.Fees)
}

func (*Analysis) Encode(encode []byte) string {

	return hex.EncodeToString(encode)
}

func (*Analysis) CalculateNum(str string) (float64, error) {

	return strconv.ParseFloat(str, 10)
}

func (*Analysis) CalculateTotalBalance(str1, str2 float64) float64 {
	return str1 + str2
}

type Analysis struct{}
