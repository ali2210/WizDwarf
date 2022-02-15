/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	binaries "github.com/ali2210/wizdwarf/other/genetic/binary"

	// CloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
	// DBModel "github.com/ali2210/wizdwarf/db/model"
	structs "github.com/ali2210/wizdwarf/other"
	bio "github.com/ali2210/wizdwarf/other/bioinformatics"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/bucket"
	"github.com/ali2210/wizdwarf/other/bucket/proto"
	cryptos "github.com/ali2210/wizdwarf/other/crypto"
	genetics "github.com/ali2210/wizdwarf/other/genetic"
	genome "github.com/ali2210/wizdwarf/other/genetic/binary"
	"github.com/ali2210/wizdwarf/other/jsonpb"
	"github.com/ali2210/wizdwarf/other/jsonpb/jsonledit"
	"github.com/ali2210/wizdwarf/other/proteins"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
	"github.com/ali2210/wizdwarf/piplines"

	// Shop "github.com/ali2210/wizdwarf/other/cart"
	// coin "github.com/ali2210/wizdwarf/other/coinbaseApi"
	weather "github.com/ali2210/wizdwarf/other/openweather"
	"github.com/ali2210/wizdwarf/other/paypal/handler"
	"github.com/pusher/pusher-http-go"


	// wizSdk "github.com/ali2210/wizdwarf/other/transaction"

	"github.com/ali2210/wizdwarf/other/users"
	. "github.com/ali2210/wizdwarf/piplines"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	pay "github.com/logpacker/PayPal-Go-SDK"
)

// Variables

var (
	emailexp    string               = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp     string               = "([A-Z][a-z]*[0-9])*"
	addressexp  string               = "(^0x[0-9a-fA-F]{40}$)"
	AppName     *firestore.Client    = SetDBClientRef()
	Cloud       users.DBFirestore    = SetDBCollect()
	digitalCode users.CreditCardInfo = users.NewClient()
	// vault          DBModel.Private           = DBModel.New()
	// ledger         db.PublicLedger           = db.NewCollectionInstance()
	paypalMini     handler.PaypalClientLevel = handler.PaypalClientGo()
	userSessions   *sessions.CookieStore     = nil //user level
	clientInstance *ethclient.Client         = nil
	ledgerPubcKeys string                    = ""
	ledgerBits     string                    = ""
	// Firestore_Rf      string                    = ""
	openReadFile      string         = ""
	publicAddress     string         = ""
	edit              bio.LevenTable = SetEditParameters()
	algo              info.Levenshtein
	visualizeReport   weather.DataVisualization = weather.DataVisualization{}
	accountID         string                    = " "
	accountKey        string                    = " "
	accountVisitEmail string                    = " "
	signed_msg        string                    = " "
	address_wallet    string                    = " "
	File              string                    = ""
	sumof int64 = 0
	sb []string
	mss []float64
	occ []int64
	count int = 0
	aminochain []*binary.Micromolecule	
	transactWeb structs.ParserObject = structs.ParserObject{}
	profiler    *users.Visitors      = &users.Visitors{}
)

// Constants

const (
	//ProjectID      string = "htickets-cb4d0"
	// mainNet       string = "https://mainnet.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	// rinkebyClient string = "https://rinkeby.infura.io/v3/95d9986e9c8f46c788fba46a2f513e0a"
	GEO_Index_KEY    string = "7efdb33c59a74e09352479b21657aee8"
	APP_CHANNEL_KEY string 				   = "65993b3c66b5317411a5"
	APP_CHANNEL_ID string = "1265511"
	APP_CHANNEL_SCRECT string = "4f8bf3faf121d9c8dadf"
	APP_CHANNEL_CLUSTER_ID string = "mt1"
)

func main() {

	// Server
	
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	wizDir := os.Getenv("WIZ_VOLUME_DIR")
	
	

	// allocate port and host
	if host == "" {

		if port == " " && wizDir == " "{
			
			fmt.Println(" Firebase Configuration complete ... [wel-done]")
			fmt.Println(" Static IP-Address Configuration ... [FAIL]")
			fmt.Println(" Protocol-Buffer v3 Configuration complete ... [wel-done]")
			fmt.Println(" UI Webboard Configuration ... [FAIL]")
			fmt.Println(" Application Configuration ... [FAIL]")
			fmt.Println(" Application Default IP-Address Allocate ... [FAIL]")
			fmt.Println(" Application Persistance Storage Configuration completed ...  [FAIL]")
			log.Fatalln("Make sure volume mount")
			panic("Application fail to started because no port is specified and we do not have writing permission on your disk ")
		} 
	} else {
		fmt.Println(" Firebase Configuration complete ... [wel-done]")
		fmt.Println(" Static IP-Address Configuration complete ... [wel-done]")
		fmt.Println(" Protocol-Buffer v3 Configuration complete ... [wel-done]")
		fmt.Println(" Channel communication Configuration complete ... [wel-done]")
		fmt.Println(" Data Events are encrypted. All the user choices secure with mathematical functions  ... ")
		fmt.Println(" UI Webboard Configuration complete ... [wel-done]")
		fmt.Println(" Application started ... [wel-done](All process have completed)")
		fmt.Println(" Application Persistance Storage Configuration completed ...  [wel-done]")
		
		if (port != "127.0.0.1:5000"){
			fmt.Println(" Application IP Address generated. The webboard started with this address at ", port)
		}
		fmt.Println(" Application Default IP-Address Allocate ... Default IP-Address allow user to access UI Webboard on your browser. The webboard started with this address 127.0.0.1:5000/ [wel-done]")
		fmt.Println(" ***************************************************************************************")                                                                                 
  		fmt.Println("  @@@  @@@  @@@ @@@ @@@@@@@@ @@@@@@@  @@@  @@@  @@@  @@@@@@  @@@@@@@  @@@@@@@@  @@@@@@ ") 
		fmt.Println("   @@!  @@!  @@! @@!      @@! @@!  @@@ @@!  @@!  @@! @@!  @@@ @@!  @@@ @@!      !@@     ")
		fmt.Println("   @!!  !!@  @!@ !!@    @!!   @!@  !@! @!!  !!@  @!@ @!@!@!@! @!@!!@!  @!!!:!    !@@!!  ")
		fmt.Println("    !:  !!:  !!  !!:  !!:     !!:  !!!  !:  !!:  !!  !!:  !!! !!: :!!  !!:          !:! ")
		fmt.Println("    ::.:  :::   :   :.::.: : :: :  :    ::.:  :::    :   : :  :   : :  :       ::.: :  ") 
		fmt.Println(" ****************************************************************************************")
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
	routing.HandleFunc("/dashboard/setting/pay/credit/add", credit)
	routing.HandleFunc("/dashbaord/setting/pay/credit/delete", deleteCard)
	routing.HandleFunc("/logout", logout)
	routing.HandleFunc("/feedback", customerViews)
	routing.HandleFunc("/terms", terms)
	// routing.HandleFunc("/open", wallet)
	routing.HandleFunc("/transact", transacts)
	routing.HandleFunc("/treasure", treasure)
	routing.HandleFunc("/phenylalanine", phenylalanine)
	routing.HandleFunc("/leucine", leucine)
	routing.HandleFunc("/isoleucine", isoleucine)
	routing.HandleFunc("/methionine", methionine)
	routing.HandleFunc("/valine", valine)
	routing.HandleFunc("/serine", serine)
	routing.HandleFunc("/proline", proline)
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
	routing.HandleFunc("/stop", stop_codon)
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
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	temp := template.Must(template.ParseFiles("index.html"))

	// route actions
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "home")
	}

}

func newUser(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	// webpage route
	temp := template.Must(template.ParseFiles("register.html"))
	
	// user query interface
	user := users.Visitors{
		Name:       "",
		LastName:   "",
		Eve:        false,
		Address:    "",
		Appartment: "",
		Zip:        "",
		City:       "",
		Country:    "",
		Email:      "",
		Password:   "",
		PhoneNo:    "",
		Twitter:    "",
	}

	// route actions
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Regsiter")
	} else if r.Method == "POST" {
		r.ParseForm()

		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		// html form
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

		// user email and email pattern both are same, means email created according to email rule.
		regex_Email, err := regexp.MatchString(emailexp, user.Email)
		if err != nil {
			log.Fatal("[Fail] Auto email pattern  ", err)
			return
		}

		// fmt.Println("Regex:", regex_Email)
		// user password and password pattern both are same, means password created according to password rule.
		regex_Pass, err := regexp.MatchString(passexp, user.Password)
		if err != nil {
			log.Fatal("[Fail] Password is very week ", err)
			return
		}

		// fmt.Println("Regex:", regex_Pass)

		// encrypted user information
		hash, encrypted := Presence(w, r, regex_Email, regex_Pass, user)
		if !hash {
			log.Fatal("[Fail] Week encryption", hash)
			return

		}

		// fmt.Println("Hash:", hash, "encrypted:", encrypted)
		// fmt.Println("Profile:", user)
		// in case your account have been created ....
		if docs, ok, err := AddNewProfile(w, r, user, encrypted.Reader); ok && err == nil {
			log.Println("account created successfully", docs)
			
			// route return Ok callback
			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			existing(w, r)
		} else {

			log.Println("account failed ")
			
			// route return BAD callback
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func existing(w http.ResponseWriter, r *http.Request) {

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")


	// webpage route
	temp := template.Must(template.ParseFiles("login.html"))
	
	// initalization of user interface
	user := users.Visitors{}
	
	// route actions
	if r.Method == "GET" {
		
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Login")
	} else if r.Method == "POST" {

		// user query parameters
		r.ParseForm()
		fmt.Println("Method:", r.Method)

		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")

		log.Println("Email:", user.Email, "Password:", user.Password)

		// match email patten with email addresss
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(user.Email)
		if !ok {
			log.Fatal("[Fail] Mismatch ", ok)
			return
		}

		// match password pattern against password
		reg := regexp.MustCompile(passexp)
		okx := reg.MatchString(user.Password)
		if !okx {
			log.Fatal("[Fail] Mismatch ", !okx)
			return
		}

		// Search Data in DB
		data, err := Firebase_Gatekeeper(w, r, user)
		if err != nil && data == nil {
			log.Fatal("[Result]: No Match Found  ", err)
			return
		}
		accountID = data.Id

		// fmt.Printf("Search Data:%v", data.Id)

		accountVisitEmail = data.Email
		accountKey = data.Password
		profiler = data

		// fmt.Println("Profile:", profiler)
		
		// account interface 
		act := structs.RouteParameter{}

		// create user session for user 
		if userSessions == nil {
			
			// initialize the web session
			userSessions = Web_Token(data.Id)

			// valid the session data
			act.SetContextSession(userSessions, w, r)
			err := act.NewToken()
			if err != nil {
				log.Fatal("[FAIL] No Token generate .. Review logs", err)
				return
			}
		}

		// route return Ok callback
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		dashboard(w, r)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// short live variable called visitor
	var member users.Visitors

	// page renderer
	temp := template.Must(template.ParseFiles("profile.html"))
	
	// to find app have information
	visit, err := Cloud.GetDocumentById(AppName, *profiler)
	if err != nil {
		log.Printf("Database query failed: %v", err.Error())
		return
	}

	// encode information with json schema
	data, err := json.Marshal(visit)
	if err != nil {
		log.Printf("json marshal: %v", err.Error())
		return
	}

	// proper encoding over data streams
	err = json.Unmarshal(data, &member)
	if err != nil {
		log.Printf("json unmarshal: %v", err.Error())
	}

	
	// web request "get"
	if r.Method == "GET" {
		log.Println("Method:", r.Method)
		log.Println("URL:", r.URL.Path)
		temp.Execute(w, member)
	} else if r.Method == "POST" {

		// user add profile picture resolution must be less 2kb
		Pictures_Stream(r, member.Id)

		// update users information
		user := users.Visitors{}

		// users information hold 
		user.Id = member.Id
		user.Password = member.Password
		user.Email = member.Email

		
		if strings.Contains(r.FormValue("name"), " ") {
			user.Name = member.Name
		} else {
			user.Name = r.FormValue("name")
		}

		if strings.Contains(r.FormValue("lastname"), " ") {
			user.LastName = member.LastName
		} else {
			user.LastName = r.FormValue("lastname")
		}

		if strings.Contains(r.FormValue("address"), " ") {
			user.Address = member.Address
		} else {
			user.Address = r.FormValue("address")
		}

		if strings.Contains(r.FormValue("appartment"), " ") {
			user.Appartment = member.Appartment
		} else {
			user.Appartment = r.FormValue("appartment")
		}
		if strings.Contains(r.FormValue("zip"), " ") {
			user.Zip = member.Zip
		} else {
			user.Zip = r.FormValue("zip")
		}

		if strings.Contains(r.FormValue("city"), " ") {
			user.City = member.City
		} else {
			user.City = r.FormValue("city")
		}

		if strings.Contains(r.FormValue("country"), " ") {
			user.Country = member.Country
		} else {
			user.Country = r.FormValue("country")
		}

		if strings.Contains(r.FormValue("tweet"), " ") {
			user.Twitter = member.Twitter
		} else {
			user.Twitter = r.FormValue("tweet")
		}

		if strings.Contains(r.FormValue("phone"), " ") {
			user.PhoneNo = member.PhoneNo
		} else {
			user.PhoneNo = r.FormValue("phone")
		}

		// user information completed and store in database
		if status_profile := UpdateProfileInfo(&user); status_profile {
			
			w.WriteHeader(http.StatusOK)
			return
		}

	}

}


func success(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("modal-success.html"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
                    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
                    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	if r.Method == "GET" {
		log.Println("[Accept]", r.URL.Path)
		temp.Execute(w, "success")
	}
}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("delete.html"))

	w.Header().Set("Access-Control-Allow-Origin", "*")
                    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
                    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

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



func setting(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("settings.html"))

	if r.Method == "GET" {
		log.Println("[Accept]", r.URL.Path)
		temp.Execute(w, "setting")
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
			ID:                 accountID,
			PayerID:            "",
			ExternalCustomerID: "",
			Number:             r.FormValue("cardNo"),
			Type:               r.FormValue("cardtype"),
			ExpireMonth:        month,
			ExpireYear:         year,
			CVV2:               r.FormValue("cvv"),
			FirstName:          r.FormValue("fholder"),
			LastName:           r.FormValue("surename"),
			BillingAddress:     &pay.Address{},
			State:              "",
			ValidUntil:         "",
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



func treasure(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	webpage := template.Must(template.ParseFiles("treasure.html"))
	
	
	// analysis data results
	algo.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" && algo.GetProbParameter() < 101{
		
		log.Println("[Path]:", r.URL.Path)
		log.Println("[Method]:", r.Method)

		// client param
		pusherClient := pusher.Client{
					AppID : APP_CHANNEL_ID,
					Key : APP_CHANNEL_KEY,
					Secret : APP_CHANNEL_SCRECT,
					Cluster : APP_CHANNEL_CLUSTER_ID,
					Secure : true,
		}

		chain := piplines.DashboardAnalytics(aminochain, sumof)

		// log.Println("chain: ", chain)

		analytics_amino := make([]map[string]interface{}, len(chain))

		for i := range chain {
			
			if !reflect.DeepEqual(reflect.ValueOf(chain[i]).Interface(),chain[0]){
					analytics_amino = append(analytics_amino,map[string]interface{}{
						"key" : piplines.Mapper(chain[i],"Symbol").(string),
						"values" : piplines.Mapper(chain[i],"Occurance").(int64),
					})
			}
			
		}

		//log.Println("whole data:", analytics_amino[len(analytics_amino)-count:])
		data := analytics_amino[len(analytics_amino)-count:]	
		err  := pusherClient.TriggerMulti([]string{"protein"},"molecule", data)
		if err != nil {
			log.Println("Error trigger event:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//w.WriteHeader(http.StatusOK)

		webpage.Execute(w, visualizeReport)
	}
}


func visualize(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	temp := template.Must(template.ParseFiles("visualize.html"))

	// log.Println("Report percentage", visualizeReport.Percentage)
	// log.Println("Report uv ", visualizeReport.UVinfo)
	// fmt.Println("Profile:", profiler)

	userProfile, err := Cloud.GetDocumentById(AppName, *profiler)
	if err != nil && userProfile != nil {
		log.Fatal("[Fail] No info  ", err)
		return
	}

	query_json, err := json.Marshal(userProfile)
	if err != nil {
		log.Fatal("query return un handle data  ", err.Error())
		return
	}

	err = json.Unmarshal(query_json, &profiler)
	if err != nil {
		log.Fatal("query return un structure data", err.Error())
		return
	}

	// algo.SetProbParameter(visualizeReport.Percentage)
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		// visualizeReport.Process = 1
		// visualizeReport.SeenBy = profiler.Name

		// firestore credentials
		genetics.Client = piplines.Firestore_Reference()

		// generate ed25519 key
		cdr, _ := cryptos.PKK25519(profiler.Id)
		genetics.Pkk = fmt.Sprintf("%x", cdr)

		// genetics object
		rece_gen := genetics.New()

		life := genome.Lifecode{}

		// genetics data string
		life.Genes = strings.Join(piplines.GetGenes(), "")
		life.Pkk = genetics.Pkk

		// create trust object ... trust verified whom that content .
		ok, err, key := piplines.TrustRequest(life.Pkk, address_wallet, signed_msg)
		if !ok && err != nil {
			log.Printf(" cryptographic trust failed %v:", err.Error())
			return
		}

		// genetics database
		status := rece_gen.AddCode(context.Background(), &life)
		if status.ErrorCode == binaries.Errors_Error {
			log.Printf(" bad request: %v", status)
			return
		}

		// get all proteins symbols
		listProteinsName := piplines.Active_Proteins(life.Genes)

		// get all amino table
		listProteins := piplines.AminoChains(life.Genes)
		ribbon := make(map[string]map[string]proteins.Aminochain)

		// read map values
		iterate := reflect.ValueOf(listProteinsName).MapRange()

		// create new marcomolecules which hold molecule state for a while
		// chains := binary.Micromolecule_List{}
		aminochain = make([]*binary.Micromolecule, len(life.Genes))
		//chains.Peplide = make([]*binary.Micromolecule, len(life.Genes))

		// iterate over map values
		for iterate.Next() {

			// store map value in other map
			ribbon[listProteinsName[iterate.Value().String()]] = listProteins

			// get polypeptide information in structured data
			extraction := piplines.Genome_Extract(ribbon, listProteinsName, iterate.Value().String())

			// if the information return void space or empty field then discard , otherwise hold that state
			if piplines.GetMoleculesState(extraction) && piplines.GetCompositionState(extraction) {
				aminochain = append(aminochain, extraction)
			}

		}

		var Ckk string = ""

		// compare provided key hold key state and they key is also generated by machine
		if !reflect.DeepEqual(key.Public(), " ") && len(key.Seed()) == 32 {
			Ckk = fmt.Sprintf("%x", key.Public())
		}

		s := make([]string, len(life.Genes))
		for j := range aminochain {
			if str := piplines.Make_string(aminochain[j]); str != " " {
				s = append(s, str)
			}
		}

		
		sb = make([]string, len(aminochain))
		occ = make([]int64, len(aminochain))
		mss = make([]float64, len(aminochain))

		// iterate_over aminochain
		for i := range aminochain {

			// check aminochain have execpted type
			if piplines.Molecular(aminochain[i]) && s[i] != " " {
				abundance, syms := piplines.Abundance(aminochain[i], strings.Join(s, ""), i)
				if reflect.DeepEqual(aminochain[i].Symbol, syms) {
					aminochain[i].Abundance = int64(abundance)
					count += 1
				}

				
				
				sb = append(sb, piplines.Symbol(aminochain[i]))
				mss = append(mss, piplines.Mass(aminochain[i]))
				occ = append(occ, piplines.Occurance(aminochain[i]))
				

				// marshal the protos message
				data, err := jsonpb.ProtojsonMarshaler(aminochain[i])
				if err != nil {
					log.Printf(" Error marshalling protos %v", err.Error())
					return
				}

				// unmarhal the protos message
				err = jsonpb.ProtojsonUnmarshaler(data, aminochain[i])
				if err != nil {
					log.Printf(" Error Un-marshalling protos %v", err.Error())
					return
				}

				// which protein exist in abudance
				sumof = sumof + proteins.Total_chain_filter(aminochain[i].Abundance)
			}
		}

		// This param either contribute in file processing or file content that want to be store in batch mode
		jsonFile := &jsonledit.FileDescriptor{}
		jsonFile.Types = ".json"
		jsonFile.Names = piplines.Generator()
		jsonFile.Molecule = aminochain
		jsonFile.Occurance = sumof

		// create a new .json file with these parameters and write data stream in file
		proteins.CreateNewJSONFile(jsonFile)

		// additional param
		context := context.Background()
		client := bucket.New_Client(&context)
		bucket.Client = piplines.Firestore_Reference()

		// each transaction param before storing in physical database.
		bucket.Key = profiler.Id
		bucket.Composite = Ckk

		// creating decentalize & dynamic links of the content
		iobject := client.New_Bucket(&proto.Object{Name: jsonFile.Names, Types: jsonFile.Types, Content: jsonFile.Molecule})
		if iobject.Istatus == proto.Object_Status_ERROR {
			log.Printf(" Error in generating link %v", iobject.Istatus)
			return
		}

		File = jsonFile.Names + jsonFile.Types

		
		
		//prev_object := client.Preview(&proto.Query{ByName: file})
		//log.Println("prev_object:", prev_object)

		//w.WriteHeader(http.StatusOK)		

		temp.Execute(w, visualizeReport)
	}else if r.Method != "POST" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		return
	}

}



func dashboard(w http.ResponseWriter, r *http.Request) {

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	// webpage route
	RouteWebpage := template.Must(template.ParseFiles("dashboard.html"))
	
	// route actions
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Url:", r.URL.Path)
		RouteWebpage.Execute(w, "Dashboard")
	} else if r.Method == "POST" {

		// user query parameters
		r.ParseForm()
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)

		// Digitalize the contents
		fname, err := Mounted(w, r, openReadFile)
		if err != nil {
			log.Fatalln("[File]:", err)
			return
		}

		choose := r.FormValue("choose")
		signed_msg = r.FormValue("status")
		address_wallet = r.FormValue("address")
		coordinates := r.FormValue("geo-marker")
		var longitude_parse float64 = 0.0
		var latitude_parse float64 = 0.0

		// compute user geostationary points
		location := Location(coordinates[0:19])
		longitude_parse, err = strconv.ParseFloat(location.Longitude_Division, 64)
		if err != nil {
			log.Printf("Error parsing longitude : %v", err.Error())
			return
		}

		latitude_parse, err = strconv.ParseFloat(location.Latituide_Division, 64)
		if err != nil {
			log.Printf("Error parsing latitude : %v", err.Error())
			return
		}

		// compute city environment level
		clientapi := weather.NewWeatherClient()

		weatherapi, err := clientapi.OpenWeather(GEO_Index_KEY)
		if err != nil {
			log.Fatalln("weather api-key :", err.Error())
			return
		}

		marker_location := clientapi.GetCoordinates(&weather.MyCoordinates{
			Longitude: longitude_parse,
			Latitude:  latitude_parse,
		})

		// fmt.Println("@marker:", marker_location)

		err = clientapi.UVCoodinates(marker_location, weatherapi)
		if err != nil {
			log.Fatalln("city weather coordinates:", err.Error())
			return
		}


		// compute city uv level
		uvinfo, err := clientapi.UVCompleteInfo(weatherapi)
		if err != nil {
			log.Fatalln("city uv tracks:", err.Error())
			return
		}

		// fmt.Println("@uv:", uvinfo)

		visualizeReport.UVinfo = uvinfo

		// compute sequence matching probabilities
		data, err := Open_SFiles("app_data/", fname)
		if err != nil {
			log.Fatalln("[No File]:", err)
			return
		}

		// following outbreak wave  
		switch choose {
		case "0":
			fmt.Fprintf(w, "Please choose any option ...")
			log.Fatalln("Choose your option")
			return
		case "1":
			var name string = "Covid-19"

			// compute predication
			err := Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				return
			}

			pattern_analysis := GetBioAlgoParameters()

			visualizeReport.Percentage = pattern_analysis.Percentage

			// route return Ok callback
			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			visualize(w, r)

		case "2":
			var name string = "FlaviDengue"
			err := Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				return
			}
			pattern_analysis := GetBioAlgoParameters()
			visualizeReport.Percentage = pattern_analysis.Percentage
			w.WriteHeader(http.StatusOK)

			r.Method = "GET"
			visualize(w, r)
		case "3":
			var name string = "KenyaEbola"
			err := Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				return
			}
			pattern_analysis := GetBioAlgoParameters()
			visualizeReport.Percentage = pattern_analysis.Percentage

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			visualize(w, r)
		case "4":
			var name string = "ZikaVirusBrazil"
			err := Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				return
			}
			pattern_analysis := GetBioAlgoParameters()
			visualizeReport.Percentage = pattern_analysis.Percentage

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			visualize(w, r)
		case "5":
			var name string = "MersSaudiaArabia"
			err := Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				return
			}
			pattern_analysis := GetBioAlgoParameters()
			visualizeReport.Percentage = pattern_analysis.Percentage
			w.WriteHeader(http.StatusOK)

			r.Method = "GET"
			visualize(w, r)

		default:
			RouteWebpage.Execute(w, "dashboard")
		}
	}

}


func customerViews(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("feedback.html"))

	// require different param
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Feedback")
	}
}

func terms(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("terms.html"))

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// route actions
	if r.Method == "GET" {
		fmt.Println("Url:", r.URL.Path)
		fmt.Println("Method:" + r.Method)
		temp.Execute(w, "Terms")
	}
}


func logout(w http.ResponseWriter, r *http.Request) {

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// route actions 
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

func stop_codon(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	// webpage route
	webpge := template.Must(template.ParseFiles("codons.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "codons")
	}
}

func phenylalanine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	// webpage route
	webpge := template.Must(template.ParseFiles("phenylalanine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "phenylalanine")
	}
}

func leucine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("leucine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "leucine")
	}
}

func isoleucine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	// webpage route  
	webpge := template.Must(template.ParseFiles("isoleucine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "isoleucine")
	}
}

func methionine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")
	
	// webpage route
	webpge := template.Must(template.ParseFiles("methionine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "methionine")
	}
}

func valine(w http.ResponseWriter, r *http.Request) {

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("valine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "valine")
	}
}

func serine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("serine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "serine")
	}
}

func proline(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("proline.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "proline")
	}
}

func threonine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("threonine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "threonine")
	}
}

func alanine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("alanine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "alanine")
	}
}

func tyrosine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("tyrosine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "tyrosine")
	}
}

func histidine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("histidine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "histidine")
	}
}
func glutamine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("glutamine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "glutamine")
	}
}

func asparagine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("asparagine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "asparagine")
	}
}

func lysine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("lysine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "lysine")
	}
}

func aspartic(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("aspartic.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "aspartic")
	}
}

func glutamic(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("glutamic.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "glutamic")
	}
}

func cysteine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("cysteine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "cysteine")
	}
}

func tryptophan(w http.ResponseWriter, r *http.Request) {
	
	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route 
	webpge := template.Must(template.ParseFiles("tryptophan.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "tryptophan")
	}
}

func arginine(w http.ResponseWriter, r *http.Request) {
	
	// user request headers 
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// page route 
	webpge := template.Must(template.ParseFiles("arginine.html"))
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "arginine")
	}
}

func glycine(w http.ResponseWriter, r *http.Request) {
	
	// User requests headers 
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// gylcine route 

	webpge := template.Must(template.ParseFiles("glycine.html"))
	
	// methods against route
	if r.Method == "GET" {
		fmt.Println("Method:" + r.Method)
		fmt.Println("Path: ", r.URL.Path)
		webpge.Execute(w, "glycine")
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