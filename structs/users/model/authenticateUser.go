package model



type Create_User struct {
	Name         string
	Fname        string
	Madam        bool
	Address      string // World Coodinates
	Address2     string // local coodinates
	Zip          string
	City         string
	Country      string
	Email        string
	Password     string
	Secure  bool
}