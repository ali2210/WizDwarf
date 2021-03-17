package handler

import (
	paypal "github.com/logpacker/PayPal-Go-SDK"
	// paypalSdk "github.com/logpacker/PayPal-Go-SDK"
)

const (

	// PAYPAL SANDBOX
	PaypalClientKey string = "AS3poQQQOrbsHIyYYpz3M_XzHhG9xlgLSj6uARAmL4CH7_BzyQYoceurrSKImPww7hq0vLJSrQ4hDesw"
	PaypalSecretKey string = "EKtlmcsXpn0_9UtwTuWDHu_jfeyfiXFoJY1l4RY71VON_mXFlxkPnm53cJd8OIPc0VpouPyXV38RNBab"
)

// Read html tag and convert into golang client

type PaypalClientLevel interface {
	NewClient() (*paypal.Client, error)
	Token(client *paypal.Client) (*paypal.TokenResponse, error)
	RetrieveCreditCardInfo(id string, client *paypal.Client) (*paypal.CreditCard, error)
	StoreCreditCardInfo(c paypal.CreditCard, client *paypal.Client) (*paypal.CreditCard, error)
	RemoveCard(id string, client *paypal.Client) error
	PaypalPayout(id, serviceID, email, value string, client *paypal.Client) (*paypal.PayoutResponse, error)
	GetPayout(id string, client *paypal.Client) (*paypal.PayoutResponse, error)
}

type PaypalMiniVersion struct{}

func PaypalClientGo() PaypalClientLevel {
	return &PaypalMiniVersion{}
}

func (p *PaypalMiniVersion) NewClient() (*paypal.Client, error) {
	return paypal.NewClient(PaypalClientKey, PaypalSecretKey, paypal.APIBaseSandBox)
}

func (p *PaypalMiniVersion) Token(client *paypal.Client) (*paypal.TokenResponse, error) {

	return client.GetAccessToken()
}

func (p *PaypalMiniVersion) RemoveCard(id string, client *paypal.Client) error {
	return client.DeleteCreditCard(id)
}

func (p *PaypalMiniVersion) RetrieveCreditCardInfo(id string, client *paypal.Client) (*paypal.CreditCard, error) {

	return client.GetCreditCard(id)
}

func (p *PaypalMiniVersion) StoreCreditCardInfo(c paypal.CreditCard, client *paypal.Client) (*paypal.CreditCard, error) {

	return client.StoreCreditCard(c)
}

func (p *PaypalMiniVersion) PaypalPayout(id, serviceID, email, value string, client *paypal.Client) (*paypal.PayoutResponse, error) {

	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			EmailSubject:  "You recceive Service Confirm notification on the behave of WIZ-DWARFS",
			SenderBatchID: id},
		Items: []paypal.PayoutItem{{
			RecipientType: "Email",
			Receiver:      email,
			SenderItemID:  serviceID,
			Note:          "Thank you for service" + id,
			Amount:        &paypal.AmountPayout{Value: value, Currency: "USD"},
		}},
	}

	return client.CreateSinglePayout(payout)

}

func (p *PaypalMiniVersion) GetPayout(id string, client *paypal.Client) (*paypal.PayoutResponse, error) {
	return client.GetPayout(id)
}
