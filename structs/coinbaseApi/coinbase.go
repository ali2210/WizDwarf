package coinbaseApi

import (
	coin "github.com/fabioberger/coinbase-go"
)

// const (
// 	CoinbaseKey    string = "uGJWOhYrm7X2njjC"
// 	CoinbaseSecret string = "U3D0pf9uwDGMAniaFyV17t2cd2ODHwVc"
// )

type (
	CoinbaseThirdPartyAccess interface {
		NewClient(key, secret string) coin.Client
		GetEthIndex(from, to string, client coin.Client) (float64, error)
		GetEthValue(f, value float64) float64
	}

	StaticWallet struct {
		EthPrice  float64
		EthPrice2 float64
		EthPrice3 float64
	}

	Permission struct{}
)

func (*Permission) NewClient(key, secret string) coin.Client {

	if key == "" && secret == "" {
		key = "uGJWOhYrm7X2njjC"
		secret = "U3D0pf9uwDGMAniaFyV17t2cd2ODHwVc"
	}
	return coin.ApiKeyClient(key, secret)

}

func (*Permission) GetEthIndex(from, to string, client coin.Client) (float64, error) {
	return client.GetExchangeRate(from, to)
}

func (*Permission) GetEthValue(f, value float64) float64 {
	return value * f
}

func New() CoinbaseThirdPartyAccess {
	return &Permission{}
}
