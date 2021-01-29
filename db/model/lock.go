package model


import "errors"

type CredentialsPrivate struct{
	
	PublicAddress string
	PrvteKey string
	
}

var pairvalue map[int]string

type Private interface{
	
	SetCryptoDB(key , value string)error
	GetCryptoDB(key string)(string, string)
}

type MapsValue struct{}

func New() Private{
	return &MapsValue{}
}

func(*MapsValue) SetCryptoDB(key , value string) error{
	pairvalue = make(map[int]string)
	if key == "" && value == ""{
		return errors.New("Empty data")
	}else{
		pairvalue[0] = value
	}
	return nil
}
func(*MapsValue) GetCryptoDB(key string)(string,string){
	pairvalue = make(map[int]string)
	if key != "" {
		return pairvalue[0], key
	}
	return "",  ""
}