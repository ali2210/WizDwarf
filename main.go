package main



import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"	
	"log"
	"io/ioutil"
	"os"
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
	"google.golang.org/api/option"
	 "./db"
	 "encoding/json"
	 // "github.com/golang/gddo/httputil/header"
	 // "errors"
	 // "io"
	 // "strings"
	// "cloud.google.com/go/storage"
	// cloudkms "google.golang.org/api/cloudkms/v1"
	// "google.golang.org/api/iterator"
	 

)


	// Struts 

	type Response struct{
		id int 
		flag bool
	}

	type Create_User struct{
		name string 
		fname string
		sir bool
		madam bool 
		address string   // World Coodinates
		address2 string  // local coodinates
		zip string
		city string
		country string
		check_me_out bool
		email string
		password string
	}
	type SignedKey struct{
		reader string
		signed string
	}


	// Variables 

	var (
		emailexp string = "([A-Z][a-z]|[0-9])*[@][a-z]*"
		passexp string = "([A-Z][a-z]*[0-9])*"
		AppName *firebase.App = SetFirestoreCredentials() // Google_Cloud [Firestore_Reference] 
		cloud db.DBFirestore = db.NewCloudInstance()
		
	)

	const(
				projectId string = "htickets-cb4d0"	
				Google_Credentials string = "/home/ali/Desktop/htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	)

	// Functions 

func main(){

	// Routing
	routing := mux.NewRouter()
	
	routing.HandleFunc("/{title}/home", Home)
	routing.HandleFunc("/{title}/signup", NewUser)
	routing.HandleFunc("/{title}/login", Existing)
	// routing.HandleFunc("/{title}/action", addVistor)
	routing.HandleFunc("/{title}/data", getVistor)
	routing.HandleFunc("/dummy", Dump)
		 // DB_Client()
	
	log.Println("Listening at 9101 ... please wait...")
	http.ListenAndServe(":9101",routing)


}

func Home(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("index.html"))
	if r.Method  ==  "GET"{
		fmt.Println("Method:" + r.Method)
		temp.Execute(w,"Home")
	}else{
		temp := template.Must(template.ParseFiles("dump.html"))
		    r.ParseForm()
			fmt.Println("Url:", r.URL.Path)
			fmt.Println("Method:" + r.Method)
			

			// FILE Upload ....
			file := UploadFiles(r); if file != nil{
				println(file) // user file upload
				choose := r.FormValue("choose")
				println("choose I make:", choose)
				switch choose{
				case "0":
					fmt.Fprintf(w, "Please choose any option ...")
					temFile := template.Must(template.ParseFiles("index.html"))
					temFile.Execute(w,"Home")
				case "1":
					var name string = "Covid-19"
					svrFile := FileReadFromDisk(name)
					println("Please Wait", svrFile.Name(), "...")
					SequenceAligmentTable(file, svrFile)			

				case "2":
					var name string = "FlaviDengue"
					svrFile := FileReadFromDisk(name)
					SequenceAligmentTable(file,svrFile)

				case "3":
					var name string = "KenyaEbola"
					svrFile := FileReadFromDisk(name)
					println("Please Wait", svrFile.Name(), "...")
					SequenceAligmentTable(file,svrFile)

				case "4":
					var name string = "ZikaVirusBrazil"
					svrFile := FileReadFromDisk(name)
					println("Please Wait", svrFile.Name(), "...")
					SequenceAligmentTable(file,svrFile)

				case "5":
					var name string = "MersSaudiaArabia"
					svrFile := FileReadFromDisk(name)
					println("Please Wait", svrFile.Name(), "...")
					SequenceAligmentTable(file,svrFile)

				default:
					temFile := template.Must(template.ParseFiles("index.html"))
					temFile.Execute(w,"Home")
				}
			}else{
				print("size must be less than 5KB")
				serverResponse := Response {0, true}
				println("Server Response:", serverResponse.id, serverResponse.flag)
				temp.Execute(w,serverResponse)
			} 

	}

}

func NewUser (w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("register.html"))
	user := Create_User{}
		if r.Method  ==  "GET"{
			fmt.Println("Method:" + r.Method)
			temp.Execute(w,"Regsiter")
		}else{
			r.ParseForm()
			fmt.Println("Url:", r.URL.Path)
			fmt.Println("Method:" + r.Method)
			user.name = r.FormValue("uname")
			user.fname = r.FormValue("ufname")
			user.address = r.FormValue("address")
			user.address2 = r.FormValue("address2")
			user.city = r.FormValue("city")
			user.country = r.FormValue("country")
			user.zip = r.FormValue("zip")
			user.email = r.FormValue("email")
			user.password = r.FormValue("password")
			if r.FormValue("sir") == "on"{
				user.sir = true 
			}else if r.FormValue("madam") == "on"{
				user.madam = true
			} else{
				fmt.Fprintf(w,"Select any option")
				user.sir = false
				user.madam = false
				temp.Execute(w,"Regsiter")
			}

			matchE,err := regexp.MatchString(emailexp, user.email); if err != nil{
				println("invalid regular expression", err)
			}
			println("regexp_email:", matchE)
			matchP,err := regexp.MatchString(passexp, user.password); if err != nil{
				println("invalid regular expression", err)
			}
			println("regexp_pass:", matchP)

			
			// security 
			hashRet, encrypted := MessageToHash(matchE,matchP,user); if hashRet == false{
				fmt.Fprintf(w, "Sorry provided data must not match with rules\n. Email must be in Upper or Lower case or some digits, while password must contain Uppercase Letter , lowercase letter")
				temp.Execute(w,"Regsiter")
			}
			 println("encryted data", encrypted.reader)


			println("Gender:" , user.sir)
			println("Gender2:", user.madam)
			if r.FormValue("check") == "on"{
					user.check_me_out = true
			}else{
				user.check_me_out = false
			}
			
			println("check:" , user.check_me_out)
			println("User record:" , user.name, user.email)
			addVistor(w,r,&user, encrypted.reader)
			// temp.Execute(w,"Regsiter")
		}

}

func Existing(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("login.html"))
	temp.Execute(w,"Login")
}


func Dump(w http.ResponseWriter, r *http.Request){
	temp := template.Must(template.ParseFiles("dump.html"))
	temp.Execute(w,"Dump")
}

func UploadFiles(r *http.Request)(*os.File){
	println("request body",r.Body)
	r.ParseMultipartForm(10 << 50)
			file , handler, err := r.FormFile("fileSeq"); if err != nil{
				fmt.Println("Error failed.... retry",err)
				return nil
			}
			defer file.Close()
				if(handler.Size <= (50 * 1024)){
					fmt.Println("File name:" + handler.Filename)
					if _, err := os.Stat(handler.Filename);os.IsExist(err){
						fmt.Println("File not exist ", err)
					}
					upldFile , err := ioutil.TempFile("user_data", handler.Filename+"-*.txt"); if err != nil{
					fmt.Println("Error received while uploading!", err)
				}
				defer upldFile.Close()
				// file convert into bytes
				bytesFile , err := ioutil.ReadAll(file); if err != nil{
					fmt.Println("Error received while reading!", err)
				}
				
				upldFile.Write(bytesFile)
				fmt.Println("File added on server")
					return upldFile
				}
				return nil
}


// func implicit(){
// 	project := "project-id"
// 	ctx := context.Background()

// 	storageClient , err := storage.NewClient(ctx); if err != nil{
// 		log.Fatal("Error", err)
// 	}
// 	it := storageClient.Buckets(ctx, project)
// 	for{
// 		BucktsAtrr , err := it.Next(); if err == iterator.Done {
// 			break
// 		}
// 		if err != nil{
// 			log.Fatal("Error in BucktsAtrr", err)
// 		}
// 		fmt.Println(BucktsAtrr.Name)
// 	}

// 	kmsService, err := cloudkms.NewService(ctx)
//         if err != nil {
//                 log.Fatal(err)
//         }

//         _ = kmsService

// }

func FileReadFromDisk(filename string)os.FileInfo{
	f , err := os.OpenFile(filename + ".txt", os.O_RDWR | os.O_CREATE, 0755); if err != nil{
		println("FILE Open Error ... " , err)
	}
	println("File Exist...", f)
	finfo, err := f.Stat() ; if err != nil{
		println("File Info not found" , err)
	}
	println("File Info" , finfo.Name())
	return finfo
}

func MessageToHash(matchE, matchP bool , user Create_User) (bool,  *SignedKey){
	code := SignedKey{}
	if matchE && matchP{
				h := sha256.New()
				// h.Write([]byte(user.email))
				hashe := h.Sum([]byte(user.email))
				fmt.Println("email:",hex.EncodeToString(hashe))

				h1 := sha256.New()
				// h1.Write([]byte(user.password))
				hashp := h1.Sum([]byte(user.password))
				fmt.Println("pass:",hex.EncodeToString(hashp))
				code.reader , code.signed = Key(hex.EncodeToString(hashe), hex.EncodeToString(hashp))
				println("data get :", code.reader, code.signed)
				return true, &code
	}
	return false, &code
}

func Key(h1 , h2 string)(string , string){

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader); if err != nil {
		panic(err)
	}

	msg := h1 + h2
	hash := sha256.Sum256([]byte(msg))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:]);if err != nil {
		panic(err)
	}
	fmt.Printf("signature : (0x%x 0x%x)\n", r, s)
	return fmt.Sprintf("0x%x", r), fmt.Sprintf("0x%x", s)
	
}
func SetFirestoreCredentials()*firebase.App{

	// set credentials
	conf := &firebase.Config{ProjectID:projectId}
	opt := option.WithCredentialsFile(Google_Credentials)
	app , err := firebase.NewApp(context.Background(), conf, opt); if err != nil{
		log.Fatal("Error in Connection with Firestore", err)
	}
	println("Connected... Welcome to Firestore")
	return app
}




func getVistor(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type", "application/json")
	visitor , err := cloud.FindAllData(AppName); if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" :"Error getting visitor result"}`))
	}

	// response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(visitor)

}

func addVistor(response http.ResponseWriter, request *http.Request ,user *Create_User, im string){
	response.Header().Set("Content-Type", "application/json")
	var member db.Vistors 

	err := json.NewDecoder(request.Body).Decode(&member); if err == nil{
		fmt.Printf("Error %v: " , err )
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error" :"Error marshal "}`))
	}
	println("Member", &member)
	 member.Id = im
	 member.Name = user.name
	 member.Email= user.email
	 member.Password = user.password
	 cloud.SaveData(&member,AppName)
	 // response.WriteHeader(http.StatusOK)
	 json.NewEncoder(response).Encode(member)


	//println("Vistors:" , p.Id)


	// if request.Header.Get("Content-Type") != ""{
	// 	value , _ := header.ParseValueAndParams(request.Header, "Content-Type")
	// 	println("Value:", value)
	// 	if value != "application/json"{
	// 		msg := "Content-type header is not application/json"
	// 		http.Error(response, msg , http.StatusUnsupportedMediaType)	
	// 	}
	// }

	// var data = []byte(`[
	// 	{
	// 	"Id" : "00x", 
	// 	"Name" : "Ali", 
	// 	"Email" : "alideveloper95@gmail.com", 
	// 	"Password" : "0000"
	// }
	// ]`)

	// type Profile struct{
	// 	Id string
	// 	Name string
	// 	Email string
	// 	Password string
	// }
	// var p []Profile
	// 
	// fmt.Printf("id%+v:name%v", p.Id, p.Name)




	// request.Body = http.MaxBytesReader(response, request.Body, 1048576)
	// unknown := json.NewDecoder(request.Body)
	 // _ , err := ioutil.ReadAll(request.Body); if err != nil{
		// println("Error report:", err)
	 // }

	// unknown.DisallowUnknownFields() 
	// var vistor db.Vistors
	// err = unknown.Decode(&vistor); if err != nil{ 
	// 	println("error :" , err)
	// 	var syntxError *json.SyntaxError
	// 	var unmarshalTypeError *json.UnmarshalTypeError
	// 	switch {
	// 	case errors.As(err, &syntxError):
	// 		msg := fmt.Sprintf("maclious formed json body%d", syntxError.Offset)
	// 		http.Error(response, msg, http.StatusBadRequest)
	// 	case errors.Is(err, io.ErrUnexpectedEOF):
	// 		msg:= fmt.Sprintf("Request contain invalid value for the field%q,%d",unmarshalTypeError.Field, unmarshalTypeError.Offset)
	// 		http.Error(response, msg, http.StatusBadRequest)
	// 	case strings.HasPrefix(err.Error(), "json: unknown field"):
	// 		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
	// 		msg := fmt.Sprintf("Request body contain unknown field %s", fieldName)
	// 		http.Error(response, msg, http.StatusBadRequest)
	// 	case errors.Is(err,io.EOF):
	// 		msg := fmt.Sprintf("Request body must not be empty")
	// 		http.Error(response, msg, http.StatusBadRequest)
	// 		fmt.Printf("Error:%v", err)
	// 	case err.Error() == "http: request body too large":
	// 		msg := "Request body must not larger than 1 MB"
	// 		http.Error(response, msg, http.StatusRequestEntityTooLarge)
	// 	default:
	// 		log.Println(err.Error())
	// 		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	}
	// 	return 
	// }
	// println("Data:", vistor.Id)

	

	

	// println("Body:", req)
	
	//response.WriteHeader(http.StatusInternalServerError)
	// 	// response.Write([]byte(`{"error" :"Error unmarshal Data}`))
	// 	http.Error(response, err.Error(), http.StatusBadRequest)
	// 	return 
	// }
}

func ReadSequence(filename string)([]byte, error){
	body ,err := ioutil.ReadFile(filename); if err != nil{
		return nil, err
	}
	// fmt.Printf("content %s:", body)
	return []byte(body), nil
}

func SequenceAligmentTable(serverFile *os.File, userFile os.FileInfo){ 
	
		//local variables but large life span
	// list, listDna := GoStructs.List{}, GoStructs.List{}
	// queue, queueDna := GoStructs.QueueList{}, GoStructs.QueueList{} 

		// local variable liitle scoope
	seq , err := ReadSequence(userFile.Name()); if err != nil{
		println("Error in read file", err)
	}
	// fmt.Printf("Seq string:%s\n", seq)
	Useq , err := ReadSequence(serverFile.Name()); if err != nil{
		println("Error in read file", err)
	}


	println("Virus Dna sequence :")
	
	for _, v := range seq{
		// fmt.Printf("Seq:%v \t",  v ) // print bytes of array
		space := DoAscii(v); if space == "---"{
			 fmt.Printf("%s\t", space)
			// queueDna = queueDna.Enque(space)
		}
		// queueDna = queueDna.Enque(space)
		// listDna = listDna.Add(queueDna)
		// println("List Ref:",&listDna)
		// queueDna.Print()
		 // fmt.Printf("%s\t", space)
	}
	println("Your Dna sequence :")
	for _, v := range Useq{
		uDna := DoAscii(v); if uDna == "---"{
			fmt.Printf("%s", uDna)
			// queue = queue.Enque(uDna)
		}
		// fmt.Printf("%s\t", uDna)
		// queue  = queue.Enque(uDna)
		// list=list.Add(queue)
		// println("List Ref:",&list,  leng)
		// queue.Print()
		// ele := queue.DeQueue()
		// fmt.Println("Data we get :", ele )
	}                    
}

func DoAscii(seq byte) string{
		if seq >= 65 && seq < 91{
		 return string (seq)
		}
		return "---"
}



