/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// Package or module
package piplines

// Libaries
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
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	error_codes "github.com/ali2210/wizdwarf/errors_codes"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/bucket/fireclient"
	dbucket "github.com/ali2210/wizdwarf/other/bucket/storj_bucket"
	dbucketerror "github.com/ali2210/wizdwarf/other/bucket/storj_bucket/bucket"
	"github.com/sethvargo/go-diceware/diceware"

	// "github.com/ali2210/wizdwarf/other/bucket/fireclient"
	"github.com/ali2210/wizdwarf/other/cloudmedia"
	"github.com/ali2210/wizdwarf/other/cloudmedia/media"
	"github.com/ali2210/wizdwarf/other/crypto"
	cryptos "github.com/ali2210/wizdwarf/other/crypto"
	wizdate "github.com/ali2210/wizdwarf/other/date_time"
	"github.com/ali2210/wizdwarf/other/parser"
	imglib "github.com/ali2210/wizdwarf/other/parser/parse_image"
	biosubtypes "github.com/ali2210/wizdwarf/other/proteins"
	user "github.com/ali2210/wizdwarf/other/users/register"
	"github.com/biogo/biogo/alphabet"
	"github.com/gorilla/sessions"
	linkcid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
	"github.com/nfnt/resize"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// Error message
const Errors = "Operation Failed"

// These attributes are used to controle execution of processes

// Firestore_Reference hold database refernece
var Firestore_Rf string

// Calendar hold Today time
var calendar string

// Parse_Date Today date
var parse_date string

//  User ID whom will get logged in
var user_Id string

// Meta-picture information
var pic_time string
var pic_tags string

var chain map[string]string
var cdr map[string]string = make(map[string]string, 1)
var ImageMeta *media.IMAGE_METADATA
var link linkcid.Cid

//  Molecular data; hold genomes sequence value
var genes []string

// session cache
//var profiler user.New_User

// Signed key hold paticular state of an object called lock
type SignedKey struct {
	Reader string
	Signed string
	Tx     *ecdsa.PrivateKey
}

// User will add their profile picture on the profile. User shared picture have different parameters such as upload time ; resize image etc.
// User currently add jpeg, png and gif to the profile picture. Profile store on private blockchain. Each Shared Picture key attached to the keys;
// Keys are generated one time; All the keys store in your account meta-wallet .
// @param responsewriter and request object
// @return string and error

func AvatarUpload(r *http.Request, user_id string) (string, error) {

	//  Set the buffer size for the picture file contents
	r.ParseMultipartForm(10 << 50)

	// Get the picture file from HtmlContent
	File, Handler, err := r.FormFile("profile-input")

	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_NOT_FOUND)
		return " ", err
	}
	defer File.Close()

	// Get today year, month and date . This help to generate image metadata which is helpful when images store in the collections.
	// Get Date from html Form
	today := r.FormValue("date")

	// Parse calendar format
	year, err := wizdate.Year(today)
	if err != nil {
		log.Println(error_codes.Operation_ERROR_CODE__DUMP)
		return " ", err
	}

	month, err := wizdate.Month(today)
	if err != nil {
		log.Println("Error parsing month: ", err)
		return " ", err
	}

	date, err := wizdate.Date(today)
	if err != nil {
		log.Println("Error parsing date: ", err)
		return " ", err
	}

	// Calendar Get Today function hold information about picture such when will last avatar changed
	calendar = wizdate.GetToday(year, time.Month(month), date).String()

	// All avatars craeted according to utc timezone
	time_utc := strings.Trim(calendar, "+0000 UTC")

	// As calendar return date, month, year, timezone and time (hour, second & minute). Avatars created in unqiue timespace
	hr, err := strconv.Atoi(time_utc[11:13])
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
		return " ", err
	}

	mns, err := strconv.Atoi(time_utc[14:16])
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
		return " ", err
	}

	// store avatars timespace instance
	create_pic_time := fmt.Sprintf("%d:%d", (hr), (mns))

	// meta-information about picture
	// pic_src = fileHandle.Filename
	parse_date = calendar
	user_Id = user_id
	pic_time = create_pic_time
	var pic_tag string

	// Seed is an upper bound on the number of timeframes available for a avatar
	// rdn.Seed(time.Now().UnixNano())

	// parse_num := strconv.Itoa(rdn.Intn(512))

	if str := strings.Join(parser.ParseTags(Handler.Filename), " "); str != "" {
		if n := strings.Compare(str, " "); n != -1 {
			pic_tag = str
		}
	}

	//hash_color := ""
	// var result interface{}
	// var status int
	var _temp_avatar *os.File
	var width, height int = 200, 200

	// shared document have typesystem or discard shared data
	ok, err := regexp.MatchString(".[a-z]+", Handler.Filename)
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_NAMING_CONVERSION)
		return " ", err
	}

	// scale-up level (SUL); how much scale is needed for shared content
	kvalue := r.FormValue("aspect-ratio")

	value, err := strconv.Atoi(kvalue)
	if err != nil {

		log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
		return " ", err
	}

	if strings.Contains(Handler.Filename, ".png") && ok {

		// unique digital print for shared resources
		_ = imglib.PNG_Color_Hash(&File)

		// store contents in memory
		_temp_avatar, err = parser.CreateFile([]string{Handler.Filename}...)
		if err != nil {

			log.Fatalln(error_codes.File_BAD_REQUEST_CODE_WRITE_FAILED)
			return " ", err
		}

		// re-create the avatar
		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		err := encoder.Encode(_temp_avatar, imglib.GetImageDecoder())
		if err != nil {

			log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
			return " ", err
		}

		// resize avatar image
		resizer := resize.Resize(uint(width+value), uint(height+value), imglib.GetImageDecoder(), resize.Lanczos3)

		encoder = png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(_temp_avatar, resizer)
		if err != nil {

			log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
			return " ", err
		}

		// check whether application have user sessions credentials
		if reflect.DeepEqual(user_id, " ") {
			log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_INPUT)
			return " ", errors.New("user session expired")
		}

		// store in ledger
		// generator := imglib.Metadata(_temp_avatar.Name(), hash_color, user_id, Firestore_Reference())

		// if reflect.DeepEqual(generator, bucket.Err) {

		// 	log.Fatalln(" Error metadata is not created for image")
		// 	return
		// }

		// // read document from ledger
		// result, status = imglib.GetMetadata(hash_color, user_id, Firestore_Reference())

		// var null_interface interface{}
		// if reflect.DeepEqual(result, null_interface) {

		// 	log.Fatalln(" This content already in your bucket! ")
		// 	return
		// }

		// if !reflect.DeepEqual(result, null_interface) && status == 1 {

		// 	log.Fatalln(" Your Avatar metadata is already created", status)
		// 	return
		// }

		words, err := diceware.Generate(12)
		if err != nil {
			return " ", err
		}

		log.Println("Your Passphrase is: ", words, "please save it somewhere on your computer. It'll generated by application one time only (OTP)")

		ImageMeta = &media.IMAGE_METADATA{
			Name:      _temp_avatar.Name(),
			Type:      "png",
			Created:   create_pic_time,
			Tokens:    media.TOKEN_CATEGORY_FUNGIABLE_TOKEN,
			Timeline:  time_utc,
			Tags:      pic_tag,
			MyProfile: true,
			Signature: words,
		}

	} else if strings.Contains(Handler.Filename, ".jpeg") && ok {

		_ = imglib.JPEG_Color_Hash(File)

		_temp_avatar, err = parser.CreateFile([]string{Handler.Filename}...)
		if err != nil {

			log.Fatalln(error_codes.File_BAD_REQUEST_CODE_WRITE_FAILED)
			return " ", err
		}

		op1 := float64(value)

		err = jpeg.Encode(_temp_avatar, imglib.GetImageDecoder(), &jpeg.Options{Quality: int(math.Mod(op1, jpeg.DefaultQuality))})
		if err != nil {

			log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
			return " ", err
		}

		resizer := resize.Resize(uint(width+value), uint(height+value), imglib.GetImageDecoder(), resize.Lanczos3)

		err = jpeg.Encode(_temp_avatar, resizer, &jpeg.Options{Quality: int(math.Mod(op1, jpeg.DefaultQuality))})
		if err != nil {

			log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
			return " ", err
		}

		if reflect.DeepEqual(user_id, " ") {
			log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_INPUT)
			return " ", errors.New("user session expired")
		}

		// generator := imglib.Metadata(_temp_avatar.Name(), hash_color, user_id, Firestore_Reference())

		// if reflect.DeepEqual(generator, bucket.Err) {

		// 	log.Fatalln(" Error metadata is not created for image")
		// 	return
		// }

		// result, status = imglib.GetMetadata(hash_color, user_id, Firestore_Reference())

		// var null_interface interface{}
		// if reflect.DeepEqual(result, null_interface) {

		// 	log.Fatalln(" This content already in your bucket! ")
		// 	return
		// }

		// if !reflect.DeepEqual(result, null_interface) && status == 1 {

		// 	log.Fatalln(" Your Avatar metadata is already created", status)
		// 	return
		// }

		words, err := diceware.Generate(12)
		if err != nil {
			return " ", err
		}

		log.Println("Your Passphrase is: ", words, "please save it somewhere on your computer. It'll generated by application one time only (OTP)")

		ImageMeta = &media.IMAGE_METADATA{
			Name:      _temp_avatar.Name(),
			Type:      "jpeg",
			Created:   create_pic_time,
			Tokens:    media.TOKEN_CATEGORY_FUNGIABLE_TOKEN,
			Timeline:  time_utc,
			Tags:      pic_tag,
			MyProfile: true,
			Signature: words,
		}

	} else if strings.Contains(Handler.Filename, ".gif") && ok {

		_ = imglib.GIF_Color_Hash(File)

		_temp_avatar, err = parser.CreateFile([]string{Handler.Filename}...)
		if err != nil {

			log.Fatalln(error_codes.File_BAD_REQUEST_CODE_WRITE_FAILED)
			return " ", err
		}

		err = gif.Encode(_temp_avatar, imglib.GetImageDecoder(), &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
		if err != nil {

			log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
			return " ", err
		}

		resizer := resize.Resize(uint(width+value), uint(height+value), imglib.GetImageDecoder(), resize.Lanczos3)

		err = gif.Encode(_temp_avatar, resizer, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
		if err != nil {

			log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
			return " ", err
		}

		if reflect.DeepEqual(user_id, " ") {

			log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_INPUT)
			return " ", errors.New("user session expired")
		}

		// generator := imglib.Metadata(_temp_avatar.Name(), hash_color, user_id, Firestore_Reference())

		// if reflect.DeepEqual(generator, bucket.Err) {

		// 	log.Fatalln(" Error metadata is not created for image")
		// 	return
		// }

		// result, status = imglib.GetMetadata(hash_color, user_id, Firestore_Reference())

		// var null_interface interface{}
		// if reflect.DeepEqual(result, null_interface) {

		// 	log.Fatalln(" This content already in your bucket! ")
		// 	return
		// }

		// if !reflect.DeepEqual(result, null_interface) && status == 1 {

		// 	log.Fatalln(" Your Avatar metadata is already created", status)
		// 	return
		// }

		words, err := diceware.Generate(12)
		if err != nil {
			return " ", err
		}

		log.Println("Your Passphrase is: ", words, "please save it somewhere on your computer. It'll generated by application one time only (OTP)")

		ImageMeta = &media.IMAGE_METADATA{
			Name:      _temp_avatar.Name(),
			Type:      "gif",
			Created:   create_pic_time,
			Tokens:    media.TOKEN_CATEGORY_FUNGIABLE_TOKEN,
			Timeline:  time_utc,
			Tags:      pic_tag,
			MyProfile: true,
			Signature: words,
		}

	} else {

		log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
		return " ", err
	}

	OpenSkyChain(ImageMeta)

	defer _temp_avatar.Close()
	return _temp_avatar.Name(), nil
}

// Sky Data Center have following properties. It's provide best solution for Decentralize Content storage ; just like ipfs & pinta;
// Each content store in public ledger which means every content access through key signature called CDR-link.
// CDR_Link is similar to URL , however cdr-link generated based on content that user want to located
// @param skynet client ; file name & file type
// @return boolean
func SkyDataCenter(media_file *media.IMAGE_METADATA) bool {

	// application store user picture in the app_data directory

	if strings.Contains(media_file.Type, " ") && strings.Contains(media_file.Name, " ") {

		log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_INPUT)
		return false
	}

	// if reflect.DeepEqual(client, skynet.SkynetClient{}) {
	// 	log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_INPUT)
	// 	return false
	// }

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	iobject := dbucket.New_Bucket(ctx, media_file.Name, "avatars")

	errs, access := iobject.StoreJCredentials(strings.Join(media_file.Signature, " "), []string{"chief inner hint orient crane mobile pattern rude moon approve train cheap"}...)
	if reflect.DeepEqual(errs, dbucketerror.Bucket_Error_Category_Error) && access != nil {
		log.Fatal(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return false
	}

	var str string
	if strings.HasPrefix(media_file.Name, "app_data/") {
		str = strings.Trim(media_file.Name, "app_data/")
	}

	errs = iobject.StoreObject(strings.Join(media_file.GetSignature(), " "), media_file.Created, str, "", []string{"chief inner hint orient crane mobile pattern rude moon approve train cheap"}...)
	if reflect.DeepEqual(errs, dbucketerror.Bucket_Error_Category_Error) {
		log.Fatal(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return false
	}

	// errs, url := cloudmedia.NewDlinkObject(&client, "app_data/").Generate(file, []string{filetype}...)
	// if errs != dlink.Errors_NONE {
	// 	log.Fatalln(errs, error_codes.Router_ERROR_CODE_EMPTY_RESPONSE)
	// 	return false
	// }

	// high order function this function takes cdr convert into bytes with netwiork identifical code and then apply hash sha-256
	// Result -- breakable
	hash_data := sha256.Sum256([]byte(media_file.Created))

	// this function also high order function which convert hash of bytes into encoded string
	// fog(c) = f(g(x)) mathematical notion of high order function
	// encoded string in hex format which mus be decode as string in hex format
	decoder, err := hex.DecodeString(hex.EncodeToString(hash_data[:]))
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
		return false
	}

	// EncodeName function takes decoder which is already in hex format and then apply x11 crypto algorithm.
	// with x11 breakable signature into unbreakable
	encodetype, err := multihash.EncodeName(decoder, "x11")
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
		return false
	}

	// This is an higher order function encodeName value as encode string in hex format.
	// multihash hex string apply on encoded string in hex format.
	encodex11, err := multihash.FromHexString(hex.EncodeToString(encodetype))
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
		return false
	}

	// generate new cid.. The specification of this function require two parameters (codeType & other one is hash algorithm)
	// merkel tree (dag) data serilaization (protocol buffer [https://en.wikipedia.org/wiki/Protocol_Buffers])
	// & hash algorithm

	cid := linkcid.NewCidV1(linkcid.DagProtobuf, encodex11)

	// check whether cid version is 0. For this application cid version must be 1
	if version := cid.Version(); version != 1 {
		log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
		return false
	}

	// create protocol buffer map object
	// map key which we had calculated before cid because key must be in string
	// value should be cdr link

	cdr = make(map[string]string, 1)

	cdr[cid.String()] = cid.Hash().B58String()
	Set_cdr(cid.String())

	link = cid

	ImageMeta.Cdr = make(map[string]string, 1)
	ImageMeta.Cdr = cdr

	log.Println("Signature created .....", cid.String()[5:10]+"****")

	return true
}

// Download Content is a special function that will get copy of your content
// @param File , client Bucket service
// func Download_Content_ownership(File string, client bucket.Bucket_Service) *proto.QState {

// 	return client.Download(&proto.Query{ByName: File})
// }

// Mapper take stream as interface , key is used to decode the kv value.

/* MapRange
	* iter := reflect.ValueOf(m).MapRange()
	*	for iter.Next() {
	* k := iter.Key()
	* v := iter.Value()
	* ...
    * }
*/

// Mapper is a specified function that will return value aginst paticular key
// @param map set and key
// @return interface (any)
func Mapper(stream map[string]interface{}, key string) interface{} {

	var occ interface{}

	it := reflect.ValueOf(stream).MapRange()
	for it.Next() {
		occ = stream[key]
	}
	return occ
}

// This function provide gateway services. User credentils update through gateway.
// @param Protocol Updated User Message
// @return boolean .
func UpdateProfileInfo(member user.Updated_User) bool {

	_, err := cloud.UpdateUserDetails(GetDBClientRef(), member)
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return false
	}
	return true
}

var CDRL string = ""

// CDR_Link
func Set_cdr(c string) { CDRL = c }

func Get_cdr() string { return CDRL }

// Trust Request ensure that both parties are trusted and ready to complete the pending transaction.
// @param message , validator and request command (string)
// @return edsca public and error
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
			log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
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
			log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
			return false, &AleKey, err
		}

		// bind message with your public key
		bindMessage := cryptos.BSED25519(message)

		// bind message verification against key
		if verify := cryptos.BVED25519(BbKey, bindMessage, []byte(message)); verify {
			log.Fatalln(error_codes.Operation_ERROR_CODE_GARBAGE_VALUE)
			return verify, &AleKey, nil
		}

		// bind message verification failed
		return false, &AleKey, errors.New("error verification error")
	}
}

// This function specified codons that exist in user genetic material
// @param string
// @return map set string
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

// This function specified codons that exist in user genetic material
// @param string
// @return map set Aminochain
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

// This function ensure whole codon either hold three proteins or none
// @param string and int
// @return boolean
func threepairs(str string, i int) bool {
	return !strings.Contains(str[i:i+1], " ") && !strings.Contains(str[i+1:i+2], " ") && !strings.Contains(str[i+2:i+3], " ")
}

// This function act as gateway services which allow to connect application with database . However , this method will only allow search functionality.
// @param responsewriter, user credentials and request object
// @return user credentials  and error message
func Firebase_Gatekeeper(w http.ResponseWriter, r *http.Request, member user.New_User) (*user.New_User, *user.Updated_User, error) {

	docIterator, err := cloud.SearchUser(GetDBClientRef(), member)
	if err != nil && !docIterator.Doc_status {
		log.Fatalln(error_codes.Router_ERROR_CODE_EMPTY_RESPONSE)
		return &user.New_User{}, &user.Updated_User{}, errors.New("no account on our server")
	}

	if !reflect.DeepEqual(docIterator.Update, user.Updated_User{}) {
		return &user.New_User{}, &docIterator.Update, nil
	}

	return &docIterator.Profile, &user.Updated_User{}, nil
}

// This function is another gateway service which allow to create new credentials on our database
// @param responsewriter, request and user credentials
func AddNewProfile(response http.ResponseWriter, request *http.Request, add user.New_User) (*firestore.DocumentRef, bool, error) {

	// var member user.New_User
	var document *firestore.DocumentRef
	var err error

	// user data accrording to json schema
	data, err := json.Marshal(add)
	if err != nil {
		log.Fatalln(error_codes.JSON_CODE_MARSHAL_ERROR)
		return &firestore.DocumentRef{}, false, err
	}

	err = json.Unmarshal(data, &add)
	if err != nil {
		log.Fatalln(error_codes.JSON_CODE_UNMARSHAL_ERROR)
		return &firestore.DocumentRef{}, false, err
	}

	// search data if there is
	candidate, _, err := Firebase_Gatekeeper(response, request, add)
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return &firestore.DocumentRef{}, false, err
	}

	// search data doesn't exist
	if reflect.DeepEqual(candidate, &user.New_User{}) {

		document, _, err = cloud.AddUser(GetDBClientRef(), add)
		if err != nil {
			log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
			return &firestore.DocumentRef{}, false, err
		}

	}

	return document, true, nil

}

// This function allow you to connect with cloud db .
// @return client

func Firestore_Reference() *firestore.Client {

	_, err := os.Stat("config/" + GetKeyFile())
	if os.IsExist(err) {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return &firestore.Client{}
	}

	Firestore_Rf = "config/" + GetKeyFile()

	firebase_connect, err := firebase.NewApp(context.Background(), &firebase.Config{ProjectID: GetProjectID()}, option.WithCredentialsFile(Firestore_Rf))
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PERMISSION_FAILED)
		return &firestore.Client{}
	}

	client, err := firebase_connect.Firestore(context.Background())
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
		return &firestore.Client{}
	}
	return client

}

// String convesion to integer
func StringInt(s string) (int, error) {

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return 0, err
	}
	return i, nil

}

// This function allow you to create session
// @param  string message
// @return session
func Web_Token(unique string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(unique))
}

// This function allow user to store file in application directory
// @param responsewriter, request and file name
func MountDisk(w http.ResponseWriter, r *http.Request, filename string) os.FileInfo {
	f, err := os.OpenFile(filename+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return nil
	}

	finfo, err := f.Stat()
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PATH_ERROR)
		return nil
	}
	return finfo
}

// This function will generate salt bit which will used in validate or transaction
// @param responserwriter, request and messages
// @return string messages and edsca key
func Signx(w http.ResponseWriter, r *http.Request, h1, h2 string) (string, string, *ecdsa.PrivateKey) {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE__DUMP)
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
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return "", "", &ecdsa.PrivateKey{}
	}

	return fmt.Sprintf("0x%x", reader), fmt.Sprintf("0x%x", s), privateKey

}

// @param message string
// @return byte vector and error
func RFiles(filename string) ([]byte, error) {

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return []byte{}, err
	}

	return []byte(body), nil
}

// @param responsewriter, request and user credentials
// @return boolean and signature
func Presence(w http.ResponseWriter, r *http.Request, regexp_emal, regexp_pss bool, add user.New_User) (bool, *SignedKey) {

	code := SignedKey{}

	if !regexp_emal && !regexp_pss {
		return false, &SignedKey{}
	}

	code.Reader, code.Signed, code.Tx = Signx(w, r, hex.EncodeToString([]byte(add.Email)), hex.EncodeToString([]byte(add.Password)))
	return true, &code
}

// Mounted will store content in the app directory to complete futher processing
// @param responsewriter & request
// @return string message & error message

func Mounted(w http.ResponseWriter, r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 50)

	var upldFile *os.File = nil
	file, handler, err := r.FormFile("fileSeq")
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_NOT_FOUND)
		return "", err
	}
	defer file.Close()

	if handler.Size >= (500000 * 1024) {
		log.Fatalln(error_codes.Operation_ERROR_CODE_OVERFLOW_STATE)
		return "", err
	}

	if _, err := os.Stat(handler.Filename); os.IsExist(err) {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PATH_ERROR)
		return "", err
	}

	path, err := os.Stat("app_data/")
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PATH_ERROR)
		return "", err
	}

	if !path.IsDir() {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_DIRECTORY_NOT_FOUND)
		return "", err
	}

	// upload file by user...
	upldFile, err = ioutil.TempFile(filepath.Dir("app_data/"), "apps-"+"*.txt")
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return "", err
	}

	defer upldFile.Close()

	_, err = upldFile.Stat()
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PATH_ERROR)
		return "", err
	}

	openReadFile := upldFile.Name()

	// file convert into bytes
	bytesFile, err := ioutil.ReadAll(file)
	if err != nil {

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return "", err
	}

	_, err = upldFile.Write(bytesFile)
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_WRITE_FAILED)
		return "", err
	}

	return openReadFile, nil

}

/*
*  To keep user privacy , Encrypted-Channels accept stream of data.
*  Onces the data transfer complete this path completely closed.
 */

//  @param file string
// @return boolean
func OpenSkyChain(media_file *media.IMAGE_METADATA) bool {

	/*
	* Create serialized message object [@Protocol-Buffers is used here].
	* This Object create metadata & image file store in database. [@Firestore]
	 */

	// Store user-picture file in the storage directory
	if ok := SkyDataCenter(media_file); ok {

		// picture metadata
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := cloudmedia.NewDc_1(ctx, Firestore_Reference()).PutData(ImageMeta, user_Id); err != nil {
			log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE, err)
			return false
		}

		if bucket_state := fireclient.New(ctx, Firestore_Reference()).Store(link.String(), ImageMeta.Cdr[Get_cdr()], user_Id); bucket_state != 0 {
			log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE, bucket_state)
			return false
		}

		return true
	}

	return false
}

// @param user genetic content and matcher file
// @return string vectors and error message

func ReadAllow(serverFile *os.File, userFile os.FileInfo) ([]string, []string, error) {

	seq, err := RFiles(userFile.Name())
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
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

		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
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

// @param byte message
// @return string message
func ToRunes(seq byte) string {

	if seq >= 65 && seq < 91 {
		return string(alphabet.Letter(seq))
	}
	return string(alphabet.Letter(seq))
}

// calculate risk analysis on the user shared data
// @param responsewriter, request, choices , data and risk algorithm params
// @return error message
func Data_Predicition(w http.ResponseWriter, r *http.Request, fname, choose string, file *os.File, algo info.Levenshtein) error {

	i, err := strconv.Atoi(choose)
	if err != nil {
		log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
		return err
	}
	if (i > 0 && i < 6) && (fname != " ") {

		// data have peristance location address
		svrFile := MountDisk(w, r, fname)

		// read document and convert into managable format for processing
		Usr, Virus, err := ReadAllow(file, svrFile)
		if err != nil {
			log.Fatalln(error_codes.Operation_ERROR_CODE_UNEXPECTED_STATE)
			return err
		}

		// Gene store in the memory
		SetGenes(Usr)

		// attributes properties
		SetEditParameters()

		// calculate matching probability
		distance := GetEditParameters().EditDistanceStrings(Virus, Usr)

		SetBioAlgoParameters(algo.Result(distance, len(Usr)), fname, algo.CalcualtePercentage(algo.GetProbParameter()))

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

// @param directory path and file name message
// @return File and error message
func Open_SFiles(path, filename string) (*os.File, error) {

	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PATH_ERROR)
		return nil, err
	}
	if !fileinfo.IsDir() {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_DIRECTORY_NOT_FOUND)
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return nil, err
	}
	defer file.Close()
	return file, nil
}

// GetDocuments is a specified algorithm that will copy image from public ledger
// @param vector string of session
// @return string and interface
func GetDocuments(session_id ...string) (string, interface{}, []string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var newInter interface{}
	var mapper map[string]interface{}
	var err error

	mapper, err = cloudmedia.NewDc_1(ctx, Firestore_Reference()).GetData(&media.IMAGE_METADATA{
		Name:      "",
		Type:      "",
		Created:   "",
		Tokens:    media.TOKEN_CATEGORY_FUNGIABLE_TOKEN,
		MyProfile: false,
		Cdr:       cdr,
		Timeline:  "",
		Tags:      pic_tags,
		Signature: []string{},
	}, session_id[0])
	if err != nil {
		log.Fatalln(error_codes.Router_ERROR_CODE_EMPTY_RESPONSE)
		return " ", newInter, []string{}
	}

	if reflect.DeepEqual(mapper, map[string]interface{}{}) {
		log.Fatalln(error_codes.Router_ERROR_CODE_EMPTY_RESPONSE)
		return " ", newInter, []string{}
	}

	lists := reflect.ValueOf(mapper).MapRange()
	var name string = ""
	var dlinks interface{}
	sign := make([]string, 12)

	for lists.Next() {

		if reflect.DeepEqual(lists.Key().String(), "Name") {
			name = lists.Value().Elem().String()
		}

		if reflect.DeepEqual(lists.Key().String(), "CDR_LINK") {

			dlinks = lists.Value().Interface()
		}

		if reflect.DeepEqual(lists.Key().String(), "Passphrase") {

			sign, _ = lists.Value().Interface().([]string)
		}
	}

	return name, dlinks, sign

}

// Get Documents return in interface which is used only when there a function which will extract values from set. Here ReflectMaps come in action
// @param interface
// @return string values
func ReflectMaps(i interface{}) (string, string) {

	var nullify interface{}
	if reflect.DeepEqual(i, nullify) {
		log.Fatalln(error_codes.Operation_ERROR_CODE_EMPTY_OUTPUT)
		return "", ""
	}

	reflex := reflect.ValueOf(i).MapRange()
	var key, value string = "", ""

	for reflex.Next() {

		key = reflex.Key().Interface().(string)
		value = reflex.Value().Interface().(string)
	}

	return key, value
}

// Read File size
// @param string as final
// @return string message
func GetFileSize(filename ...string) string {

	_, err := os.ReadDir("app_data/")
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_DIRECTORY_NOT_FOUND)
		return " "
	}

	properties, err := os.Stat(filename[0])
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
		return " "
	}

	return fmt.Sprintf("%d", properties.Size()/1024) + "KiB"
}

func GetFileCreationTime(filename string) string {

	var properties fs.FileInfo
	var err error

	files, err := os.ReadDir("app_data/")
	if err != nil {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_DIRECTORY_NOT_FOUND)
		return ""
	}

	for entries, _ := range files {

		if strings.Contains("app_data/"+files[entries].Name(), filename) {

			properties, err = os.Stat("app_data/" + files[entries].Name())
			if err != nil {
				log.Fatalln(error_codes.File_BAD_REQUEST_CODE_READ_FAILED)
				return ""
			}
		}
	}

	if reflect.DeepEqual(properties, nil) {
		log.Fatalln(error_codes.File_BAD_REQUEST_CODE_FILE_PATH_ERROR)
	}
	return properties.ModTime().String()
}
