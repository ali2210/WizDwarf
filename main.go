/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package
package main

// Libraries
import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ali2210/wizdwarf/other/bucket/fireclient"
	cryptos "github.com/ali2210/wizdwarf/other/crypto"
	genetics "github.com/ali2210/wizdwarf/other/genetic"
	"github.com/ali2210/wizdwarf/other/jsonpb"
	"github.com/ali2210/wizdwarf/other/jsonpb/jsonledit"
	"github.com/ali2210/wizdwarf/other/parser"
	"github.com/ali2210/wizdwarf/other/proteins"
	secrets "github.com/ali2210/wizdwarf/other/secrets/secretsparam"

	dbucketerror "github.com/ali2210/wizdwarf/other/bucket/storj_bucket/bucket"
	"github.com/sethvargo/go-diceware/diceware"

	"cloud.google.com/go/firestore"
	"github.com/SkynetLabs/go-skynet/v2"
	error_codes "github.com/ali2210/wizdwarf/errors_codes"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/bucket"
	dbucket "github.com/ali2210/wizdwarf/other/bucket/storj_bucket"
	logcache "github.com/ali2210/wizdwarf/other/cache_logs"
	"github.com/ali2210/wizdwarf/other/cloudmedia"
	"github.com/ali2210/wizdwarf/other/cloudmedia/dlink"
	"github.com/ali2210/wizdwarf/other/cloudmedia/media"
	timer "github.com/ali2210/wizdwarf/other/date_time"
	genome "github.com/ali2210/wizdwarf/other/genetic/binary"
	"github.com/ali2210/wizdwarf/other/logformatter"
	"github.com/ali2210/wizdwarf/other/molecules"
	"github.com/ali2210/wizdwarf/other/proteins/binary"
	websession "github.com/ali2210/wizdwarf/other/session"
	user "github.com/ali2210/wizdwarf/other/users/register"
	"github.com/ali2210/wizdwarf/piplines"
	connection "github.com/alimasyhur/is-connect"
	"github.com/briandowns/openweathermap"
	"github.com/emojisum/emojisum/emoji"
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
	"github.com/hashicorp/vault/api"

	"github.com/ali2210/wizdwarf/other/geo"
	weather "github.com/ali2210/wizdwarf/other/openweather"
	"github.com/pusher/pusher-http-go"

	"github.com/ali2210/wizdwarf/other/users"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// constants

const Secret_path string = "kv/appsecret"
const VAULT_ADDRESS string = "http://127.0.0.1:8200"

// Variables

var (
	emailexp string = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	passexp  string = "([A-Z][a-z]*[0-9])*"
	// addressexp   string                = "(^0x[0-9a-fA-F]{40}$)"
	userSessions *sessions.CookieStore = nil //user level
	// publicAddress     string                = ""
	// edit              bio.LevenTable        = SetEditParameters()
	algo            info.Levenshtein
	visualizeReport weather.DataVisualization = weather.DataVisualization{}
	// accountID              string                    = " "
	// accountKey             string = " "
	// accountVisitEmail      string = " "
	signed_msg             string = " "
	address_wallet         string = " "
	File                   string = ""
	sumof                  int64  = 0
	sb                     []string
	mss                    []float64
	occ                    []int64
	count                  int = 0
	aminochain             []*binary.Micromolecule
	GEO_Index_KEY          string = ""
	APP_CHANNEL_KEY        string = " "
	APP_CHANNEL_ID         string = " "
	APP_CHANNEL_SECRET     string = " "
	APP_CHANNEL_CLUSTER_ID string = " "
	SECRET_TOKEN           string = " "
	_start                 time.Time
	CONNECTIVITY           string = " "
	client                 *api.Client
	hcldeclare             piplines.HCLDeclaration
	evolve                 int64  = 1
	updateInfo             bool   = false
	_secret                string = " "
	_email                 string = " "
)

var (
	cacheObject                    = logcache.New(logcache.GetBigcached_config())
	logformat                      = logformatter.New()
	AppName      *firestore.Client = piplines.SetDBClientRef()
	Cloud        users.DBFirestore = piplines.SetDBCollect()
	updated_info user.Updated_User = user.Updated_User{}
)

func main() {

	// Server

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	wizDir := os.Getenv("WIZ_VOLUME_DIR")

	ast := hclparse.NewParser()

	body, errs := ast.ParseHCLFile("creds.hcl")

	if errs.HasErrors() {

		log.Fatalln(" Error :", errs)

		return
	}

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

	emoji_keys := []int64{196, 3, 50, 7, 28, 9, 119, 18, 30, 240, 175}
	emoji_values := make([]string, 11)

	for _, value := range emoji_keys {

		emoji_map_value := emoji.Map(byte(value))
		emoji_values = append(emoji_values, emoji.CodepointToUnicode(emoji_map_value[1]))
	}

	// whether port or host will be empty.
	if reflect.DeepEqual(strings.Compare(CONNECTIVITY, "OFFLINE"), 0) && !reflect.DeepEqual(host, " ") {

		log.Fatalln(emoji_values[11], " Internet Status:", CONNECTIVITY)
		log.Fatalln(emoji_values[12], " Firebase Configuration complete ... [well-done]")
		log.Fatalln(emoji_values[13], " Private IP-Address Configuration ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[14], " Protocol-Buffer v3 Configuration complete ... [well-done]")
		log.Fatalln(emoji_values[16], " UI Webboard Configuration ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[17], " Application Configuration ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[18], " Application Default IP-Address Allocate ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[19], " Application Persistance Storage Configuration completed ...  [fail]", rand.Intn(100), "%")

		panic("Application fail to started because no port is specified and we do not have writing permission on your disk ")
	}

	if reflect.DeepEqual(strings.Contains(port, ":5000"), false) && reflect.DeepEqual(strings.Contains(wizDir, "/app/app_data/apps.txt"), false) {

		log.Fatalln(emoji_values[11], " Internet Status:", CONNECTIVITY)
		log.Fatalln(emoji_values[12], " Firebase Configuration complete ... [well-done]")
		log.Fatalln(emoji_values[13], " Private IP-Address Configuration ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[14], " Protocol-Buffer v3 Configuration complete ... [well-done]")
		log.Fatalln(emoji_values[16], " UI Webboard Configuration ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[17], " Application Configuration ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[18], " Application Default IP-Address Allocate ... [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[19], " Application Persistance Storage Configuration completed ...  [fail]", rand.Intn(100), "%")
		log.Fatalln(emoji_values[20], "Make sure volume mount", rand.Intn(100), "%")
		panic("Application fail to started because no port is specified and we do not have writing permission on your disk ")
	} else {

		log.Println(emoji_values[15], " Internet Status:", CONNECTIVITY)
		log.Println(emoji_values[12], " Firebase Configuration complete ... [well-done]")
		log.Println(emoji_values[13], " Private IP-Address Configuration complete ... [well-done]")
		log.Println(emoji_values[14], " Protocol-Buffer v3 Configuration complete ... [well-done]")
		log.Println(emoji_values[16], " Salt and Pepper added ... [well-done]")
		log.Println(emoji_values[17], " Channel communication Configuration complete ... [well-done]")
		log.Println(emoji_values[18], " Data Events are encrypted.   ... [well-done]")
		log.Println(emoji_values[19], " UI Webboard Configuration complete ... [well-done]")
		log.Println(emoji_values[20], " Application started ... [well-done](All process have completed)")
		log.Println(emoji_values[21], " Application Persistance Storage Configuration completed ...  [well-done]")

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

	if reflect.DeepEqual(hcldeclare.Token_Auth, " ") {

		GEO_Index_KEY = piplines.Extractor(hcldeclare.Weatherapi, hcldeclare.Weatherapi[0:40])[3:]
		APP_CHANNEL_KEY = piplines.Extractor(hcldeclare.Channel_key, hcldeclare.Channel_key[0:40])[3:]
		APP_CHANNEL_SECRET = piplines.Extractor(hcldeclare.Secret, hcldeclare.Secret[30:36])[43:]
		APP_CHANNEL_ID = piplines.Extractor(hcldeclare.Channel_id, hcldeclare.Channel_id[0:2])
		APP_CHANNEL_CLUSTER_ID = piplines.Extractor(hcldeclare.Cluster_ID, hcldeclare.Cluster_ID[0:4])

	} else {
		geoApi := piplines.Extractor(hcldeclare.Weatherapi, hcldeclare.Weatherapi[0:40])[3:]
		channelKey := piplines.Extractor(hcldeclare.Channel_key, hcldeclare.Channel_key[0:40])[3:]
		channelSecret := piplines.Extractor(hcldeclare.Secret, hcldeclare.Secret[30:36])[43:]
		channelId := piplines.Extractor(hcldeclare.Channel_id, hcldeclare.Channel_id[0:2])
		channelCluster := piplines.Extractor(hcldeclare.Cluster_ID, hcldeclare.Cluster_ID[0:4])
		SECRET_TOKEN = hcldeclare.Token_Auth

		log.Println(emoji_values[15], "Developer Mode initating .....", rand.Intn(100), "%")

		var err error
		client, err = api.NewClient(&api.Config{Address: VAULT_ADDRESS, HttpClient: &http.Client{Timeout: time.Second * 30}})
		if err != nil {
			log.Fatalln(" connecting Error ... :", err)
			return
		}

		client.SetToken(SECRET_TOKEN)

		err = piplines.PutKV(&piplines.HCLDeclaration{
			Weatherapi:  geoApi,
			Channel_key: channelKey,
			Channel_id:  channelId,
			Secret:      channelSecret,
			Cluster_ID:  channelCluster,
			Token_Auth:  "",
		}, Secret_path, client)
		if err != nil {
			log.Fatalln(" Error :", err)
			return
		}

		mapper, err := piplines.GetKV(Secret_path)
		if err != nil {
			log.Fatalln(" Error :", err)
			return
		}

		var _credentials piplines.HCLDeclaration

		cred, err := json.Marshal(mapper)
		if err != nil {
			log.Fatalln("Error data marshal: ...", err)
			return
		}

		err = json.Unmarshal(cred, &_credentials)
		if err != nil {
			log.Fatalln("Error data unmarshal: ...", err)
			return
		}

		GEO_Index_KEY = _credentials.Weatherapi
		APP_CHANNEL_KEY = _credentials.Channel_key
		APP_CHANNEL_SECRET = _credentials.Secret
		APP_CHANNEL_ID = _credentials.Channel_id
		APP_CHANNEL_CLUSTER_ID = _credentials.Cluster_ID
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
	routing.HandleFunc("/dashboard/profile/edit", update)
	routing.HandleFunc("/dashboard/profile/view", view)
	routing.HandleFunc("/logout", logout)
	routing.HandleFunc("/messages", messages)
	routing.HandleFunc("/error", distorted)
	routing.HandleFunc("/visualize", visualize)
	routing.HandleFunc("/dashboard/dvault", dvault)
	// routing.HandleFunc("/feedback", customerViews)
	routing.HandleFunc("/terms", terms)
	routing.HandleFunc("/analysis", analysis)
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
	// routing.HandleFunc("/modal/success", success)

	// Static Files
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	routing.PathPrefix("/images/").Handler(images)
	css := http.StripPrefix("/css/", http.FileServer(http.Dir("./css")))
	routing.PathPrefix("/css/").Handler(css)
	js := http.StripPrefix("/js/", http.FileServer(http.Dir("./js")))
	routing.PathPrefix("/js/").Handler(js)
	gltf := http.StripPrefix("/models/", http.FileServer(http.Dir("./models")))
	routing.PathPrefix("/models/").Handler(gltf)

	docker_files := http.StripPrefix("/app_data/", http.FileServer(http.Dir("./app_data")))
	routing.PathPrefix("/app_data/").Handler(docker_files)

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
	user := user.New_User{
		Name:            "",
		Email:           "",
		Password:        "",
		Lastname:        "",
		Address:         "",
		Phone:           "",
		Zip:             "",
		City:            "",
		State:           "",
		Gender:          0,
		ID:              "",
		Friends:         0,
		Inspire:         0,
		Lead:            0,
		SocialEvolution: 0,
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
		user.Lastname = r.FormValue("ufname")
		user.Address = r.FormValue("address") + r.FormValue("add")
		user.City = r.FormValue("inputCity")
		user.State = r.FormValue("co")
		user.Zip = r.FormValue("inputZip")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		if r.FormValue("gender") == "on" {
			user.Gender = 1
		} else {
			user.Gender = 0
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
		ok, encrypted := piplines.Presence(w, r, regex_Email, regex_Pass, user)
		if !ok {
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

		user.Friends = int64(0)
		user.Inspire = int64(0)
		user.Lead = int64(0)
		user.SocialEvolution = int64(2)
		user.ID = encrypted.Reader

		// in case your account have been created ....
		if _, ok, err := piplines.AddNewProfile(w, r, user); ok && err == nil {

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
	add := user.New_User{}

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

		add.Email = r.FormValue("email")
		add.Password = r.FormValue("password")

		// match email patten with email addresss
		exp := regexp.MustCompile(emailexp)
		ok := exp.MatchString(add.Email)
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
		okx := reg.MatchString(add.Password)
		if !okx && len(add.Password) >= 7 {

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

		_secret = add.Password
		_email = add.Email

		// Search Data in DB
		data, update, err := piplines.Firebase_Gatekeeper(w, r, add)
		var ID string
		if !reflect.DeepEqual(data, &user.New_User{}) && err != nil {

			ID = data.ID

		} else {

			ID = update.ID
		}

		// account interface
		act := websession.Cookies{}

		// create user session for user
		if userSessions == nil {

			// initialize the web session
			userSessions = piplines.Web_Token(ID)

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

	// page renderer
	temp := template.Must(template.ParseFiles("profile.html"))

	// to find app have information
	data, update, err := piplines.Firebase_Gatekeeper(w, r, user.New_User{Email: _email, Password: _secret})
	if err != nil {
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			return
		}

		logformat.Error(value)
		return
	}

	var ID string

	// web request "get"
	if r.Method == "GET" {
		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			return
		}

		logformat.Trace(value)
		temp.Execute(w, "Profile")

	} else if r.Method == "POST" {

		if !reflect.DeepEqual(update, user.Updated_User{}) {

			ID = update.ID

		} else {

			ID = data.ID
		}

		// user add profile picture resolution must be less 2kb
		if filename, err := piplines.AvatarUpload(r, ID); err != nil && strings.Contains(filename, " ") {
			cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

			value, err := cacheObject.Get_Key("Route_Path:")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)
				return
			}

			logformat.Trace(value)
		}

		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		dashboard(w, r)
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

func analysis(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	webpage := template.Must(template.ParseFiles("analysis.html"))

	if reflect.DeepEqual(visualizeReport.UVinfo, openweathermap.UVIndexInfo{}) {

		log.Fatalln(" Error : no information")
		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)
		cacheObject.Set_Key("Bad:", "Bad Request")

		value, err := cacheObject.Get_Key("Bad:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	data, update, err := piplines.Firebase_Gatekeeper(w, r, user.New_User{Email: _email, Password: _secret})
	if err != nil {
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
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

		logformat.Error(value)
		return
	}

	var ID string

	if !reflect.DeepEqual(update, user.Updated_User{}) {

		ID = update.ID

	} else {

		ID = data.ID
	}

	files, err := os.ReadDir("app_data/")
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_DIRECTORY_NOT_FOUND)
		cacheObject.Set_Key("Path:", "No Directory Found")

		value, err := cacheObject.Get_Key("Path:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return

	}

	meta, mapData, _ := piplines.GetDocuments([]string{ID}...)
	contKey, contVae := piplines.ReflectMaps(mapData)

	for list := range files {

		if strings.Contains("app_data/"+files[list].Name(), meta) {

			visualizeReport.Avatar_Path = "/" + meta

			break
		}

		if !strings.Contains(meta, "app_data/"+files[list].Name()) {

			// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			// defer cancel()

			log.Println("Key:", contKey, "Value:", contVae)

			// inter, num := fireclient.New(ctx, piplines.Firestore_Reference()).Get(contKey, []string{contVae}...)

			// if num != 0 {
			// 	log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
			// 	return
			// }

			// _, resource := piplines.ReflectMaps(inter)

			// log.Println("Value:", inter, "Resources:", resource)

			// if strings.Contains(contKey, " ") && !strings.Contains(contVae, resource) {

			// 	log.Fatalln(error_codes.Operation_ERROR_CODE_MISMATCH_STATE)
			// 	return
			// }

			// iobject := dbucket.New_Bucket(ctx, meta, "avatars")

			// errs, access := iobject.StoreJCredentials(strings.Join(signature, " "), []string{"chief inner hint orient crane mobile pattern rude moon approve train cheap"}...)
			// if reflect.DeepEqual(errs, dbucketerror.Bucket_Error_Category_Error) && access != nil {
			// 	log.Fatalln(error_codes.CHANNEL_ERROR_CHANNEL_CLOSED)
			// 	return
			// }

			// if strings.Contains(meta, ".png") {

			// 	str := strings.Trim(meta, ".png")
			// 	str = strings.Trim(meta, "app_data/")
			// 	err := iobject.DownloadObject(iobject.GetUplinkProject(), " ", []string{str}...)
			// 	if err != nil {

			// 		log.Fatalln(error_codes.DATABASE_ERRORS_DOCUMENT_READ_ERROR)
			// 		return
			// 	}
			// }

		}
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
	var ID string
	var NAME string
	var Email string
	_start = time.Now()

	data, update, err := piplines.Firebase_Gatekeeper(w, r, user.New_User{Email: _email, Password: _secret})
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

	life := genome.Lifecode{}

	if !reflect.DeepEqual(update, &user.Updated_User{}) {

		ID = update.ID
		NAME = update.Name
		Email = update.Email
	} else {

		ID = data.ID
		NAME = data.Name
		Email = data.Email
	}

	visualizeReport.Record = weather.Man{Name: NAME, Email: Email}

	// firestore credentials
	genetics.Client = piplines.Firestore_Reference()

	// generate ed25519 key
	cdr, _ := cryptos.PKK25519(update.ID)
	genetics.Pkk = fmt.Sprintf("%x", cdr)

	// genetics object
	// rece_gen := genetics.New()

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
			protoc, err := jsonpb.ProtojsonMarshaler(aminochain[i])
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
			err = jsonpb.ProtojsonUnmarshaler(protoc, aminochain[i])
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

	cdrs, cid := cryptos.FilePrints([]string{(*jsonFile).Names}...)

	// additional param
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if !reflect.DeepEqual(update, user.Updated_User{}) {

		// each transaction param before storing in physical database.
		bucket.Key = update.ID
		bucket.Composite = Ckk

	} else {

		// each transaction param before storing in physical database.
		bucket.Key = data.ID
		bucket.Composite = Ckk

	}

	iobject := dbucket.New_Bucket(ctx, (*jsonFile).Names, "amino-chemical")

	words, err := diceware.Generate(12)
	if err != nil {
		return
	}

	log.Println("Your Passphrase is: ", words, "please save it somewhere on your computer. It'll generated by application one time only (OTP)")

	errs, access := iobject.StoreJCredentials(strings.Join(words, " "), []string{"heavy cancel window wild supply replace oppose until canvas lava lamp muffin"}...)
	if reflect.DeepEqual(errs, dbucketerror.Bucket_Error_Category_Error) && access != nil {

		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)

		cacheObject.Set_Key("Internal:", errs.String())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	errs = iobject.StoreObject(strings.Join(words, " "), life.Pkk, jsonFile.Names, jsonFile.Types, []string{"heavy cancel window wild supply replace oppose until canvas lava lamp muffin"}...)
	if reflect.DeepEqual(errs, dbucketerror.Bucket_Error_Category_Error) {

		w.WriteHeader(http.StatusBadRequest)
		distorted(w, r)

		cacheObject.Set_Key("Internal:", errs.String())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	// additional param
	ctxs, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if bucket_state := fireclient.New(ctxs, piplines.Firestore_Reference()).Store(cid.String(), cdrs[cid.String()], ID); bucket_state != 0 {
		log.Fatalln("Error:", error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE, bucket_state)
		return
	}

	// additional param
	context, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = cloudmedia.NewMediaDescriptor(context, piplines.Firestore_Reference()).AddMediaFile(&media.MediaStream{
		Name:         (*jsonFile).Names,
		Datecreated:  piplines.GetFileCreationTime((*jsonFile).Names),
		IdentityCode: ID,
		Category:     media.Descriptor_Category_Text,
		Path:         "app_data/",
		Signature:    words,
		Cdrlink:      cdrs[cid.String()],
	})
	if err != nil {
		log.Fatalln("Error : ", error_codes.DATABASE_ERRORS_DOCUMENT_CREATE_ERROR)
		return
	}

	if r.Method == "GET" {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)

		temp.Execute(w, visualizeReport)
	}

	if r.Method != "GET" {
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
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
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
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			return
		}

		logformat.Trace(value)

		act := websession.Cookies{}

		act.SetContextSession(userSessions, w, r)

		if !reflect.DeepEqual(hcldeclare.Token_Auth, " ") {
			client.ClearToken()
			client.ClientTimeout()
		}

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

func update(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	if r.Method == "POST" {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)

		r.ParseForm()

		updated_info.Name = r.FormValue("name")

		if strings.Contains(r.FormValue("sufix"), "Mr") {
			updated_info.Suffix = user.Respecful_Mr
			updated_info.Gender = user.Traits_MAN
		} else if strings.Contains(r.FormValue("sufix"), "mr") {
			updated_info.Suffix = user.Respecful_Mr
			updated_info.Gender = user.Traits_MAN
		} else if strings.Contains(r.FormValue("sufix"), "Mrs") {
			updated_info.Suffix = user.Respecful_Mrs
			updated_info.Gender = user.Traits_WOMEN
		} else if strings.Contains(r.FormValue("sufix"), "mrs") {
			updated_info.Suffix = user.Respecful_Mrs
			updated_info.Gender = user.Traits_WOMEN
		} else if strings.Contains(r.FormValue("sufix"), "Ms") {
			updated_info.Suffix = user.Respecful_Ms
			updated_info.Gender = user.Traits_WOMEN
		} else if strings.Contains(r.FormValue("sufix"), "ms") {
			updated_info.Suffix = user.Respecful_Ms
			updated_info.Gender = user.Traits_WOMEN
		} else {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			log.Fatalln(error_codes.Operation_ERROR_CODE_MISMATCH_STATE)
			return
		}

		updated_info.Lastname = r.FormValue("family")
		updated_info.Address = r.FormValue("home")
		updated_info.Zip = r.FormValue("zip")

		state := strings.Split(r.FormValue("city"), ",")
		updated_info.City, updated_info.State = state[0], state[1]

		if strings.Contains(r.FormValue("status"), "Single") {
			updated_info.Status = user.Relationship_Single
		} else if strings.Contains(r.FormValue("status"), "single") {
			updated_info.Status = user.Relationship_Single
		} else if strings.Contains(r.FormValue("status"), "Maried") {
			updated_info.Status = user.Relationship_Maried
		} else if strings.Contains(r.FormValue("status"), "maried") {
			updated_info.Status = user.Relationship_Maried
		} else if strings.Contains(r.FormValue("status"), "Widow") {
			updated_info.Status = user.Relationship_Widow
		} else if strings.Contains(r.FormValue("status"), "widow") {
			updated_info.Status = user.Relationship_Widow
		} else if strings.Contains(r.FormValue("status"), "Divoced") {
			updated_info.Status = user.Relationship_Divoced
		} else if strings.Contains(r.FormValue("status"), "divoced") {
			updated_info.Status = user.Relationship_Divoced
		} else {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			log.Fatalln(error_codes.Operation_ERROR_CODE_MISMATCH_STATE)
			return
		}

		updated_info.Zip = r.FormValue("zip")
		updated_info.University = r.FormValue("alumni")
		updated_info.Circulum = r.FormValue("programme")

		num, _ := strconv.Atoi(r.FormValue("research"))
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
			return
		}

		updated_info.Published = int64(num)
		updated_info.Citation = r.FormValue("citation")
		updated_info.Company = r.FormValue("work")
		updated_info.Date = r.FormValue("joined")
		updated_info.Email = r.FormValue("email")
		updated_info.Phone = r.FormValue("phone")
		updated_info.Achievements = r.FormValue("achievement")

		data, update, err := piplines.Firebase_Gatekeeper(w, r, user.New_User{Email: updated_info.Email, Password: _secret})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			distorted(w, r)
			cacheObject.Set_Key("Internal:", err.Error())

			value, err := cacheObject.Get_Key("Internal:")
			if err != nil {
				return
			}

			logformat.Error(value)
			log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
			return
		}

		if reflect.DeepEqual(update, &user.Updated_User{}) {

			updated_info.ID = data.ID
			updated_info.Friends = data.Friends
			updated_info.Inspire = data.Inspire
			updated_info.Lead = data.Lead
			updated_info.Password = data.Password
			updated_info.SocialEvolution = (evolve + data.SocialEvolution)
			log.Println("INFO :", updated_info, update)

			if !reflect.DeepEqual(update, &updated_info) {

				ok := piplines.UpdateProfileInfo(updated_info)
				updateInfo = false
				log.Println("Account information update successdully ..... ")
				if !ok {
					w.WriteHeader(http.StatusBadRequest)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					return
				}
			}
		}

		if !reflect.DeepEqual(update, &user.Updated_User{}) {

			updated_info.ID = update.ID
			updated_info.Friends = update.Friends
			updated_info.Inspire = update.Inspire
			updated_info.Lead = update.Lead
			updated_info.Password = update.Password
			updated_info.SocialEvolution = (evolve + update.SocialEvolution)

			ok := piplines.UpdateProfileInfo(updated_info)
			updateInfo = true
			log.Println("Account update successdully ..... ")
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				distorted(w, r)
				cacheObject.Set_Key("Internal:", err.Error())

				value, err := cacheObject.Get_Key("Internal:")
				if err != nil {
					return
				}

				logformat.Error(value)
				log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
				return
			}

		}
		w.WriteHeader(http.StatusOK)
		r.Method = "GET"
		view(w, r)

	}
}

func view(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("viewprofile.html"))

	data, update, err := piplines.Firebase_Gatekeeper(w, r, user.New_User{Email: _email, Password: _secret})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return

	}

	files, err := os.ReadDir("app_data/")
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_DIRECTORY_NOT_FOUND)
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		distorted(w, r)
		cacheObject.Set_Key("Internal:", err.Error())

		value, err := cacheObject.Get_Key("Internal:")
		if err != nil {
			return
		}

		logformat.Error(value)
		return
	}

	if !reflect.DeepEqual(update, &user.Updated_User{}) && updateInfo {

		meta, mapData, _ := piplines.GetDocuments([]string{update.ID}...)
		contKey, contVae := piplines.ReflectMaps(mapData)

		for list := range files {

			if strings.Contains("app_data/"+files[list].Name(), meta) {

				update.Link = "/" + meta
				break
			}

			if !strings.Contains(meta, "app_data/"+files[list].Name()) {

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()

				inter, num := fireclient.New(ctx, piplines.Firestore_Reference()).Get(contKey, contVae, update.ID)

				if num != 0 {
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					w.WriteHeader(http.StatusBadRequest)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				_, resource := piplines.ReflectMaps(inter)

				if strings.Contains(contKey, " ") && !strings.Contains(contVae, resource) {

					log.Fatalln(error_codes.Operation_ERROR_CODE_MISMATCH_STATE)
					w.WriteHeader(http.StatusBadRequest)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				client := skynet.New()

				if dlinkerr := cloudmedia.NewDlinkObject(&client, "app_data/").Get(meta, contVae); dlinkerr != dlink.Errors_NONE {

					log.Fatalln(dlinkerr)
					w.WriteHeader(http.StatusBadRequest)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				log.Println("File download successfully... open mounted directory", meta)
				update.Link = "/" + meta
				break
			}
		}

		if r.Method == "GET" {

			cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

			value, err := cacheObject.Get_Key("Route_Path:")
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

			logformat.Trace(value)
			webpge.Execute(w, update)
		}

	} else {

		meta, mapData, _ := piplines.GetDocuments([]string{data.ID}...)
		contKey, contVae := piplines.ReflectMaps(mapData)

		for list := range files {

			if strings.Contains("app_data/"+files[list].Name(), meta) {

				data.Link = "/" + meta

				break
			}

			if !strings.Contains(meta, "app_data/"+files[list].Name()) {

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()

				inter, num := fireclient.New(ctx, piplines.Firestore_Reference()).Get(contKey, contVae, data.ID)

				if num != 0 {
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					w.WriteHeader(http.StatusBadRequest)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				_, resource := piplines.ReflectMaps(inter)

				if strings.Contains(contKey, " ") && !strings.Contains(contVae, resource) {

					log.Fatalln(error_codes.Operation_ERROR_CODE_MISMATCH_STATE)
					w.WriteHeader(http.StatusBadRequest)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				client := skynet.New()

				if dlinkerr := cloudmedia.NewDlinkObject(&client, "app_data/").Get(meta, contVae); dlinkerr != dlink.Errors_NONE {

					log.Fatalln(dlinkerr)
					w.WriteHeader(http.StatusBadRequest)
					log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
					distorted(w, r)
					cacheObject.Set_Key("Internal:", err.Error())

					value, err := cacheObject.Get_Key("Internal:")
					if err != nil {
						return
					}

					logformat.Error(value)
					return
				}

				log.Println("File download successfully... open mounted directory", meta)
				data.Link = "/" + meta
				break
			}
		}

		if r.Method == "GET" {

			cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

			value, err := cacheObject.Get_Key("Route_Path:")
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

			logformat.Trace(value)

			webpge.Execute(w, data)
		}

	}

	if r.Method != "GET" {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
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
		logformat.Trace(value)
		return
	}

}

func dvault(w http.ResponseWriter, r *http.Request) {

	// user request headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length")

	// webpage route
	webpge := template.Must(template.ParseFiles("dvault.html"))

	data, update, err := piplines.Firebase_Gatekeeper(w, r, user.New_User{Email: _email, Password: _secret})
	if err != nil {
		return
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	// const docname string = "ContentAddress"
	// list, err := manager.NewSecretsInstance(ctx, piplines.Firestore_Reference()).GetAllDocuments([]string{docname}...)
	// if err != nil {
	// 	log.Fatalln("Error:", error_codes.DATABASE_ERRORS_DOCUMENT_READ_ERROR)
	// 	return
	// }

	var ID string
	var imagecred []*piplines.DocumentCredentials
	// var contentcred *piplines.DocumentCredentials

	old, updated := piplines.GetID(update, data)
	if !reflect.DeepEqual(update, &user.Updated_User{}) {
		ID = old.ID

	} else {

		ID = updated.ID
	}

	files, err := os.ReadDir("app_data/")
	if err != nil {
		return
	}

	counter, jcounter := 0, 0

	// Keys directly proportional to files in app_data directory
	for file := range files {

		if strings.Contains(files[file].Name(), ".png") {

			counter += 1
		} else if strings.Contains(files[file].Name(), ".gif") {
			counter += 1
		} else if strings.Contains(files[file].Name(), ".jpeg") {
			counter += 1
		} else if strings.Contains(files[file].Name(), ".json") {
			jcounter += 1
		} else {
			continue
		}
	}

	for i := 0; i < counter; i++ {

		cred, err := piplines.ImagesCryptoSignature([]string{ID}...)
		if err != nil {
			return
		}

		imagecred = append(imagecred, cred)

	}

	wallet := make([]*secrets.Vault_Params, jcounter)

	contentcred, ranger, err := piplines.ProteinsCryptoSignature(int64(jcounter))
	if err != nil {
		return
	}

	for item := range imagecred {

		for _, index := range imagecred[item].Filename {

			if !reflect.DeepEqual((*imagecred[item]), nil) && !strings.Contains(index, ".json") {

				for i := 0; i < counter; i++ {

					wallet = append(wallet, &secrets.Vault_Params{

						CDR_LINK:    (*imagecred[item]).ReflectKey[i],
						TEXT:        string((*imagecred[item]).TextView[i]),
						IsImage:     true,
						SizeOf:      (*imagecred[item]).SizeOf[i],
						Access:      "public",
						Objects:     fmt.Sprintf("%d", counter),
						ImagePath:   (*imagecred[item]).Filename[i],
						ContentView: []string{""},
					})

					// pictures = append(pictures, &secrets.Content_Vault{SharedImages: (*imagecred[item]).Filename})
				}

			}

		}

	}

	var word []string
	var str []string

	for i := range (*contentcred).TextView {

		str = append(str, string((*contentcred).TextView[i][1000:]), "*")

		word = append(word, strings.Join(str, " "))
	}

	for index := range wallet {

		for item := range imagecred {

			if !reflect.DeepEqual((*imagecred[item]).ReflectKey, (*contentcred).ReflectKey) && !reflect.DeepEqual(imagecred, nil) && index >= 0 && index < int(ranger) {

				wallet = append(wallet, &secrets.Vault_Params{

					CDR_LINK:  (*contentcred).ReflectKey[index],
					IsImage:   false,
					SizeOf:    (*contentcred).SizeOf[index],
					Access:    "private",
					Objects:   fmt.Sprintf("%d", jcounter),
					ImagePath: "",
					// ContentView: word,
				})

			}

		}

	}

	list_wallet := &secrets.List_Vault{
		List: wallet,
		ID:   ID,
	}

	// console.log("content: ...",(document.getElementsByClassName('info-content-1')[0].children[4].children[1].innerHTML).substring((document.getElementsByClassName('info-content-1')[0].children[4].children[1].innerHTML.length)-5000, (document.getElementsByClassName('info-content-1')[0].children[4].children[1].innerHTML.length)-1),"file:", document.getElementsByClassName('info-content-1')[0].children[4].children[0].innerHTML)

	pusherCred := pusher.Client{
		AppID:   APP_CHANNEL_ID,
		Key:     APP_CHANNEL_KEY,
		Secret:  APP_CHANNEL_SECRET,
		Cluster: APP_CHANNEL_CLUSTER_ID,
		Secure:  true,
	}

	if err := pusherCred.TriggerMulti([]string{"keygen"}, "tnxs", map[string]interface{}{
		"keys":   list_wallet.ID,
		"values": list_wallet.List,
	}); err != nil {
		log.Fatalln("Erorr:", error_codes.CHANNEL_ERROR_CHANNEL_CLOSED)
		return
	}

	if r.Method == "GET" {

		cacheObject.Set_Key("Route_Path:", "%"+r.URL.Path+"%"+r.Method)

		value, err := cacheObject.Get_Key("Route_Path:")
		if err != nil {
			return
		}

		logformat.Trace(value)

		webpge.Execute(w, "Keygen")
		return
	}

	if !(r.Method == "HEAD") {

		log.Println("Method:", r.Method)

		webpge.Execute(w, "head")

		return
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
