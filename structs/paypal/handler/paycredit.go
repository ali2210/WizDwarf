package handler


import (
	paypalSdk "github.com/logpacker/PayPal-Go-SDK"
)

const(
	
	// PAYPAL SANDBOX
	PaypalClientKey string = "AS3poQQQOrbsHIyYYpz3M_XzHhG9xlgLSj6uARAmL4CH7_BzyQYoceurrSKImPww7hq0vLJSrQ4hDesw"
	PaypalSecretKey string = "EKtlmcsXpn0_9UtwTuWDHu_jfeyfiXFoJY1l4RY71VON_mXFlxkPnm53cJd8OIPc0VpouPyXV38RNBab"
	
)
// Read html tag and convert into golang client


type PaypalClientLevel interface{
	NewClient()(*paypalSdk.Client, error)
	Token(client *paypalSdk.Client)(*paypalSdk.TokenResponse, error)
	RetrieveCreditCardInfo(id string, client *paypalSdk.Client)(*paypalSdk.CreditCard, error)
	StoreCreditCardInfo(c paypalSdk.CreditCard, client *paypalSdk.Client)(*paypalSdk.CreditCard, error)
	RemoveCard(id string, client *paypalSdk.Client) error
}

type PaypalMiniVersion struct{}

func PaypalClientGo() PaypalClientLevel  {
	return &PaypalMiniVersion{}
}

func (p *PaypalMiniVersion) NewClient()(*paypalSdk.Client, error)  {
	client , err := paypalSdk.NewClient(PaypalClientKey, PaypalSecretKey,paypalSdk.APIBaseSandBox)
	return client, err
}

func (p *PaypalMiniVersion) Token(client *paypalSdk.Client) (*paypalSdk.TokenResponse, error) {
	
		token , err := client.GetAccessToken()
		return token, err
}

func (p *PaypalMiniVersion) RemoveCard(id string, client *paypalSdk.Client,) error  {
	return client.DeleteCreditCard(id)
}

func (p *PaypalMiniVersion) RetrieveCreditCardInfo(id string, client *paypalSdk.Client)(*paypalSdk.CreditCard, error)  {
		
	c, err := client.GetCreditCard(id)
	return c, err
}

func (p *PaypalMiniVersion) StoreCreditCardInfo(c paypalSdk.CreditCard, client *paypalSdk.Client)(*paypalSdk.CreditCard, error)  {
	cc , err := client.StoreCreditCard(c)
	return cc, err
}
