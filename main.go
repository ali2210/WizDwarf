package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"cloud.google.com/go/firestore"
	// firebase "firebase.google.com/go"
	"github.com/ali2210/wizdwarf/db"
	."github.com/ali2210/wizdwarf/piplines"
	CloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
	DBModel "github.com/ali2210/wizdwarf/db/model"
	"github.com/ali2210/wizdwarf/structs"
	bio "github.com/ali2210/wizdwarf/structs/bioinformatics"
	info "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
	Shop "github.com/ali2210/wizdwarf/structs/cart"
	coin "github.com/ali2210/wizdwarf/structs/coinbaseApi"
	weather "github.com/ali2210/wizdwarf/structs/openweather"
	"github.com/ali2210/wizdwarf/structs/paypal/handler"
	wizSdk "github.com/ali2210/wizdwarf/structs/transaction"
	"github.com/ali2210/wizdwarf/structs/users"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	pay "github.com/logpacker/PayPal-Go-SDK"
	"golang.org/x/crypto/sha3"
	"golang.org/x/net/context"
)

// Variables

var (
	emailexp          string                    = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp           string                    = "([A-Z][a-z]*[0-9])*"
	addressexp        string                    = "(^0x[0-9a-fA-F]{40}$)"
	AppName           *firestore.Client         = SetDBClientRef() 
	Cloud             users.DBFirestore         = SetDBCollect()
	digitalCode       users.CreditCardInfo      = users.NewClient()
	vault             DBModel.Private           = DBModel.New()
	ledger            db.PublicLedger           = db.NewCollectionInstance()
	paypalMini        handler.PaypalClientLevel = handler.PaypalClientGo()
	userSessions      *sessions.CookieStore     = nil //user level
	clientInstance    *ethclient.Client         = nil
	ledgerPubcKeys    string                    = ""
	ledgerBits        string                    = ""
	// Firestore_Rf      string                    = ""
	openReadFile      string                    = ""
	publicAddress     string                    = ""
	edit              bio.LevenTable            = SetEditParameters()
	algo              info.Levenshtein
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
		SenderBatchID:            "",
		SenderPrivateKey:         &ecdsa.PrivateKey{},
		EthBlockHeader:           "",
		EthNewPublicKeyGenerator: eth,
		EthNewPublic:             &ecdsa.PublicKey{},
		EthAddress:               [20]byte{},
		EthNonceAtStatus:         0,
		EthGasUnits:              &big.Int{},
		EthReciptAddress:         [20]byte{},
	}
	// js gopher.EmptyDomObject = gopher.EmptyDomObject{}
	coinbaseClient coin.Permission      = coin.Permission{}
	staticData     coin.StaticWallet    = coin.StaticWallet{}
	transactWeb    structs.ParserObject = structs.ParserObject{}
	profiler 	   *users.Visitors = &users.Visitors{}
	// Https_port	   string = os.Getenv("Http_Port")
)

// Constants

const (
	//ProjectID      string = "htickets-cb4d0"
	mainNet        string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	rinkebyClient  string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	geocodeAPI     string = "7efdb33c59a74e09352479b21657aee8"
)

func main() {

	// Server
	log.Println("[OK] Wiz-Dwarfs starting")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	wizDir := os.Getenv("WIZ_VOLUME_DIR")
	if wizDir == "" {
		log.Fatalln("Make sure volume mount", wizDir)
		return
	}

	if host == "" {

		// env port setting

		if port == " "{
			log.Fatalln("[Fail] No Application port allocated", port)
		} else {
			if port != "5000" && host == "wizdwarfs" {
				// any Listening PORT {heroku}
				log.Println("[Open] Application Port", port)
				log.Println("[Open] Application host", host)
				
			} else {
				// specfic port allocated {docker}
				port = "5000"
				host = "wizdwarfs"
				log.Println("[New] Application Default port", port)
				log.Println("[Host] Explicit Host ", host)
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

	})
	
	routing.HandleFunc("/home", home)
	routing.HandleFunc("/signup", newUser)
	routing.HandleFunc("/login", existing)
	routing.HandleFunc("/dashboard", dashboard)
	routing.HandleFunc("/dashbaord/setting", setting)
	routing.HandleFunc("/dashboard/profile", profile)
	routing.HandleFunc("/dashbaord/setting/about", aboutMe)
	routing.HandleFunc("/dashboard/setting/pay/credit/add", credit)
	routing.HandleFunc("/dashbaord/setting/pay/credit/delete", deleteCard)
	routing.HandleFunc("/logout", logout)
	routing.HandleFunc("/createwallet", createWallet)
	routing.HandleFunc("/terms", terms)
	routing.HandleFunc("/open", wallet)
	routing.HandleFunc("/transact", transacts)
	routing.HandleFunc("/treasure", treasure)
	routing.HandleFunc("/phenylalanine", phenylalanine)
	routing.HandleFunc("/leucine", leucine)
	routing.HandleFunc("/isoleucine", isoleucine)
	routing.HandleFunc("/methionine", methionine)
	routing.HandleFunc("/valine", valine)
	routing.HandleFunc("/serine", serine)
	routing.HandleFunc("/proline",proline)
	routing.HandleFunc("/threonine", threonine)
	routing.HandleFunc("/alanine", alanine)
	routing.HandleFunc("/tyrosine", tyrosine)
	routing.HandleFunc("/histidine", histidine)
	routing.HandleFunc("/glutamine", glutamine)
	routing.HandleFunc("/asparagine", asparagine)
	routing.HandleFunc("/lysine", lysine)
	routing.HandleFunc("/aspartic", aspartic)
	routing.HandleFunc("/glutamic", glutamic)
	routing.HandleFunc("/cysteine", cysteine)
	routing.HandleFunc("/tryptophan", tryptophan)
	routing.HandleFunc("/arginine", arginine)
	routing.HandleFunc("/glycine", glycine)
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


	// Static Files
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	routing.PathPrefix("/images/").Handler(images)
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
	routing.PathPrefix("/css/").Handler(css)
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	routing.PathPrefix("/js/").Handler(js)


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

		accountNum := r.FormValue("account")
		

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

		if accountNum != "" && Card_Verification(accountNum, ret.Number) {
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

	// temp := template.Must(template.ParseFiles("about.html"))

	// if r.Method == "GET" {
	// 	log.Println("[Accept]", r.URL.Path)

		

	// 	userProfile, err := Cloud.GetProfile(AppName, accountID, accountVisitEmail)
	// 	if err != nil && userProfile != nil {
	// 		log.Fatal("[Fail] No info  ", err)
	// 		return
	// 	}
		

	// 	key, address := vault.GetCryptoDB(publicAddress)
	// 	access := DBModel.CredentialsPrivate{
	// 		PublicAddress: address,
	// 		PrvteKey:      key,
	// 	}

	// 	log.Println("Ledger Info:", access)

	// 	client, err := paypalMini.NewClient()
	// 	if err != nil {
	// 		log.Fatalln("[Fail] Client Operation:", err)
	// 		return
	// 	}

	// 	_, err = paypalMini.Token(client)
	// 	if err != nil {
	// 		log.Fatalln("[Fail] Token Operation:", err)
	// 		return
	// 	}

	// 	ret, err := paypalMini.RetrieveCreditCardInfo(digitalCode.GetAuthorizeStoreID(), client)
	// 	if err != nil {
	// 		log.Fatalln("[Fail] Retrieve Card info Operation:", err)
	// 		return
	// 	}

	// 	myProfile := users.DigialProfile{
	// 		Public:      access.PublicAddress,
	// 		Private:     access.PrvteKey,
	// 		Name:        userProfile.FirstName,
	// 		FName:       userProfile.LastName,
	// 		Email:       userProfile.Email,
	// 		Address:     userProfile.HouseAddress1,
	// 		LAddress:    userProfile.HouseAddress2,
	// 		City:        userProfile.City,
	// 		Zip:         userProfile.Zip,
	// 		Country:     userProfile.Country,
	// 		Phone:       userProfile.PhoneNo,
	// 		Twitter:     userProfile.Twitter,
	// 		Number:      ret.Number,
	// 		ExpireMonth: ret.ExpireMonth,
	// 		ExpireYear:  ret.ExpireYear,
	// 		Type:        ret.Type,
	// 	}

	// 	log.Println("[Accept] Profile", myProfile)
	// 	temp.Execute(w, myProfile)

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
	var member users.Visitors
	visit ,err := Cloud.GetDocumentById(AppName, *profiler)
	if err != nil {
		log.Printf("Database query failed: %v", err.Error())
		return
	}

	data , err := json.Marshal(visit); if err != nil {
		log.Printf("json marshal: %v", err.Error())
		return
	}

	err = json.Unmarshal(data, &member); if err != nil {
		log.Printf("json unmarshal: %v", err.Error())
	}

	fmt.Println("Member :", member)

	if r.Method == "GET"{
		log.Println("Method:", r.Method)
		log.Println("URL:", r.URL.Path)
		temp.Execute(w, member)
	}else{
		
				
		// user add profile picture resolution must be less 2kb
		Pictures_Stream(r)
		
		// update users information 
		user := users.Visitors{Name: r.FormValue("name"), LastName: r.FormValue("lastname"),
			 Address: r.FormValue("address"), Appartment: r.FormValue("apartment"),  
			 Zip: r.FormValue("zip"), City: r.FormValue("city"), Country: r.FormValue("country"),
			 Twitter : r.FormValue("tweet"), PhoneNo : r.FormValue("phone")}
		fmt.Println("edit profile:", user)
		
		
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

	// temp := template.Must(template.ParseFiles("payee.html"))
	// var check users.Analysis = users.Analysis{}
	// user, err := Cloud.GetProfile(AppName, accountID, accountVisitEmail)
	// if err != nil {
	// 	log.Fatalln("[Fail] Operation..", err)
	// 	return
	// }

	// key, address := vault.GetCryptoDB(publicAddress)
	// credentials := DBModel.CredentialsPrivate{
	// 	PublicAddress: address,
	// 	PrvteKey:      key,
	// }

	// client, err := paypalMini.NewClient()
	// if err != nil {
	// 	log.Fatalln("[Fail] Client Operation:", err)
	// 	return
	// }

	// _, err = paypalMini.Token(client)
	// if err != nil {
	// 	log.Fatalln("[Fail] Token Operation:", err)
	// 	return
	// }

	// ret, err := paypalMini.RetrieveCreditCardInfo(digitalCode.GetAuthorizeStoreID(), client)
	// if err != nil {
	// 	log.Fatalln("[Fail] Retrieve Card info Operation:", err)
	// 	return
	// }

	// order := cart.GetItemsFromCart()
	// batchID := order.TypeofService + user.ID
	// resp, err := paypalMini.GetPayout(batchID, client)
	// if err != nil {
	// 	return
	// }

	// amountStr, err := check.MarshalJSONAmount(resp)
	// if err != nil {
	// 	return
	// }
	// feesStr, err := check.MarshalJSONFees(resp)
	// if err != nil {
	// 	return
	// }
	// amountValue := check.Encode(amountStr)
	// feesValue := check.Encode(feesStr)
	// amount, err := check.CalculateNum(amountValue)
	// if err != nil {
	// 	return
	// }
	// fees, err := check.CalculateNum(feesValue)
	// if err != nil {
	// 	return
	// }
	// total := check.CalculateTotalBalance(amount, fees)
	// sTotal := fmt.Sprintf("%f", total)
	// // blockchains.SenderBatchID = ethAddrressGenerate
	// if blockchains.SenderBatchID == " " {
	// 	log.Fatal("Public Address:", err)
	// 	return
	// }
	// err = userCredit()
	// if err != nil {
	// 	fmt.Println("Error:")
	// }
	// balance.SetTransactionWiz(user.FirstName, fmt.Sprintf("%v", blockchains.Balance), sTotal, sTotal, credentials.PublicAddress)

	// if r.Method == "GET" {
	// 	log.Println("[Accept] Path :", r.URL.Path)
	// 	log.Println("Method :", r.Method)
	// 	temp.Execute(w, balance.GetTransactionWiz())
	// } else {
	// 	log.Println("[Accept] Path :", r.URL.Path)
	// 	log.Println("Method :", r.Method)
	// 	res, err := paypalMini.PaypalPayout(ret.ID, batchID, accountVisitEmail, order.Price, client)
	// 	if err != nil {
	// 		return
	// 	}

	// 	amountStr, err := check.MarshalJSONAmount(res)
	// 	if err != nil {
	// 		return
	// 	}
	// 	feesStr, err := check.MarshalJSONFees(res)
	// 	if err != nil {
	// 		return
	// 	}
	// 	amountValue := check.Encode(amountStr)
	// 	feesValue := check.Encode(feesStr)
	// 	amount, err := check.CalculateNum(amountValue)
	// 	if err != nil {
	// 		return
	// 	}
	// 	fees, err := check.CalculateNum(feesValue)
	// 	if err != nil {
	// 		return
	// 	}
	// 	total := check.CalculateTotalBalance(amount, fees)
	// 	sTotal := fmt.Sprintf("%f", total)
	// 	balance.SetTransactionWiz(user.FirstName, fmt.Sprintf("%v", blockchains.Balance), sTotal, sTotal, credentials.PublicAddress)
	// 	temp.Execute(w, balance.GetTransactionWiz())
	// }
}

func treasure(w http.ResponseWriter, r *http.Request) {
	webpage := template.Must(template.ParseFiles("treasure.html"))
	// userProfile, err := Cloud.GetDocumentById(AppName, *profiler)
	// if err != nil && userProfile != nil {
	// 	log.Fatal("[Fail] No info  ", err)
	// 	return
	// }

	// query_json, err := json.Marshal(userProfile); if err != nil{
	// 	log.Fatal("query return un handle data  ", err.Error())
	// 	return		
	// } 
	// err = json.Unmarshal(query_json, &profiler); if err != nil{
	// 	log.Fatal("query return un structure data", err.Error())
	// 	return
	// }
	log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
	fmt.Println("@Percentage:", visualizeReport.Percentage)
	fmt.Println("@prob parameter:", algo.GetProbParameter())
	algo.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" && algo.GetProbParameter() != -1.0{
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)
		webpage.Execute(w, visualizeReport)
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
	

}

func tKernel(w http.ResponseWriter, r *http.Request) {
	// webpage := template.Must(template.ParseFiles(""))
	// var check users.Analysis = users.Analysis{}
	// user, err := Cloud.GetProfile(AppName, accountID, accountVisitEmail)
	// if err != nil {
	// 	log.Fatalln("[Fail] Operation..", err)
	// 	return
	// }

	// key, address := vault.GetCryptoDB(publicAddress)
	// credentials := DBModel.CredentialsPrivate{
	// 	PublicAddress: address,
	// 	PrvteKey:      key,
	// }

	// client, err := paypalMini.NewClient()
	// if err != nil {
	// 	log.Fatalln("[Fail] Client Operation:", err)
	// 	return
	// }

	// _, err = paypalMini.Token(client)
	// if err != nil {
	// 	log.Fatalln("[Fail] Token Operation:", err)
	// 	return
	// }

	// btx, err := send()
	// if err != nil {
	// 	return
	// }
	// log.Println("Block Header:", btx)

	
	// blockchains.Nonce = bitInterface.EthNonceAtStatus
	// blockchains.GasLimit = uint64(21000)

	// bitInterface.EthGasUnits, err = eth.BTCGasConsumerPrice()
	// if err != nil {
	// 	return
	// }

	// order := cart.GetItemsFromCart()
	// batchID := order.TypeofService + user.ID
	// resp, err := paypalMini.GetPayout(batchID, client)
	// if err != nil {
	// 	return
	// }

	// amountStr, err := check.MarshalJSONAmount(resp)
	// if err != nil {
	// 	return
	// }
	// feesStr, err := check.MarshalJSONFees(resp)
	// if err != nil {
	// 	return
	// }
	// amountValue := check.Encode(amountStr)
	// feesValue := check.Encode(feesStr)
	// amount, err := check.CalculateNum(amountValue)
	// if err != nil {
	// 	return
	// }
	// fees, err := check.CalculateNum(feesValue)
	// if err != nil {
	// 	return
	// }
	// total := check.CalculateTotalBalance(amount, fees)
	// sTotal := fmt.Sprintf("%f", total)
	// if blockchains.SenderBatchID == "" {
	// 	log.Fatalln("Public Address", err)
	// 	return
	// }
	// // blockchains.SenderBatchID = blockchains.SenderBatchID
	// err = userCredit()
	// if err != nil {
	// 	log.Fatal("[Fail] Block Balance should not be zero  ", blockchains.Balance)
	// 	return
	// }
	
	// balance.SetTransactionWiz(user.FirstName, fmt.Sprintf("%v", blockchains.Balance), sTotal, sTotal, credentials.PublicAddress)

	// if r.Method == "GET" {
	// 	log.Println("[Path]:", r.URL.Path)
	// 	log.Println("[Method]:", r.Method)
	// 	webpage.Execute(w, balance.GetTransactionWiz())
	// } else {

	// 	log.Println("[Path]:", r.URL.Path)
	// 	log.Println("[Method]:", r.Method)
	// 	order := cart.GetItemsFromCart()
	// 	value, err := strconv.ParseUint(order.Price, 2, 10)
	// 	if err != nil {
	// 		return
	// 	}
	// 	blockchains.Checkout = value
	// 	blockchains.Amount = new(big.Int).SetUint64(uint64(blockchains.Checkout))
	// 	switch order.TypeofService {
	// 	case "Kernel":
	// 		n, m := int64(2), int64(60)
	// 		automatedNetworkFees(n, m)
	// 	case "Cluster":
	// 		n, m := int64(3), int64(110)
	// 		automatedNetworkFees(n, m)
	// 	case "Multi-Cluster":
	// 		n, m := int64(5), int64(550)
	// 		automatedNetworkFees(n, m)
	// 	default:
	// 		w.WriteHeader(http.StatusPreconditionFailed)
	// 		r.Method = "GET"
	// 		transacts(w, r)
	// 	}
	// }
	// // Send Transaction
	// blockchains.RecieveBatchID = "0x55057eb78fDbF783C961b4AAd6A5f8BC60cab44B"
	// bitInterface.EthReciptAddress = eth.BTCAddressHex(blockchains.RecieveBatchID)

	// // Network ID
	// chainID, err := clientInstance.NetworkID(context.Background())
	// bitInterface.EthTransaction = eth.BTCNewTransactions(blockchains, bitInterface)
	// bitInterface.FingerPrint, err = eth.BTCTransactionSignature(chainID, bitInterface)
	// if err != nil {
	// 	log.Fatal("[Fail] Signed Transaction", err)
	// 	return
	// }

	// // Send Transaction
	// err = eth.TransferBTC(bitInterface.FingerPrint)
	// if err != nil {
	// 	log.Fatalln("[Fail] Transaction", err)
	// 	return
	// }

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
	fmt.Println("Profile:", profiler)
	userProfile, err := Cloud.GetDocumentById(AppName, *profiler)
	if err != nil && userProfile != nil {
		log.Fatal("[Fail] No info  ", err)
		return
	}

	query_json, err := json.Marshal(userProfile); if err != nil{
		log.Fatal("query return un handle data  ", err.Error())
		return		
	} 
	err = json.Unmarshal(query_json, &profiler); if err != nil{
		log.Fatal("query return un structure data", err.Error())
		return
	}

	// algo.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		visualizeReport.Process = 1
		visualizeReport.SeenBy = profiler.Name

		temp.Execute(w, visualizeReport)
	}

}

func userCredit() error {

	var err error
	//blockchains.SenderBatchID = id
	blockchains.Balance, err = genesis.GetLastTransaction(blockchains)
	if err != nil {
		fmt.Println("Error:")
	}

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
		fname, err := Mounted(w, r, openReadFile)

		if err != nil {
			log.Fatalln("[File]:", err)
			return
		}
		
		choose := r.FormValue("choose")
		coordinates := r.FormValue("geo-marker")
		var longitude_parse float64 = 0.0
		var latitude_parse float64 = 0.0
		fmt.Println("@length:", len(coordinates))
		if (len(coordinates) > 0 && len(coordinates) <= 15){
			longitude_parse , err = strconv.ParseFloat(coordinates[:5],10)
			if err != nil {
				log.Fatalln("parse longituide value #0 :", err.Error())
				return 
			}
		
			fmt.Println("@longitude # 0:", longitude_parse)
			latitude_parse , err = strconv.ParseFloat(coordinates[9:],10)
			if err != nil {
				log.Fatalln("parse latitude value # 0 :", err.Error())
				return 
			}

			fmt.Println("@latitude:", latitude_parse)
		}else if (len(coordinates) > 15 && len(coordinates) <= 25) {
			longitude_parse , err := strconv.ParseFloat(coordinates[:8],10)
			if err != nil {
				log.Fatalln("parse longituide value # 1 :", err.Error())
				return 
			}
		
			fmt.Println("@longitude # 1:", longitude_parse)
			latitude_parse , err = strconv.ParseFloat(coordinates[12:],10)
			if err != nil {
				log.Fatalln("parse latitude value # 1 :", err.Error())
				return 
			}

			fmt.Println("@latitude # 1:", latitude_parse)
		}else if(len(coordinates) > 25 && len(coordinates) < 40){
			longitude_parse , err := strconv.ParseFloat(coordinates[:18],10)
			if err != nil {
				log.Fatalln("parse longituide value :", err.Error())
				return 
			}
		
			fmt.Println("@longitude # 2:", longitude_parse)
			latitude_parse , err := strconv.ParseFloat(coordinates[21:],10)
			if err != nil {
				log.Fatalln("parse latitude value # 2 :", err.Error())
				return 
			}

			fmt.Println("@latitude: # 2", latitude_parse)
		}else if len(coordinates) > 15 && len(coordinates) == 23 {
			longitude_parse , err := strconv.ParseFloat(coordinates[:8],10)
			if err != nil {
				log.Fatalln("parse longituide value # 1 :", err.Error())
				return 
			}
		
			fmt.Println("@longitude # 3:", longitude_parse)
			latitude_parse , err := strconv.ParseFloat(coordinates[13:],10)
			if err != nil {
				log.Fatalln("parse latitude value # 1 :", err.Error())
				return 
			}

			fmt.Println("@latitude # 3:", latitude_parse)
		}else {
			log.Fatalln("Empty Field value :", err.Error())
		}
		
		clientapi := weather.NewWeatherClient()
		
		weatherapi ,err := clientapi.OpenWeather(geocodeAPI); if err != nil {
			log.Fatalln("weather api-key :", err.Error())
			return 
		}

		marker_location := clientapi.GetCoordinates(&weather.MyCoordinates{
			Longitude : longitude_parse,
			Latitude : latitude_parse,
		})

		fmt.Println("@marker:", marker_location)

		err = clientapi.UVCoodinates(marker_location,weatherapi); if err != nil {
			log.Fatalln("city weather coordinates:", err.Error())
			return 
		}

		uvinfo, err := clientapi.UVCompleteInfo(weatherapi); if err != nil {
			log.Fatalln("city uv tracks:", err.Error())
			return 
		}

		fmt.Println("@uv:", uvinfo)
		
		visualizeReport.UVinfo = uvinfo
		
		data, err := Open_SFiles("app_data/", fname)
		if err != nil {
			log.Fatalln("[No File]:", err)
			return
		}
		
		log.Println("File Name:", data.Name())
		
		switch choose {
		case "0":
		 	fmt.Fprintf(w, "Please choose any option ...")
		 	log.Fatalln("Choose your option")
			return
		 case "1":
		 	var name string = "Covid-19"
		 	
			 err := Data_Predicition(w, r, name, choose, data,algo)
		 	if err != nil {
		 		return
		 	}
			
			 pattern_analysis := GetBioAlgoParameters()
	 		
			visualizeReport.Percentage = pattern_analysis.Percentage

		 	w.WriteHeader(http.StatusOK)
		// 	// LifeCode = genome
			r.Method = "GET"
		 	visualize(w, r)
		// 	// fmt.Println("Virus:", capsid)

		 case "2":
		 	var name string = "FlaviDengue"
		 	err := Data_Predicition(w, r, name, choose, data, algo)
		 	if err != nil {
		 		return
		 	}
			pattern_analysis := GetBioAlgoParameters()
		 	visualizeReport.Percentage = pattern_analysis.Percentage
		// 	// v :=  infectedUv()
		// 	// v.UVinfo = uvslice
		 	w.WriteHeader(http.StatusOK)

		 	r.Method = "GET"
		 	visualize(w, r)
		// 	// Wallet(w,r)
		// 	// fmt.Println("Virus:", capsid)
		 case "3":
		 	var name string = "KenyaEbola"
		 	err := Data_Predicition(w, r, name, choose, data,algo)
		 	if err != nil {
		 		return
		 	}
		 	pattern_analysis := GetBioAlgoParameters()
		 	visualizeReport.Percentage = pattern_analysis.Percentage
		// 	// v :=  infectedUv()
		// 	// v.UVinfo = uvslice

		 	w.WriteHeader(http.StatusOK)
		 	r.Method = "GET"
		 	visualize(w, r)

		// 	// fmt.Println("Virus:", capsid)
		 case "4":
		 	var name string = "ZikaVirusBrazil"
		 	err := Data_Predicition(w, r, name, choose, data, algo)
		 	if err != nil {
		 		return
		 	}
		 	pattern_analysis := GetBioAlgoParameters()
		 	visualizeReport.Percentage = pattern_analysis.Percentage
		// 	// v :=  infectedUv()
		// 	// openStreet.Country = r.FormValue("country")
		// 	// openStreet.PostalCode = r.FormValue("postal")
		// 	// openStreet.City = r.FormValue("city")
		// 	// openStreet.State = r.FormValue("state")
		// 	// openStreet.StreetAddress = r.FormValue("street")
		// 	// i, err := strconv.Atoi(r.FormValue("route")); if err != nil {
		// 	// 	return
		// 	// }

		// 	//v.UVinfo = uvslice

		 	w.WriteHeader(http.StatusOK)
		 	r.Method = "GET"
		 	visualize(w, r)

		// 	// fmt.Println("Virus:", capsid)
		 case "5":
		 	var name string = "MersSaudiaArabia"
		 	err := Data_Predicition(w, r, name, choose, data,algo)
		 	if err != nil {
		 		return
		 	}
		 	pattern_analysis := GetBioAlgoParameters()
		 	visualizeReport.Percentage = pattern_analysis.Percentage
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
		// 	// fmt.Println("Virus:", capsid)

		 default:
		 	 RouteWebpage.Execute(w, "dashboard")
		 }
	}

}

func send() (structs.BitsBlocks, error) {


	var err error

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
			return
		}

		// private key into bytes
		PrvateKyByte := crypto.FromECDSA(privateKey)

		key := hexutil.Encode(PrvateKyByte)[2:]

		pblicKey := privateKey.Public()

		pbcKey, ok := pblicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("[Fail] Public Key from Private Key  ", err)
			return
		}

		publicBytes := crypto.FromECDSAPub(pbcKey)
		

		PublicKey := crypto.PubkeyToAddress(*pbcKey).Hex()

		acc.PubKey = PublicKey[:8]
		acc.PrvteKey = key[:8]

		// hash to ethereum
		hshCode := sha3.NewLegacyKeccak256()
		hshCode.Write(publicBytes[1:])
		ethereum := hexutil.Encode(hshCode.Sum(nil)[12:])

		acc.EthAddress = ethereum[:8]

		// valid address
		// valid := IsEvm(acc.EthAddress)
		// if valid {

		// 	// smart contract address
		// 	log.Println("[Feature] Smart Address", valid)
		// 	w.WriteHeader(http.StatusForbidden)
		// 	w.Write([]byte("Thank-you for your response! , This feature will added upcoming build... Sorry for inconvenience"))
		// 	return

		// }

		myWallet := CloudWallet.EthereumWalletAcc{}

		signWallet, err := json.Marshal(myWallet)
		if err != nil {
			return

		}

		err = json.Unmarshal(signWallet, &myWallet)
		if err != nil {
			log.Fatal("[Fail] Data JSON FORMAT ERROR ", err)
			return
		}

		
		// ok, ethAdd := Retrieve_Crypto(&acc, ledger)
		// if ok && ethAdd != nil {
		// 	log.Fatal("[Replicate] Already Data exist  ", err)
		// 	return

		// }

		myWallet.Email = acc.Email
		myWallet.Password = acc.Password
		myWallet.EthAddress = acc.EthAddress
		myWallet.Terms = acc.Terms
		myWallet.Allowed = acc.Allowed
		myWallet.PrvteKey = acc.PrvteKey

		// merchant, err := ledger.CreatePublicAddress(&myWallet, AppName)
		// if err != nil {
		// 	log.Fatal("[Fail] Wallet Don't have Public Accessibity  ", err)
		// 	return

		// }

		clientInstance = nil

		// log.Println("[Accept] Welcome ! Your Account Has been created", &merchant)
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		existing(w, r)
	}
}

func transacts(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("transact.html"))
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Transacts")
	} else {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		r.ParseForm()
		doc, err := transactWeb.ReadContent("transact.html")
		if err != nil {
			fmt.Println("Fail:", err)
			return
		}
		fmt.Print("@oaram:", &doc)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Thank-you for your response! , This feature will added upcoming build... Sorry for inconvenience"))
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
			return

		}

		log.Println("[Accept] Firestore Cloud Database Connection secure", client)

		clientInstance = client

		log.Println("[Accept] Your Account Details:", acc, "Client api Reference: ", clientInstance)

		myWallet := CloudWallet.EthereumWalletAcc{}

		signWallet, err := json.Marshal(myWallet)
		if err != nil {
			log.Fatal("[Fail] Data JSON FORMAT ERROR ", err)
			return

		}

		err = json.Unmarshal(signWallet, &myWallet)
		if err != nil {
			log.Fatal("[Fail] JSON DATA RETURN ERROR", err)
			return

		}
		// addr, ok := MyEthAddress(&acc)
		// if !ok {

		// 	log.Fatal("[Fail] No Ethereum Account ", !ok)
		// 	return

		// }
		// if addr != nil {

		// 	acc.EthAddress = addr.EthAddress
		// 	publicAddress = acc.EthAddress

		// 	// Secure Key
		// 	ledgerBits = addr.PrvteKey
		// 	vault.SetCryptoDB(acc.EthAddress, ledgerBits)

			
		// 	blockchains.SenderBatchID = acc.EthAddress
		// 	log.Println("Your Wallet:", acc)


		// 	//dataabse -- Retrieve_Crypto
		// 	secureWallet, ok := FindEthWallet(&acc)
		// 	if !ok && secureWallet != nil {
		// 		log.Fatal("[Fail] No crypto wallet found against your account ", !ok)
		// 		return

		// 	}
		// 	log.Println("[Accept] Your Ethereum Wallet Info:", secureWallet)

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			dashboard(w, r)
		}
	}
//}

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
	user := users.Visitors{
		Name:     "",
		LastName:    "",
		Eve:    false,
		Address:  "",
		Appartment: "",
		Zip:      "",
		City:     "",
		Country:  "",
		Email:    "",
		Password: "",
		PhoneNo: "",
		Twitter : "",
	}

	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Regsiter")
	} else {
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		user.Name = r.FormValue("uname")
		user.LastName = r.FormValue("ufname")
		user.Address = r.FormValue("address")
		user.Appartment = r.FormValue("add")
		user.City = r.FormValue("inputCity")
		user.Country = r.FormValue("co")
		user.Zip = r.FormValue("inputZip")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		if r.FormValue("gender") == "on" {
			user.Eve = true
		} else {
			user.Eve = false
		}

		regex_Email, err := regexp.MatchString(emailexp, user.Email)
		if err != nil {
			log.Fatal("[Fail] Auto email pattern  ", err)
			return
		}

		fmt.Println("Regex:", regex_Email)

		regex_Pass, err := regexp.MatchString(passexp, user.Password)
		if err != nil {
			log.Fatal("[Fail] Password is very week ", err)
			return
		}
		
		fmt.Println("Regex:", regex_Pass)
		
		hash, encrypted := Presence(w, r, regex_Email, regex_Pass, user)
		if !hash {
			log.Fatal("[Fail] Week encryption", hash)
			return

		}

		fmt.Println("Hash:", hash, "encrypted:", encrypted)
		fmt.Println("Profile:", user)
		AddNewProfile(w, r, user, encrypted.Reader)
	}

}

func existing(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("login.html"))
	user := users.Visitors{
		Name:     "",
		LastName:    "",
		Eve:    false,
		Address:  "",
		Appartment: "",
		Zip:      "",
		City:     "",
		Country:  "",
		Email:    "",
		Password: "",
		Id : "",
		PhoneNo : "",
		Twitter : "",
	}
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Login")
	} else {
		// Parse Form
		r.ParseForm()
		fmt.Println("Method:\n", r.Method)
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		// if r.FormValue("check") == "on" {
		// 	user.Remember = true
		// } else {
		// 	user.Remember = false
		// }
		log.Println("Login form data[", user, "]")

		// Valid Data for processing
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(user.Email)
		if !ok {
			log.Fatal("[Fail] Mismatch ", ok)
			
			return

		}

		reg := regexp.MustCompile(passexp)
		okx := reg.MatchString(user.Password)
		if !okx {
			log.Fatal("[Fail] Mismatch password", !okx)
			return
		}
		println("regexp_pass:", okx)

		// Search Data in DB
		data, err := Firebase_Gatekeeper(w, r, user)
		
		if err != nil && data == nil {
			log.Fatal("[Result]: No Match Found  ", err)
			return
		}
			accountID = data.Id
			fmt.Printf("Search Data:%v", data.Id)
			accountVisitEmail = data.Email
			accountKey = data.Password
			profiler = data
			fmt.Println("Profile:", profiler)
			act := structs.RouteParameter{}
			
			if userSessions == nil {
				userSessions = Web_Token(data.Id)

				act.SetContextSession(userSessions, w, r)
				err := act.NewToken()
				if err != nil {
					log.Fatal("[FAIL] No Token generate .. Review logs", err)
					return

				}
			}

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			dashboard(w, r)
	}
}



func logout(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		act := structs.RouteParameter{}

		log.Println("[Access] ", r.URL.Path)
		act.SetContextSession(userSessions, w, r)
		err := act.ExpireToken()
		if err != nil {
			log.Fatal("[Fail] No Token Expire  ", err)			
			return
		}
		existing(w, r)
	}
}

func phenylalanine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("Phenylalanine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "phenylalanine")
	}	
}

func leucine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("leucine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "leucine")
	}	
}

func isoleucine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("isoleucine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "isoleucine")
	}
}

func methionine(w http.ResponseWriter,r *http.Request)  {
	webpge := template.Must(template.ParseFiles("methionine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "methionine")
	}
}

func valine(w http.ResponseWriter, r *http.Request) {
	webpge := template.Must(template.ParseFiles("valine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "valine")
	}
}

func serine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("serine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "serine")
	}
}

func proline(w http.ResponseWriter, r *http.Request){
	webpge := template.Must(template.ParseFiles("proline.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "proline")
	}
}

func threonine(w http.ResponseWriter, r *http.Request){
	webpge := template.Must(template.ParseFiles("threonine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "threonine")
	}
}

func alanine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("alanine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "alanine")
	}
}

func tyrosine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("tyrosine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "tyrosine")
	}
}

func histidine(w http.ResponseWriter, r * http.Request)  {
	webpge := template.Must(template.ParseFiles("histidine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "histidine")
	}
}
func glutamine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("glutamine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "glutamine")
	}
}

func asparagine(w http.ResponseWriter, r *http.Request){
	webpge := template.Must(template.ParseFiles("asparagine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "asparagine")
	}
}

func lysine(w http.ResponseWriter, r * http.Request)  {
	webpge := template.Must(template.ParseFiles("lysine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "lysine")
	}
}

func aspartic(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("aspartic.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "aspartic")
	}
}

func glutamic(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("glutamic.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "glutamic")
	}
}

func cysteine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("cysteine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "cysteine")
	}
}

func tryptophan(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("tryptophan.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "tryptophan")
	}
}

func arginine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("arginine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "arginine")
	}
}

func glycine(w http.ResponseWriter, r *http.Request)  {
	webpge := template.Must(template.ParseFiles("glycine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "glycine")
	}
}