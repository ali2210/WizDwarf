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
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	templates "text/template"
	"time"

	firebase "firebase.google.com/go"
	"github.com/ali2210/wizdwarf/db"
	cloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
	DBModel "github.com/ali2210/wizdwarf/db/model"
	"github.com/ali2210/wizdwarf/structs"
	bio "github.com/ali2210/wizdwarf/structs/bioinformatics"
	info "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
	Shop "github.com/ali2210/wizdwarf/structs/cart"
	weather "github.com/ali2210/wizdwarf/structs/openweather"
	"github.com/ali2210/wizdwarf/structs/paypal/handler"
	wizSdk "github.com/ali2210/wizdwarf/structs/transaction"
	"github.com/ali2210/wizdwarf/structs/users"
	"github.com/ali2210/wizdwarf/structs/users/model"
	"github.com/biogo/biogo/alphabet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	pay "github.com/logpacker/PayPal-Go-SDK"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// Variables

var (
	emailexp          string                    = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp           string                    = "([A-Z][a-z]*[0-9])*"
	addressexp        string                    = "(^0x[0-9a-fA-F]{40}$)"
	appName           *firebase.App             = SetFirestoreCredentials() // Google_Cloud [Firestore_Reference]
	cloud             users.DBFirestore         = users.NewCloudInstance()
	digitalCode       users.CreditCardInfo      = users.NewClient()
	vault             DBModel.Private           = DBModel.New()
	ledger            db.PublicLedger           = db.NewCollectionInstance()
	paypalMini        handler.PaypalClientLevel = handler.PaypalClientGo()
	userSessions      *sessions.CookieStore     = nil //user level
	clientInstance    *ethclient.Client         = nil
	ledgerPubcKeys    string                    = ""
	ledgerBits        string                    = ""
	googleCredentials string                    = ""
	openReadFile      string                    = ""
	publicAddress     string                    = ""
	edit              bio.LevenTable            = bio.NewMatch()
	algo              info.Levenshtein          = info.Levenshtein{}
	visualizeReport   weather.DataVisualization = weather.DataVisualization{}
	accountID         string                    = " "
	accountKey        string                    = " "
	accountVisitEmail string                    = " "
	checkout          Shop.Shopping             = Shop.Shopping{
		Price:         "",
		TypeofService: "",
		PaymentMethod: "",
		Description:   "",
	}
	payment wizSdk.BankRecord = wizSdk.BankRecord{
		Name:       "",
		Btc:        "",
		CreditCard: "",
		TotalCash:  "",
		Public:     publicAddress,
	}
	cart        Shop.Items         = Shop.Items{}
	balance     wizSdk.FingerPrint = wizSdk.FingerPrint{}
	blockchains structs.Block      = structs.Block{
		Balance:        &big.Int{},
		SenderBatchID:  "",
		RecieveBatchID: "",
		Amount:         &big.Int{},
		Nonce:          0,
		GasPrice:       &big.Int{},
		GasLimit:       0,
		DataBlock:      structs.DataBlock{},
	}
	genesis      structs.BlockTransactionGateway = structs.BlockTransactionGateway{}
	eth          structs.EthToken                = structs.EthToken{}
	bitInterface structs.BitsBlocks              = structs.BitsBlocks{
		SenderBatchID:    "",
		SenderPrivateKey: &ecdsa.PrivateKey{},
	}
)

// Constants

const (
	projectID      string = "htickets-cb4d0"
	configFilename string = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	mainNet        string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	rinkebyClient  string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	geocodeAPI     string = "7efdb33c59a74e09352479b21657aee8"
	// serviceID0     string = "kernel"
	// serviceID1     string = "cluster"
	// serviceID2     string = "multicluster"
)

func main() {

	// Server
	log.Println("[OK] Wiz-Dwarfs starting")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	wizDir := os.Getenv("WIZ_VOLUME_DIR")

	if wizDir == "" {
		log.Fatalln("Make sure volume mount", wizDir)
		panic(errors.New("Fail to mount"))
	}

	if host == "" {

		// env port setting

		if port == " " {
			log.Fatalln("[Fail] No Application port allocated", port)
		} else {
			if port != "5000" && host == "wizdwarfs" {
				// any Listening PORT {heroku}
				log.Println("[Open] Application Port", port)
				log.Println("[Open] Application host", host)
				// url = "https://" + "127.0.0.1:" + port + "/"
			} else {
				// specfic port allocated {docker}
				port = "5000"
				host = "wizdwarfs"
				log.Println("[New] Application Default port", port)
				log.Println("[Host] Explicit Host ", host)
				// url = "https://" + host + ".io" + "/"
			}

		}
	} else {
		log.Println("[Accept] Application hostname allocated", host)
	}

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
	routing.HandleFunc("/transact/pay/paypal/kernel", kernel)
	routing.HandleFunc("/transact/pay/paypal/cluster", cluster)
	routing.HandleFunc("/transact/pay/paypal/multicluster", multicluster)
	routing.HandleFunc("/transact/pay/crypto/kernel", tKernel)
	routing.HandleFunc("/transact/pay/crypto/cluster", tCluster)
	routing.HandleFunc("/transact/pay/crypto/multicluster", tMulticluster)

	// routing.HandleFunc("/transact/send", send)
	// routing.HandleFunc("/transact/userCredit", userCredit)
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

	// open browser
	// err := openBrowser(url)
	// if err != nil {
	// 	panic(err)
	// }

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
		} else {
			w.WriteHeader(http.StatusBadRequest)
			r.Method = "GET"
			deleteCard(w, r)
		}
	}

}

func aboutMe(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("about.html"))

	if r.Method == "GET" {
		log.Println("[Accept]", r.URL.Path)

		// cardInfoID :=  digitalCode.GetAuthorizeStoreID()

		userProfile, err := cloud.GetProfile(appName, accountID, accountVisitEmail)
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

		myProfile := model.DigialProfile{
			Public:      access.PublicAddress,
			Private:     access.PrvteKey,
			Name:        userProfile.FirstName,
			FName:       userProfile.LastName,
			Email:       userProfile.Email,
			Address:     userProfile.HouseAddress1,
			LAddress:    userProfile.HouseAddress2,
			City:        userProfile.City,
			Zip:         userProfile.Zip,
			Country:     userProfile.Country,
			Phone:       userProfile.PhoneNo,
			Twitter:     userProfile.Twitter,
			Number:      ret.Number,
			ExpireMonth: ret.ExpireMonth,
			ExpireYear:  ret.ExpireYear,
			Type:        ret.Type,
		}

		log.Println("[Accept] Profile", myProfile)
		temp.Execute(w, myProfile)

	}
}

func setting(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("settings.html"))

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
	detailsAcc, err := cloud.GetProfile(appName, accountID, accountVisitEmail)
	if err != nil {
		log.Fatalln("[Fail] Operation..", err)
		return
	}

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
			Id:           accountID,
			FirstName:    r.FormValue("uname"),
			LastName:     r.FormValue("ufname"),
			Phone:        r.FormValue("phone"),
			HouseAddress: r.FormValue("address"),
			SubAddress:   r.FormValue("inputAddress2"),
			Country:      r.FormValue("country"),
			Zip:          r.FormValue("inputZip"),
			Email:        r.FormValue("email"),
			Twitter:      r.FormValue("tweet"),
			City:         r.FormValue("city"),
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

		// store credit card information.
		client, err := paypalMini.NewClient()
		if err != nil {
			log.Fatalln("[Fail] Client Operation:", err)
			return
		}

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

func cluster(w http.ResponseWriter, r *http.Request) {

}

func kernel(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("payee.html"))
	var check users.Analysis = users.Analysis{}
	user, err := cloud.GetProfile(appName, accountID, accountVisitEmail)
	if err != nil {
		log.Fatalln("[Fail] Operation..", err)
		return
	}

	key, address := vault.GetCryptoDB(publicAddress)
	credentials := DBModel.CredentialsPrivate{
		PublicAddress: address,
		PrvteKey:      key,
	}

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

	order := cart.GetItemsFromCart()
	batchID := order.TypeofService + user.ID
	resp, err := paypalMini.GetPayout(batchID, client)
	if err != nil {
		return
	}

	amountStr, err := check.MarshalJSONAmount(resp)
	if err != nil {
		return
	}
	feesStr, err := check.MarshalJSONFees(resp)
	if err != nil {
		return
	}
	amountValue := check.Encode(amountStr)
	feesValue := check.Encode(feesStr)
	amount, err := check.CalculateNum(amountValue)
	if err != nil {
		return
	}
	fees, err := check.CalculateNum(feesValue)
	if err != nil {
		return
	}
	total := check.CalculateTotalBalance(amount, fees)
	sTotal := fmt.Sprintf("%f", total)
	// blockchains.SenderBatchID = ethAddrressGenerate
	if blockchains.SenderBatchID == " " {
		log.Fatal("Public Address:", err)
		return
	}
	err = userCredit()
	if err != nil {
		fmt.Println("Error:")
	}
	balance.SetTransactionWiz(user.FirstName, fmt.Sprintf("%v", blockchains.Balance), sTotal, sTotal, credentials.PublicAddress)

	if r.Method == "GET" {
		log.Println("[Accept] Path :", r.URL.Path)
		log.Println("Method :", r.Method)
		temp.Execute(w, balance.GetTransactionWiz())
	} else {
		log.Println("[Accept] Path :", r.URL.Path)
		log.Println("Method :", r.Method)
		res, err := paypalMini.PaypalPayout(ret.ID, batchID, accountVisitEmail, order.Price, client)
		if err != nil {
			return
		}

		amountStr, err := check.MarshalJSONAmount(res)
		if err != nil {
			return
		}
		feesStr, err := check.MarshalJSONFees(res)
		if err != nil {
			return
		}
		amountValue := check.Encode(amountStr)
		feesValue := check.Encode(feesStr)
		amount, err := check.CalculateNum(amountValue)
		if err != nil {
			return
		}
		fees, err := check.CalculateNum(feesValue)
		if err != nil {
			return
		}
		total := check.CalculateTotalBalance(amount, fees)
		sTotal := fmt.Sprintf("%f", total)
		balance.SetTransactionWiz(user.FirstName, fmt.Sprintf("%v", blockchains.Balance), sTotal, sTotal, credentials.PublicAddress)
		temp.Execute(w, balance.GetTransactionWiz())
	}
}

func multicluster(w http.ResponseWriter, r *http.Request) {

}

func tCluster(w http.ResponseWriter, r *http.Request) {
	webpage := template.Must(template.ParseFiles(""))

	if r.Method == "GET" {
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		webpage.Execute(w, "Cluster")
	} else {
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		order := cart.GetItemsFromCart()
		value, err := strconv.ParseUint(order.Price, 2, 10)
		if err != nil {
			return
		}
		blockchains.Checkout = value
		blockchains.Amount = new(big.Int).SetUint64(uint64(blockchains.Checkout))
		switch order.TypeofService {
		case "Kernel":
			n, m := int64(2), int64(60)
			automatedNetworkFees(n, m)
		case "Cluster":
			n, m := int64(3), int64(110)
			automatedNetworkFees(n, m)
		case "Multi-Cluster":
			n, m := int64(5), int64(550)
			automatedNetworkFees(n, m)
		default:
			w.WriteHeader(http.StatusPreconditionFailed)
			r.Method = "GET"
			transacts(w, r)
		}
	}
	// Send Transaction
	blockchains.RecieveBatchID = "0x55057eb78fDbF783C961b4AAd6A5f8BC60cab44B"
	bitInterface.EthReciptAddress = eth.BTCAddressHex(blockchains.RecieveBatchID)

	// Network ID
	chainID, err := clientInstance.NetworkID(context.Background())
	bitInterface.EthTransaction = eth.BTCNewTransactions(blockchains, bitInterface)
	bitInterface.FingerPrint, err = eth.BTCTransactionSignature(chainID, bitInterface)
	if err != nil {
		log.Fatal("[Fail] Signed Transaction", err)
		return
	}

	// Send Transaction
	err = eth.TransferBTC(bitInterface.FingerPrint)
	if err != nil {
		log.Fatalln("[Fail] Transaction", err)
		return
	}
	response := structs.Response{
		Flag:    false,
		Message: "",
		Links:   "",
	}
	temp := server(w, r)
	_ = response.ClientRequestHandle(false, "[Operation] Successful , be proceed ", "/login", w, r)
	response.ClientLogs()
	err = response.Run(temp)
	if err != nil {
		log.Println("[Error]: checks logs...", err)
		return
	}

}

func tMulticluster(w http.ResponseWriter, r *http.Request) {
	webpage := template.Must(template.ParseFiles(""))
	btx, err := send()
	if err != nil {
		return
	}

	if r.Method == "GET" {
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		webpage.Execute(w, btx)
	} else {
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		order := cart.GetItemsFromCart()
		value, err := strconv.ParseUint(order.Price, 2, 10)
		if err != nil {
			return
		}
		blockchains.Checkout = value
		blockchains.Amount = new(big.Int).SetUint64(uint64(blockchains.Checkout))
		switch order.TypeofService {
		case "Kernel":
			n, m := int64(2), int64(60)
			automatedNetworkFees(n, m)
		case "Cluster":
			n, m := int64(3), int64(110)
			automatedNetworkFees(n, m)
		case "Multi-Cluster":
			n, m := int64(5), int64(550)
			automatedNetworkFees(n, m)
		default:
			w.WriteHeader(http.StatusPreconditionFailed)
			r.Method = "GET"
			transacts(w, r)
		}
	}
	// Send Transaction
	blockchains.RecieveBatchID = "0x55057eb78fDbF783C961b4AAd6A5f8BC60cab44B"
	bitInterface.EthReciptAddress = eth.BTCAddressHex(blockchains.RecieveBatchID)

	// Network ID
	chainID, err := clientInstance.NetworkID(context.Background())
	bitInterface.EthTransaction = eth.BTCNewTransactions(blockchains, bitInterface)
	bitInterface.FingerPrint, err = eth.BTCTransactionSignature(chainID, bitInterface)
	if err != nil {
		log.Fatal("[Fail] Signed Transaction", err)
		return
	}

	// Send Transaction
	err = eth.TransferBTC(bitInterface.FingerPrint)
	if err != nil {
		log.Fatalln("[Fail] Transaction", err)
		return
	}
	response := structs.Response{
		Flag:    false,
		Message: "",
		Links:   "",
	}
	temp := server(w, r)
	_ = response.ClientRequestHandle(false, "[Operation] Successful , be proceed ", "/login", w, r)
	response.ClientLogs()
	err = response.Run(temp)
	if err != nil {
		log.Println("[Error]: checks logs...", err)
		return
	}

}

func tKernel(w http.ResponseWriter, r *http.Request) {
	webpage := template.Must(template.ParseFiles(""))
	var check users.Analysis = users.Analysis{}
	user, err := cloud.GetProfile(appName, accountID, accountVisitEmail)
	if err != nil {
		log.Fatalln("[Fail] Operation..", err)
		return
	}

	key, address := vault.GetCryptoDB(publicAddress)
	credentials := DBModel.CredentialsPrivate{
		PublicAddress: address,
		PrvteKey:      key,
	}

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

	btx, err := send()
	if err != nil {
		return
	}
	log.Println("Block Header:", btx)

	// ret, err := paypalMini.RetrieveCreditCardInfo(digitalCode.GetAuthorizeStoreID(), client)
	// if err != nil {
	// 	log.Fatalln("[Fail] Retrieve Card info Operation:", err)
	// 	return
	// }
	blockchains.Nonce = bitInterface.EthNonceAtStatus
	blockchains.GasLimit = uint64(21000)

	bitInterface.EthGasUnits, err = eth.BTCGasConsumerPrice()
	if err != nil {
		return
	}

	order := cart.GetItemsFromCart()
	batchID := order.TypeofService + user.ID
	resp, err := paypalMini.GetPayout(batchID, client)
	if err != nil {
		return
	}

	amountStr, err := check.MarshalJSONAmount(resp)
	if err != nil {
		return
	}
	feesStr, err := check.MarshalJSONFees(resp)
	if err != nil {
		return
	}
	amountValue := check.Encode(amountStr)
	feesValue := check.Encode(feesStr)
	amount, err := check.CalculateNum(amountValue)
	if err != nil {
		return
	}
	fees, err := check.CalculateNum(feesValue)
	if err != nil {
		return
	}
	total := check.CalculateTotalBalance(amount, fees)
	sTotal := fmt.Sprintf("%f", total)
	if blockchains.SenderBatchID == "" {
		log.Fatalln("Public Address", err)
		return
	}
	// blockchains.SenderBatchID = blockchains.SenderBatchID
	err = userCredit()
	if err != nil {
		log.Fatal("[Fail] Block Balance should not be zero  ", blockchains.Balance)
		return
	}
	balance.SetTransactionWiz(user.FirstName, fmt.Sprintf("%v", blockchains.Balance), sTotal, sTotal, credentials.PublicAddress)

	if r.Method == "GET" {
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		webpage.Execute(w, balance.GetTransactionWiz())
	} else {

		// btx, err := send()
		// if err != nil {
		// 	return
		// }
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		order := cart.GetItemsFromCart()
		value, err := strconv.ParseUint(order.Price, 2, 10)
		if err != nil {
			return
		}
		blockchains.Checkout = value
		blockchains.Amount = new(big.Int).SetUint64(uint64(blockchains.Checkout))
		switch order.TypeofService {
		case "Kernel":
			n, m := int64(2), int64(60)
			automatedNetworkFees(n, m)
		case "Cluster":
			n, m := int64(3), int64(110)
			automatedNetworkFees(n, m)
		case "Multi-Cluster":
			n, m := int64(5), int64(550)
			automatedNetworkFees(n, m)
		default:
			w.WriteHeader(http.StatusPreconditionFailed)
			r.Method = "GET"
			transacts(w, r)
		}
	}
	// Send Transaction
	blockchains.RecieveBatchID = "0x55057eb78fDbF783C961b4AAd6A5f8BC60cab44B"
	bitInterface.EthReciptAddress = eth.BTCAddressHex(blockchains.RecieveBatchID)

	// Network ID
	chainID, err := clientInstance.NetworkID(context.Background())
	bitInterface.EthTransaction = eth.BTCNewTransactions(blockchains, bitInterface)
	bitInterface.FingerPrint, err = eth.BTCTransactionSignature(chainID, bitInterface)
	if err != nil {
		log.Fatal("[Fail] Signed Transaction", err)
		return
	}

	// Send Transaction
	err = eth.TransferBTC(bitInterface.FingerPrint)
	if err != nil {
		log.Fatalln("[Fail] Transaction", err)
		return
	}

}

func automatedNetworkFees(n, m int64) {
	fee := new(big.Int)
	result := new(big.Int)
	blockchains.GasPrice = bitInterface.EthGasUnits
	fee.SetInt64(n)
	result.Mul(blockchains.GasPrice, fee)
	result.Add(result, blockchains.Amount)
	y := new(big.Int).SetInt64(m)
	if result.Cmp(y) == 0 || result.Cmp(y) == 1 {
		blockchains.Amount = result
	} else {
		k := 0.5
		k -= float64(n)
		automatedNetworkFees((int64(k)), m)
	}
}

func visualize(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("visualize.html"))
	log.Println("Report percentage", visualizeReport.Percentage)
	log.Println("Report uv ", visualizeReport.UVinfo)
	userProfile, err := cloud.FindAllData(appName, accountVisitEmail, accountKey)
	if err != nil && userProfile != nil {
		log.Fatal("[Fail] No info  ", err)
		response := structs.Response{
			Flag:    false,
			Message: "",
			Links:   "",
		}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! No Information ", "/login", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return
		}

	}
	algo.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" && algo.GetProbParameter() != -1.0 {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		visualizeReport.Process = 1
		visualizeReport.SeenBy = userProfile.Name

		temp.Execute(w, visualizeReport)
	}

}

func userCredit() error {
	// temp := template.Must(template.ParseFiles("userCredit.html"))
	// acc := structs.Static{}
	// block := structs.Block{}
	var err error
	//blockchains.SenderBatchID = id
	blockchains.Balance, err = genesis.GetLastTransaction(blockchains)
	if err != nil {
		fmt.Println("Error:")
	}

	// block.TxSen = r.FormValue("send")
	// block.TxRec = r.FormValue("rece")

	return err
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	RouteWebpage := template.Must(template.ParseFiles("dashboard.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Url:", r.URL.Path)
		RouteWebpage.Execute(w, "Dashboard")
	} else {
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		// FILE Upload ....
		fname, err := UploadFiles(w, r)

		if err != nil {
			log.Fatalln("[File]:", err)
			return
		}
		choose := r.FormValue("choose")
		data, err := fileGet("seqDir/", fname)
		if err != nil {
			log.Fatalln("[No File]:", err)
			return
		}
		log.Println("File Name:", data.Name())
		switch choose {
		case "0":
			fmt.Fprintf(w, "Please choose any option ...")
			log.Fatalln("Choose your option")
			panic(err)
		case "1":
			var name string = "Covid-19"
			err := choosePattern(w, r, name, choose, data)
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
			err := choosePattern(w, r, name, choose, data)
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
			err := choosePattern(w, r, name, choose, data)
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
			err := choosePattern(w, r, name, choose, data)
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
			err := choosePattern(w, r, name, choose, data)
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
			RouteWebpage.Execute(w, "dashboard")
		}
	}

}

func send() (structs.BitsBlocks, error) {

	// temp := template.Must(template.ParseFiles("server.html"))

	var err error

	// if r.Method == "POST" {

	// 	fmt.Println("Url:", r.URL.Path)
	// 	fmt.Println("Method:" + r.Method)
	// 	r.ParseForm()
	// block.TxSen = r.FormValue("sendAdd")
	// block.TxRec = r.FormValue("add")
	// choice := r.FormValue("transact")
	// amount := r.FormValue("amount")
	// blockchains.SenderBatchID = id
	blockchains.BlockSenderID = structs.HeaderBlock
	// Block number
	blockchains.BlockHeaderID, err = clientInstance.BlockByNumber(context.Background(), blockchains.BlockSenderID)
	if err != nil {
		log.Fatal("[Fail] Header Number   ", err)
		return structs.BitsBlocks{}, err
	}
	bitInterface.SenderPrivateKey, err = eth.BTCECDSAHEX(blockchains)
	if err != nil {
		log.Fatal("[Fail] Secure Wallet Key  ", err)
		return structs.BitsBlocks{}, err

	}

	bitInterface.EthBlockHeader = eth.BTCHeaderBlockerID(blockchains)
	log.Println("Header :", bitInterface.EthBlockHeader)
	bitInterface.EthNewPublicKeyGenerator = eth.BTCECDSAPublic(bitInterface.SenderPrivateKey)

	// Convert Public key
	bitInterface.EthNewPublic = eth.BTCCryptoToKey(bitInterface.EthNewPublicKeyGenerator)

	bitInterface.EthAddress = eth.BTCKeyToAddress(bitInterface.EthNewPublic)

	// nonce pending
	bitInterface.EthNonceAtStatus, err = eth.BTCNoncePendingStatus(bitInterface.EthAddress)
	if err != nil {

		log.Fatal("[Fail] Current Pending Nonce Status  ", err)
		return structs.BitsBlocks{}, err
	}

	return bitInterface, err
	// charge, err := StringToInt(blockchains.Amount)
	// if err != nil {
	// 	log.Fatal("[Fail] charge must be String  ", err)
	// 	response := structs.Response{}
	// 	temp := server(w, r)
	// 	_ = response.ClientRequestHandle(true, "Sorry ! Compuatation Error   ", "/transact", w, r)
	// 	response.ClientLogs()
	// 	err := response.Run(temp)
	// 	if err != nil {
	// 		log.Println("[Error]: checks logs...", err)
	// 		return
	// 	}

	// }

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

		// btx, err := send()
		// if err != nil {
		// 	return
		// }

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
	// acc := structs.Static{}
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		// acc.Eth = ethAddrressGenerate
		// acc.Balance = GetBalance(&acc)
		// if acc.Balance == nil {
		// 	log.Fatal("[Fail] Connection Reject ", acc.Balance)
		// 	response := structs.Response{}
		// 	temp := server(w, r)
		// 	_ = response.ClientRequestHandle(true, "Sorry ! Connectivity Issue   ", "/transact", w, r)
		// 	response.ClientLogs()
		// 	err := response.Run(temp)
		// 	if err != nil {
		// 		log.Println("[Error]: checks logs...", err)
		// 	}
		// }
		// fmt.Println("Details:", acc)
		temp.Execute(w, "Transaction")
	} else {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		r.ParseForm()
		if r.FormValue("method") != " " {
			cart.PlaceItemsInCart(r.FormValue("price"), r.FormValue("typeClass"), r.FormValue("method"), r.FormValue("describe"))
		}
		cart.PlaceItemsInCart(r.FormValue("price"), r.FormValue("typeClass"), r.FormValue("method1"), r.FormValue("describe"))
		log.Println("Cart-info:", cart)
		temp.Execute(w, "Transacts")
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
			blockchains.SenderBatchID = acc.EthAddress
			log.Println("Your Wallet:", acc)

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
			// response.WriteHeader(http.StatusOK)
			// request.Method = "GET"
			// createWallet(response, request)

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

// func GetBalance(account *structs.Static) *big.Int {

// 	wallet := common.HexToAddress(account.Eth)
// 	balnce, err := clientInstance.BalanceAt(context.Background(), wallet, nil)
// 	if err != nil {
// 		log.Fatalln("[Fail] Balance reading issue/ connectivity issue")
// 		return nil
// 	}
// 	account.Balance = balnce
// 	return account.Balance
// }

// func ReadBalanceFromBlock(acc *structs.Block) *big.Int {
// 	wallet := common.HexToAddress(acc.TxRec)
// 	balnce, err := clientInstance.BalanceAt(context.Background(), wallet, nil)
// 	if err != nil {
// 		log.Fatalln("[Fail] Connectivity issue", err)
// 		return nil
// 	}
// 	acc.Balance = balnce
// 	return acc.Balance

// }

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

func UploadFiles(w http.ResponseWriter, r *http.Request) (string, error) {
	// println("request body", r.Body)
	r.ParseMultipartForm(10 << 50)

	var upldFile *os.File = nil
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
			return "", err
		}
		return "", err
	}
	defer file.Close()
	if handler.Size >= (500000 * 1024) {
		log.Fatalln((500000 * 1024) - handler.Size)
		return "", err
	}
	fmt.Println("File name:"+handler.Filename, "Size:", handler.Size)
	if _, err := os.Stat(handler.Filename); os.IsExist(err) {
		log.Fatal("[Fail] Already have  this file ", err)
		response := structs.Response{
			Flag:    false,
			Message: "",
			Links:   "",
		}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! Two files doesn't have same name  ", "/dashboard", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return "", err
		}
		return "", err

	}
	// err = os.Mkdir("appdir/", 0755)
	// if err != nil {
	// 	log.Fatalln("[Fail] Directory ", err)
	// 	return upldFile, err
	// }
	path, err := os.Stat("seqDir/")
	if err != nil {
		log.Fatalln("[Error] In directory", err)
		return "", err
	}
	if !path.IsDir() {
		log.Fatalln("[Error] Reading File", err)
		return "", err
	}

	path, err = os.Stat("seqDir/")
	if err != nil {
		log.Fatalln("[Error] In directory", err)
		return "", err
	}
	if !path.IsDir() {
		log.Fatalln("[Error] Reading File", err)
		return "", err
	}
	// upload file by user...
	upldFile, err = ioutil.TempFile(filepath.Dir("seqDir/"), "seqFile-"+"*.txt")
	if err != nil {
		log.Fatal("[Fail] Temporary File ", err)
		response := structs.Response{
			Flag:    false,
			Message: "",
			Links:   "",
		}
		temp := server(w, r)
		_ = response.ClientRequestHandle(true, "Sorry ! Computation Error   ", "/dashboard", w, r)
		response.ClientLogs()
		err := response.Run(temp)
		if err != nil {
			log.Println("[Error]: checks logs...", err)
			return "", err
		}
		return "", err

	}
	defer upldFile.Close()
	_, err = upldFile.Stat()
	if err != nil {
		log.Fatalln("[Fail] File Stats", err)
		return "", err
	}

	openReadFile = upldFile.Name()
	log.Println(openReadFile)

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
			return "", err
		}
	}
	n, err := upldFile.Write(bytesFile)
	if err != nil {
		return "", err
	}
	log.Println("[Result] = File added on server", upldFile.Name(), "Size:", n)
	return openReadFile, nil

}

func sequenceFile(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error) {

	seq, err := ReadSequence(userFile.Name())
	if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}

	var gen []string
	for _, v := range seq {
		space := DoAscii(v)
		if space == "" {
			gen = append(gen, "")
		}
		gen = append(gen, space)
	}

	pathogen, err := ReadSequence(serverFile.Name())
	if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}

	var genV []string
	for _, v := range pathogen {
		space := DoAscii(v)
		if space == "" {
			genV = append(genV, "")
		}
		genV = append(genV, space)
	}

	return gen, genV, nil

}

func DoAscii(seq byte) string {

	if seq >= 65 && seq < 91 {
		return string(alphabet.Letter(seq))
	}
	return string(alphabet.Letter(seq))
}

func blockSession(id int) *sessions.CookieStore {

	return sessions.NewCookieStore([]byte(strconv.Itoa(id)))
}

func choosePattern(w http.ResponseWriter, r *http.Request, fname, choose string, file *os.File) error {

	i, err := strconv.Atoi(choose)
	if err != nil {
		log.Fatalln("[Fail] Sorry there is some issue report!", err)
		return err
	}
	if (i > 0 && i < 6) && (fname != " ") {
		svrFile := FileReadFromDisk(w, r, fname)
		Usr, Virus, err := sequenceFile(file, svrFile)
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

func fileGet(path, filename string) (*os.File, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fileinfo.IsDir() {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return file, nil
}

// func openBrowser(url string) error {
// 	var err error
// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd := exec.Command("/usr/bin/sensible-browser", url)
// 		err = cmd.Start()
// 		if err != nil {
// 			log.Fatalln("Error", err)
// 		}
// 		err = cmd.Run()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	default:
// 		err = fmt.Errorf("Operation fail")
// 	}
// 	return err

// }
