package main

import (
	"bytes"
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
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	templates "text/template"
	"time"

	firebase "firebase.google.com/go"
	"github.com/ali2210/wizdwarf/db"
	cloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
	DBModel "github.com/ali2210/wizdwarf/db/model"
	"github.com/ali2210/wizdwarf/structs"
	weather "github.com/ali2210/wizdwarf/structs/OpenWeather"
	"github.com/ali2210/wizdwarf/structs/amino"
	bio "github.com/ali2210/wizdwarf/structs/bioinformatics"
	info "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
	"github.com/ali2210/wizdwarf/structs/paypal/handler"
	"github.com/ali2210/wizdwarf/structs/users"
	"github.com/ali2210/wizdwarf/structs/users/model"
	"github.com/biogo/biogo/alphabet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fogleman/ribbon/pdb"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	pay "github.com/logpacker/PayPal-Go-SDK"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// Variables

var (
	emailexp            string                    = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp             string                    = "([A-Z][a-z]*[0-9])*"
	addressexp          string                    = "(^0x[0-9a-fA-F]{40}$)"
	appName             *firebase.App             = SetFirestoreCredentials() // Google_Cloud [Firestore_Reference]
	cloud               users.DBFirestore         = users.NewCloudInstance()
	digitalCode         users.CreditCardInfo      = users.NewClient()
	vault               DBModel.Private           = DBModel.New()
	ledger              db.PublicLedger           = db.NewCollectionInstance()
	paypalMini          handler.PaypalClientLevel = handler.PaypalClientGo()
	userSessions        *sessions.CookieStore     = nil //user level
	clientInstance      *ethclient.Client         = nil
	ethAddrressGenerate string                    = ""
	ledgerPubcKeys      string                    = ""
	ledgerBits          string                    = ""
	googleCredentials   string                    = ""
	openReadFile        string                    = ""
	publicAddress       string                    = ""
	edit                bio.LevenTable            = bio.NewMatch()
	algo                info.Levenshtein          = info.Levenshtein{}
	visualizeReport     weather.DataVisualization = weather.DataVisualization{}
	accountID           string                    = " "
	accountKey          string                    = " "
	accountVisitEmail   string                    = " "
)

// Constants

const (
	projectID      string = "htickets-cb4d0"
	configFilename string = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	//Google_Credentials string = "/home/ali/Desktop/htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	// Main application
	mainNet string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	// Rickeby for test purpose
	rinkebyClient string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	geocodeAPI    string = "7efdb33c59a74e09352479b21657aee8"
)

func main() {

	// Server
	log.Println("[OK] Wiz-Dwarfs starting")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	//  env host
	if host == "" {

		// env port setting

		if port == " " {
			log.Fatalln("[Fail] No Application port allocated", port)
			log.Fatalln("[Fail] No Application hostname allocated", host)
		} else {
			if port != "5000" {
				// any Listening PORT {heroku}
				log.Println("[Open] Application Port", port)
				log.Println("[Open] Application host", host)
			} else {
				// specfic port allocated {docker}
				port = "5000"
				log.Println("[New] Application Default port", port)
			}

		}
	} else {
		log.Println("[Accept] Application hostname allocated", host)
	}

	log.Println("[OK] Application Explicit Credentials :", host+":", "Port", port)
	// Routing
	routing := mux.NewRouter()

	// Links
	routing.HandleFunc("/", func(arg1 http.ResponseWriter, arg2 *http.Request) {

		temp := template.Must(template.ParseFiles("initial.html"))

		if arg2.Method == "GET" {
			log.Println("[OK] URL :", arg2.URL.Path)
			temp.Execute(arg1, "MainPage")
		}
		_ = ProcessWaiit()

	})
	routing.HandleFunc("/home", home)
	routing.HandleFunc("/signup", newUser)
	routing.HandleFunc("/login", existing)
	routing.HandleFunc("/dashboard", dashboard)
	routing.HandleFunc("/dashbaord/setting", setting)
	routing.HandleFunc("/dashbaord/setting/profile", profile)
	routing.HandleFunc("/dashbaord/setting/about", aboutMe)
	routing.HandleFunc("/dashboard/setting/pay/credit/add", credit)
	routing.HandleFunc("/dashbaord/setting/pay/credit/delete", deleteCard)
	routing.HandleFunc("/logout", logout)
	routing.HandleFunc("/createwallet", createWallet)
	routing.HandleFunc("/terms", terms)
	routing.HandleFunc("/open", wallet)
	routing.HandleFunc("/transact", transacts)
	routing.HandleFunc("/transact/send", send)
	routing.HandleFunc("/transact/treasure", treasure)
	routing.HandleFunc("/visualize", visualize)
	routing.HandleFunc("/modal/success", success)

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
	err := http.ListenAndServe(net.JoinHostPort(host, port), routing)
	if err != nil {
		log.Println("Listening Error: ", err)
		panic(err)
	}

}

// Routes Handle

func home(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("index.html"))

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "home")
	}

}

func success(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("modal-success.html"))
	if r.Method == "GET" {
		log.Println("[Accept]", r.URL.Path)
		temp.Execute(w, "success")
	}
}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("delete.html"))

	if r.Method == "GET" {

		log.Println("[Accept]", r.URL.Path)
		temp.Execute(w, "DeleteForm")
	} else {

		log.Println("[Accept]", r.URL.Path)
		log.Println("Method:" + r.Method)

		r.ParseForm()

		//ccv := r.FormValue("prefixInside")
		accountNum := r.FormValue("account")
		//cardInfoID :=  digitalCode.GetAuthorizeStoreID()

		//log.Println("card info:", cardInfoID)
		// hashcode := AutoKeyGenerate(ccv)

		client, err := paypalMini.NewClient()
		if err != nil {
			log.Fatalln("[Fail] Client Operation:", err)
			return
		}

		token, err := paypalMini.Token(client)
		if err != nil {
			log.Fatalln("[Fail]Token Operation:", err)
			return
		}

		ret, err := paypalMini.RetrieveCreditCardInfo(digitalCode.GetAuthorizeStoreID(), client)
		if err != nil {
			log.Fatalln("[Fail] CreditCard Info Operation:", err)
			return
		}

		if accountNum != "" && cardNumberValid(accountNum, ret.Number) {
			err := paypalMini.RemoveCard(ret.ID, client)
			if err != nil {
				log.Fatalln("[Fail] Remove card operation", err)
				return
			}

			log.Println("[Accept:]", client, token)
			temp.Execute(w, "Complete")
			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			success(w, r)
		}
	}

}

func aboutMe(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("about.html"))

	if r.Method == "GET" {
		log.Println("[Accept]", r.URL.Path)

		// cardInfoID :=  digitalCode.GetAuthorizeStoreID()

		userProfile, err := cloud.FindAllData(appName, accountVisitEmail, accountKey)
		if err != nil && userProfile != nil {
			log.Fatal("[Fail] No info  ", err)
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry ! No Information ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}

		}
		// log.Println("User profile:", userProfile)

		key, address := vault.GetCryptoDB(publicAddress)
		access := DBModel.CredentialsPrivate{
			PublicAddress: address,
			PrvteKey:      key,
		}

		log.Println("Ledger Info:", access)

		client, err := paypalMini.NewClient()
		if err != nil {
			log.Fatalln("[Fail] Client Operation:", err)
			return
		}

		_, err = paypalMini.Token(client)
		if err != nil {
			log.Fatalln("[Fail] Token Operation:", err)
			return
		}

		ret, err := paypalMini.RetrieveCreditCardInfo(digitalCode.GetAuthorizeStoreID(), client)
		if err != nil {
			log.Fatalln("[Fail] Retrieve Card info Operation:", err)
			return
		}

		// err = paypalMini.RemoveCard(ret.ID,client); if err != nil {
		// 	log.Fatalln("[Fail] Remove Card operation:", err)
		// 	return
		// }

		myProfile := model.DigialProfile{

			Name:     userProfile.Name,
			FName:    userProfile.FName,
			Email:    userProfile.Email,
			Address:  userProfile.Address,
			LAddress: userProfile.LAddress,
			City:     userProfile.City,
			Zip:      userProfile.Zip,
			Country:  userProfile.Country,

			Public:  access.PublicAddress,
			Private: access.PrvteKey,

			Number:      ret.Number,
			Type:        ret.Type,
			ExpireMonth: ret.ExpireMonth,
			ExpireYear:  ret.ExpireYear,
		}

		log.Println("[Accept] Profile", myProfile)
		temp.Execute(w, myProfile)

	}
}

func setting(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("settings.html"))

	// bankProfile , _ := paypalMini.RetrieveCreditCardInfo(accountID)

	if r.Method == "GET" {
		log.Println("[Accept]", r.URL.Path)
		temp.Execute(w, "setting")
	}

}

func profile(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("profile.html"))

	if accountID == "" {
		log.Fatal("[Error ] Please login  ")
		response := structs.Response{}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry Session expire   ", "/login", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return
		}
		return
	}
	detailsAcc, err := cloud.ToFindByGroupSet(accountID, accountVisitEmail, appName)
	if err != nil {
		log.Fatalln("[Fail] Operation..", err)
		return
	}
	log.Println(detailsAcc)

	if r.Method == "GET" {
		log.Println("[Accept] Method:", r.Method)
		log.Println("[Accept]", r.URL.Path)

		log.Println("Details:", detailsAcc)
		temp.Execute(w, detailsAcc)
	} else {
		log.Println("[Accept] Method:", r.Method)
		log.Println("[Accept] Path:", r.URL.Path)

		r.ParseForm()

		// save value in db
		MyProfile := model.UpdateProfile{
			Email:        r.FormValue("email"),
			Phone:        r.FormValue("phone"),
			FirstName:    r.FormValue("uname"),
			LastName:     r.FormValue("ufname"),
			HouseAddress: r.FormValue("inputAddress"),
			SubAddress:   r.FormValue("inputAddress2"),
			Country:      r.FormValue("country"),
			Zip:          r.FormValue("inputZip"),
		}

		if accountID == "" {
			log.Fatal("[Error ] Please login  ")
			response := structs.Response{}
			temp := server(w, r)
			_ = response.ClientRequestHandle(true, "Sorry Session expire   ", "/login", w, r)
			response.ClientLogs()
			err := response.Run(temp)
			if err != nil {
				log.Println("[Error]: checks logs...", err)
				return
			}
		}
		MyProfile.Id = accountID

		male := r.FormValue("gender")
		if male == "on" {
			MyProfile.Male = true
		}
		MyProfile.Male = false

		profile, err := cloud.UpdateProfiles(appName, &MyProfile)
		if err != nil {
			log.Fatalln("[Fail] Operation..", err)
			return
		}

		log.Println("[Accept] Profile updated... ", profile)
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		success(w, r)

	}
}

func credit(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("credit.html"))
	if r.Method == "GET" {
		log.Println("[Accept] Path :", r.URL.Path)
		log.Println("Method :", r.Method)
		temp.Execute(w, "Credit")
	} else {
		log.Println("[Accept ]Path :", r.URL.Path)
		log.Println("Method :", r.Method)
		r.ParseForm()
		calender := r.FormValue("expire")
		sliceByte := []byte(calender)
		year := string(sliceByte[:04])
		month := string(sliceByte[5:])

		card := pay.CreditCard{
			FirstName:   r.FormValue("fholder"),
			LastName:    r.FormValue("surename"),
			Number:      r.FormValue("cardNo"),
			CVV2:        r.FormValue("cvv"),
			Type:        r.FormValue("cardtype"),
			ExpireMonth: month,
			ExpireYear:  year,
		}

		//  card.ID =  AutoKeyGenerate(card.CVV2)
		//  cardInfoID = card.ID
		//  log.Println("Id generated:" , card)

		// store credit card information.
		//mini := handler.PaypalMiniVersion{}
		client, err := paypalMini.NewClient()
		if err != nil {
			log.Fatalln("[Fail] Client Operation:", err)
			return
		}
		//mini.Client = client

		token, err := paypalMini.Token(client)
		if err != nil {
			log.Fatalln("[Fail] Token Operation:", err)
			return
		}

		store, err := paypalMini.StoreCreditCardInfo(card, client)
		if err != nil {
			log.Fatalln("[Fail] CreditCard Operation:", err, card)
			return
		}

		digitalCode.SetAuthorizeStoreID(store.ID)

		ret, err := paypalMini.RetrieveCreditCardInfo(digitalCode.GetAuthorizeStoreID(), client)
		if err != nil {
			log.Fatalln("[Fail] Retrieve Card info Operation:", err)
			return
		}

		log.Println("[Accept] Token issue:", token, "retInfo:", ret)
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		success(w, r)

	}

}

func visualize(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("visualize.html"))
	log.Println("Report percentage", visualizeReport.Percentage)
	log.Println("Report uv ", visualizeReport.UVinfo)
	algo.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" && algo.GetProbParameter() != -1.0 {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		temp.Execute(w, visualizeReport)
	}
	// err := SessionExpire(w,r); if err != nil {
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	// r.Method = "GET"
	// dashboard(w,r)
}

func treasure(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("treasure.html"))
	acc := structs.Static{}
	block := structs.Block{}

	if r.Method == "GET" {

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		acc.Eth = ethAddrressGenerate
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

			chainID, err := clientInstance.NetworkID(context.Background())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// get recipt address
			message, err := tx.AsMessage(types.NewEIP155Signer(chainID))
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

func dashboard(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp := template.Must(template.ParseFiles("dashboard.html"))
		fmt.Println("Method:" + r.Method)
		fmt.Println("Url:", r.URL.Path)
		temp.Execute(w, "dashboard")
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
				err := ChoosePattern(w, r, "", choose, file)
				if err != nil {
					return
				}
			case "1":
				var name string = "Covid-19"
				err := ChoosePattern(w, r, name, choose, file)
				if err != nil {
					return
				}
				visualizeReport.Percentage = algo.Percentage
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
				visualize(w, r)
				// fmt.Println("Virus:", capsid)

			case "2":
				var name string = "FlaviDengue"
				err := ChoosePattern(w, r, name, choose, file)
				if err != nil {
					return
				}
				visualizeReport.Percentage = algo.Percentage
				// v :=  infectedUv()
				// v.UVinfo = uvslice
				w.WriteHeader(http.StatusOK)

				r.Method = "GET"
				visualize(w, r)
				// Wallet(w,r)
				// fmt.Println("Virus:", capsid)
			case "3":
				var name string = "KenyaEbola"
				err := ChoosePattern(w, r, name, choose, file)
				if err != nil {
					return
				}
				visualizeReport.Percentage = algo.Percentage
				// v :=  infectedUv()
				// v.UVinfo = uvslice

				w.WriteHeader(http.StatusOK)
				r.Method = "GET"
				visualize(w, r)

				// fmt.Println("Virus:", capsid)
			case "4":
				var name string = "ZikaVirusBrazil"
				err := ChoosePattern(w, r, name, choose, file)
				if err != nil {
					return
				}
				visualizeReport.Percentage = algo.Percentage
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
				visualize(w, r)

				// fmt.Println("Virus:", capsid)
			case "5":
				var name string = "MersSaudiaArabia"
				err := ChoosePattern(w, r, name, choose, file)
				if err != nil {
					return
				}
				visualizeReport.Percentage = algo.Percentage
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
				visualize(w, r)
				// fmt.Println("Virus:", capsid)

			default:
				temFile := template.Must(template.ParseFiles("dashboard.html"))
				temFile.Execute(w, "dashboard")
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

func send(w http.ResponseWriter, r *http.Request) {

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
		// fmt.Println("Wallet kEY:", ledgerBits)

		// private key to public address
		secure, err := crypto.HexToECDSA(ledgerBits)
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
		chainID, err := clientInstance.NetworkID(context.Background())

		var nofield []byte

		tx := types.NewTransaction(block.Nonce, transfer, block.Amount, block.GasLimit, block.GasPrice, nofield)

		// Signed Transaction
		sign, err := types.SignTx(tx, types.NewEIP155Signer(chainID), secure)
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
		visualize(w, r)

	}
}

func createWallet(w http.ResponseWriter, r *http.Request) {

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

		client, err := ethclient.Dial(mainNet)
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

		acc.PubKey = PublicKey[:8]
		acc.PrvteKey = key[:8]

		// hash to ethereum
		hshCode := sha3.NewLegacyKeccak256()
		hshCode.Write(publicBytes[1:])
		ethereum := hexutil.Encode(hshCode.Sum(nil)[12:])

		acc.EthAddress = ethereum[:8]

		// valid address
		valid := isYourPublcAdresValid(acc.EthAddress)
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

		log.Println("[Accept] Welcome ! Your Account Has been created", &merchant)
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		existing(w, r)
	}
}

func transacts(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("transact.html"))
	acc := structs.Static{}
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		acc.Eth = ethAddrressGenerate
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

func wallet(w http.ResponseWriter, r *http.Request) {

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

		client, err := ethclient.Dial(mainNet)
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
			publicAddress = acc.EthAddress

			// Secure Key
			ledgerBits = addr.PrvteKey
			fmt.Println("Wallet key:", ledgerBits)
			vault.SetCryptoDB(acc.EthAddress, ledgerBits)

			// variable address for futher processing
			ethAddrressGenerate = acc.EthAddress
			log.Println("Your Wallet:", ethAddrressGenerate)

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
				_ = response.ClientRequestHandle(true, "Sorry ! NO CRYPTO WALLET   ", "/createWallet", w, r)
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
			dashboard(w, r)
		}
	}
}

func terms(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("terms.html"))
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Terms")
	}
}

func newUser(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("register.html"))
	user := model.Create_User{}

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
		log.Println("[Accept] Review your logs.. ")
		// println("phase:", KeyTx)
		addVistor(w, r, &user, encrypted.Reader)
	}

}

func existing(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("login.html"))
	user := model.Create_User{}

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
			// existing(w,r)
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
			accountVisitEmail = data.Email
			accountKey = data.Password
			act := structs.RouteParameter{}
			//complex.AddByName(data.Name)
			// User Session
			if userSessions == nil {
				userSessions = SessionsInit(data.Id)

				act.SetContextSession(userSessions, w, r)
				err := act.NewToken()
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
			dashboard(w, r)
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

func logout(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		act := structs.RouteParameter{}

		log.Println("[Access] ", r.URL.Path)
		act.SetContextSession(userSessions, w, r)
		err := act.ExpireToken()
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
		existing(w, r)
	}
}

//  Advance Functions

func SearchDB(w http.ResponseWriter, r *http.Request, email, pass string) (*model.Vistors, error) {

	var data *model.Vistors
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
		log.Println("[Accept] Your Request results :")
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

func addVistor(response http.ResponseWriter, request *http.Request, user *model.Create_User, im string) {

	// var err error
	//response.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		fmt.Println("Method:" + request.Method)
	} else {
		var member model.Vistors
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

			log.Println("Records:", record)
			response.WriteHeader(http.StatusOK)
			request.Method = "GET"
			createWallet(response, request)

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
	conf := &firebase.Config{ProjectID: projectID}
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

func MessageToHash(w http.ResponseWriter, r *http.Request, matchE, matchP bool, user model.Create_User) (bool, *structs.SignedKey) {
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
		openReadFile = upldFile.Name()
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

	for i := range sq {

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
	for i := range s {
		bases = append(bases, s[i])
	}

	proteins := amino.AminoClass{}

	ls := proteins.Bases(bases)

	return ls
}

func blockSession(id int) *sessions.CookieStore {

	return sessions.NewCookieStore([]byte(strconv.Itoa(id)))
}

func ChoosePattern(w http.ResponseWriter, r *http.Request, fname, choose string, file *os.File) error {

	i, err := strconv.Atoi(choose)
	if err != nil {
		log.Fatalln("[Fail] Sorry there is some issue report!", err)
		return err
	}
	if (i > 0 && i < 6) && (fname != " ") {
		svrFile := FileReadFromDisk(w, r, fname)
		Usr, Virus, err := SequenceFile(file, svrFile)
		if err != nil {
			log.Fatalln("[Fail] Sequence DataFile Error", err)
			return err
		}
		log.Println("Genome:", len(Usr), "virus:", len(Virus))
		distance := edit.EditDistanceStrings(Usr, Virus)
		algo.Probablity = algo.Result(distance)
		algo.Name = fname
		algo.Percentage = algo.CalcualtePercentage(algo.Probablity)
		return err
	} else if i == 0 {
		temFile := template.Must(template.ParseFiles("dashboard.html"))
		temFile.Execute(w, "Dashbaord")
	}
	return err

}

func ProcessWaiit() bool {
	clockHand := time.NewTimer(3 * time.Second)
	<-clockHand.C
	stop := clockHand.Stop()
	return stop
}

// func AutoKeyGenerate(s1 string) string{
// 	h0 := sha256.New()
// 	h1 := h0.Sum([]byte(s1)) // hash of string-1
// 	e := hex.EncodeToString([]byte(h1))
// 	h := hex.EncodeToString([]byte(e))[:8]
// 	return h
// }

func cardNumberValid(s1, s2 string) bool {

	m, n := []byte(s1), []byte(s2)
	res := bytes.Compare(m, n)
	if res == 0 {
		return true
	}
	return false

}
