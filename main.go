package main

import (
	contxt "context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
	pay "github.com/logpacker/PayPal-Go-SDK"
	templates "text/template"
	firebase "firebase.google.com/go"
	cloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
	"github.com/ali2210/wizdwarf/structs"
	"github.com/biogo/biogo/alphabet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"github.com/ali2210/wizdwarf/db"
	"github.com/ali2210/wizdwarf/structs/paypal/handler"
	"github.com/ali2210/wizdwarf/structs/users"
	// "strings"
	"github.com/ali2210/wizdwarf/structs/amino"
	
	weather "github.com/ali2210/wizdwarf/structs/OpenWeather"
	"github.com/fogleman/ribbon/pdb"
)

// Variables

var (
	emailexp           string                = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp            string                = "([A-Z][a-z]*[0-9])*"
	addressexp         string                = "(^0x[0-9a-fA-F]{40}$)"
	appName            *firebase.App         = SetFirestoreCredentials() // Google_Cloud [Firestore_Reference]
	cloud              users.DBFirestore        = users.NewCloudInstance()
	ledger             db.PublicLedger       = db.NewCollectionInstance()
	paypalMini         handler.PaypalClientLevel  =handler.PaypalClientGo()
	userSessions       *sessions.CookieStore = nil //user level
	cryptoSessions     *sessions.CookieStore = nil // crypto level
	clientInstance     *ethclient.Client     = nil
	ETHAddressInstance string                = ""
	WalletPubKey       string                = ""
	WalletSecureKey    string                = ""
	// LifeCode []amino.AminoClass
	configFilename    string              = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	googleCredentials string              = ""
	FILENAME          string              = ""
	edit              structs.Levenshtein = structs.Levenshtein{}
	visualizeReport weather.DataVisualization = weather.DataVisualization{}
	accountID string 				= " "
	
	
)

// Constants

const (
	projectId string = "htickets-cb4d0"
	//Google_Credentials string = "/home/ali/Desktop/htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	// Main application
	EtherMainClientUrl string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	// Rickeby for test purpose
	RinkebyClientUrl string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	geocodeAPI string = "7efdb33c59a74e09352479b21657aee8"
)

// Functions

type Location struct{
	x string
	y string
}

func main() {

	// Server
	log.Println("[OK] Wiz-Dwarfs starting")
	port := os.Getenv("PORT")

	// env port setting

	if port == " " {
		log.Println("[Fail] No Application port allocated")
		return
	} else {
		if port != "5000" {
			// any Listening PORT {heroku}
			log.Println("[Open] Application Port", port)
		} else {
			// specfic port allocated {docker}
			port = "5000"
			log.Println("[New] Application Default port")
		}

	}

	log.Println("[OK] Application :", port+" Port")
	// Routing
	routing := mux.NewRouter()

	// Links
	routing.HandleFunc("/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		
		temp := template.Must(template.ParseFiles("initial.html"))
		
		if arg2.Method == "GET" {
			log.Println("[OK] URL :", arg2.URL.Path)
			temp.Execute(arg1, "MainPage")
		}
		flag := ProcessWaiit()
		if !flag{
			arg1.WriteHeader(http.StatusOK)
			arg2.Method = "GET"
			Home(arg1, arg2)
		}
	})
	routing.HandleFunc("/home", Home)
	routing.HandleFunc("/signup", NewUser)
	routing.HandleFunc("/login", Existing)
	routing.HandleFunc("/dashboard", Dashboard)
	routing.HandleFunc("/dashbaord/setting", Setting)
	routing.HandleFunc("/dashbaord/setting/profile", Profile)
	routing.HandleFunc("/logout", Logout)
	routing.HandleFunc("/createWallet", CreateWallet)
	routing.HandleFunc("/terms", Terms)
	routing.HandleFunc("/open", Wallet)
	// routing.HandleFunc("/open/setting", WalletSettingMenu)
	routing.HandleFunc("/open/setting/credit", Credit)
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

	// routing.HandleFunc("/dummy", server)*templates.Template

	// tcp connection
	err := http.ListenAndServe(":"+port, routing)
	if err != nil {
		log.Println("Listening Error: ", err)
		panic(err)
	}

}

// Routes Handle

func Home(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html"))

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Home")
	}

}

// func WalletSettingMenu(w http.ResponseWriter, r *http.Request)  {
// 	temp := template.Must(template.ParseFiles("walletsetting.html"))

// 	if r.Method == "GET" {
// 		fmt.Println("Method:" + r.Method)
// 		temp.Execute(w, "WalletSetting")
// 	}
// }

func Setting(w http.ResponseWriter, r *http.Request)  {
	
	temp := template.Must(template.ParseFiles("settings.html"))
	
	ret , err := paypalMini.RetrieveCreditCardInfo(accountID); if err != nil {
		log.Fatalln("[Fail] Operation:", err)
		return 
	}
	
	if r.Method == "GET" {
		log.Println("[Accept]" , r.URL.Path)
		temp.Execute(w,ret)
	}

}

func Profile(w http.ResponseWriter, r *http.Request)  {

	temp := template.Must(template.ParseFiles("profile.html"))
	if accountID == ""{	
		log.Fatal("[Error ] Please login  ")
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry Session expire   ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp); if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
	}

	detailsAcc , err := cloud.FindDataByID(accountID, appName); if err != nil {
		log.Fatalln("[Fail] Operation..", err)
		return 
	}

	if r.Method == "GET" {
		log.Println("[Accept]" , r.URL.Path)
		temp.Execute(w,detailsAcc)
	}else{
		log.Println("[Accept] Method:", r.Method)
		log.Println("[Accept] Path:", r.URL.Path)

		r.ParseForm()

		// save value in db
		MyProfile := users.UpdateProfile{
			Email : r.FormValue("email"),
			Phone : r.FormValue("phone"),
			FirstName : r.FormValue("uname"),
			LastName : r.FormValue("ufname"),
			HouseAddress : r.FormValue("inputAddress"),
			SubAddress : r.FormValue("inputAddress2"),
			Country : r.FormValue("country"),
			Zip : r.FormValue("inputZip"),
		}

		if accountID == ""{
			log.Fatal("[Error ] Please login  ")
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry Session expire   ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp); if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
		}
		MyProfile.Id = accountID
		
		male := r.FormValue("gender")
		if male == "on"{
			MyProfile.Male = true 
		}
		MyProfile.Male = false

		profile, err  := cloud.UpdateProfiles(appName, &MyProfile); if err != nil {
			log.Fatalln("[Fail] Operation..", err)
			return 
		} 
		 
		log.Println("[Accept] Profile updated... ", profile)

	}
}



func Credit(w http.ResponseWriter, r *http.Request)  {
	temp := template.Must(template.ParseFiles("credit.html"))
	if r.Method == "GET" {
		log.Println("[Accept] Path :" , r.URL.Path)
		temp.Execute(w,"Credit")
	}else{
		log.Println("[Accept ]Path :" , r.URL.Path)
		log.Println("Method :", r.Method)
		r.ParseForm()
		card := pay.CreditCard{
			FirstName : r.FormValue("fholder"),
			LastName : r.FormValue("surename"),
			Number : r.FormValue("cardNo"),
			ExpireMonth : r.FormValue("expire"),
		 	CVV2 : r.FormValue("cvv"),
		}

		 card.ID =  AutoKeyGenerate(card.CVV2) 
		 log.Println("Id generated:" , card)

		// store credit card information.
		mini := handler.PaypalMiniVersion{}
	  	client, err := paypalMini.NewClient(); if err != nil {
			log.Fatalln("[Fail] Operation:", err)
			  return
		}
		mini.Client = client
		
		token , err := paypalMini.Token();if err != nil {
			log.Fatalln("[Fail] Operation:", err)
			return 
		}
		store, err := paypalMini.StoreCreditCardInfo(card); if err != nil {
			log.Fatalln("[Fail] Operation:", err)
			return 
		}
		ret , err := paypalMini.RetrieveCreditCardInfo(store.ID); if err != nil {
			log.Fatalln("[Fail] Operation:", err)
			return 
		}

		log.Println("[Accept] Token issue:", token, "retInfo:", ret)



	}

}

func Visualize(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("visualize.html"))
	log.Println("Report percentage", visualizeReport.Percentage)
	log.Println("Report uv ", visualizeReport.UVinfo)
	edit.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" && edit.GetProbParameter() != -1.0 {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		temp.Execute(w, visualizeReport)
	}
	// err := SessionExpire(w,r); if err != nil {
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// r.Method = "GET"
	// Dashboard(w,r)
}

func Treasure(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("treasure.html"))
	acc := structs.Static{}
	block := structs.Block{}

	if r.Method == "GET" {

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		acc.Eth = ETHAddressInstance
		acc.Balance = GetBalance(&acc)
		if acc.Balance == nil {
			fmt.Println("Error:")
		}
		fmt.Println("Details:", acc)
		temp.Execute(w, acc)

	} else {

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		r.ParseForm()
		block.TxSen = r.FormValue("send")
		block.TxRec = r.FormValue("rece")

		fmt.Println("block:", block)

		// Block Number
		blockNumber := big.NewInt(6677972)
		bNum, err := clientInstance.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Length:", len(bNum.Transactions()))
		fmt.Println("Hash:", bNum.Hash().Hex())
		fmt.Println("Block:", bNum.Number().Uint64())

		for _, tx := range bNum.Transactions() {
			fmt.Println("To:", tx.To().Hex())
			fmt.Println("Block_Hash:", tx.Hash().Hex())

			// ChainId

			chainId, err := clientInstance.NetworkID(context.Background())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// get recipt address
			message, err := tx.AsMessage(types.NewEIP155Signer(chainId))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Message From:", message.From().Hex())

			recp, err := clientInstance.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Status:", recp.Status)
		}

		txs := common.HexToHash(bNum.Hash().Hex())
		fmt.Println("Tx:", txs.Hex())

		// Number of Transaction
		count, err := clientInstance.TransactionCount(context.Background(), txs)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Num #:", count)

		for i := uint(0); i < count; i++ {
			Tx, err := clientInstance.TransactionInBlock(context.Background(), txs, i)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Tx Hash:", Tx.Hash().Hex())

			txHash := common.HexToHash(Tx.Hash().Hex())

			// Transaction status
			tx, isPending, err := clientInstance.TransactionByHash(context.Background(), txHash)
			if err != nil {
				fmt.Println("Error:", err)
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

	if r.Method == "GET" {
		temp := template.Must(template.ParseFiles("dashboard.html"))
		fmt.Println("Method:" + r.Method)
		fmt.Println("Url:", r.URL.Path)
		temp.Execute(w, "Dashboard")
	} else {
		// temp := template.Must(template.ParseFiles("server.html"))
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		// FILE Upload ....
		file := UploadFiles(w, r)
		if file != nil {

			choose := r.FormValue("choose")

			switch choose {
			case "0":
				fmt.Fprintf(w, "Please choose any option ...")
				_, err := ChoosePattern(w,r,"", choose,file); if err != nil {
					return 
				}
			case "1":
				var name string = "Covid-19"
				e , err := ChoosePattern(w,r, name, choose,file); if err != nil {
					return 
				}
				visualizeReport.Percentage = e.Percentage
				// v :=  infectedUv()
				
				// openStreet.Country = r.FormValue("country")
				// openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				
				// v.UVinfo = uvslice

				w.WriteHeader(http.StatusOK)
				// LifeCode = genome
				r.Method = "GET"
				// Wallet(w,r)
				Visualize(w, r)
				// fmt.Println("Virus:", capsid)

			case "2":
				var name string = "FlaviDengue"
				e , err := ChoosePattern(w,r, name, choose,file); if err != nil {
					return
				}
				visualizeReport.Percentage = e.Percentage
				// v :=  infectedUv()				
				// v.UVinfo = uvslice	
				w.WriteHeader(http.StatusOK)
				
				r.Method = "GET"
				Visualize(w, r)
				// Wallet(w,r)
				// fmt.Println("Virus:", capsid)
			case "3":
				var name string = "KenyaEbola"
				e , err := ChoosePattern(w,r, name, choose,file); if err != nil {
					return 
				}
				visualizeReport.Percentage = e.Percentage
				// v :=  infectedUv()
				// v.UVinfo = uvslice
				
				w.WriteHeader(http.StatusOK)
				r.Method = "GET"
				Visualize(w, r)

				// fmt.Println("Virus:", capsid)
			case "4":
				var name string = "ZikaVirusBrazil"
				e , err := ChoosePattern(w,r, name, choose,file); if err != nil {
					return 
				}
				visualizeReport.Percentage = e.Percentage
				// v :=  infectedUv()
				// openStreet.Country = r.FormValue("country")
				// openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				
				//v.UVinfo = uvslice

				w.WriteHeader(http.StatusOK)
				r.Method = "GET"
				Visualize(w, r)

				// fmt.Println("Virus:", capsid)
			case "5":
				var name string = "MersSaudiaArabia"
				e , err := ChoosePattern(w,r, name, choose,file); if err != nil {
					return 
				}
				visualizeReport.Percentage = e.Percentage
				// v :=  infectedUv()				//  openStreet.Country = r.FormValue("country")
				//  openStreet.PostalCode = r.FormValue("postal")
				// openStreet.City = r.FormValue("city")
				// openStreet.State = r.FormValue("state")
				// openStreet.StreetAddress = r.FormValue("street")
				// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
				// 	return
				// }
				// v.UVinfo = uvslice	
				w.WriteHeader(http.StatusOK)
				
				r.Method = "GET"
				Visualize(w, r)
				// fmt.Println("Virus:", capsid)

			default:
				temFile := template.Must(template.ParseFiles("dashboard.html"))
				temFile.Execute(w, "Dashboard")
			}
		} else {
			log.Fatal("[Fail] Size Limit reached 512MB  ")
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry Size Limit reached   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
			
		}

	}

}

func Send(w http.ResponseWriter, r *http.Request) {

	// temp := template.Must(template.ParseFiles("server.html"))

	block := structs.Block{}

	if r.Method == "POST" {

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		r.ParseForm()
		block.TxSen = r.FormValue("sendAdd")
		block.TxRec = r.FormValue("add")
		choice := r.FormValue("transact")
		amount := r.FormValue("amount")
		block.Balance = ReadBalanceFromBlock(&block)
		if block.Balance == nil {
			log.Fatal("[Fail] Block Balance must be zero  ", block.Balance)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
			return
		}
		// Block number
		header, err := clientInstance.HeaderByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal("[Fail] Header Number   ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		// fmt.Println("Block Num :\n", header.Number.String())
		// fmt.Println("Wallet kEY:", WalletSecureKey)

		// private key to public address
		secure, err := crypto.HexToECDSA(WalletSecureKey)
		if err != nil {
			log.Fatal("[Fail] Secure Wallet Key  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		log.Println("Header :", header.Number.String())
		walletPublicKey := secure.Public()

		// Convert Public key
		ecdsaPubKey, ok := walletPublicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("[Fail] Your Wallet Public Key  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		ethAdd := crypto.PubkeyToAddress(*ecdsaPubKey)

		// nonce pending
		noncePending, err := clientInstance.PendingNonceAt(context.Background(), ethAdd)
		if err != nil {

			log.Fatal("[Fail] Current Pending Nonce Status  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! 	Connectivity Issue   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		// block number
		/*blockNumber := big.NewInt(6677972)*/

		block.Nonce = noncePending

		/*block.Nonce = r.FormValue("nonce")*/

		// gas
		gasLImit := uint64(21000)
		block.GasLimit = gasLImit

		gasPrice, err := clientInstance.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal("[Fail] Gas Price  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Connectivity issue   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		// fmt.Println("Gas:", gasPrice)

		//Conversion
		charge, err := StringToInt(amount)
		if err != nil {
			log.Fatal("[Fail] charge must be String  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		gwei := new(big.Int).SetInt64(int64(charge))
		block.Amount = gwei

		fee := new(big.Int)
		result := new(big.Int)

		switch choice {
		case "Normal":
			block.GasPrice = gasPrice
			fee.SetInt64(2)
			result.Mul(block.GasPrice, fee)
		case "Fair":
			block.GasPrice = gasPrice
			fee.SetInt64(3)
			result.Mul(block.GasPrice, fee)
		case "Blink":
			block.GasPrice = gasPrice
			fee.SetInt64(5)
			result.Mul(block.GasPrice, fee)

		default:
			log.Println(" Choice [1-5]")
		}

		log.Println("[Accept] Block Info:", block)
		// Send Transaction

		transfer := common.HexToAddress(block.TxSen)

		// Network ID
		chainId, err := clientInstance.NetworkID(context.Background())

		var nofield []byte

		tx := types.NewTransaction(block.Nonce, transfer, block.Amount, block.GasLimit, block.GasPrice, nofield)

		// Signed Transaction
		sign, err := types.SignTx(tx, types.NewEIP155Signer(chainId), secure)
		if err != nil {
			log.Fatal("[Fail] Signed Transaction", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Transaction is signed by your wallet   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		// Send Transaction
		err = clientInstance.SendTransaction(context.Background(), sign)

		if err != nil {
			log.Fatal("[Fail] Operation Fail  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Insuffient Balance   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}
			return
		}
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(false, "Hurrah  ! OPERATION SUCCESSED   ", "/dashbaord", w, r)
		response.ClientLogs()
		err = response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		Visualize(w, r)

	}
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("seed.html"))
	acc := structs.Acc{}

	if r.Method == "GET" {

		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Seed")
	} else {

		// temp := template.Must(template.ParseFiles("server.html"))
		fmt.Println("Method:" + r.Method)
		r.ParseForm()
		acc.Email = r.FormValue("email")
		acc.Password = r.FormValue("password")

		if r.FormValue("agreeTerms") == "on" {
			acc.Terms = true
		} else {
			acc.Terms = false
		}

		if r.FormValue("allow") == "on" {
			acc.Allowed = true
		} else {
			acc.Allowed = false
		}

		client, err := ethclient.Dial(EtherMainClientUrl)
		if err != nil {
			log.Fatal("[Fail] Request Failed  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Connectivity Issue   ", "/dashboard", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
			return
		}

		log.Println("[Accept] Connection accepted", client)
		clientInstance = client

		// private key
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal("[Fail] Public Key generate  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/dashboard", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
			return
		}

		// private key into bytes
		PrvateKyByte := crypto.FromECDSA(privateKey)

		key := hexutil.Encode(PrvateKyByte)[2:]

		pblicKey := privateKey.Public()

		pbcKey, ok := pblicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("[Fail] Public Key from Private Key  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/dashboard", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}
			return
		}

		publicBytes := crypto.FromECDSAPub(pbcKey)
		// fmt.Println("Public_Hash :", hexutil.Encode(publicBytes)[4:])

		PublicKey := crypto.PubkeyToAddress(*pbcKey).Hex()

		acc.PubKey = PublicKey
		acc.PrvteKey = key

		// hash to ethereum
		hshCode := sha3.NewLegacyKeccak256()
		hshCode.Write(publicBytes[1:])
		ethereum := hexutil.Encode(hshCode.Sum(nil)[12:])

		acc.EthAddress = ethereum

		// valid address
		valid := isYourPublcAdresValid(ethereum)
		if valid {

			// smart contract address
			log.Println("[Feature] Smart Address", valid)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Smart Contract Address added in future release ", "/dashboard", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		myWallet := cloudWallet.EthereumWalletAcc{}

		signWallet, err := json.Marshal(myWallet)
		if err != nil {
			log.Fatal("[Fail] Data JSON FORMAT ERROR ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! POOR DATA FORMAT   ", "/createWallet", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		err = json.Unmarshal(signWallet, &myWallet)
		if err != nil {
			log.Fatal("[Fail] Data JSON FORMAT ERROR ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! POOR DATA FORMAT   ", "/createWallet", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		//dataabse -- FindAddress
		ok, ethAdd := FindAddress(&acc)
		if ok && ethAdd != nil {
			log.Fatal("[Replicate] Already Data exist  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry !  Replicate not allowed  ", "/createWallet", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}

		}

		myWallet.Email = acc.Email
		myWallet.Password = acc.Password
		myWallet.EthAddress = acc.EthAddress
		myWallet.Terms = acc.Terms
		myWallet.Allowed = acc.Allowed
		myWallet.PrvteKey = acc.PrvteKey

		merchant, err := ledger.CreatePublicAddress(&myWallet, appName)
		if err != nil {
			log.Fatal("[Fail] Wallet Don't have Public Accessibity  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Conenctivity issue   ", "/createWallet", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		clientInstance = nil

		log.Println("[Accept] Welcome ! Your Account Has been created", merchant)
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		Existing(w, r)
	}
}

func Transacts(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("transact.html"))
	acc := structs.Static{}
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		acc.Eth = ETHAddressInstance
		acc.Balance = GetBalance(&acc)
		if acc.Balance == nil {
			log.Fatal("[Fail] Connection Reject ", acc.Balance)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Connectivity Issue   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}
		}
		fmt.Println("Details:", acc)
		temp.Execute(w, acc)
	}

}

func Wallet(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("wallet.html"))
	acc := structs.Acc{}

	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Wallet")

	} else {

		fmt.Println("Method:" + r.Method)
		r.ParseForm()

		acc.Email = r.FormValue("email")
		acc.Password = r.FormValue("password")

		client, err := ethclient.Dial(EtherMainClientUrl)
		if err != nil {
			log.Fatal("[Fail] Connection Reject ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Connectivity Issue   ", "/open", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		log.Println("[Accept] Firestore Cloud Database Connection secure", client)

		clientInstance = client

		log.Println("[Accept] Your Account Details:", acc, "Client api Reference: ", clientInstance)

		myWallet := cloudWallet.EthereumWalletAcc{}

		signWallet, err := json.Marshal(myWallet)
		if err != nil {
			log.Fatal("[Fail] Data JSON FORMAT ERROR ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! POOR DATA FORMAT   ", "/open", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

		err = json.Unmarshal(signWallet, &myWallet)
		if err != nil {
			log.Fatal("[Fail] JSON DATA RETURN ERROR", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Poor Data Format   ", "/open", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		addr, ok := MyEthAddress(&acc)
		if !ok {

			log.Fatal("[Fail] No Ethereum Account ", !ok)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! NO Ethereum Account   ", "/open", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		if addr != nil {

			acc.EthAddress = addr.EthAddress
			// Secure Key

			WalletSecureKey = addr.PrvteKey

			// variable address for futher processing
			ETHAddressInstance = acc.EthAddress
			log.Println("Your Wallet:", ETHAddressInstance)

			// read file and add swarm
			// if add.Allowed {
			//
			// 	AddFilesEthereumSwarm()
			// }
			// add this address in html page as static.

			//dataabse -- FindAddress
			secureWallet, ok := FindEthWallet(&acc)
			if !ok && secureWallet != nil {
				log.Fatal("[Fail] No crypto wallet found against your account ", !ok)
				response := structs.Response{}
				temp := server(w, r)
				_ = response.ClientRequestHandle(true, "Sorry ! NO CRYPTOWALLET   ", "/createWallet", w, r)
				response.ClientLogs()
				err := response.Run(temp)
				if err != nil {
					log.Println("[Error]: checks logs...", err)
					return
				}

			}
			log.Println("[Accept] Your Ethereum Wallet Info:", secureWallet)

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			// Transacts(w, r)
		}
	}
}

func Terms(w http.ResponseWriter, r *http.Request) {

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
		} else {
			user.Madam = false
		}

		matchE, err := regexp.MatchString(emailexp, user.Email)
		if err != nil {
			log.Fatal("[Fail] Auto email pattern  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Email is not in format  ", "/signup", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		// println("regexp_email:", matchE)
		matchP, err := regexp.MatchString(passexp, user.Password)
		if err != nil {
			log.Fatal("[Fail] Password is very week ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Password is too short   ", "/signup", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		// println("regexp_pass:", matchP)

		// security
		hashRet, encrypted := MessageToHash(w, r, matchE, matchP, user)
		if !hashRet {
			log.Fatal("[Fail] Week encryption", hashRet)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Data discard because of your data in not our standard.   ", "/signup", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		log.Println("[Accept] Review your logs.. ", user)
		// println("phase:", KeyTx)
		addVistor(w, r, &user, encrypted.Reader)
	}

}

func Existing(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("login.html"))
	user := structs.Create_User{}

	if r.Method == "GET" {
		fmt.Printf("\nMethod:%v", r.Method)
		temp.Execute(w, "Login")
	} else {
		// Parse Form
		r.ParseForm()
		fmt.Println("Method:\n", r.Method)
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		if r.FormValue("check") == "on" {
			user.Secure = true
		} else {
			user.Secure = false
		}
		log.Println("Login form data[", user, "]")

		// Valid Data for processing
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(user.Email)
		if !ok {
			log.Fatal("[Fail] Mismatch ", ok)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Mismatch   ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
			}

		}

		reg := regexp.MustCompile(passexp)
		okx := reg.MatchString(user.Password)
		if !okx {
			log.Fatal("[Fail] Mismatch password", !okx)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Mismatch password ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

			// r.Method = "GET"
			// Existing(w,r)
		}
		println("regexp_pass:", okx)

		// Search Data in DB
		data, err := SearchDB(w, r, user.Email, user.Password)
		if err != nil {
			log.Fatal("[Result]: No Match Found  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! No information Retry ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		if data != nil {
			accountID = data.Id
			fmt.Printf("Search Data:%v", data.Id)
			act := structs.RouteParameter{}
			//complex.AddByName(data.Name)
			// User Session
			if userSessions == nil {
				userSessions = SessionsInit(data.Id)

				err := act.NewToken(userSessions, w, r)
				if err != nil {
					log.Fatal("[FAIL] No Token generate .. Review logs", err)
					response := structs.Response{}
					temp := server(w, r)
					_ = response.ClientRequestHandle(true, "Sorry ! Access Denied, ", "/login", w, r)
					response.ClientLogs()
					err := response.Run(temp)
					if err != nil {
						log.Println("[Error]: checks logs...", err)
						return
					}

				}
				println("Id :", data.Id, "user:", userSessions)
			}

			// Login page
			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			Dashboard(w, r)
		} else {
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Access Denied, Plesse Register ", "/signup", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}

	}
}

func server(w http.ResponseWriter, r *http.Request) *templates.Template {
	// temp := template.Must(template.ParseFiles("server.html"))
	// temp.Execute(w, "server")
	myResponse := structs.Response{}
	temp := myResponse.ClientHTMLRequest("server")
	return temp
	// err := myResponse.Run() ; if err != nil{
	// 	log.Println("[Error]: ", err)
	//}

}

func Logout(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		act := structs.RouteParameter{}

		log.Println("[Access] ", r.URL.Path)

		err := act.ExpireToken(userSessions, w, r)
		if err != nil {
			log.Fatal("[Fail] No Token Expire  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Token Expire   ", "/transact", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		Existing(w, r)
	}
}

//  Advance Functions

func SearchDB(w http.ResponseWriter, r *http.Request, email, pass string) (*users.Vistors, error) {

	var data *users.Vistors
	var err error
	// w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
	} else {
		fmt.Println("Method:" + r.Method)
		data, err = cloud.FindAllData(appName, email, pass)
		if err != nil && data != nil {
			log.Fatal("[Fail] No info  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! No Information ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return nil, err
			}

		}
		log.Println("[Accept] Your Request results :", data)
	}
	return data, nil
}

// func getVistorData(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("Content-Type", "application/json")
// 	visitor, err := cloud.FindAllData(appName)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{"error" :"Error getting visitor result"}`))
// 		return
// 	}
// 	fmt.Printf("Vistors array%v", visitor)
//
// 	// response.WriteHeader(http.StatusOK)
// 	json.NewEncoder(response).Encode(visitor)
//
// }

func addVistor(response http.ResponseWriter, request *http.Request, user *structs.Create_User, im string) {

	// var err error
	//response.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		fmt.Println("Method:" + request.Method)
	} else {
		var member users.Vistors
		data, err := json.Marshal(member)
		if err != nil {
			log.Fatal("[Fail] Poor DATA JSON FORMAT  ", err)
			r := structs.Response{}
			temp := server(response, request)
			_ = r.ClientRequestHandle(true, "Sorry ! Data in poor format", "/signup", response, request)
			r.ClientLogs()
			err := r.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		err = json.Unmarshal(data, &member)
		if err != nil {
			log.Fatal("[Fail] Poor Format  ", err)
			r := structs.Response{}
			temp := server(response, request)
			_ = r.ClientRequestHandle(true, "Sorry ! Data in poor format ", "/signup", response, request)
			r.ClientLogs()
			err := r.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		candidate, err := SearchDB(response, request, user.Email, user.Password)
		if err == nil && candidate == nil {

			member.Id = im
			member.Name = user.Name
			member.Email = user.Email
			member.Password = user.Password
			member.FName = user.Fname
			if user.Madam {
				member.Eve = user.Madam
			} else {
				member.Eve = user.Madam
			}
			member.Address = user.Address
			member.LAddress = user.Address2
			member.City = user.City
			member.Zip = user.Zip
			member.Country = user.Country
			record, err := cloud.SaveData(&member, appName)
			if err != nil {
				log.Fatal("[Fail] Data has not saved", err)
				r := structs.Response{}
				temp := server(response, request)
				_ = r.ClientRequestHandle(true, "Sorry ! Internal issue   ", "/signup", response, request)
				r.ClientLogs()
				err := r.Run(temp)
				if err != nil {
					log.Println("[Error]: checks logs...", err)
					return
				}

			}

			log.Println("Records:", record, "Info: ", candidate)
			response.WriteHeader(http.StatusOK)
			request.Method = "GET"
			CreateWallet(response,request)

		}

		log.Fatal("[Fail] Already HAVE  ")
		r := structs.Response{}
		temp := server(response, request)
		_ = r.ClientRequestHandle(true, "Sorry ! This information already in system   ", "/signup", response, request)
		r.ClientLogs()
		err = r.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return
		}

	}

}

// Functions

func SetFirestoreCredentials() *firebase.App {

	_, err := os.Stat("config/" + configFilename)
	if os.IsExist(err) {
		fmt.Println("File Doesn't exist...", err)
		return nil
	}
	//fmt.Println("Filename:", file.Name())

	googleCredentials = "config/" + configFilename

	// set credentials
	conf := &firebase.Config{ProjectID: projectId}
	if googleCredentials != " " {
		opt := option.WithCredentialsFile(googleCredentials)
		app, err := firebase.NewApp(context.Background(), conf, opt)
		if err != nil {
			log.Fatal("[Fail] Connection Reject", err)
			return nil
		}
		log.Println("[Accept] Concection build with cloud firestore database")
		return app
	}
	return nil
}

func StringToInt(s string) (int, error) {

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("[Fail] Conversion", err)
		return 0, err
	}
	return i, nil

}

func GetBalance(account *structs.Static) *big.Int {

	wallet := common.HexToAddress(account.Eth)
	balnce, err := clientInstance.BalanceAt(context.Background(), wallet, nil)
	if err != nil {
		log.Fatalln("[Fail] Balance reading issue/ connectivity issue")
		return nil
	}
	account.Balance = balnce
	return account.Balance
}

func ReadBalanceFromBlock(acc *structs.Block) *big.Int {
	wallet := common.HexToAddress(acc.TxRec)
	balnce, err := clientInstance.BalanceAt(context.Background(), wallet, nil)
	if err != nil {
		log.Fatalln("[Fail] Connectivity issue", err)
		return nil
	}
	acc.Balance = balnce
	return acc.Balance

}

func FindAddress(w *structs.Acc) (bool, *cloudWallet.EthereumWalletAcc) {

	ethAcc, err := ledger.FindMyPublicAddress(w, appName)
	if err != nil {
		log.Fatalln("[Fail] Ledger ahve no Information / internal issue ", err)
		return false, nil
	}
	if ethAcc != nil {
		return false, nil
	}
	return true, ethAcc

}

func MyEthAddress(w *structs.Acc) (*cloudWallet.EthereumWalletAcc, bool) {

	acc, err := ledger.FindMyAddressByEmail(w, appName)
	if err != nil {
		log.Fatalln("[Fail] Configuration issue", err)
		return nil, false
	}
	if acc == nil {
		return nil, false
	}
	return acc, true
}

func FindEthWallet(w *structs.Acc) (*cloudWallet.EthereumWalletAcc, bool) {

	acc, err := ledger.FindMyPublicAddress(w, appName)
	if err != nil {
		log.Fatalln("[Fail] Configuration issue", err)
		return nil, false
	}
	return acc, true
}

func isYourPublcAdresValid(hash string) bool {

	expression := regexp.MustCompile(addressexp)
	v := expression.MatchString(hash)

	address := common.HexToAddress(hash)
	bytecode, err := clientInstance.CodeAt(contxt.Background(), address, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	contract := len(bytecode) > 0
	log.Println("[Accept] Contract Address: ", contract, "Result:", v)
	return contract
}

func SessionsInit(unique string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(unique))
}

func FileReadFromDisk(w http.ResponseWriter, r *http.Request, filename string) os.FileInfo {
	f, err := os.OpenFile(filename+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("[Fail] No File Exist  ", err)
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! No Information    ", "/dashboard", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return nil
		}
		return nil
	}

	finfo, err := f.Stat()
	if err != nil {
		log.Fatal("[Fail] Application Access   ", err)
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! Permission Denied   ", "/dashboard", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return nil
		}

	}
	return finfo
}

func Key(w http.ResponseWriter, r *http.Request, h1, h2 string) (string, string, *ecdsa.PrivateKey) {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal("[Fail] Key generate   ", err)
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! Computation Error   ", "/signup", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return "", "", nil
		}

	}

	// 0x40fa6d8c32594a971b692c44c0c56b19c32613deb1c6200c26ea4fe33d34a5fd

	msg := h1 + h2
	hash := sha256.Sum256([]byte(msg))

	reader, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {

		log.Fatal("[Fail] Key signed   ", err)
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! Computation Error   ", "/signup", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return "", "", privateKey
		}

	}

	return fmt.Sprintf("0x%x", reader), fmt.Sprintf("0x%x", s), privateKey

}

func ReadSequence(filename string) ([]byte, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("[Fail] File  Access ", err)
		return nil, err
	}
	return []byte(body), nil
}

func MessageToHash(w http.ResponseWriter, r *http.Request, matchE, matchP bool, user structs.Create_User) (bool, *structs.SignedKey) {
	code := structs.SignedKey{}
	if matchE && matchP {
		h := sha256.New()
		hashe := h.Sum([]byte(user.Email))

		h1 := sha256.New()
		// h1.Write([]byte(user.password))
		hashp := h1.Sum([]byte(user.Password))

		fmt.Println("pass:", hex.EncodeToString(hashp))
		code.Reader, code.Signed, code.Tx = Key(w, r, hex.EncodeToString(hashe), hex.EncodeToString(hashp))
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

		log.Fatal("[Fail] Error in upload   ", err)
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! Computation Error {type of file must be MIME}  ", "/dashboard", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return nil
		}

	}
	defer file.Close()
	if handler.Size <= (500000 * 1024) {
		fmt.Println("File name:"+handler.Filename, "Size:", handler.Size)
		if _, err := os.Stat(handler.Filename); os.IsExist(err) {
			log.Fatal("[Fail] File Access   ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Permission Denied  ", "/dashboard", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return nil
			}
			return nil

		}

		// upload file by user...
		upldFile, err := ioutil.TempFile("user_data", handler.Filename+".txt")
		/*fmt.Println("file:", upldFile.Name())*/
		FILENAME = upldFile.Name()
		if err != nil {
			log.Fatal("[Fail] Temporary File    ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Computation Error   ", "/dashboard", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return nil
			}
			return nil

		}
		defer upldFile.Close()
		// file convert into bytes
		bytesFile, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal("[Fail] File Reading Permission   ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! Computation Error   ", "/dashnaord", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return nil
			}
			return nil

		}

		upldFile.Write(bytesFile)
		fmt.Println("File added on server")
		return upldFile
	}
	return nil
}

// []amino.AminoClass, []amino.AminoClass
func SequenceFile(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error) {

	seq, err := ReadSequence(userFile.Name())
	if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}

	var gen []string
	for _, v := range seq {
		space := DoAscii(v)
		if space == "" {
			fmt.Printf("Gap%v:\t", space)
		}
		gen = append(gen, space)
	}
	fmt.Println("Gen:{", gen, "}")

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

	pathogen, err := ReadSequence(serverFile.Name())
	if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}

	var genV []string
	for _, v := range pathogen {
		space := DoAscii(v)
		if space == "" {
			fmt.Printf("Gap%v:\t", space)
		}
		genV = append(genV, space)
	}
	fmt.Println("Genes:{", genV, "}")

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
	return gen, genV, nil

}

func DoAscii(seq byte) string {

	if seq >= 65 && seq < 91 {
		return string(alphabet.Letter(seq))
	}
	return string(alphabet.Letter(seq))
}

func RNASequence(sq []string) []string {

	var k []string

	for i, _ := range sq {

		if sq[i] == "T" {
			sq[i] = "U"
		}
		k = append(k, sq[i])
	}

	return k

}

func bioChemRecord(st2 string) structs.MolecularBio {

	molecule := structs.MolecularBio{}
	// helx record
	hlix := *pdb.ParseHelix(st2)
	fmt.Println("Serial:", hlix.Serial)
	fmt.Println("Id:", hlix.HelixID)
	fmt.Println("ResName+:", hlix.InitResName)
	fmt.Println("ChainId+:", hlix.InitChainID)
	fmt.Println("SeqNum+:", hlix.InitSeqNum)
	fmt.Println("Icode+:", hlix.InitICode)
	fmt.Println("ResName-:", hlix.EndResName)
	fmt.Println("ChainId-:", hlix.EndChainID)
	fmt.Println("SeqNum-:", hlix.EndSeqNum)
	fmt.Println("Icode-:", hlix.EndICode)
	fmt.Println("HelixClass:", hlix.HelixClass)
	fmt.Println("Length:", hlix.Length)
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

func RNAToAminoAcids(s []string) []amino.AminoClass {

	bases := []string{}
	for i, _ := range s {
		bases = append(bases, s[i])
	}

	proteins := amino.AminoClass{}

	ls := proteins.Bases(bases)

	return ls
}

func blockSession(id int) *sessions.CookieStore {

	return sessions.NewCookieStore([]byte(strconv.Itoa(id)))
}


func ChoosePattern(w http.ResponseWriter, r *http.Request, fname , choose string, file *os.File) (structs.Levenshtein, error) {
	
	
	i  , err := strconv.Atoi(choose); if err != nil {
		log.Fatalln("[Fail] Sorry there is some issue report!", err)
		return edit, err
	}
	if (i > 0 && i < 6) && (fname != " ") {
		svrFile := FileReadFromDisk(w, r, fname)		
		Usr, Virus, err := SequenceFile(file, svrFile); if err != nil {
					log.Fatalln("[Fail] Sequence DataFile Error", err)
					return  edit, err
		}
		log.Println("Genome:", len(Usr), "virus:", len(Virus))
		distance := structs.EditDistanceStrings(Usr,Virus)
		edit.Probablity = edit.Result(distance)
		edit.Name = fname
		edit.Percentage = edit.CalcualtePercentage(edit.Probablity)
		return edit, err
	}else if i == 0{
		temFile := template.Must(template.ParseFiles("dashboard.html"))
		temFile.Execute(w, "Dashbaord")
	}
	return edit, err


}

func ProcessWaiit() bool {
	clockHand := time.NewTimer(3 * time.Second)
	<-clockHand.C
	stop := clockHand.Stop()
	return stop
}

func AutoKeyGenerate(s1 string) string{
	h0 := sha256.New()
	h1 := h0.Sum([]byte(s1)) // hash of string-1 
	e := hex.EncodeToString([]byte(h1))
	h := hex.EncodeToString([]byte(e))[:8]
	return h
}


