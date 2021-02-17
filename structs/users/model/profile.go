package model

type UpdateProfile struct {
	Id           string
	FirstName    string
	LastName     string
	Phone        string
	HouseAddress string
	SubAddress   string
	Country      string
	Zip          string
	Male         bool
	Email        string
	Twitter      string
	City         string
}

type Vistors struct {
	Id       string `json:"Id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	FName    string `json:"FName"`
	City     string `json:"City"`
	Zip      string `json:"Zip"`
	Address  string `json:"Address"`
	LAddress string `json:"LAddress"`
	Country  string `json:"Country"`
	Eve      bool   `json:"Eve"`
}
