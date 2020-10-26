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
	"github.com/ali2210/wizdwarf/db"
	"github.com/ali2210/wizdwarf/structs"
	cloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
	"encoding/json"
	"google.golang.org/api/option"
	"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	contxt "context"
	"github.com/biogo/biogo/alphabet"
	"strconv"
	"math/big"
	// "strings"
	"github.com/fogleman/ribbon/pdb"
	"github.com/ali2210/wizdwarf/structs/amino"

)


// Variables

var (
	emailexp string         = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp  string         = "([A-Z][a-z]*[0-9])*"
	addressexp string       = "(^0x[0-9a-fA-F]{40}$)"
	appName  *firebase.App  = SetFirestoreCredentials() // Google_Cloud [Firestore_Reference]
	cloud   db.DBFirestore = db.NewCloudInstance()
	ledger  db.PublicLedger = db.NewCollectionInstance()
	userSessions *sessions.CookieStore = nil //user level
	cryptoSessions *sessions.CookieStore = nil // crypto level
	clientInstance *ethclient.Client = nil
	ETHAddressInstance string = ""
	WalletPubKey string = ""
	WalletSecureKey string  = ""
	// LifeCode []amino.AminoClass
	configFilename string = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	googleCredentials string = ""
	FILENAME string = ""
	edit structs.Levenshtein = structs.Levenshtein{}
	// openStreet structs.Address = structs.Address{}
	visualizeReport structs.DataVisualization  = structs.DataVisualization{}

	/*_, b, _, _ = runtime.Caller(0)
    basepath   = filepath.Dir(b)*/
)


// Constants

const (
	projectId          string = "htickets-cb4d0"
	//Google_Credentials string = "/home/ali/Desktop/htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	// Main application
	EtherMainClientUrl  string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	// Rickeby for test purpose
	RinkebyClientUrl  	string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	openwizweather string = "7efdb33c59a74e09352479b21657aee8"

)

// Functions

func main() {

		// Server
		fmt.Println("[OK] Wiz-Dwarfs starting")
		port := os.Getenv("PORT")

		if port == " "{
			log.Println("[Fail] No Application port allocated")
			return
		}else{
			if port != "5000"{
				// any Listening PORT {heroku}
				port =  " "
				log.Println("[Open] Application Port")
			}else{
				// specfic port allocated {docker}
				port = "5000"
				log.Println("[New] Application Default port")
			}

		}

		log.Println("[OK] Application :" , port + " Port")
		// Routing
		routing := mux.NewRouter()

		// Links
		routing.HandleFunc("/home", Home)
		routing.HandleFunc("/signup", NewUser)
		routing.HandleFunc("/login", Existing)
		routing.HandleFunc("/dashboard",Dashboard)
		routing.HandleFunc("/logout", Logout)
		routing.HandleFunc("/createWallet",CreateWallet)
		routing.HandleFunc("/terms",Terms)
		routing.HandleFunc("/open", Wallet)
		routing.HandleFunc("/transact", Transacts)
		routing.HandleFunc("/transact/send", Send)
		routing.HandleFunc("/transact/treasure", Treasure)
		routing.HandleFunc("/visualize", Visualize)
		/*routing.HandleFunc("/transact/advance-fileoption", Blocks)*/

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



		err := http.ListenAndServe(":"+port, routing); if err != nil{
			log.Println("Listening Error: ", err)
			panic(err)
		}

}

// Routes Handle

func Home(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("index.html"))

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Home")
	}

}


func Visualize(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("visualize.html"))
	fmt.Println("report : " , visualizeReport.Percentage , visualizeReport.UVinfo)
	if r.Method == "GET" && edit.Probablity >= 0.0 {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		 temp.Execute(w,visualizeReport)
	}
	// err := SessionExpire(w,r); if err != nil {
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// r.Method = "GET"
	// Dashboard(w,r)
}


func Treasure(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("treasure.html"))
	acc := structs.Static{}
	block := structs.Block{}

	if r.Method == "GET" {

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		acc.Eth = ETHAddressInstance
		acc.Balance = GetBalance(&acc); if acc.Balance == nil{
			fmt.Println("Error:")
		}
		fmt.Println("Details:", acc )
		temp.Execute(w,acc)

	}else{

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		r.ParseForm()
		block.TxSen = r.FormValue("send")
		block.TxRec = r.FormValue("rece")

		fmt.Println("block:", block)

		// Block Number
		blockNumber := big.NewInt(6677972)
		bNum , err := clientInstance.BlockByNumber(context.Background(), blockNumber); if err != nil {
				fmt.Println("Error:",err)
				return
		}

		fmt.Println("Length:", len(bNum.Transactions()))
		fmt.Println("Hash:", bNum.Hash().Hex())
		fmt.Println("Block:", bNum.Number().Uint64())


		for _ , tx := range bNum.Transactions(){
			fmt.Println("To:", tx.To().Hex())
			fmt.Println("Block_Hash:", tx.Hash().Hex())

			// ChainId

			chainId, err := clientInstance.NetworkID(context.Background()); if err != nil {
				fmt.Println("Error:",err)
				return
			}


			// get recipt address
			message , err := tx.AsMessage(types.NewEIP155Signer(chainId)); if err != nil {
				fmt.Println("Error:",err)
				return
			}

			fmt.Println("Message From:", message.From().Hex())

			recp, err := clientInstance.TransactionReceipt(context.Background(), tx.Hash()); if err != nil {
				fmt.Println("Error:",err)
				return
			}

			fmt.Println("Status:", recp.Status)
		}

 		txs := common.HexToHash(bNum.Hash().Hex())
		fmt.Println("Tx:", txs.Hex())

		// Number of Transaction
		count , err := clientInstance.TransactionCount(context.Background(), txs); if err != nil {
				fmt.Println("Error:",err)
				return
		}
		fmt.Println("Num #:", count)

		for i := uint(0); i < count; i++{
			Tx , err := clientInstance.TransactionInBlock(context.Background(), txs, i); if err != nil {
				fmt.Println("Error:",err)
				return
	 		}

	 		fmt.Println("Tx Hash:", Tx.Hash().Hex())

	 		txHash := common.HexToHash(Tx.Hash().Hex())

	 		// Transaction status
	 		tx , isPending, err := clientInstance.TransactionByHash(context.Background(), txHash); if err != nil {
				fmt.Println("Error:",err)
				return
			}

			fmt.Println("hash:", tx.Hash().Hex())
			fmt.Println("Transaction_Pending:", isPending)

	 	}

	}
}


/*func Blocks(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("swarm.html"))

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Url:", r.URL.Path)
		temp.Execute(w, "Blocks")
	}
}*/

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("dashboard.html"))

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Url:", r.URL.Path)
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
				meGenome ,virusGenome , err := SequenceFile(file, svrFile); if err != nil{
					fmt.Println("Error:", err)
					return
				}
				fmt.Println("Genome:", len(meGenome), "virusGenome:", len(virusGenome))
				distance := structs.EditDistanceStrings(meGenome,virusGenome)
				edit.Probablity = edit.Result(distance)
				edit.Name = name
				edit.Percentage = edit.CalcualtePercentage(edit.Probablity)
				visualizeReport.Percentage = edit.Percentage
				// openStreet.Country = r.FormValue("country")
				// openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				// openStreet.RouteNum  = i
				// fmt.Println("openStreet:", openStreet)
				//  street , err := openStreet.CurrentLocationByPostalAddress(openStreet);if err != nil {
				// 	fmt.Println("Error:", err)
				// 	return
				//  }
				//  fmt.Println("Street:", street)
				//  uv ,err := visualizeReport.OpenWeather(openwizweather); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  loc := visualizeReport.GetCoordinates(street)
				//
				//  if err := uv.Current(loc); err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  info , err := uv.UVInformation(); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  visualizeReport.UVinfo = info
				w.WriteHeader(http.StatusOK)
				// LifeCode = genome
	    		r.Method = "GET"
	    		// Wallet(w,r)
	    		Visualize(w,r)
				// fmt.Println("Virus:", capsid)

			case "2":
				var name string = "FlaviDengue"
				svrFile := FileReadFromDisk(w, name)
				meGenome ,virusGenome,  err := SequenceFile(file, svrFile); if err != nil{
					fmt.Println("Error:", err)
					return
				}
				distance := structs.EditDistanceStrings(meGenome,virusGenome)
				edit.Probablity = edit.Result(distance )
				edit.Name = name
				edit.Percentage = edit.CalcualtePercentage(edit.Probablity)
				//  openStreet.Country = r.FormValue("country")
				//  openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				// openStreet.RouteNum  = i
				// fmt.Println("state:", openStreet)
				// street , err := openStreet.CurrentLocationByPostalAddress(openStreet);if err != nil {
				// 	fmt.Println("Error:", err)
				// 	return
				//  }
				//  fmt.Println("Street:", street)
				//  uv ,err := visualizeReport.OpenWeather(openwizweather); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  loc := visualizeReport.GetCoordinates(street)
				//
				//  if err := uv.Current(loc); err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  info , err := uv.UVInformation(); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  visualizeReport.UVinfo = info
				w.WriteHeader(http.StatusOK)
	    		r.Method = "GET"
				Visualize(w,r)
				// Wallet(w,r)
				// fmt.Println("Virus:", capsid)
			case "3":
				var name string = "KenyaEbola"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				meGenome ,virusGenome, err := SequenceFile(file, svrFile); if err != nil{
					fmt.Println("Error:", err)
					return
				}
				distance := structs.EditDistanceStrings(meGenome,virusGenome)
				edit.Probablity = edit.Result(distance )
				edit.Name = name
				edit.Percentage = edit.CalcualtePercentage(edit.Probablity)
				//  openStreet.Country = r.FormValue("country")
				//  openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				// openStreet.RouteNum  = i
				// fmt.Println("state:", openStreet)
				//  street , err := openStreet.CurrentLocationByPostalAddress(openStreet);if err != nil {
				// 	fmt.Println("Error:", err)
				// 	return
				//  }
				//  fmt.Println("Street:", street)
				//  uv ,err := visualizeReport.OpenWeather(openwizweather); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  loc := visualizeReport.GetCoordinates(street)
				//
				//  if err := uv.Current(loc); err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  info , err := uv.UVInformation(); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  visualizeReport.UVinfo = info
				w.WriteHeader(http.StatusOK)
	    		r.Method = "GET"
				Visualize(w,r)

				// fmt.Println("Virus:", capsid)
			case "4":
				var name string = "ZikaVirusBrazil"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				meGenome ,virusGenome, err := SequenceFile(file, svrFile); if err != nil{
					fmt.Println("Error:", err)
					return
				}
				distance := structs.EditDistanceStrings(meGenome,virusGenome)
				edit.Probablity = edit.Result(distance )
				edit.Name = name
				edit.Percentage = edit.CalcualtePercentage(edit.Probablity)
				// openStreet.Country = r.FormValue("country")
				// openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				// openStreet.RouteNum  = i
				// fmt.Println("state:", openStreet)
				//  street , err := openStreet.CurrentLocationByPostalAddress(openStreet);if err != nil {
				// 	fmt.Println("Error:", err)
				// 	return
				//  }
				//  fmt.Println("Street:", street)
				//  uv ,err := visualizeReport.OpenWeather(openwizweather); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  loc := visualizeReport.GetCoordinates(street)
				//
				//  if err := uv.Current(loc); err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  info , err := uv.UVInformation(); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  visualizeReport.UVinfo = info
				w.WriteHeader(http.StatusOK)
	    		r.Method = "GET"
				Visualize(w,r)

				// fmt.Println("Virus:", capsid)
			case "5":
				var name string = "MersSaudiaArabia"
				svrFile := FileReadFromDisk(w, name)
				println("Please Wait", svrFile.Name(), "...")
				meGenome ,virusGenome, err  := SequenceFile(file, svrFile); if err != nil{
					fmt.Println("Error:", err)
					return
				}
				distance := structs.EditDistanceStrings(meGenome,virusGenome)
				edit.Probablity = edit.Result(distance )
				edit.Name = name
				edit.Percentage = edit.CalcualtePercentage(edit.Probablity)
				//  openStreet.Country = r.FormValue("country")
				//  openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				// openStreet.RouteNum  = i
				// fmt.Println("state:", openStreet)
				//  street , err := openStreet.CurrentLocationByPostalAddress(openStreet);if err != nil {
				// 	fmt.Println("Error:", err)
				// 	return
				//  }
				//  fmt.Println("Street:", street)
				//  uv ,err := visualizeReport.OpenWeather(openwizweather); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  loc := visualizeReport.GetCoordinates(street)
				//
				//  if err := uv.Current(loc); err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  info , err := uv.UVInformation(); if err != nil {
				// 	 fmt.Println("Error:", err)
				// 	return
				//  }
				//  visualizeReport.UVinfo = info

				w.WriteHeader(http.StatusOK)
	    		r.Method = "GET"
				Visualize(w,r)
				// fmt.Println("Virus:", capsid)

			default:
				temFile := template.Must(template.ParseFiles("dashboard.html"))
				temFile.Execute(w, "Dashboard")
			}
		} else {
			print("size must be less than 512MB")
			Repon := structs.Response{true,"Error in Upload File {size must not exceed with 512MB}", "WizDawrf/dashboard"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
		}

	}

}


func Send(w http.ResponseWriter, r *http.Request){

	temp := template.Must(template.ParseFiles("server.html"))

	block := structs.Block{}

	if(r.Method == "POST"){

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		r.ParseForm()
		block.TxSen = r.FormValue("sendAdd")
		block.TxRec = r.FormValue("add")
		choice :=  r.FormValue("transact")
		amount := r.FormValue("amount")
		block.Balance = ReadBalanceFromBlock(&block); if block.Balance == nil{
			Repon := structs.Response{true,"Some Issue ; [Balance]", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			fmt.Println("Error:")
			return
		}
		// Block number
		header, err := clientInstance.HeaderByNumber(context.Background(), nil); if err != nil{
			fmt.Println("Error:", err)
			Repon := structs.Response{true,"Error {HeaderByNumber}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			return
		}

		fmt.Println("Block Num :\n" , header.Number.String())
		fmt.Println("Wallet kEY:", WalletSecureKey)

		// private key to public address
		secure , err := crypto.HexToECDSA(WalletSecureKey); if err != nil {
			fmt.Println("Error:", err)
			Repon := structs.Response{true,"Error {Your Account dont have any Priavte key, use valid address}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			return
		}
		walletPublicKey :=  secure.Public()


		// Convert Public key
		ecdsaPubKey , ok := walletPublicKey.(*ecdsa.PublicKey); if !ok{
			fmt.Println("Error:", err)
			Repon := structs.Response{true,"Error {No Public Key}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			return
		}
		ethAdd := crypto.PubkeyToAddress(*ecdsaPubKey)
		fmt.Println("Your Adddress:", ethAdd)

		fmt.Println("EthAddress:", block.TxRec)


		// nonce pending
		 noncePending , err := clientInstance.PendingNonceAt(context.Background(), ethAdd); if err != nil {
		 	Repon := structs.Response{true,"Error {Get Nonce}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
		 	fmt.Println("Error:", err)
		 	return
		 }
		 fmt.Println("Pending Nonce:", noncePending)


		 // block number
		 /*blockNumber := big.NewInt(6677972)*/

		 block.Nonce = noncePending


		/*block.Nonce = r.FormValue("nonce")*/
		fmt.Println("choice:", choice)

		// gas
		gasLImit := uint64(21000)
		block.GasLimit = gasLImit

		gasPrice , err := clientInstance.SuggestGasPrice(context.Background()); if err != nil {
				fmt.Println("Error:", err)
				Repon := structs.Response{true,"Error :{Gas Price Error}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
				return
			}
			fmt.Println("Gas:", gasPrice)

			//Conversion
		charge , err := StringToInt(amount); if err != nil {
			fmt.Println("Error:", err)
			Repon := structs.Response{true,"Error :{Conversion}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			return
		}

		gwei := new(big.Int).SetInt64(int64(charge))
		block.Amount = gwei

		fee := new(big.Int)
		result := new(big.Int)

		switch choice{
			case "Normal":
				block.GasPrice = gasPrice
				fee.SetInt64(2)
				result.Mul(block.GasPrice , fee)
				fmt.Println("Block:" , block)
			case "Fair":
				block.GasPrice = gasPrice
				fee.SetInt64(3)
				result.Mul(block.GasPrice , fee)
				fmt.Println("Block:" , block)
			case "Blink":
				block.GasPrice = gasPrice
				fee.SetInt64(5)
				result.Mul(block.GasPrice , fee)
				fmt.Println("Block:" , block)
			default:
				fmt.Println("No choice")
		}

	// Send Transaction

		transfer := common.HexToAddress(block.TxSen)

		// Network ID
		chainId , err := clientInstance.NetworkID(context.Background())

		var nofield []byte

		tx := types.NewTransaction(block.Nonce, transfer,block.Amount, block.GasLimit, block.GasPrice, nofield)

		// Signed Transaction
		sign , err := types.SignTx(tx, types.NewEIP155Signer(chainId), secure); if err != nil {
			fmt.Println("Error:", err)
			Repon := structs.Response{true,"Error in Upload File", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			return
		}

		// Send Transaction
		err = clientInstance.SendTransaction(context.Background(), sign); if err != nil {
			fmt.Println("Error:", err)
			session := SessionExpire(w,r); if err != nil {
				fmt.Println("EXpire :", session)
			}
			Repon := structs.Response{true,"Error {Transaction Failed , Insufficent Balance}", "/transact"}
			println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			temp.Execute(w, Repon)
			return
		}

		fmt.Println("Send:", sign.Hash().Hex())
		w.WriteHeader(http.StatusOK)
	    r.Method = "GET"
		Visualize(w,r)

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

		if r.FormValue("allow")  == "on"{
			acc.Allowed = true
		}else{
			acc.Allowed = false
		}

		client , err := ethclient.Dial(EtherMainClientUrl); if err != nil {
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

		key := hexutil.Encode(PrvateKyByte)[2:]

		fmt.Println("Private_Key :" , key)

		pblicKey := privateKey.Public()

		pbcKey , ok := pblicKey.(*ecdsa.PublicKey); if !ok{
			println("Instaniate error {public key}")
			return
		}

		publicBytes := crypto.FromECDSAPub(pbcKey)
		fmt.Println("Public_Hash :" , hexutil.Encode(publicBytes)[4:])

		PublicKey := crypto.PubkeyToAddress(*pbcKey).Hex()
		fmt.Println("PublicKey:" , PublicKey)

		acc.PubKey = PublicKey
		acc.PrvteKey = key


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
				Repon := structs.Response{true,"Sorry! This is Smart Contact Adddress , We will handle in future", "/dashboard"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			fmt.Println("smart address:" , valid)

		myWallet := cloudWallet.EthereumWalletAcc{}

		signWallet , err := json.Marshal(myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := structs.Response{true,"Sorry! JSON Marshal Stream ", "/createWallet"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}

		err = json.Unmarshal(signWallet, &myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := structs.Response{true,"Sorry! JSON Unmarshal Stream", "/createWallet"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}

		//dataabse -- FindAddress
		ok , ethAdd := FindAddress(&acc); if ok && ethAdd != nil {
				Repon := structs.Response{true,"Sorry! Data Already register ", "/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}
		fmt.Println("Eth_Add:" , ethAdd)
		// ethAdd.SetPrivateKey()

		myWallet.Email = acc.Email
		myWallet.Password = acc.Password
		myWallet.EthAddress = acc.EthAddress
		myWallet.Terms = acc.Terms
		myWallet.Allowed = acc.Allowed
		myWallet.PrvteKey = acc.PrvteKey


		merchant , err := ledger.CreatePublicAddress(&myWallet, appName); if err != nil{
				fmt.Println("Error:", err)
				Repon := structs.Response{true,"Sorry! Invalid Ethereum Account ", "/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
		}

		fmt.Println("merchant:" , merchant)
		clientInstance = nil
			// Server response
		/*Repon := structs.Response{false,"Account Created!!! , Please don't share your key  & click on the link for futher...", "/dashboard"}
		println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
		temp.Execute(w, Repon)*/
		w.WriteHeader(http.StatusOK)
	    r.Method = "GET"
		Dashboard(w,r)
	}
}

func Transacts(w http.ResponseWriter, r *http.Request){


		temp := template.Must(template.ParseFiles("transact.html"))
		acc := structs.Static{}
		if r.Method == "GET" {
			fmt.Println("Url:", r.URL.Path)
			fmt.Println("Method:" + r.Method)

			acc.Eth = ETHAddressInstance
			acc.Balance = GetBalance(&acc); if acc.Balance == nil{
			fmt.Println("Error:")
		}
		fmt.Println("Details:", acc )
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


		client , err := ethclient.Dial(EtherMainClientUrl); if err != nil {
			fmt.Println("Error :" , err)
			return
		}

		fmt.Printf("Connection successfull ....%v\n", client)

			clientInstance = client
			println("Email:"+ acc.Email + "Password:"+ acc.Password)


			myWallet := cloudWallet.EthereumWalletAcc{}

			signWallet , err := json.Marshal(myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := structs.Response{true,"Sorry! JSON Marshal Stream ", "/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}

			err = json.Unmarshal(signWallet, &myWallet); if err != nil{
				fmt.Println("Error:", err)
				Repon := structs.Response{true,"Sorry! JSON Unmarshal Stream", "/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			add, ok := MyEthAddress(&acc); if !ok {
				Repon := structs.Response{true,"Sorry! No Account Exist ", "/open"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			if add != nil{
				fmt.Println("Address:" , add)
				acc.EthAddress = add.EthAddress
				// Secure Key
				/*if add.GetPrivateKey() != ""{
					WalletSecureKey = add.GetPrivateKey()
				}*/
				WalletSecureKey = add.PrvteKey
				fmt.Println("Secure:", WalletSecureKey)

				// variable address for futher processing
				ETHAddressInstance = acc.EthAddress
				fmt.Println("myWallet:", ETHAddressInstance)

				// read file and add swarm
				if add.Allowed {

					AddFilesEthereumSwarm()
				}
			 	// add this address in html page as static.

				//dataabse -- FindAddress
			secureWallet, ok := FindEthWallet(&acc); if !ok && secureWallet != nil {
				fmt.Println("Error", err)
				Repon := structs.Response{true,"Sorry! No Account Exist ", "/createWallet"}
				println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
				temp.Execute(w, Repon)
				return
			}
			fmt.Println("MyEthAddress Details:", secureWallet)



			// fmt.Println("Session:", cryptoSessions)

			// cryptoSessions = blockSession(LifeCode[0].Id)
			// err := Authentication(w, r); if err != nil {
			// 		Repon := structs.Response{true,"Sorry! No Account Exist ", "/createWallet"}
			// 		println("Server Response:", Repon.Flag,Repon.Message,Repon.Links)
			// 		temp.Execute(w, Repon)
			// 		return
			// }

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
	user := structs.Create_User{}

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Regsiter")
	} else {
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		user.Name = r.FormValue("uname")
		user.Fname = r.FormValue("ufname")
		user.Address = r.FormValue("address")
		user.Address2 = r.FormValue("add")
		user.City = r.FormValue("inputCity")
		user.Country = r.FormValue("co")
		user.Zip = r.FormValue("inputZip")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		if r.FormValue("gender") == "on" {
			user.Madam = true
		} else{
			user.Madam = false
		}

		// println("Gender:", user.sir)
		// println("Gender2:", user.madam)


		matchE, err := regexp.MatchString(emailexp, user.Email)
		if err != nil {
			println("invalid regular expression", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Data must be valid", "/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("regexp_email:", matchE)
		matchP, err := regexp.MatchString(passexp, user.Password)
		if err != nil {
			println("invalid regular expression", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Data must be valid", "/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("regexp_pass:", matchP)

		// security
		hashRet, encrypted := MessageToHash(w, matchE, matchP, user)
		if hashRet == false {
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Data must be valid", "/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("encryted data", encrypted.Reader)
		println("FamilyName:", user.Fname)
		println("Address", user.Address)
		println("Address2", user.Address2)
		println("City", user.City)
		println("Zip", user.Zip)
		println("Female", user.Madam)
		println("Country", user.Country)
		// println("check:", user.check_me_out)
		println("User record:", user.Name, user.Email)
		// println("phase:", KeyTx)
		addVistor(w, r, &user, encrypted.Reader)
	}

}

func Existing(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("login.html"))
	user := structs.Create_User{}

	if r.Method == "GET"{
		fmt.Printf("\nMethod:%v", r.Method)
		temp.Execute(w, "Login")
	}else{
		// Parse Form
		r.ParseForm()
		fmt.Println("Method:\n", r.Method)
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		if r.FormValue("check") == "on"{
			user.Secure = true
		}else{
			user.Secure = false
		}
		println("Login form data[", user.Email, user.Password, user.Secure,"]")

		// Valid Data for processing
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(user.Email)
		if !ok {
			println("invalid regular expression", !ok)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Data must be valid", "/login"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
		}
		println("regexp_email:", ok)
		reg := regexp.MustCompile(passexp)
		okx := reg.MatchString(user.Password)
		if !okx {
			println("invalid regular expression", !okx)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Data must be valid", "/login"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return
			// r.Method = "GET"
			// Existing(w,r)
		}
		println("regexp_pass:", okx)

		// Search Data in DB
		 data, err := SearchDB(w, r, user.Email,user.Password); if err != nil{
		 	// log.Fatal("Error", err)
		 	// w.Write([]byte(`{error: No Result Found }`))
		 	temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "No Record Exist", "/login"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
		 	return
		 }
		 	if data != nil{
		 		fmt.Printf("Search Data:%v", data.Id)

		 		//complex.AddByName(data.Name)
		 		// User Session
		 		if userSessions == nil {
		 			userSessions = SessionsInit(data.Id)
		 			sessId , _ := userSessions.Get(r, "session-name")
		 			sessId.Values["authenticated"] = true
		 			err = sessId.Save(r,w); if err != nil{
		 			// log.Fatal("Error", err)
		 			temp := template.Must(template.ParseFiles("server.html"))
					Res := structs.Response{true, "Sorry We have no Record, Please Regsiter", "/signup"}
					println("Server Response:", Res.Flag,Res.Message,Res.Links)
					temp.Execute(w, Res)
		 				return
		 		}
		 		println("Id :", sessId, "user:", userSessions)
		 }


		 // Login page
		w.WriteHeader(http.StatusOK)
	    r.Method = "GET"
		Dashboard(w,r)
		}else{
				temp := template.Must(template.ParseFiles("server.html"))
				Res := structs.Response{true, "Sorry We donot have any record , please register", "/signup"}
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
				Res := structs.Response{true, "Sorry We have no Record, Please Regsiter", "/signup"}
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

func addVistor(response http.ResponseWriter, request *http.Request, user *structs.Create_User, im string) {

	// var err error
	//response.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		fmt.Println("Method:" + request.Method)
	} else {
		var member db.Vistors
		data, err  := json.Marshal(member); if err != nil{
			fmt.Printf("Error in Marshal%v\n", err)
				temp := template.Must(template.ParseFiles("server.html"))
				Res := structs.Response{true, "Sorry Data must be in Format", "/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(response, Res)
			return
		}
		err = json.Unmarshal(data, &member); if err != nil{
			fmt.Printf("Error%v\n", err)
				temp := template.Must(template.ParseFiles("server.html"))
				Res := structs.Response{true, "Sorry Data Format Issue", "/signup"}
				println("Server Response:", Res.Flag,Res.Message,Res.Links)
				temp.Execute(response, Res)
			return
		}
		candidate , err := SearchDB(response, request, user.Email,user.Password); if err == nil && candidate == nil{
		 	// log.Fatal("Error", err)
		 	// w.Write([]byte(`{error: No Result Found }`))
		member.Id = im
		member.Name = user.Name
		member.Email = user.Email
		member.Password = user.Password
		member.FName = user.Fname
		if user.Madam {
			member.Eve = user.Madam
		}else{
			member.Eve = user.Madam
		}
		member.Address = user.Address
		member.LAddress	= user.Address2
		member.City = user.City
		member.Zip = user.Zip
		member.Country = user.Country
		record ,err := cloud.SaveData(&member, appName); if err != nil {
			fmt.Printf("Error%v\n", err)
				temp := template.Must(template.ParseFiles("server.html"))
				Res := structs.Response{true, "Sorry Data is not save yet", "/signup"}
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
			Res := structs.Response{true, "Your Data Already in DB", "/signup"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(response, Res)
	}

}



// Functions

func SetFirestoreCredentials() *firebase.App {

	_ , err := os.Stat("config/"+configFilename); if os.IsExist(err){
		fmt.Println("File Doesn't exist...", err)
		return nil
	}
	//fmt.Println("Filename:", file.Name())

	googleCredentials = "config/"+configFilename

	// set credentials
	conf := &firebase.Config{ProjectID: projectId}
	if googleCredentials != " "{
		opt := option.WithCredentialsFile(googleCredentials)
		app, err := firebase.NewApp(context.Background(), conf, opt)
		if err != nil {
			println("Error in Connection with Firestore", err)
			return nil
		}
		println("Connected... Welcome to Firestore")
		return app
	}
	return nil
}

func StringToInt(s string)(int, error){

	i, err := strconv.Atoi(s); if err != nil {
		fmt.Println("Error:", err)
		return 0 , err
	}
	return i, nil

}

func GetBalance(account *structs.Static)(*big.Int){

	wallet :=  common.HexToAddress(account.Eth)
	balnce , err := clientInstance.BalanceAt(context.Background(), wallet, nil); if err != nil{
		fmt.Println("Error:", err)
		return nil
	}
	account.Balance = balnce
	return account.Balance
}

func ReadBalanceFromBlock(acc *structs.Block)(*big.Int){
	wallet :=  common.HexToAddress(acc.TxRec)
	balnce , err := clientInstance.BalanceAt(context.Background(), wallet, nil); if err != nil{
		fmt.Println("Error:", err)
		return nil
	}
	acc.Balance = balnce
	return acc.Balance

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
		Res := structs.Response{true, "Sorry Error in open File", "/dashboard"}
		println("Server Response:", Res.Flag,Res.Message,Res.Links)
		temp.Execute(w, Res)
		return nil
	}
	println("File Exist...", f)
	finfo, err := f.Stat()
	if err != nil {
		println("File Info not found", err)
		temp := template.Must(template.ParseFiles("server.html"))
		Res := structs.Response{true, "Sorry, Server have NO INFORMATION", "/dashboard"}
		println("Server Response:", Res.Flag,Res.Message,Res.Links)
		temp.Execute(w, Res)
		return nil
	}
	println("File Info", finfo.Name())
	return   finfo
}

func Key(w http.ResponseWriter, h1, h2 string) (string, string, *ecdsa.PrivateKey) {


		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			// panic(err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Sorry Error in encrytion", "/signup"}
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
			Res := structs.Response{true, "Sorry Error in encrytion", "/signup"}
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
		return nil, err
	}
	return []byte(body), nil
}

func MessageToHash(w http.ResponseWriter,matchE, matchP bool, user structs.Create_User) (bool, *structs.SignedKey) {
	code := structs.SignedKey{}
	if matchE && matchP {
		h := sha256.New()
		// h.Write([]byte(user.email))
		hashe := h.Sum([]byte(user.Email))
		fmt.Println("email:", hex.EncodeToString(hashe))

		h1 := sha256.New()
		// h1.Write([]byte(user.password))
		hashp := h1.Sum([]byte(user.Password))
		fmt.Println("pass:", hex.EncodeToString(hashp))
		code.Reader, code.Signed, code.Tx = Key(w,hex.EncodeToString(hashe), hex.EncodeToString(hashp))
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
		 Res := structs.Response{true, "Sorry Error in Upload File {TYPE : MIME}", "/dashboard"}
		 println("Server Response:", Res.Flag,Res.Message,Res.Links)
		temp.Execute(w, Res)
		return nil
	}
	defer file.Close()
	if handler.Size <= (500000  * 1024) {
		fmt.Println("File name:" + handler.Filename, "Size:", handler.Size)
		if _, err := os.Stat(handler.Filename); os.IsExist(err) {
			fmt.Println("File not exist ", err)
			 temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Sorry Error , No Such Directory", "/dashboard"}
			println("Server Response:", Res.Flag,Res.Message,Res.Links)
			temp.Execute(w, Res)
			return nil
		}

		// upload file by user...
		upldFile, err := ioutil.TempFile("user_data", handler.Filename+"-*.txt")
		/*fmt.Println("file:", upldFile.Name())*/
		FILENAME = upldFile.Name()

		if err != nil {
			fmt.Println("Error received while uploading!", err)
			temp := template.Must(template.ParseFiles("server.html"))
			Res := structs.Response{true, "Sorry Upload File must have MIME TYPE: ", "/dashboard"}
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
			Res := structs.Response{true, "Error in Reading File", "/dashboard"}
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
// []amino.AminoClass, []amino.AminoClass
func SequenceFile(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error){

	seq, err := ReadSequence(userFile.Name());if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}


	var gen []string
	for _, v := range seq {
		space := DoAscii(v)
		if space == ""{
			fmt.Printf("Gap%v:\t", space)
		}
		gen = append(gen, space)
	}
	fmt.Println("Gen:{", gen , "}")

	// Dna to rna

	// rna35 := RNASequence(gen)
	// fmt.Println("single:", rna35)
	//
	// st1 := rna35
	// st2 := strings.Join(st1, "")
	//
	//
	// helixOrgan := bioChemRecord(st2)
	// fmt.Println("Helix Organsim:", helixOrgan)
	// proteins := RNAToAminoAcids(rna35)

	pathogen , err := ReadSequence(serverFile.Name()); ;if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}

	var genV []string
	for _, v := range pathogen {
		space := DoAscii(v)
		if space == ""{
			fmt.Printf("Gap%v:\t", space)
		}
		genV = append(genV, space)
	}
	fmt.Println("Genes:{", genV , "}")

	// Dna to rna

	// rnaVirus := RNASequence(genV)
	// fmt.Println("single:", rnaVirus)
	//
	//  st := rnaVirus
	// st21 := strings.Join(st, "")
	//
	//
	// helixVirus := bioChemRecord(st21)
	// fmt.Println("helix Virus:", helixVirus)
	// caspidProteins := RNAToAminoAcids(rnaVirus)


	// return proteins, caspidProteins	, nil
	return gen, genV , nil

}

func DoAscii(seq byte) string {

	if seq >= 65 && seq < 91{
		return string(alphabet.Letter(seq))
	}
	return string(alphabet.Letter(seq))
}

func RNASequence(sq []string) []string{

	var  k []string


	for i , _ := range sq{

		if sq[i] == "T"{
			sq[i] = "U"
		}
	  k = append(k , sq[i])
	}

	return k

}



func bioChemRecord(st2 string) structs.MolecularBio{

	molecule := structs.MolecularBio{}
	// helx record
	hlix := *pdb.ParseHelix(st2)
	fmt.Println("Serial:" , hlix.Serial)
	fmt.Println("Id:" , hlix.HelixID)
	fmt.Println("ResName+:" , hlix.InitResName)
	fmt.Println("ChainId+:" , hlix.InitChainID)
	fmt.Println("SeqNum+:" , hlix.InitSeqNum)
	fmt.Println("Icode+:" , hlix.InitICode)
	fmt.Println("ResName-:" , hlix.EndResName)
	fmt.Println("ChainId-:" , hlix.EndChainID)
	fmt.Println("SeqNum-:" , hlix.EndSeqNum)
	fmt.Println("Icode-:" , hlix.EndICode)
	fmt.Println("HelixClass:" , hlix.HelixClass)
	fmt.Println("Length:" , hlix.Length)
	// parseTree()

	//strand records
	stand := *pdb.ParseStrand(st2)
	fmt.Println("Strand:", stand.Strand)
	fmt.Println("Num:", stand.NumStrands)
	fmt.Println("Atom+:", stand.CurAtom)


	molecule.HelixA = hlix
	molecule.StrandB = stand

	return molecule
}

func RNAToAminoAcids(s []string) []amino.AminoClass{

	bases := []string{}
	for i , _ := range s{
		bases = append(bases, s[i])
	}

	proteins := amino.AminoClass{}

		ls := proteins.Bases(bases)

		return ls
}

func blockSession(id int) *sessions.CookieStore{

	return sessions.NewCookieStore([]byte(strconv.Itoa(id)))
}


func Authentication(w http.ResponseWriter, r * http.Request) error{

	sessId , _ := cryptoSessions.Get(r, "session-name")
	sessId.Values["authenticated"] = true
	err := sessId.Save(r,w); if err != nil{
		return err
	}
	return nil
}

func SessionExpire(w http.ResponseWriter , r *http.Request)error{
		sessId , _ := cryptoSessions.Get(r, "session-name")
		 	sessId.Values["authenticated"] = false
		 	err := sessId.Save(r,w); if err != nil{
		 		// log.Fatal("Error", err)
		 		return err
		 	}
		 	return nil
}

func AddFilesEthereumSwarm(){

	_ , err := os.Stat(FILENAME) ; if os.IsExist(err){
		fmt.Println("Error:", err)
		return
	}



}
