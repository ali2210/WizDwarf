package handler


import (
	paypalSdk "github.com/logpacker/PayPal-Go-SDK"
)

const(
	
	// PAYPAL SANDBOX
	PaypalClientKey string = "AVWpbSSab2hTMsrGZ3oAeM6ALkXFKOVhYgbjb4nKN6Eu0sG-AnVRmPeGPbkPWBnF3n9w5by3JP6p2oGv"
	PaypalSecretKey string = "EFbCIMBM5s-shGaXx536oQbu0_QAzutn-_y1usOInfYrIvFNEIwhH_bR6KLv2wZrkZyJhjUk3qI17yp-"
	PaypalSandboxApi string = "sb-ylmhq4096831@business.example.com"
	
)
// Read html tag and convert into golang client


type PaypalClientLevel interface{
	NewClient()(*paypalSdk.Client, error)
	Token()(*paypalSdk.TokenResponse, error)
	RetrieveCreditCardInfo(id string)(*paypalSdk.CreditCard, error)
	StoreCreditCardInfo(c paypalSdk.CreditCard)(*paypalSdk.CreditCard, error)
	RemoveCard(id string) error
}

type PaypalMiniVersion struct{	
	
	Client *paypalSdk.Client
}

func PaypalClientGo() PaypalClientLevel  {
	return &PaypalMiniVersion{}
}

func (p *PaypalMiniVersion) NewClient()(*paypalSdk.Client, error)  {
	client , err := paypalSdk.NewClient(PaypalClientKey, PaypalSecretKey,PaypalSandboxApi)
	(*p).Client = client
	return (*p).Client, err
}

func (p *PaypalMiniVersion) Token() (*paypalSdk.TokenResponse, error) {
	
		token , err := (*p).Client.GetAccessToken()
		return token, err
}

func (p *PaypalMiniVersion) RemoveCard(id string) error  {
	return (*p).Client.DeleteCreditCard(id)
}

func (p *PaypalMiniVersion) RetrieveCreditCardInfo(id string)(*paypalSdk.CreditCard, error)  {
		
	c, err := (*p).Client.GetCreditCard(id)
	return c, err
}

func (p *PaypalMiniVersion) StoreCreditCardInfo(c paypalSdk.CreditCard)(*paypalSdk.CreditCard, error)  {
	cc , err := (*p).Client.StoreCreditCard(c)
	return cc, err
}
