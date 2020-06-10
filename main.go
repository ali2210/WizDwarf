package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"./db"
	"./structs"
	cloudWallet "./db/cloudwalletclass"
	"encoding/json"
	"google.golang.org/api/option"
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	contxt "context"
)

// Struts

type Response struct {
	Flag bool
	Message string
	Links string
}

type Create_User struct {
	name         string
	fname        string
	madam        bool
	address      string // World Coodinates
	address2     string // local coodinates
	zip          string
	city         string
	country      string
	email        string
	password     string
	secure  bool
}
type SignedKey struct {
	reader string
	signed string
	tx *ecdsa.PrivateKey
}

// Variables
 
var (
	emailexp string         = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp  string         = "([A-Z][a-z]*[0-9])*"
	addressexp string       = "(^0x[0-9a-fA-F]{40}$)"
	appName  *firebase.App  = SetFirestoreCredentials() // Google_Cloud [Firestore_Reference]
	cloud   db.DBFirestore = db.NewCloudInstance()
	ledger  db.PublicLedger = db.NewCollectionInstance()
	userSessions *sessions.CookieStore = nil
	clientInstance *ethclient.Client = nil
	ETHAddressInstance string = ""
)


// Constants

const (
	projectId          string = "htickets-cb4d0"
	Google_Credentials string = "/home/ali/Desktop/htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	// Main application
	EtherMainClientUrl  string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	// Rickeby for test purpose
	RinkebyClientUrl  	string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"	
)

// Functions

func main() {

	// Routing
	routing := mux.NewRouter()

	// Links 
	routing.HandleFunc("/{title}/home", Home)
	routing.HandleFunc("/{title}/signup", NewUser)
	routing.HandleFunc("/{title}/login", Existing)
	routing.HandleFunc("/{title}/dashboard",Dashboard)
	routing.HandleFunc("/{title}/logout", Logout)
	routing.HandleFunc("/{title}/createWallet",CreateWallet)
	routing.HandleFunc("/{title}/terms",Terms)
	routing.HandleFunc("/{title}/open", Wallet)
	routing.HandleFunc("/{title}/transact", Transacts)
	routing.HandleFunc("/{title}/transact/send", Send)
	// routing.HandleFunc("/{title}/transact/receive", Receive)

		// Static Files
	// routing.HandleFunc("/{title}/action", addVistor)
	// routing.HandleFunc("/{title}/data", getVistorData)
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	routing.PathPrefix("/images/").Handler(images)
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
	routing.PathPrefix("/css/").Handler(css)
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	routing.PathPrefix("/js/").Handler(js)
	routing.HandleFunc("/dummy", server)

		// Server
	log.Println("Listening at 9101 ... please wait...")
	http.ListenAndServe(":9101", routing)

}

// Routes Handle

func Home(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("index.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)																										
		temp.Execute(w, "Home")
	}

}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("dashboard.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Dashboard")
	} else {
		temp := template.Must(template.ParseFiles("server.html"))
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		// FILE Upload ....
		file := UploadFiles(w, r)
		if file != nil {
			println(file) // user file upload
			choose := r.FormValue("choose")
			println("choose I make:", choose)
			switch choose {
			case "0":
				fmt.Fprintf(w, "Please choose any option ...")
				temFile := template.Must(template.ParseFiles("dashboard.html"))
				temFile.Execute(w, "Home")
			case "1":
				var name string = "Covid-19"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				SequenceAligmentTable(file, svrFile)

			case "2":
				var name string = "FlaviDengue"
				svrFile := FileReadFromDisk(w, name)
				SequenceAligmentTable(file, svrFile)

			case "3":
				var name string = "KenyaEbola"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				SequenceAligmentTable(file, svrFile)

			case "4":
				var name string = "ZikaVirusBrazil"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				SequenceAligmentTable(file, svrFile)

			case "5":
				var name string = "MersSaudiaArabia"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				SequenceAligmentTable(file, svrFile)

			default:
				temFile := template.Must(template.ParseFiles("dashboard.html"))
				temFile.Execute(w, "Dashboard")
			}
		} else {
			print("size must be less than 5KB")
			Repon := Response{true,"Error in Upload File", "WizDawrf/dashboard"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
		}

	}

}


func Send(w http.ResponseWriter, r *http.Request){

	// temp := template.Must(template.ParseFiles("server.html"))
	
	block := structs.Block{}

	if(r.Method == "POST"){

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		r.ParseForm()
		block.TxSen = r.FormValue("sendAdd")
		block.TxRec = r.FormValue("add")
		choice :=  r.FormValue("transact")
		fmt.Println("Block:" , block)
		fmt.Println("choice:", choice)

	}
}


func CreateWallet(w http.ResponseWriter, r*http.Request){
	
	temp := template.Must(template.ParseFiles("seed.html"))
	acc := structs.Acc{} 

	if r.Method == "GET" {
		
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Seed")
	}else{
		
		temp := template.Must(template.ParseFiles("server.html"))
		fmt.Println("Method:"+ r.Method)
		r.ParseForm()
		acc.Email = r.FormValue("email")
		acc.Password = r.FormValue("password")
		
		if r.FormValue("agreeTerms") == "on"{
			acc.Terms = true
		}else{
			acc.Terms = false
		}
		
		client , err := ethclient.Dial(RinkebyClientUrl); if err != nil {
			fmt.Println("Error :" , err)
			return
		}
		
		fmt.Printf("Connection successfull ....%v\n", client)
		clientInstance = client
		
		println("Email:"+ acc.Email + "Password:"+ acc.Password)

		
		// private key 
		privateKey ,err := crypto.GenerateKey(); if err != nil{
			println("Error:" , err)
			return
		}

			// private key into bytes 
		PrvateKyByte := crypto.FromECDSA(privateKey)

		fmt.Println("Private_Key :" , hexutil.Encode(PrvateKyByte)[2:])

		pblicKey := privateKey.Public()

		pbcKey , ok := pblicKey.(*ecdsa.PublicKey); if !ok{
			println("Instaniate error {public key}")
			return
		}

		publicBytes := crypto.FromECDSAPub(pbcKey)
		fmt.Println("Public_Hash :" , hexutil.Encode(publicBytes)[4:])

		PublicKey := crypto.PubkeyToAddress(*pbcKey).Hex()
		fmt.Println("PublicKey:" , PublicKey)

		// hash 
		hshCode := sha3.NewLegacyKeccak256()
		hshCode.Write(publicBytes[1:])
		ethereum := hexutil.Encode(hshCode.Sum(nil)[12:])
		fmt.Println("Hash_sha3-256:", ethereum) //ethereum address
		acc.EthAddress = ethereum

			// valid address 
			valid := isYourPublcAdresValid(ethereum); if valid {
				// smart contract address
				fmt.Println("smart contract address :" , valid)
				Repon := Response{true,"Sorry! This is Smart Contact Adddress , We will handle in future", "WizDawrf/dashboard"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			fmt.Println("smart address:" , valid)

		myWallet := cloudWallet.EthereumWalletAcc{} 

		signWallet , err := json.Marshal(myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := Response{true,"Sorry! JSON Marshal Stream ", "WizDawrf/dashboard"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}

		err = json.Unmarshal(signWallet, &myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := Response{true,"Sorry! JSON Unmarshal Stream", "WizDawrf/dashboard"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}		

		//dataabse -- FindAddress 
		ok , ethAdd := FindAddress(&acc); if ok && ethAdd != nil {
				Repon := Response{true,"Sorry! Data Already register ", "WizDawrf/dashboard"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}
		fmt.Println("Eth_Add:" , ethAdd)

		myWallet.Email = acc.Email
		myWallet.Password = acc.Password
		myWallet.EthAddress = acc.EthAddress
		myWallet.Terms = acc.Terms

		merchant , err := ledger.CreatePublicAddress(&myWallet, appName); if err != nil{
				fmt.Println("Error:", err)
				Repon := Response{true,"Sorry! Invalid Ethereum Account ", "WizDawrf/dashboard"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}

		fmt.Println("merchant:" , merchant)
		clientInstance = nil
			// Server response
			// Repon := Response{false,acc.EthAddress, "WizDawrf/dashboard"}
			// println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			// temp.Execute(w, Repon)
		w.WriteHeader(http.StatusOK)
	    r.Method = "GET"
		Wallet(w,r)
	}
}

func Transacts(w http.ResponseWriter, r *http.Request){
	
	temp := template.Must(template.ParseFiles("transact.html"))	
	acc := structs.Acc{}
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		acc.EthAddress = ETHAddressInstance
		temp.Execute(w,acc)
	
	}
}


func Wallet(w http.ResponseWriter, r *http.Request){

	temp := template.Must(template.ParseFiles("wallet.html"))
	acc := structs.Acc{}

	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w,"Wallet")
	
	}else{

		temp := template.Must(template.ParseFiles("server.html"))
		// whtml := template.Must(template.ParseFiles("wallet.html"))

		fmt.Println("Method:"+ r.Method)
		r.ParseForm()

		acc.Email = r.FormValue("email")
		acc.Password = r.FormValue("password")
		
		client , err := ethclient.Dial(RinkebyClientUrl); if err != nil {
			fmt.Println("Error :" , err)
			return
		}
		
			fmt.Printf("Connection successfull ....%v\n", client)
		
		
			println("Email:"+ acc.Email + "Password:"+ acc.Password)

			myWallet := cloudWallet.EthereumWalletAcc{} 

			signWallet , err := json.Marshal(myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := Response{true,"Sorry! JSON Marshal Stream ", "WizDawrf/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}

			err = json.Unmarshal(signWallet, &myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := Response{true,"Sorry! JSON Unmarshal Stream", "WizDawrf/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}		
			add, ok := MyEthAddress(&acc); if !ok {
				Repon := Response{true,"Sorry! No Account Exist ", "WizDawrf/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			if add != nil{ 
				fmt.Println("Address:" , add)
				acc.EthAddress = add.EthAddress
				
				// variable address for futher processing
				ETHAddressInstance = acc.EthAddress
				fmt.Println("myWallet:", ETHAddressInstance)

			 // add this address in html page as static. 

		//dataabse -- FindAddress 
			secureWallet, ok := FindEthWallet(&acc); if !ok && secureWallet != nil {
				fmt.Println("Error", err)
				Repon := Response{true,"Sorry! No Account Exist ", "WizDawrf/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			fmt.Println("MyEthAddress Details:", secureWallet)


			w.WriteHeader(http.StatusOK)
	  	   	r.Method = "GET"
	     	Transacts(w,r)			
		}
	}
}


func Terms(w http.ResponseWriter, r *http.Request){

	temp := template.Must(template.ParseFiles("terms.html"))
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Terms")
	}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("register.html"))
	user := Create_User{}
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Regsiter")
	} else {
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		user.name = r.FormValue("uname")
		user.fname = r.FormValue("ufname")
		user.address = r.FormValue("address")
		user.address2 = r.FormValue("add")
		user.city = r.FormValue("inputCity")
		user.country = r.FormValue("co")
		user.zip = r.FormValue("inputZip")
		user.email = r.FormValue("email")
		user.password = r.FormValue("password")
		if r.FormValue("gender") == "on" {
			user.madam = true
		} else{
			user.madam = false
		}

		// println("Gender:", user.sir)
		// println("Gender2:", user.madam)


		matchE, err := regexp.MatchString(emailexp, user.email)
		if err != nil {
			println("invalid regular expression", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Data must be valid", "WizDawrf/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("regexp_email:", matchE)
		matchP, err := regexp.MatchString(passexp, user.password)
		if err != nil {
			println("invalid regular expression", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Data must be valid", "WizDawrf/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("regexp_pass:", matchP)

		// security
		hashRet, encrypted := MessageToHash(w, matchE, matchP, user)
		if hashRet == false {
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Data must be valid", "WizDawrf/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("encryted data", encrypted.reader)
		println("FamilyName:", user.fname)
		println("Address", user.address)
		println("Address2", user.address2)
		println("City", user.city)
		println("Zip", user.zip)
		println("Female", user.madam)
		println("Country", user.country)
		// println("check:", user.check_me_out)
		println("User record:", user.name, user.email)
		// println("phase:", KeyTx)
		addVistor(w, r, &user, encrypted.reader)
	}

}

func Existing(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("login.html"))
	user := Create_User{}
	if r.Method == "GET"{
		fmt.Printf("\nMethod:%v", r.Method)
		temp.Execute(w, "Login")	
	}else{
		// Parse Form
		r.ParseForm()
		fmt.Println("Method:\n", r.Method)
		user.email = r.FormValue("email")
		user.password = r.FormValue("password")
		if r.FormValue("check") == "on"{
			user.secure = true
		}else{
			user.secure = false
		}
		println("Login form data[", user.email, user.password, user.secure,"]")

		// Valid Data for processing
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(user.email)
		if !ok {
			println("invalid regular expression", !ok)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Data must be valid", "WizDawrf/login"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("regexp_email:", ok)
		reg := regexp.MustCompile(passexp)
		okx := reg.MatchString(user.password)
		if !okx {
			println("invalid regular expression", !okx)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Data must be valid", "WizDawrf/login"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
			// r.Method = "GET"
			// Existing(w,r)
		}
		println("regexp_pass:", okx)

		// Search Data in DB
		 data, err := SearchDB(w, r, user.email,user.password); if err != nil{
		 	// log.Fatal("Error", err)
		 	// w.Write([]byte(`{error: No Result Found }`))
		 	temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "No Record Exist", "WizDawrf/login"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
		 	return 
		 }
		 	if data != nil{
		 	fmt.Printf("Search Data:%v", data.Id)

		 	// User Session
		 	if userSessions == nil {
		 		userSessions = SessionsInit(data.Id)
		 		sessId , _ := userSessions.Get(r, "session-name")
		 		sessId.Values["authenticated"] = true
		 		err = sessId.Save(r,w); if err != nil{
		 		// log.Fatal("Error", err)
		 		temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry We have no Record, Please Regsiter", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(w, Res)
		 			return
		 		}
		 		println("Id :", sessId, "user:", userSessions)
		 }else{
		 	sessId , _ := userSessions.Get(r, "session-name")
		 	sessId.Values["authenticated"] = true
		 	err = sessId.Save(r,w); if err != nil{
		 		// log.Fatal("Error", err)
		 		temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry We donot have any record , please register", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(w, Res)
		 		return
		 	}
		 	println("Id :", sessId)
		 }
		
		
		 // Login page 
		 w.WriteHeader(http.StatusOK)
	    r.Method = "GET"
		Dashboard(w,r)
		}else{
				temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry We donot have any record , please register", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(w, Res)
		 		return
		}

	}
}

func server(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("server.html"))
	temp.Execute(w, "server")
}

func Logout(w http.ResponseWriter, r *http.Request){


	if r.Method == "GET"{
		println("User Session:", userSessions)
		 	sessId , _ := userSessions.Get(r, "session-name")
		 	sessId.Values["authenticated"] = false
		 	err := sessId.Save(r,w); if err != nil{
		 		// log.Fatal("Error", err)
		 		temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry We have no Record, Please Regsiter", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(w, Res)
		 		return
		 	}
		 	Existing(w,r)
	}
}


//  Advance Functions 


func SearchDB(w http.ResponseWriter, r *http.Request, email,pass string)(*db.Vistors, error){
	
	 var data *db.Vistors
	 var err error
	// w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
	} else {
		fmt.Println("Method:" + r.Method)
		data, err = cloud.FindData(email,pass, appName); if err != nil && data != nil{
			// log.Fatal("Error", err)
			println("Error:", err)
				/*temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry We have no Record, Please Regsiter", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(w, Res)*/
			return nil, err 
		}
		println("Data:", data)
	}
	return data, nil
}


func getVistorData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	visitor, err := cloud.FindAllData(appName)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" :"Error getting visitor result"}`))
		return
	}
	fmt.Printf("Vistors array%v", visitor)

	// response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(visitor)

}

func addVistor(response http.ResponseWriter, request *http.Request, user *Create_User, im string) {
	
	// var err error
	//response.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		fmt.Println("Method:" + request.Method)
	} else {
		var member db.Vistors
		data, err  := json.Marshal(member); if err != nil{
			fmt.Printf("Error in Marshal%v\n", err)
				temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry Data must be in Format", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(response, Res)
			return
		}
		err = json.Unmarshal(data, &member); if err != nil{
			fmt.Printf("Error%v\n", err)
				temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry Data Format Issue", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(response, Res)
			return
		}
		candidate , err := SearchDB(response, request, user.email,user.password); if err == nil && candidate == nil{
		 	// log.Fatal("Error", err)
		 	// w.Write([]byte(`{error: No Result Found }`))
		member.Id = im
		member.Name = user.name
		member.Email = user.email
		member.Password = user.password
		member.FName = user.fname
		if user.madam {
			member.Eve = user.madam
		}else{
			member.Eve = user.madam
		}
		member.Address = user.address
		member.LAddress	= user.address2																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																									
		member.City = user.city
		member.Zip = user.zip
		member.Country = user.country
		record ,err := cloud.SaveData(&member, appName); if err != nil {
			fmt.Printf("Error%v\n", err)
				temp := template.Must(template.ParseFiles("server.html"))
				Res := Response{true, "Sorry Data is not save yet", "WizDawrf/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(response, Res)
			return	
		}

		println("Record:", record.Id)
		response.WriteHeader(http.StatusOK)
		request.Method = "GET"
		Existing(response,request)
			return 
		}
			fmt.Printf("Search Data:%v", candidate.Email)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Your Data Already in DB", "WizDawrf/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(response, Res)
	}
	
}


// Functions

func SetFirestoreCredentials() *firebase.App {

	// set credentials
	conf := &firebase.Config{ProjectID: projectId}
	opt := option.WithCredentialsFile(Google_Credentials)
	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		println("Error in Connection with Firestore", err)
				// temp := template.Must(template.ParseFiles("server.html"))
				// Res := Response{true, "Sorry, Internet Connection Failed", ""}
				// println("Server Response:", Res.Flag,Res.Message,Res.Links)
				// temp.Execute(w, Res)
		return nil
	}
	println("Connected... Welcome to Firestore")
	return app
}

func FindAddress(w *structs.Acc)(bool, *cloudWallet.EthereumWalletAcc){

	ethAcc , err := ledger.FindMyPublicAddress(w, appName); if err != nil{
		fmt.Println("Error:", err)
		return false, nil
	}
	if ethAcc != nil{
		return false, nil	
	}
	return true, ethAcc

}

func MyEthAddress(w *structs.Acc)(*cloudWallet.EthereumWalletAcc, bool){

	acc , err := ledger.FindMyAddressByEmail(w, appName); if err != nil{
		fmt.Println("Error:", err)
		return nil, false
	}
	if acc == nil{
		return nil,false	
	}
	return acc,true		
}

func FindEthWallet(w *structs.Acc)(*cloudWallet.EthereumWalletAcc,bool){

	acc , err := ledger.FindMyPublicAddress(w, appName); if err != nil{
		fmt.Println("Error", err)
		return nil,false
	}
	return acc , true
}

func isYourPublcAdresValid(hash string) bool{


	expression := regexp.MustCompile(addressexp)
	v := expression.MatchString(hash)

	fmt.Println("Hash Valid:" , v)

	address := common.HexToAddress(hash)
	bytecode , err := clientInstance.CodeAt(contxt.Background(),address,nil); if err != nil{
		fmt.Println("Error:", err)
		return false
	}

	contract := len(bytecode)> 0
	return contract
}

func SessionsInit(unique string)(*sessions.CookieStore){
	return sessions.NewCookieStore([]byte(unique))
}

func FileReadFromDisk(w http.ResponseWriter, filename string) os.FileInfo {
	f, err := os.OpenFile(filename+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		println("FILE Open Error ... ", err)
		temp := template.Must(template.ParseFiles("server.html"))
		Res := Response{true, "Sorry Error in open File", "WizDawrf/dashboard"}
		println("Server Response:", Res.Flag,Res.Message,Res.Links)
		temp.Execute(w, Res)
		return nil
	}
	println("File Exist...", f)
	finfo, err := f.Stat()
	if err != nil {
		println("File Info not found", err)
		temp := template.Must(template.ParseFiles("server.html"))
		Res := Response{true, "Sorry, Server have NO INFORMATION", "WizDawrf/dashboard"}
		println("Server Response:", Res.Flag,Res.Message,Res.Links)
		temp.Execute(w, Res)
		return nil
	}
	println("File Info", finfo.Name())
	return finfo
}

func Key(w http.ResponseWriter, h1, h2 string) (string, string, *ecdsa.PrivateKey) {


		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			// panic(err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Sorry Error in encrytion", "WizDawrf/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return "", "", nil
		}

		// 0x40fa6d8c32594a971b692c44c0c56b19c32613deb1c6200c26ea4fe33d34a5fd

		println("PrivateKey", privateKey)
		msg := h1 + h2
		hash := sha256.Sum256([]byte(msg))


		fmt.Println("hash:",hash)
		r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
		println("Reader_reg:", rand.Reader)
		if err != nil {
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Sorry Error in encrytion", "WizDawrf/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			// panic(err)
			return "", "", privateKey
		}
		fmt.Printf("signature : (0x%x 0x%x)\n", r, s)
		return fmt.Sprintf("0x%x", r), fmt.Sprintf("0x%x", s),privateKey

}

func ReadSequence(filename string) ([]byte, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		// w.Write([]byte(`{error: Unable to Read File }`))
		return nil, err
	}
	// fmt.Printf("content %s:", body)
	return []byte(body), nil
}

func MessageToHash(w http.ResponseWriter,matchE, matchP bool, user Create_User) (bool, *SignedKey) {
	code := SignedKey{}
	if matchE && matchP {
		h := sha256.New()
		// h.Write([]byte(user.email))
		hashe := h.Sum([]byte(user.email))
		fmt.Println("email:", hex.EncodeToString(hashe))

		h1 := sha256.New()
		// h1.Write([]byte(user.password))
		hashp := h1.Sum([]byte(user.password))
		fmt.Println("pass:", hex.EncodeToString(hashp))
		code.reader, code.signed, code.tx = Key(w,hex.EncodeToString(hashe), hex.EncodeToString(hashp))
		// println("data get :", code.reader, code.signed)
		return true, &code
	}
	return false, &code
}

func UploadFiles(w http.ResponseWriter, r *http.Request) *os.File {
	// println("request body", r.Body)
	r.ParseMultipartForm(10 << 50)
	file, handler, err := r.FormFile("fileSeq")
	if err != nil {
		fmt.Println("Error failed.... retry", err)
		 temp := template.Must(template.ParseFiles("server.html"))
		 Res := Response{true, "Sorry Error in Upload File", "WizDawrf/dashboard"}
		 println("Server Response:", Res.Flag,Res.Message,Res.Links)
		temp.Execute(w, Res)
		return nil
	}
	defer file.Close()
	if handler.Size <= (50   * 1024) {
		fmt.Println("File name:" + handler.Filename)
		if _, err := os.Stat(handler.Filename); os.IsExist(err) {
			fmt.Println("File not exist ", err)
			 temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Sorry Error , No Such Directory", "WizDawrf/dashboard"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return nil
		}
		upldFile, err := ioutil.TempFile("user_data", handler.Filename+"-*.txt")
		if err != nil {
			fmt.Println("Error received while uploading!", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Sorry Upload File must have .txt extension ", "WizDawrf/dashboard"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return nil 
		}
		defer upldFile.Close()
		// file convert into bytes
		bytesFile, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error received while reading!", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := Response{true, "Error in Reading File", "WizDawrf/dashboard"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return nil
		}

		upldFile.Write(bytesFile)
		fmt.Println("File added on server")
		return upldFile
	}
	return nil
}

func SequenceAligmentTable(serverFile *os.File, userFile os.FileInfo) {


	// local variable liitle scoope
	seq, err := ReadSequence(userFile.Name())
	if err != nil {
		println("Error in read file", err)
		// temp := template.Must(template.ParseFiles("server.html"))
		// 	Res := Response{true, "Error in File", ""}
		// 	println("Server Response:", Res.Flag,Res.Message,Res.Links)
		// 	temp.Execute(w, Res)
		return
	}
  // fmt.Printf("Seq string:%s\n", seq)
	Useq, err := ReadSequence(serverFile.Name())
	if err != nil {
		println("Error in read file", err)
		// temp := template.Must(template.ParseFiles("server.html"))
		// 	Res := Response{true, "Error in File", ""}
		// 	println("Server Response:", Res.Flag,Res.Message,Res.Links)
		// 	temp.Execute(w, Res)
		return
	}

	println("Virus Dna sequence :")

	for _, v := range seq {
		// fmt.Printf("Seq:%v \t",  v ) // print bytes of array
		space := DoAscii(v)
		if space == "---" {
			fmt.Printf("%s\t", space)
		}
		fmt.Printf("%s\t", space)
	}
	println("Your Dna sequence :")
	for _, v := range Useq {
		uDna := DoAscii(v)
		if uDna == "---" {
			fmt.Printf("%s", uDna)
			
		}
		fmt.Printf("%s\t", uDna)

	}
}

func DoAscii(seq byte) string {
	if seq >= 65 && seq < 91 {
		return string(seq)
	}
	return "---"
}
