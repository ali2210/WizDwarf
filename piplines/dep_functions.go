/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package piplines

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	rdn "math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	skynet "github.com/SkynetLabs/go-skynet/v2"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/bucket"
	"github.com/ali2210/wizdwarf/other/bucket/proto"
	"github.com/ali2210/wizdwarf/other/collection"
	"github.com/ali2210/wizdwarf/other/crypto"
	cryptos "github.com/ali2210/wizdwarf/other/crypto"
	datetime "github.com/ali2210/wizdwarf/other/date_time"
	"github.com/ali2210/wizdwarf/other/parser"
	biosubtypes "github.com/ali2210/wizdwarf/other/proteins"
	"github.com/ali2210/wizdwarf/other/users"
	"github.com/biogo/biogo/alphabet"
	"github.com/gorilla/sessions"
	linkcid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
	pusher "github.com/pusher/pusher-http-go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var (
	Firestore_Rf string
	calendar     string
	parse_date   string
	user_Id      string
	pic_src      string
	pic_time     string
	pic_tags     string
	pic_id       string
	chain        map[string]string
)

var cdr map[string]string = make(map[string]string, 1)
var genes []string

type SignedKey struct {
	Reader string
	Signed string
	Tx     *ecdsa.PrivateKey
}

const Errors = "Operation Failed"

func Pictures_Stream(r *http.Request, user_id string) {

	//  Set the buffer size for the picture file contents
	r.ParseMultipartForm(10 << 50)

	// Get the picture file from HtmlContent
	file, fileHandle, err := r.FormFile("profile-input")

	if err != nil {
		return
	}
	defer file.Close()

	// Image file accessible to application
	if _, err := os.Stat(fileHandle.Filename); os.IsExist(err) {
		return
	}

	// application store user picture in the app_data directory
	path, err := os.Stat("app_data/")
	if err != nil {
		return
	}

	// Application storage path
	if !path.IsDir() {
		return
	}

	// Store user-picture file in the storage directory
	imageFile, err := ioutil.TempFile(filepath.Dir("app_data/"), "img-*-"+fileHandle.Filename)
	if err != nil {
		return
	}
	defer imageFile.Close()

	// Read data from image-file
	readBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	// In case of error all proceding stop, otherwise file content copy into new file
	if _, err := imageFile.Write(readBytes); err != nil {
		return
	}

	// Get today year, month and date . This help to generate image metadata which is helpful when images store in the collections.
	// Get Date from html Form
	today := r.FormValue("date")
	str := strings.Trim(today, " ")

	// Parse calendar format
	date := datetime.Date(str)
	month := datetime.Month(str)
	year := datetime.Year(str)

	// parse format in string format but the typo of year format is integer
	yrs, err := strconv.Atoi(year)
	if err != nil {
		return
	}

	// parse format in string format but the typo of month format is integer
	mnths, err := strconv.Atoi(month)
	if err != nil {
		return
	}

	// parse format in string format but the typo of date format is integer
	dtes, err := strconv.Atoi(date)
	if err != nil {
		return
	}

	// calendar have time typo for encoding data (data serialization) , data typo must be string
	calendar = datetime.GetToday(yrs, time.Month(mnths), dtes).String()

	// Now Time format according american time format and some metainformation which i discarded by using [Trim func]
	time_utc := strings.Trim(calendar, "+0000 UTC")

	// After parsing American Format and remove some metainformation. We get hr and minutes. These variable again in american or international format.
	// So that i slice american time format and then parse the string
	hr, err := strconv.Atoi(time_utc[11:13])
	if err != nil {
		return
	}

	mns, err := strconv.Atoi(time_utc[14:16])
	if err != nil {
		return
	}

	// convert according to asian time format
	create_pic_time := fmt.Sprintf("%d:%d", (hr), (mns))

	// metainformation about picture
	pic_src = fileHandle.Filename
	parse_date = calendar
	user_Id = user_id
	pic_time = create_pic_time

	// parse file name which user shared. (Remove image file format which access to shared images namespace)
	// This arose a new problem that is namespace in array object and there may be at most 2% probability that
	// file don't have any file format. Compare shared resources namespace must not be empty
	rdn.Seed(1024)

	parse_num := strconv.Itoa(rdn.Intn(512))

	if str := strings.Join(parser.ParseTags(fileHandle.Filename), " "); str != "" {
		if n := strings.Compare(str, " "); n != -1 {
			pic_tags = str
			pic_id = str + "-" + parse_num
		}
	}

	// encrypted stream channel return status about content.
	// here it's not necessary
	Encrypted_Stream_Channel(imageFile.Name())

}

func SiaObjectStorage(client skynet.SkynetClient, file string) bool {

	// application store user picture in the app_data directory
	path, err := os.Stat("app_data/")
	if err != nil {
		return false
	}

	// Application storage path
	if !path.IsDir() {
		return false
	}

	// skynet portal option
	options := skynet.DefaultUploadOptions
	options.APIKey = "skynetdwarfs"
	options.CustomUserAgent = "Sia-Agent"

	// upload file to storage
	sia_object_url, err := client.UploadFile(file, options)
	if err != nil {
		return false
	}

	// high order function this function takes cdr convert into bytes with netwiork identifical code and then apply hash sha-256
	// Result -- breakable
	hash_data := sha256.Sum256([]byte(strings.Trim(sia_object_url, "sia://")))

	// this function also high order function which convert hash of bytes into encoded string
	// fog(c) = f(g(x)) mathematical notion of high order function
	// encoded string in hex format which mus be decode as string in hex format
	decoder, err := hex.DecodeString(hex.EncodeToString(hash_data[:]))
	if err != nil {
		return false
	}

	// EncodeName function takes decoder which is already in hex format and then apply x11 crypto algorithm.
	// with x11 breakable signature into unbreakable
	encodetype, err := multihash.EncodeName(decoder, "x11")
	if err != nil {
		log.Printf("Error Signature : %v", err.Error())
		return false
	}

	// This is an higher order function encodeName value as encode string in hex format.
	// multihash hex string apply on encoded string in hex format.
	encodex11, err := multihash.FromHexString(hex.EncodeToString(encodetype))
	if err != nil {
		return false
	}

	// generate new cid.. The specification of this function require two parameters (codeType & other one is hash algorithm)
	// merkel tree (dag) data serilaization (protocol buffer [https://en.wikipedia.org/wiki/Protocol_Buffers])
	// & hash algorithm

	cid := linkcid.NewCidV1(linkcid.DagProtobuf, encodex11)

	// check whether cid version is 0. For this application cid version must be 1
	if version := cid.Version(); version != 1 {
		return false
	}

	// create protocol buffer map object
	// map key which we had calculated before cid because key must be in string
	// value should be cdr link

	cdr = make(map[string]string, 1)
	cdr[cid.String()] = sia_object_url
	return true
}

func Download_Content_ownership(File string, client bucket.Bucket_Service) *proto.QState {

	return client.Download(&proto.Query{ByName: File})

}

func Mapper(stream map[string]interface{}, key string) interface{} {

	var occ interface{}
	it := reflect.ValueOf(stream).MapRange()
	for it.Next() {
		occ = stream[key]
	}
	return occ
}

func UpdateProfileInfo(member *users.Visitors) bool {
	err := cloud.UpdateUserDetails(GetDBClientRef(), *member)
	if err != nil {
		return false
	}
	return true
}

func TrustRequest(message, verifier, request string) (bool, *ed25519.PrivateKey, error) {

	// contain check whether request have pass-key ,
	// address contain address of trusted user wallet
	//  & message must not be empty

	if strings.Contains(request, "signed") && !strings.Contains(verifier, " ") && !strings.Contains(message, " ") {

		// generate keys for message
		BbKey, AleKey := cryptos.PKK25519(message)

		// bind keys with message
		bind_message, err := crypto.ASED25519(message, AleKey)
		if err != nil {
			log.Printf(" Error message binding fail %v", err.Error())
			return false, &AleKey, err
		}

		// key signature verified
		if verified := cryptos.AVED25519(message, bind_message, AleKey, BbKey); verified {
			return verified, &AleKey, nil
		}

		// key verification failed.
		return false, &AleKey, errors.New("verification failed")
	} else {

		// generate keys
		BbKey, AleKey, err := cryptos.BKED25519()
		if err != nil {
			log.Printf(" Error generating key: %v", err.Error())
			return false, &AleKey, err
		}

		// bind message with your public key
		bindMessage := cryptos.BSED25519(message)

		// bind message verification against key
		if verify := cryptos.BVED25519(BbKey, bindMessage, []byte(message)); verify {
			return verify, &AleKey, nil
		}

		// bind message verification failed
		return false, &AleKey, errors.New("error verification error")
	}
}

func Active_Proteins(str string) map[string]string {
	i, j := 0, 3

	chain = make(map[string]string, len(str))

	for u := 0; u <= len(str); u++ {
		if strings.Contains(str, str[i:j]) && u != len(str) {

			if !strings.Contains(biosubtypes.Class(str, i, j), " ") {
				chain[str[i:j]] = biosubtypes.Class(str, i, j)
			}
			i = i + 3
			j = j + 3
			if j >= len(str) {
				break
			}
		}
	}
	return chain
}

func AminoChains(str string) map[string]biosubtypes.Aminochain {

	// initalization & declaration of local attributes
	chain := make(map[string]biosubtypes.Aminochain, len(str))
	i, j := 0, 3

	// get amino map chain
	for u := 0; u <= len(str); u++ {

		if strings.Contains(str, str[i:j]) && u != len(str) {

			// in case amino chain return protein symbol and store back to the local attribute
			if !strings.Contains(biosubtypes.GetAmino(str, i, j).Symbol, " ") && threepairs(str, i) {
				chain[str[i:j]] = biosubtypes.GetAmino(str, i, j)
			}

			// protein patterns
			i = i + 3
			j = j + 3

			// "j" indicator make sure that iteration won't panic
			if j >= len(str) {
				break
			}
		}
	}
	return chain
}

func threepairs(str string, i int) bool {
	return !strings.Contains(str[i:i+1], " ") && !strings.Contains(str[i+1:i+2], " ") && !strings.Contains(str[i+2:i+3], " ")
}

func Firebase_Gatekeeper(w http.ResponseWriter, r *http.Request, member users.Visitors) (*users.Visitors, error) {

	var mapper map[string]interface{}
	var profile users.Visitors
	data, err := cloud.SearchUser(GetDBClientRef(), member)
	if reflect.DeepEqual(data, mapper) && err != nil {
		log.Println("no account on our server")
		return &profile, errors.New("no account on our server")
	}

	query, err := json.Marshal(data)
	if err != nil {
		return &profile, err
	}

	err = json.Unmarshal(query, &profile)
	if err != nil {
		return &profile, err
	}
	return &profile, nil
}

func AddNewProfile(response http.ResponseWriter, request *http.Request, user users.Visitors, im string) (*firestore.DocumentRef, bool, error) {

	var member users.Visitors
	var replicate *firestore.DocumentRef

	// user data accrording to json schema
	data, err := json.Marshal(user)
	if err != nil {
		return replicate, false, err
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return replicate, false, err
	}

	// search data if there is
	candidate, err := Firebase_Gatekeeper(response, request, user)
	if err != nil {
		return replicate, false, err

	}

	// search data doesn't exist
	if reflect.DeepEqual(candidate, &member) {

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

		// add user data in your dataabse
		document, _, err := cloud.AddUser(GetDBClientRef(), member)
		if err != nil {
			return replicate, false, err
		}

		replicate = document
		return document, true, nil
	} else {
		return replicate, false, err
	}
}

func Firestore_Reference() *firestore.Client {

	_, err := os.Stat("config/" + GetKeyFile())
	if os.IsExist(err) {
		return &firestore.Client{}
	}

	Firestore_Rf = "config/" + GetKeyFile()

	firebase_connect, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: GetProjectID()}, option.WithCredentialsFile(Firestore_Rf))
	if err != nil {
		return &firestore.Client{}
	}

	client, err := firebase_connect.Firestore(context.Background())
	if err != nil {
		return &firestore.Client{}
	}
	return client

}

func StringInt(s string) (int, error) {

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil

}

// func IsEvm(hash, addressexp string, clientInstance *ethclient.Client) bool {

// 	expression := regexp.MustCompile(addressexp)
// 	v := expression.MatchString(hash)

// 	address := common.HexToAddress(hash)
// 	bytecode, err := clientInstance.CodeAt(contxt.Background(), address, nil)
// 	if err != nil {
// 		return
// 	}

// 	contract := len(bytecode) > 0
// 	return contract
// }

func Web_Token(unique string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(unique))
}

func MountDisk(w http.ResponseWriter, r *http.Request, filename string) os.FileInfo {
	f, err := os.OpenFile(filename+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil
	}

	finfo, err := f.Stat()
	if err != nil {
		return nil
	}
	return finfo
}

func Signx(w http.ResponseWriter, r *http.Request, h1, h2 string) (string, string, *ecdsa.PrivateKey) {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", &ecdsa.PrivateKey{}
	}

	// 0x40fa6d8c32594a971b692c44c0c56b19c32613deb1c6200c26ea4fe33d34a5fd
	hash_fist := sha256.Sum256([]byte(h1))
	hash_sec := sha256.Sum256([]byte(h2))

	hash := make([]byte, len(hash_fist))
	for i := range hash_sec {
		hash = append(hash, (hash_fist[i] &^ hash_sec[i]))
	}

	reader, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", "", &ecdsa.PrivateKey{}
	}

	return fmt.Sprintf("0x%x", reader), fmt.Sprintf("0x%x", s), privateKey

}

func RFiles(filename string) ([]byte, error) {

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return []byte{}, err
	}

	return []byte(body), nil
}

func Presence(w http.ResponseWriter, r *http.Request, regexp_emal, regexp_pss bool, user users.Visitors) (bool, *SignedKey) {

	code := SignedKey{}

	if !regexp_emal && !regexp_pss {
		return false, &SignedKey{}
	}

	code.Reader, code.Signed, code.Tx = Signx(w, r, hex.EncodeToString([]byte(user.Email)), hex.EncodeToString([]byte(user.Password)))
	return true, &code
}

func Mounted(w http.ResponseWriter, r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 50)

	var upldFile *os.File = nil
	file, handler, err := r.FormFile("fileSeq")
	if err != nil {
		return "", err
	}
	defer file.Close()

	if handler.Size >= (500000 * 1024) {
		return "", err
	}

	if _, err := os.Stat(handler.Filename); os.IsExist(err) {
		return "", err
	}

	path, err := os.Stat("app_data/")
	if err != nil {
		return "", err
	}

	if !path.IsDir() {
		return "", err
	}

	path, err = os.Stat("app_data/")

	if err != nil {
		return "", err
	}

	if !path.IsDir() {
		return "", err
	}

	// upload file by user...
	upldFile, err = ioutil.TempFile(filepath.Dir("app_data/"), "apps-"+"*.txt")
	if err != nil {
		return "", err
	}
	defer upldFile.Close()

	_, err = upldFile.Stat()
	if err != nil {
		return "", err
	}

	openReadFile := upldFile.Name()

	// file convert into bytes
	bytesFile, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	_, err = upldFile.Write(bytesFile)
	if err != nil {
		return "", err
	}

	return openReadFile, nil

}

/*
*  To keep user privacy , Encrypted-Channels accept stream of data.
*  Onces the data transfer complete this path completely closed.
 */
func Encrypted_Stream_Channel(file string) bool {

	/*
	* Create serialized message object [@Protocol-Buffers is used here].
	* This Object create metadata & image file store in database. [@Firestore]
	 */

	// Create firestore object
	gallery_firestore_object := &collection.Gallery_Stream_Server{}

	// Get Firestore client object
	collection.Firestore_Picture_Client = Firestore_Reference()

	// create decentralize skynet storage object
	client := skynet.New()

	// Store user-picture file in the storage directory
	if ok := SiaObjectStorage(client, file); ok {
		pusher_credentials := pusher.Client{
			AppID:   "1265511",
			Key:     "65993b3c66b5317411a5",
			Secret:  "4f8bf3faf121d9c8dadf",
			Cluster: "mt1",
			Secure:  true,
		}

		// picture metadata
		storage := &collection.Pictures{}
		storage.PicTime = pic_time
		storage.PicSrc = pic_src
		storage.PicId = pic_id
		storage.PicDate = parse_date
		storage.UserAgentId = user_Id
		storage.PicTags = pic_tags
		storage.CDR = make(map[string]string, 1)
		storage.CDR = cdr

		firestore_collection := gallery_firestore_object.NewPictures(context.Background(), storage)

		pusher_credentials.Trigger("encrypted-photo-stream", "encrypted-pic", firestore_collection)
		return true
	}
	return false
}

func ReadAllow(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error) {

	seq, err := RFiles(userFile.Name())
	if err != nil {
		return []string{}, []string{}, err
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
		return []string{}, []string{}, err
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

// func blockSession(id int) *sessions.CookieStore {

// 	return sessions.NewCookieStore([]byte(strconv.Itoa(id)))
// }

func Data_Predicition(w http.ResponseWriter, r *http.Request, fname, choose string, file *os.File, algo info.Levenshtein) error {

	i, err := strconv.Atoi(choose)
	if err != nil {
		return err
	}
	if (i > 0 && i < 6) && (fname != " ") {

		// data have peristance location address
		svrFile := MountDisk(w, r, fname)

		// read document and convert into managable format for processing
		Usr, Virus, err := ReadAllow(file, svrFile)
		if err != nil {
			return err
		}

		// Gene store in the memory
		SetGenes(Usr)

		// calculate matching probability
		distance := GetEditParameters().EditDistanceStrings(Virus, Usr)
		SetBioAlgoParameters(algo.Result(distance), fname, algo.CalcualtePercentage(algo.Probablity))

		return err
	} else if i == 0 {
		// reload dashboard page
		temFile := template.Must(template.ParseFiles("dashboard.html"))
		temFile.Execute(w, "Dashbaord")
	}
	return err

}

// set genes
func SetGenes(gene []string) {
	genes = append(genes, gene...)
}

// get genes
func GetGenes() []string { return genes }

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
