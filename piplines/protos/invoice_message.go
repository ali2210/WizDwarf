package protos

// counter_request hold counter call value
var counter_request *Pricing_Call_Request = &Pricing_Call_Request{}

//  number of keys
var Tnx_Keys int64 = 0

// Pricing Call Service Interface ensure function definition validation at module level.
type Pricing_Call_Service interface {
	Initial(value *Pricing_Call_Request)
	Inc() int64
}

// Set initial call request value
func Initial(value *Pricing_Call_Request) {

	counter_request.CRequest = value.CRequest
}

// Inc incement call request value
func Inc() int64 {

	return counter_request.CRequest + 1
}

// Incement key counter issue ticket key number ; so that data remain in order
func Incement() {

	Tnx_Keys += 1
}

// Get Keys returns last counter key number
func GetKeys() int64 {
	return Tnx_Keys
}
