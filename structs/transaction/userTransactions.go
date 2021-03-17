package transaction

type BankRecord struct {
	Name       string
	Btc        string
	CreditCard string
	TotalCash  string
	Public     string
}

var checkout BankRecord = BankRecord{
	Name:       "",
	Btc:        "",
	CreditCard: "",
	TotalCash:  "",
	Public:     "",
}

type FingerPrint struct{}

type wizDwarfsInterface interface {
	SetTransactionWiz(name, btc, credit, total, address string)
	GetTransactionWiz() BankRecord
}

func PlaceYourFinger() wizDwarfsInterface { return &FingerPrint{} }

func (*FingerPrint) SetTransactionWiz(name, btc, credit, total, address string) {
	checkout.Name = name
	checkout.Btc = btc
	checkout.CreditCard = credit
	checkout.TotalCash = total
	checkout.Public = address
}

func (*FingerPrint) GetTransactionWiz() BankRecord {
	return checkout
}
