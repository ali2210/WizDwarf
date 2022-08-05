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
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"

	"cloud.google.com/go/firestore"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/bucket"
	"github.com/ali2210/wizdwarf/other/bucket/proto"
	logcache "github.com/ali2210/wizdwarf/other/cache_logs"
	cryptos "github.com/ali2210/wizdwarf/other/crypto"
	timer "github.com/ali2210/wizdwarf/other/date_time"
	genetics "github.com/ali2210/wizdwarf/other/genetic"
	genome "github.com/ali2210/wizdwarf/other/genetic/binary"
	"github.com/ali2210/wizdwarf/other/jsonpb"
	"github.com/ali2210/wizdwarf/other/jsonpb/jsonledit"
	"github.com/ali2210/wizdwarf/other/logformatter"
	"github.com/ali2210/wizdwarf/other/molecules"
	"github.com/ali2210/wizdwarf/other/parser"
	"github.com/ali2210/wizdwarf/other/proteins"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
	websession "github.com/ali2210/wizdwarf/other/session"
	"github.com/ali2210/wizdwarf/piplines"
	connection "github.com/alimasyhur/is-connect"
	"github.com/briandowns/openweathermap"

	"github.com/ali2210/wizdwarf/other/geo"
	weather "github.com/ali2210/wizdwarf/other/openweather"
	"github.com/pusher/pusher-http-go"

	"github.com/ali2210/wizdwarf/other/users"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// Variables

var (
	emailexp     string                = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp      string                = "([A-Z][a-z]*[0-9])*"
	addressexp   string                = "(^0x[0-9a-fA-F]{40}$)"
	userSessions *sessions.CookieStore = nil //user level
	// publicAddress     string                = ""
	// edit              bio.LevenTable        = SetEditParameters()
	algo                   info.Levenshtein
	visualizeReport        weather.DataVisualization = weather.DataVisualization{}
	accountID              string                    = " "
	accountKey             string                    = " "
	accountVisitEmail      string                    = " "
	signed_msg             string                    = " "
	address_wallet         string                    = " "
	File                   string                    = ""
	sumof                  int64                     = 0
	sb                     []string
	mss                    []float64
	occ                    []int64
	count                  int = 0
	aminochain             []*binary.Micromolecule
	profiler               *users.Visitors = &users.Visitors{}
	GEO_Index_KEY          string          = ""
	APP_CHANNEL_KEY        string          = " "
	APP_CHANNEL_ID         string          = " "
	APP_CHANNEL_SECRET     string          = " "
	APP_CHANNEL_CLUSTER_ID string          = " "
	SECRET_TOKEN           string          = " "
	_start                 time.Time
	CONNECTIVITY           string = ""
)

var (
	cacheObject                   = logcache.New(logcache.GetBigcached_config())
	logformat                     = logformatter.New()
	AppName     *firestore.Client = piplines.SetDBClientRef()
	Cloud       users.DBFirestore = piplines.SetDBCollect()
)

type HCLDeclaration struct {
	Weatherapi  string `hcl:"Weatherapi"`
	Channel_key string `hcl:"Channel_key"`
	Channel_id  string `hcl:"Channel_id"`
	Secret      string `hcl:"Secret"`
	Cluster_ID  string `hcl:"Cluster_ID"`
	Token_Auth  string `hcl:"Token_Auth"`
}

func main() {

	// Server

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	wizDir := os.Getenv("WIZ_VOLUME_DIR")

	// GEO_Index_KEY = os.Getenv("GEOCOORDINATE")
	// APP_CHANNEL_KEY = os.Getenv("Registry_PUSHER_KEY")
	// APP_CHANNEL_ID = os.Getenv("Registry_CHANNEL_ID")
	// APP_CHANNEL_SCRECT = os.Getenv("Registry_CHANNEL_SCRECT")
	// APP_CHANNEL_CLUSTER_ID = os.Getenv("Registry_CHANNEL_CLUSTER_ID")

	ast := hclparse.NewParser()

	body, errs := ast.ParseHCLFile("creds.hcl")

	if errs.HasErrors() {

		log.Fatalln(" Error :", errs)

		return
	}

	var hcldeclare HCLDeclaration

	errs = gohcl.DecodeBody(body.Body, nil, &hcldeclare)

	if errs.HasErrors() {

		log.Fatalln(" Error :", errs)
		return
	}

	// Makesure application will connect with internet

	if reflect.DeepEqual(connection.IsOnline(), false) {

		CONNECTIVITY = "OFFLINE"

	} else {

		CONNECTIVITY = "ONLINE"
	}

	if reflect.DeepEqual(hcldeclare.Token_Auth, "") {

		GEO_Index_KEY = piplines.Extractor(hcldeclare.Weatherapi, hcldeclare.Weatherapi[0:40])[3:]
		APP_CHANNEL_KEY = piplines.Extractor(hcldeclare.Channel_key, hcldeclare.Channel_key[0:40])[3:]
		APP_CHANNEL_SECRET = piplines.Extractor(hcldeclare.Secret, hcldeclare.Secret[30:36])[43:]
		APP_CHANNEL_ID = piplines.Extractor(hcldeclare.Channel_id, hcldeclare.Channel_id[0:2])
		APP_CHANNEL_CLUSTER_ID = piplines.Extractor(hcldeclare.Cluster_ID, hcldeclare.Cluster_ID[0:4])

	} else {
		geoApi := piplines.Extractor(hcldeclare.Weatherapi, hcldeclare.Weatherapi[0:40])[3:]
		channelKey := piplines.Extractor(hcldeclare.Channel_key, hcldeclare.Channel_key[0:40])[3:]
		channelSecret := piplines.Extractor(hcldeclare.Secret, hcldeclare.Secret[0:35])[7:]
		channelId := piplines.Extractor(hcldeclare.Channel_id, hcldeclare.Channel_id[0:2])
		channelCluster := piplines.Extractor(hcldeclare.Cluster_ID, hcldeclare.Cluster_ID[0:4])
		SECRET_TOKEN = hcldeclare.Token_Auth

		log.Println(" Vault credentials:", geoApi, channelKey, channelSecret, channelId, channelCluster)
	}

	// whether port or host will be empty.
	if host == "" {

		if port == " " && wizDir == " " {
			log.Fatalln("Internet Status:", CONNECTIVITY)
			log.Fatalln(" Firebase Configuration complete ... [well-done]")
			log.Fatalln(" Private IP-Address Configuration ... [fail]")
			log.Fatalln(" Protocol-Buffer v3 Configuration complete ... [well-done]")
			log.Fatalln(" UI Webboard Configuration ... [fail]")
			log.Fatalln(" Application Configuration ... [fail]")
			log.Fatalln(" Application Default IP-Address Allocate ... [fail]")
			log.Fatalln(" Application Persistance Storage Configuration completed ...  [fail]")
			log.Fatalln("Make sure volume mount")
			panic("Application fail to started because no port is specified and we do not have writing permission on your disk ")
		}
	} else {
		fmt.Println("Internet Status:", CONNECTIVITY)
		fmt.Println(" Firebase Configuration complete ... [well-done]")
		fmt.Println(" Private IP-Address Configuration complete ... [well-done]")
		fmt.Println(" Protocol-Buffer v3 Configuration complete ... [well-done]")
		fmt.Println(" FAUNA DB Connected ... [well-done]")
		fmt.Println(" Channel communication Configuration complete ... [well-done]")
		fmt.Println(" Data Events are encrypted.   ... [well-done]")
		fmt.Println(" UI Webboard Configuration complete ... [well-done]")
		fmt.Println(" Application started ... [well-done](All process have completed)")
		fmt.Println(" Application Persistance Storage Configuration completed ...  [well-done]")

		if port != "127.0.0.1:5000" {
			fmt.Println(" The webboard started with this address at ", "127.0.0.1"+port)
		}
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

			cacheObject.Set_Key("Route_Path:", "%"+arg2.URL.Path+"%"+arg2.Method)

			value, err := cacheObject.Get_Key("Route_Path:")
			if err != nil {
				return
			}

			logformat.Trace(value)

			temp.Execute(arg1, "MainPage")
		}

	})

	routing.HandleFunc("/home", home)
	routing.HandleFunc("/signup", newUser)
	routing.HandleFunc("/login", existing)
	routing.HandleFunc("/dashboard", dashboard)
	routing.HandleFunc("/dashbaord/setting", setting)
	routing.HandleFunc("/dashboard/profile", profile)
	routing.HandleFunc("/logout", logout)
	routing.HandleFunc("/messages", messages)
	routing.HandleFunc("/error", distorted)
	// routing.HandleFunc("/feedback", customerViews)
	routing.HandleFunc("/terms", terms)
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
	// routing.HandleFunc("/modal/success", success)

	// Static Files
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	routing.PathPrefix("/images/").Handler(images)
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
	routing.PathPrefix("/css/").Handler(css)
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	routing.PathPrefix("/js/").Handler(js)

	// tcp connection
	err := http.ListenAndServe(":5000", routing)
	if err != nil {
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
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
		cacheObject.Set_Key("Route_Path:", r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "home")
	}

}

func distorted(w http.ResponseWriter, r *http.Request) {

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	temp := template.Must(template.ParseFiles("distorted.html"))
	if r.Method == "GET" {
		return
	}

	temp.Execute(w, "Distorted")

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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "Regsiter")
	} else if r.Method == "POST" {
		r.ParseForm()

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

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		// user password and password pattern both are same, means password created according to password rule.
		regex_Pass, err := regexp.MatchString(passexp, user.Password)
		if err != nil {
			cacheObject.Set_Key("Internal:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		// encrypted user information
		hash, encrypted := piplines.Presence(w, r, regex_Email, regex_Pass, user)
		if !hash {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		// in case your account have been created ....
		if _, ok, err := piplines.AddNewProfile(w, r, user, encrypted.Reader); ok && err == nil {

			// route return Ok callback
			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			existing(w, r)
		} else {

			// route return BAD callback
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			return
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "Login")
	} else if r.Method == "POST" {

		// user query parameters
		r.ParseForm()

		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")

		// match email patten with email addresss
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(user.Email)
		if !ok {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("Internal:", "Internal Server Error")

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}
			logformat.Error(value)
			return
		}

		// match password pattern against password
		reg := regexp.MustCompile(passexp)
		okx := reg.MatchString(user.Password)
		if !okx && len(user.Password) >= 7 {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Credentials", "authorization_code return error")

			value, err := cacheObject.Get_Key("Credentials")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		// Search Data in DB
		data, err := piplines.Firebase_Gatekeeper(w, r, user)
		if reflect.DeepEqual(data, &users.Visitors{Eve: false}) && err == nil {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("DB connection:", "No information about user")
			value, err := cacheObject.Get_Key("DB connection:")
			if err != nil {
				return
			}

			logformat.Error(value)

			return
		}

		accountID = data.Id

		accountVisitEmail = data.Email
		accountKey = data.Password
		profiler = data

		// account interface
		act := websession.Cookies{}

		// create user session for user
		if userSessions == nil {

			// initialize the web session
			userSessions = piplines.Web_Token(data.Id)

			// valid the session data
			act.SetContextSession(userSessions, w, r)
			err := act.NewToken()
			if err != nil {

				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)
				cacheObject.Set_Key("Token:", "Web token are not created")

				value, err := cacheObject.Get_Key("Token:")
				if err != nil {
					return
				}

				logformat.Error(value)
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

	w.WriteHeader(http.StatusBadRequest)
	distorted(w, r)
	// to find app have information
	visit, err := Cloud.GetDocumentById(AppName, *profiler)
	if err != nil {
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	// encode information with json schema
	data, err := json.Marshal(visit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	// proper encoding over data streams
	err = json.Unmarshal(data, &member)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	// web request "get"
	if r.Method == "GET" {
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		temp.Execute(w, member)
	} else if r.Method == "POST" {

		// user add profile picture resolution must be less 2kb
		piplines.AvatarUpload(r, member.Id)

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
		if status_profile := piplines.UpdateProfileInfo(&user); status_profile {

			w.WriteHeader(http.StatusOK)
			return
		}

	}

}

func messages(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("messages.html"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	if r.Method == "GET" {
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "chat page render")
	}
}

func success(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("modal-success.html"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	if r.Method == "GET" {
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "success")
	}
}

func setting(w http.ResponseWriter, r *http.Request) {

	temp := template.Must(template.ParseFiles("settings.html"))

	if r.Method == "GET" {
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "setting")
	}
}

func treasure(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	webpage := template.Must(template.ParseFiles("analysis.html"))

	if reflect.DeepEqual(visualizeReport.UVinfo, openweathermap.UVIndexInfo{}) {
		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)

		return
	}

	// Check whether genome exist or not. If yes and probablity value will be less or equal to 5 then. No worry

	if visualizeReport.Probab_Event <= 5 {

		visualizeReport.Immune_Test = "negative"

	} else {

		visualizeReport.Immune_Test = "positive"
	}

	if r.Method == "GET" && algo.GetProbParameter() < 101 {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)

		// client param
		pusherClient := pusher.Client{
			AppID:   APP_CHANNEL_ID,
			Key:     APP_CHANNEL_KEY,
			Secret:  APP_CHANNEL_SECRET,
			Cluster: APP_CHANNEL_CLUSTER_ID,
			Secure:  true,
		}

		chain := molecules.DashboardAnalytics(aminochain, sumof)

		analytics_amino := make([]map[string]interface{}, len(chain))

		for i := range chain {

			if !reflect.DeepEqual(reflect.ValueOf(chain[i]).Interface(), chain[0]) {
				analytics_amino = append(analytics_amino, map[string]interface{}{
					"key":    piplines.Mapper(chain[i], "Symbol").(string),
					"values": piplines.Mapper(chain[i], "Occurance").(int64),
				})
			}

		}

		data := analytics_amino[len(analytics_amino)-count:]
		err = pusherClient.TriggerMulti([]string{"protein"}, "molecule", data)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {

				return
			}

			logformat.Error(value)
			return
		}

		visualizeReport.Elapse = timer.Elasped(time.Now(), _start)

		webpage.Execute(w, visualizeReport)
	}
}

func visualize(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	temp := template.Must(template.ParseFiles("computation.html"))

	userProfile, err := Cloud.GetDocumentById(AppName, *profiler)
	if err != nil && userProfile != nil {

		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	query_json, err := json.Marshal(userProfile)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	err = json.Unmarshal(query_json, &profiler)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	visualizeReport.Record = weather.Man{Name: profiler.Name, Email: profiler.Email}

	if r.Method == "GET" {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)

		_start = time.Now()

		// firestore credentials
		genetics.Client = piplines.Firestore_Reference()

		// generate ed25519 key
		cdr, _ := cryptos.PKK25519(profiler.Id)
		genetics.Pkk = fmt.Sprintf("%x", cdr)

		// genetics object
		// rece_gen := genetics.New()

		life := genome.Lifecode{}

		// genetics data string
		life.Genes = strings.Join(piplines.GetGenes(), "")
		life.Pkk = genetics.Pkk

		// create trust object ... trust verified whom that content .
		ok, key, err := piplines.TrustRequest(life.Pkk, address_wallet, signed_msg)
		if !ok && err != nil {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {

				return
			}

			logformat.Error(value)
			return
		}

		// genetics database
		// status := rece_gen.AddCode(context.Background(), &life)
		// if status.ErrorCode == binaries.Errors_Error {

		// 	w.WriteHeader(http.StatusBadRequest)
		// 	distorted(w, r)
		// 	cacheObject.Set_Key("Internal:", status.ErrorCode.String())

		// 	value, err := cacheObject.Get_Key("Internal:")
		// 	if err != nil {
		// 		return
		// 	}

		// 	logformat.Error(value)
		// 	return
		// }

		// get all proteins symbols
		listProteinsName := piplines.Active_Proteins(life.Genes)

		// get all amino table
		listProteins := piplines.AminoChains(life.Genes)
		ribbon := make(map[string]map[string]proteins.Aminochain)

		// read map values
		iterate := reflect.ValueOf(listProteinsName).MapRange()

		// create new marcomolecules which hold molecule state for a while

		aminochain = make([]*binary.Micromolecule, len(life.Genes))

		// iterate over map values
		for iterate.Next() {

			// store map value in other map
			ribbon[listProteinsName[iterate.Value().String()]] = listProteins

			// get polypeptide information in structured data
			extraction := molecules.Genome_Extract(ribbon, listProteinsName, iterate.Value().String())

			// if the information return void space or empty field then discard , otherwise hold that state
			if molecules.GetMoleculesState(extraction) && molecules.GetCompositionState(extraction) {
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
			if str := molecules.Make_string(aminochain[j]); str != " " {
				s = append(s, str)
			}
		}

		sb = make([]string, len(aminochain))
		occ = make([]int64, len(aminochain))
		mss = make([]float64, len(aminochain))

		// iterate_over aminochain
		for i := range aminochain {

			// check aminochain have execpted type
			if molecules.Molecular(aminochain[i]) && s[i] != " " {
				abundance, syms := molecules.Abundance(aminochain[i], strings.Join(s, ""), i)
				if reflect.DeepEqual(aminochain[i].Symbol, syms) {
					aminochain[i].Abundance = int64(abundance)
					count += 1
				}

				sb = append(sb, molecules.Symbol(aminochain[i]))
				mss = append(mss, molecules.Mass(aminochain[i]))
				occ = append(occ, molecules.Occurance(aminochain[i]))

				// marshal the protos message
				data, err := jsonpb.ProtojsonMarshaler(aminochain[i])
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				// unmarhal the protos message
				err = jsonpb.ProtojsonUnmarshaler(data, aminochain[i])
				if err != nil {

					w.WriteHeader(http.StatusBadRequest)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {

						return
					}

					logformat.Error(value)
					return
				}

				// which protein exist in abudance
				sumof = sumof + proteins.Total_chain_filter(aminochain[i].Abundance)
			}
		}

		// This param either contribute in file processing or file content that want to be store in batch mode
		jsonFile := &jsonledit.FileDescriptor{}
		jsonFile.Types = ".json"
		jsonFile.Names = parser.Generator()
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
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("Internal:", iobject.Istatus.String())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		File = jsonFile.Names + jsonFile.Types

		//prev_object := client.Preview(&proto.Query{ByName: file})
		//log.Println("prev_object:", prev_object)

		//w.WriteHeader(http.StatusOK)

		temp.Execute(w, visualizeReport)
	} else if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())
		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {

			return
		}

		logformat.Error(value)
		panic("method not supported")
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		RouteWebpage.Execute(w, "Dashboard")
	} else if r.Method == "POST" {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		// user query parameters
		r.ParseForm()

		// Digitalize the contents
		fname, err := piplines.Mounted(w, r)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {

				return
			}

			logformat.Error(value)
			return
		}

		choose := r.FormValue("choose")
		signed_msg = r.FormValue("status")
		address_wallet = r.FormValue("address")
		coordinates := r.FormValue("geo-marker")
		var longitude_parse float64 = 0.0
		var latitude_parse float64 = 0.0

		// compute user geostationary points
		location := geo.Location(coordinates[0:19])
		longitude_parse, err = strconv.ParseFloat(location.Longitude_Division, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		latitude_parse, err = strconv.ParseFloat(location.Latituide_Division, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)

			return
		}

		// compute city environment level
		clientapi := weather.NewWeatherClient()

		weatherapi, err := clientapi.OpenWeather(GEO_Index_KEY)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Trace(value)
			return
		}

		marker_location := clientapi.GetCoordinates(&weather.MyCoordinates{
			Longitude: longitude_parse,
			Latitude:  latitude_parse,
		})

		err = clientapi.UVCoodinates(marker_location, weatherapi)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			return
		}

		// compute city uv level
		uvinfo, err := clientapi.UVCompleteInfo(weatherapi)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {

				return
			}

			logformat.Error(value)
			return
		}

		visualizeReport.UVinfo = uvinfo

		// compute sequence matching probabilities
		data, err := piplines.Open_SFiles("app_data/", fname)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)

			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
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
			err := piplines.Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)

				cacheObject.Set_Key("Internal:", err.Error())

				value, err := cacheObject.Get_Key("Internal:")
				if err != nil {
					return
				}

				logformat.Error(value)
				return
			}

			visualizeReport.Probab_Event = piplines.GetBioAlgoParameters().Probablity

			visualizeReport.Percentage = piplines.GetBioAlgoParameters().Percentage

			// route return Ok callback
			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			visualize(w, r)

		case "2":
			var name string = "FlaviDengue"
			err := piplines.Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)

				cacheObject.Set_Key("Internal:", err.Error())

				value, err := cacheObject.Get_Key("Internal:")
				if err != nil {
					return
				}

				logformat.Error(value)
				return
			}

			visualizeReport.Probab_Event = piplines.GetBioAlgoParameters().Probablity
			visualizeReport.Percentage = piplines.GetBioAlgoParameters().Percentage
			w.WriteHeader(http.StatusOK)

			r.Method = "GET"
			visualize(w, r)
		case "3":
			var name string = "KenyaEbola"
			err := piplines.Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)

				cacheObject.Set_Key("Internal:", err.Error())

				value, err := cacheObject.Get_Key("Internal:")
				if err != nil {
					return
				}

				logformat.Error(value)
				return
			}

			visualizeReport.Probab_Event = piplines.GetBioAlgoParameters().Probablity
			visualizeReport.Percentage = piplines.GetBioAlgoParameters().Percentage

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			visualize(w, r)
		case "4":
			var name string = "ZikaVirusBrazil"
			err := piplines.Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)

				cacheObject.Set_Key("Internal:", err.Error())

				value, err := cacheObject.Get_Key("Internal:")
				if err != nil {
					return
				}

				logformat.Error(value)
				return
			}

			visualizeReport.Probab_Event = piplines.GetBioAlgoParameters().Probablity
			visualizeReport.Percentage = piplines.GetBioAlgoParameters().Percentage

			w.WriteHeader(http.StatusOK)
			r.Method = "GET"
			visualize(w, r)
		case "5":
			var name string = "MersSaudiaArabia"
			err := piplines.Data_Predicition(w, r, name, choose, data, algo)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)

				cacheObject.Set_Key("Internal:", err.Error())

				value, err := cacheObject.Get_Key("Internal:")
				if err != nil {
					return
				}

				logformat.Error(value)
				return
			}

			visualizeReport.Probab_Event = piplines.GetBioAlgoParameters().Probablity
			visualizeReport.Percentage = piplines.GetBioAlgoParameters().Percentage
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		act := websession.Cookies{}

		act.SetContextSession(userSessions, w, r)
		err = act.ExpireToken()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
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
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)
		webpge.Execute(w, "glycine")
	}
}
