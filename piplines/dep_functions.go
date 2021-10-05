package piplines

import(
	"bytes"
	contxt "context"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"log"
	"cloud.google.com/go/firestore"
	"strconv"
	"path/filepath"
	"io/ioutil"
	"reflect"
	"strings"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"time"
	"regexp"
	"text/template"
	"os"
	"math"
	"github.com/gorilla/sessions"
	"github.com/biogo/biogo/alphabet"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/api/option"
	firebase "firebase.google.com/go"
	info "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
	"github.com/ali2210/wizdwarf/structs/users"
	"github.com/ali2210/wizdwarf/structs"
	"github.com/ali2210/wizdwarf/structs/collection"
	skynet "github.com/SkynetLabs/go-skynet/v2"
	multihash "github.com/multiformats/go-multihash"
	linkcid "github.com/ipfs/go-cid"
)

var(
	Firestore_Rf string
	calendar string
	parse_date string
	user_Id string
	pic_src string
	pic_time string
	pic_tags string
	pic_id string
	picture_ref map[string]string
)


func Pictures_Stream(r *http.Request, user_id string){
	
	
	//  Set the buffer size for the picture file contents
	r.ParseMultipartForm(10 << 50)
	
	// Get the picture file from HtmlContent
	file , fileHandle, err := r.FormFile("profile-input")
	
	if err != nil {
		log.Printf("Error parsing avatars: %v", err.Error())
		return 
	}
	defer file.Close()
	fmt.Println(" File Handler :", fileHandle.Filename, fileHandle.Size)
	
	// Image file accessible to application
	if _ , err := os.Stat(fileHandle.Filename);os.IsExist(err) {
		log.Println(" Error directory exists :", err.Error())
		return
	}
	
	// application store user picture in the app_data directory
	path , err := os.Stat("app_data/"); if err != nil {
		log.Println(" Error stat :", err.Error())
		return
	}


	// Application storage path 
	if !path.IsDir(){
		log.Fatalln("[Error] Reading File", err.Error())
		return
	}
	// Store user-picture file in the storage directory
	imageFile, err := ioutil.TempFile(filepath.Dir("app_data/"), "img-*-"+fileHandle.Filename)
	if err != nil {
		log.Printf("Error creating temporary image file: %v", err.Error())
		return
	}
	defer imageFile.Close()
	
	// Read data from image-file 
	readBytes , err := ioutil.ReadAll(file); if err != nil {
		log.Println("Error reading image file", err.Error())
		return
	}
	
	// In case of error all proceding stop, otherwise file content copy into new file
	if _, err := imageFile.Write(readBytes); err != nil {
		log.Println("Error writing image file", err.Error())
		return
	}

	// User image file returned name 
	fmt.Println("Image file created: ", imageFile.Name())

	// Get today year, month and date . This help to generate image metadata which is helpful when images store in the collections.
	// Get Date from html Form
	today := r.FormValue("date")
	str := strings.Trim(today, " ")
	
	// Parse calendar format 
	date := Date(str)
	month := Month(str)
	year := Year(str)
	
	fmt.Println("Today:", date, month, year)
	
	// parse format in string format but the typo of year format is integer
	yrs, err := strconv.Atoi(year); if err != nil {
		log.Println("Error year parsing:", err.Error())
		return 
	}
	fmt.Println("Year: ", yrs)

	// parse format in string format but the typo of month format is integer
	mnths , err := strconv.Atoi(month); if err != nil {
		log.Println("Error month parsing:", err.Error())
		return 
	}

	fmt.Println("months:", mnths)

	// parse format in string format but the typo of date format is integer
	dtes , err := strconv.Atoi(date); if err != nil {
		log.Println("Error date parsing:", err.Error())
		return 
	}

	fmt.Println("date:", dtes)


	// calendar have time typo for encoding data (data serialization) , data typo must be string
	calendar = GetToday(yrs, time.Month(mnths) ,dtes).String()

	// Now Time format according american time format and some metainformation which i discarded by using [Trim func]
	time_utc := strings.Trim(calendar,"+0000 UTC")
	
	// After parsing American Format and remove some metainformation. We get hr and minutes. These variable again in american or international format.
	// So that i slice american time format and then parse the string
	hr , err := strconv.Atoi(time_utc[11:13]); if err != nil {
		log.Println("Error parsing hr", err.Error())
		return
	}
	fmt.Println("Hour:", hr)
	
	mns , err := strconv.Atoi(time_utc[14:16]); if err != nil {
		log.Println("Error parsing min", err.Error())
		return
	}
	fmt.Println("min", mns)
	
	// convert according to asian time format 
	create_pic_time := fmt.Sprintf("%d:%d",(hr+21%12),(mns+11%60))
	fmt.Println("image create at :", create_pic_time)


	// metainformation about picture
	pic_src = fileHandle.Filename
	parse_date = calendar
	user_Id = user_id
	pic_time = create_pic_time

	// parse file name which user shared. (Remove image file format which access to shared images namespace)
	// This arose a new problem that is namespace in array object and there may be at most 2% probability that 
	// file don't have any file format. Compare shared resources namespace must not be empty 

	rd_num := math.Pow(2, 512) * 1.61
	num := strconv.FormatFloat(rd_num, 'E', -1, 64)

	Ints, err := strconv.Atoi(num); if err != nil {
		log.Printf("Error converting %s to int ", err.Error())
		return
	}

	parse_num := strconv.Itoa(Ints)

	if str := strings.Join(ParseTags(fileHandle.Filename), " "); str != "" {
		if n := strings.Compare(str, " "); n != -1 {
			pic_tags = str
			pic_id = str+ parse_num
		}
	}

    
	// encrypted stream channel calling
	Encrypted_Stream_Channel(imageFile.Name())
	
}

// DATE PARSE
func Date(s string) string{
	var d string 
	if (s[2:2]) == "-"{
		d = s[0:1]
		fmt.Println("date :", d)
	}else if s[3:3] == "-"{
		d = s[0:2]
		fmt.Println("date :", d)
	} 
	fmt.Println("date :", d)
	return d
}

// MONTH PARSE
func Month(s string) string{
	var m string
	if s[4:4] == "-"{
		m = s[2:3]
		fmt.Println("month [2:3]", m)
	}else if s[5:5] == "-"{
		m = s[2:4]
		fmt.Println("month [2:4]", m)
	}else{
		m = s[2:3]
		fmt.Println("mnths :", s[2:3])
	} 
	fmt.Println("m :", m)
	return m
}

// YEAR PARSE
func Year(s string) string{
	var y string
	if len(s) == 9 || len(s) == 10{
		y = s[5:9]
	}
	y = s[4:8]
	return y 
}



func GetToday(year int, month time.Month, date int) (time.Time) {
	return time.Date(year, month, date,12,30, 0,0,time.UTC)
}


func ParseTags(s string) []string {
	var tags = make([]string,len(s))
	if strings.Contains(s, ".png"){
		tags = strings.Split(s, ".png")
	}else if strings.Contains(s, ".jpeg"){
		tags = strings.Split(s, ".jpeg")
	}else if strings.Contains(s, ".img"){
		tags = strings.Split(s, ".img")
	}else if strings.Contains(s, ".gif"){
		tags = strings.Split(s, ".gif")
	}else if strings.Contains(s, ".tif"){
		tags = strings.Split(s, ".tif")
	}else{
		tags = append(tags, " ")
	}
	return tags
}


func SiaObjectStorage(client skynet.SkynetClient, file string) bool {
	
	// application store user picture in the app_data directory
	path , err := os.Stat("app_data/"); if err != nil {
		log.Println(" Error stat :", err.Error())
		return false
	}

	// Application storage path 
	if !path.IsDir(){
		log.Fatalln("[Error] Reading File", err.Error())
		return false
	}

	// skynet portal option
	options := skynet.DefaultUploadOptions
	options.APIKey = "skynetdwarfs"
	options.CustomUserAgent = "Sia-Agent"

	// upload file to storage
	sia_object_url, err := client.UploadFile(file, options); if err != nil {
		log.Printf("Error uploading %v:", err.Error())
		return false
	}
	
	fmt.Println("Sia Object:", sia_object_url)
	ky, err := multihash.FromB58String(sia_object_url); if err != nil {
		log.Printf("Error B58 string conversion:%v", err.Error())
		return false
	}

	cid := linkcid.NewCidV0(ky) 
	set_Cid := linkcid.NewSet()
	set_Cid.Add(cid)
	if ok := set_Cid.Has(cid); ok {
		fmt.Println("Key: ", set_Cid.Keys(), set_Cid.Len())
		// picture_ref[set_Cid.Keys()] = sia_object_url
	}
	return true

}


func Firebase_Gatekeeper(w http.ResponseWriter, r *http.Request, member users.Visitors) (*users.Visitors, error) {

	
	data, err := cloud.SearchUser(GetDBClientRef(), member)
	if err != nil && data != nil {
			log.Fatal("[Fail] No info ", err)
			return &users.Visitors{}, err
	}

	fmt.Println("Member:", data)
	
	query , err := json.Marshal(data); if err != nil{
		log.Fatal("Alien Format:", err.Error())
			return &users.Visitors{}, err
	}
	
	var profile users.Visitors
	err = json.Unmarshal(query, &profile)
	if err != nil{
		log.Fatal("Bash processing error:  ", err)
			return &users.Visitors{}, err
	}
	fmt.Println("Profile:", profile)
	return &profile, nil
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

func UpdateProfile(w http.ResponseWriter, r *http.Request, user users.Visitors)(*users.Visitors, error){

	var member users.Visitors
	request_member , err := json.Marshal(member); if err != nil {
		log.Println("Error marshalling", err.Error())
		return &member, err 
	}

	if err = json.Unmarshal(request_member, &member); err != nil {
		log.Println("Error unmarshalling", err.Error())
		return &member , err
	}
	fire_gateway, err := Firebase_Gatekeeper(w,r,user); if err != nil {
		log.Println("Error data request", err.Error())
		return &member, err
	}

	if !reflect.DeepEqual(fire_gateway, request_member){
		doc , result, err := cloud.AddUser(GetDBClientRef(), user); if err != nil {
			log.Println("Error update user", err.Error())
			return &member, err
		}

		fmt.Println("Document:", doc, "Result:", result)
		json_firebase_user, err := cloud.SearchUser(GetDBClientRef(), user); if err != nil {
			log.Println("Error search user", err.Error())
			return &member, err 
		}
		data, err := json.Marshal(json_firebase_user); if err != nil {
			log.Println("Error marshal user", err.Error())
			return &member, err
		}
		if err = json.Unmarshal(data, &member); err != nil {
			log.Println("Error unmarshal user", err.Error())
			return  &member, err
		}
		log.Println("Data:", string(data))
	}else{
		log.Println("Error when updating user", err.Error())
		return &member, err
	}
	return fire_gateway, nil
}

func AddNewProfile(response http.ResponseWriter, request *http.Request, user users.Visitors, im string) (*firestore.DocumentRef, error){

		var member users.Visitors
		var replicate *firestore.DocumentRef

		fmt.Println("Member:", member, "exuser:", user)

		data, err := json.Marshal(member)
		if err != nil {
			log.Fatal("[Fail] Poor DATA JSON FORMAT  ", err)
			return &firestore.DocumentRef{}, err
		}
		
		fmt.Println("json_data:", string(data))
		
		err = json.Unmarshal(data, &member)
		if err != nil {
			log.Fatal("[Fail] Poor Formating  ", err)
			return &firestore.DocumentRef{}, err
		}

		candidate, err := Firebase_Gatekeeper(response, request, user)
		
		if err != nil {
			log.Fatal("[Fail] Iterator Terminate :  ", err)
			return &firestore.DocumentRef{}, err
		}
		
		fmt.Println("Candiate :", candidate, )
		if reflect.DeepEqual(candidate, &member){

			member.Id = im
			member.Name = user.Name
			member.Email = user.Email
			member.Password = user.Password
			member.LastName = user.LastName
			if user.Eve {
				member.Eve = user.Eve
			} else {
				member.Eve = user.Eve
			}
			member.Address = user.Address
			member.Appartment = user.Appartment
			member.City = user.City
			member.Zip = user.Zip
			member.Country = user.Country
			
			document,_, err := cloud.AddUser(GetDBClientRef(), member)
			if err != nil {
				log.Fatal(" Bash Processing Error ", err.Error())
				return &firestore.DocumentRef{}, err
			}
			
			fmt.Println("Document:", document)
			replicate = document
			return document, nil
		}
		log.Fatal(" Iterator return data: ", err.Error(), replicate)
		return replicate, err
}

// Functions

func Firestore_Reference() *firestore.Client {

	_, err := os.Stat("config/" + GetKeyFile())
	if os.IsExist(err) {
		fmt.Println("File Doesn't exist...", err)
		return nil
	}

	Firestore_Rf = "config/" + GetKeyFile()

	firebase_connect , err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: GetProjectID()}, option.WithCredentialsFile(Firestore_Rf)); if err != nil{
		fmt.Println("Connection reject", err.Error())
		return &firestore.Client{}	
	}
	
	client , err := firebase_connect.Firestore(context.Background()); if err != nil{
		fmt.Println("Connection busy", err.Error())
		return &firestore.Client{}
	}
	return client
	

	
}

func StringInt(s string) (int, error) {

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

// func Retrieve_Crypto(w *structs.Acc, ledger db.PublicLedger) (bool, *cloudWallet.EthereumWalletAcc) {

// 	ethAcc, err := ledger.FindMyPublicAddress(w, GetDBClientRef())
// 	if err != nil {
// 		log.Fatalln("[Fail] Ledger ahve no Information / internal issue ", err)
// 		return false, nil
// 	}
// 	if ethAcc != nil {
// 		return false, nil
// 	}
// 	return true, ethAcc

// }

// func MyEthAddress(w *structs.Acc, ledge db.PublicLedger) (*cloudWallet.EthereumWalletAcc, bool) {

// 	acc, err := ledge.FindMyAddressByEmail(w, GetDBClientRef())
// 	if err != nil {
// 		log.Fatalln("[Fail] Configuration issue", err)
// 		return nil, false
// 	}
// 	if acc == nil {
// 		return nil, false
// 	}
// 	return acc, true
// }

// func FindEthWallet(w *structs.Acc, ledge db.PublicLedger) (*cloudWallet.EthereumWalletAcc, bool) {

// 	acc, err := ledge.FindMyPublicAddress(w, GetDBClientRef())
// 	if err != nil {
// 		log.Fatalln("[Fail] Configuration issue", err)
// 		return nil, false
// 	}
// 	return acc, true
// }

func IsEvm(hash, addressexp string, clientInstance *ethclient.Client) bool {

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

func Web_Token(unique string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(unique))
}

func MountDisk(w http.ResponseWriter, r *http.Request, filename string) os.FileInfo {
	f, err := os.OpenFile(filename+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("[Fail] No File Exist  ", err)
			return nil
	}

	finfo, err := f.Stat()
	if err != nil {
		log.Fatal("[Fail] Application Access   ", err)
		return nil

	}
	return finfo
}

func Signx(w http.ResponseWriter, r *http.Request, h1, h2 string) (string, string, *ecdsa.PrivateKey) {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal("[Fail] Key generate   ", err)
		return "", "", privateKey
	}

	// 0x40fa6d8c32594a971b692c44c0c56b19c32613deb1c6200c26ea4fe33d34a5fd
	hash_fist := sha256.Sum256([]byte(h1)) 
	hash_sec := sha256.Sum256([]byte(h2))

	hash := make([]byte, len(hash_fist))
	for i := range hash_sec{
		hash = append(hash, (hash_fist[i] &^ hash_sec[i]))
	}
	
	reader, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {

		log.Fatal("[Fail] Key signed   ", err)
		return "", "", privateKey

	}

	return fmt.Sprintf("0x%x", reader), fmt.Sprintf("0x%x", s), privateKey

}

func RFiles(filename string) ([]byte, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("[Fail] File  Access ", err)
		return nil, err
	}
	return []byte(body), nil
}

func Presence(w http.ResponseWriter, r *http.Request, regexp_emal, regexp_pss bool, user users.Visitors) (bool, *structs.SignedKey) {
	code := structs.SignedKey{}
	if !regexp_emal && !regexp_pss {
		return false, &structs.SignedKey{}	
	}
	code.Reader, code.Signed, code.Tx = Signx(w, r, hex.EncodeToString([]byte(user.Email)), hex.EncodeToString([]byte(user.Password)))
	return true, &code
}

func Mounted(w http.ResponseWriter, r *http.Request, openReadFile string) (string, error) {
	r.ParseMultipartForm(10 << 50)

	var upldFile *os.File = nil
	file, handler, err := r.FormFile("fileSeq")
	if err != nil {
		log.Fatal("[Fail] Error in upload   ", err)
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
		return "", err

	}
	
	path, err := os.Stat("app_data/")
	if err != nil {
		log.Fatalln("[Error] In directory", err)
		return "", err
	}
	
	if !path.IsDir() {
		log.Fatalln("[Error] Reading File", err)
		return "", err
	}

	path, err = os.Stat("app_data/")
	if err != nil {
		log.Fatalln("[Error] In directory", err)
		return "", err
	}
	if !path.IsDir() {
		log.Fatalln("[Error] Reading File", err)
		return "", err
	}
	// upload file by user...
	upldFile, err = ioutil.TempFile(filepath.Dir("app_data/"), "apps-"+"*.txt")
	if err != nil {
		log.Fatal("[Fail] Temporary File ", err)
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
		return "", err
	}
	n, err := upldFile.Write(bytesFile)
	if err != nil {
		return "", err
	}
	log.Println("[Result] = File added on server", upldFile.Name(), "Size:", n)
	return openReadFile, nil

}

	/*
	*  To keep user privacy , Encrypted-Channels accept stream of data. 
	*  Onces the data transfer complete this path completely closed.
	*/
func Encrypted_Stream_Channel(file string){


	/* 
	* Create serialized message object [@Protocol-Buffers is used here]. 
	* This Object create metadata & image file store in database. [@Firestore]
	*/
	
	// Create firestore object
	gallery_firestore_object := &collection.Gallery_Stream_Server{}
	collection.Firestore_Picture_Client = Firestore_Reference()
	firestore_collection := gallery_firestore_object.NewPictures(context.Background(), &collection.Pictures{
		PicTime: pic_time,
		PicSrc: pic_src,
		PicId : pic_id,
		PicDate : parse_date,
		UserAgentId : user_Id,
		PicTags : pic_tags,
	})
	fmt.Println("content:", firestore_collection)


	// create decentralize skynet storage object
	client := skynet.New()

	// Store user-picture file in the storage directory
	if ok := SiaObjectStorage(client, file); ok{
		fmt.Println("File stored successfully");
	}
}




func ReadAllow(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error) {

	seq, err := RFiles(userFile.Name())
	if err != nil {
		println("Error in rsead file", err)
		return nil, nil, err
	}

	var gen []string
	for _, v := range seq {
		space := ToRunes(v)
		if space == "" {
			gen = append(gen, "")
		}
		gen = append(gen, space)
	}

	pathogen, err := RFiles(serverFile.Name())
	if err != nil {
		println("Error in read file", err)
		return nil, nil, err
	}

	var genV []string
	for _, v := range pathogen {
		space := ToRunes(v)
		if space == "" {
			genV = append(genV, "")
		}
		genV = append(genV, space)
	}

	return gen, genV, nil

}

func ToRunes(seq byte) string {

	if seq >= 65 && seq < 91 {
		return string(alphabet.Letter(seq))
	}
	return string(alphabet.Letter(seq))
}

func blockSession(id int) *sessions.CookieStore {

	return sessions.NewCookieStore([]byte(strconv.Itoa(id)))
}

func Data_Predicition(w http.ResponseWriter, r *http.Request, fname, choose string, file *os.File, algo info.Levenshtein) error {

	i, err := strconv.Atoi(choose)
	if err != nil {
		log.Fatalln("[Fail] Sorry there is some issue report!", err)
		return err
	}
	if (i > 0 && i < 6) && (fname != " ") {
		svrFile := MountDisk(w, r, fname)
		Usr, Virus, err := ReadAllow(file, svrFile)
		if err != nil {
			log.Fatalln("Sequence data file error", err)
			return err
		}
		log.Println("Genome:", len(Virus), "virus:", len(Usr))
		distance := GetEditParameters().EditDistanceStrings(Virus, Usr)
		SetBioAlgoParameters(algo.Result(distance), fname, algo.CalcualtePercentage(algo.Probablity)) 
		
		return err
	} else if i == 0 {
		temFile := template.Must(template.ParseFiles("dashboard.html"))
		temFile.Execute(w, "Dashbaord")
	}
	return err

}



func Card_Verification(s1, s2 string) bool {

	m, n := []byte(s1), []byte(s2)
	res := bytes.Compare(m, n)
	if res == 0 {
		return true
	}
	return false

}

func Open_SFiles(path, filename string) (*os.File, error) {
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

func getValuesFromStruct(parser interface{}) []reflect.Value {
	y := reflect.ValueOf(parser).Elem()
	x := make([]reflect.Value, y.NumField())
	for i := 0; i < y.NumField(); i++ {
		x[i] = y.Field(i)
	}
	return x
}
