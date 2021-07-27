package piplines

import(
	"bytes"
	contxt "context"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"cloud.google.com/go/firestore"
	// "github.com/ali2210/wizdwarf/db"
	// "errors"
	"path/filepath"
	"io/ioutil"
	"reflect"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"strconv"
	"regexp"
	"text/template"
	"os"
	"github.com/gorilla/sessions"
	"github.com/biogo/biogo/alphabet"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/api/option"
	firebase "firebase.google.com/go"
	info "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
	"github.com/ali2210/wizdwarf/structs/users"
	"github.com/ali2210/wizdwarf/structs"
	// cloudWallet "github.com/ali2210/wizdwarf/db/cloudwalletclass"
)



func Firebase_Gatekeeper(w http.ResponseWriter, r *http.Request, member users.Visitors) (*users.Visitors, error) {

	
	data, err := Cloud.SearchUser(GetDBClientRef(), member)
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
			member.Apparment = user.Apparment
			member.City = user.City
			member.Zip = user.Zip
			member.Country = user.Country
			
			document,_, err := Cloud.AddUser(GetDBClientRef(), member)
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

func ReadAllow(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error) {

	seq, err := RFiles(userFile.Name())
	if err != nil {
		println("Error in read file", err)
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
			log.Fatalln("[Fail] Sequence DataFile Error", err)
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
