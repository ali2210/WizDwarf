package cart

type Shopping struct {
	Price         string
	TypeofService string
	PaymentMethod string
	Description   string
}

type shoppingInterface interface {
	PlaceItemsInCart(price, service, payment, description string)
	GetItemsFromCart() Shopping
}

var order Shopping = Shopping{
	Price:         "",
	TypeofService: "",
	PaymentMethod: "",
	Description:   "",
}

type Items struct{}

func NewCart() shoppingInterface {
	return &Items{}
}

func (*Items) PlaceItemsInCart(price, service, payment, description string) {
	order = Shopping{
		Price:         price,
		TypeofService: service,
		PaymentMethod: payment,
		Description:   description,
	}
}

func (*Items) GetItemsFromCart() Shopping {
	return order
}
