package handler


import (
	paypalSdk "github.com/logpacker/PayPal-Go-SDK"
)

const(
	
	// PAYPAL SANDBOX
	PaypalClientKey string = "AcnBZ7cYRt1jj5Unnv34idX2kucc1RIDJHk91-V13ftIdR1QqF1_6cajyJLb0PsNOQAQ_Ivo3MIZrjpU"
	PaypalSecretKey string = "EKCNTKJBJBuc5SncrzKJRU1iuPDBqTZiiihw3VBGzHnFPyr_xisJba76XvpDbO75LXTcQX9W5-VdvIT_"
	PaypalSandboxApi string = "sb-z7blg3773073@personal.example.com"
)
// Read html tag and convert into golang client


type PaypalClientLevel interface{
	NewClient()(*paypalSdk.Client, error)
	Token()(*paypalSdk.TokenResponse, error)
	RetrieveCreditCardInfo(id string)(*paypalSdk.CreditCard, error)
	StoreCreditCardInfo(c paypalSdk.CreditCard)(*paypalSdk.CreditCard, error)
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

func (p *PaypalMiniVersion) RetrieveCreditCardInfo(id string)(*paypalSdk.CreditCard, error)  {
		
	c, err := (*p).Client.GetCreditCard(id)
	return c, err
}

func (p *PaypalMiniVersion) StoreCreditCardInfo(c paypalSdk.CreditCard)(*paypalSdk.CreditCard, error)  {
	cc , err := (*p).Client.StoreCreditCard(c)
	return cc, err
}
