package users

import (
	// paypalSdk "github.com/logpacker/PayPal-Go-SDK"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/ali2210/wizdwarf/structs/users/model"
	paypal "github.com/logpacker/PayPal-Go-SDK"
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
