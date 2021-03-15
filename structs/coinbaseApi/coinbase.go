package coinbaseApi

import (
	coin "github.com/fabioberger/coinbase-go"
)

const (
	APICoinbase string = "uGJWOhYrm7X2njjC"
	CoinbaseKey string = "U3D0pf9uwDGMAniaFyV17t2cd2ODHwVc"
)

type (
	CoinbaseThirdPartyAccess interface {
		NewClient() coin.Client
		GetEthIndex(from, to string, client *coin.Client) (float64, error)
		GetEthValue(f, value float64) float64
	}

	StaticWallet struct {
		EthPrice  float64
		EthPrice2 float64
		EthPrice3 float64
	}

	Permission struct{}
)

func (*Permission) NewClient() coin.Client {
	return coin.ApiKeyClient(APICoinbase, CoinbaseKey)
}

func (*Permission) GetEthIndex(from, to string, client *coin.Client) (float64, error) {
	return client.GetExchangeRate(from, to)
}

func (*Permission) GetEthValue(f, value float64) float64 {
	return value * f
}

func New() CoinbaseThirdPartyAccess {
	return &Permission{}
}
