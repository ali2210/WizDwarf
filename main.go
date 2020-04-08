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
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref" 
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

	var emailexp string = "([A-Z][a-z]|[0-9])*[@][a-z]*"
	var passexp string = "([A-Z][a-z]*[0-9])*"
	var ClientDB *mongo.Client

func main(){

	routing := mux.NewRouter()
	
	DB_Client()
	//println("client update:", ClientDB)
	routing.HandleFunc("/{title}/home", Home)
	routing.HandleFunc("/{title}/signup", NewUser)
	routing.HandleFunc("/{title}/login", Existing)
	routing.HandleFunc("/dummy", Dump)
	
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
					 FileReadFromDisk(name);
				case "2":
					var name string = "FlaviDengue"
					 FileReadFromDisk(name);
				case "3":
					var name string = "KenyaEbola"
					 FileReadFromDisk(name);
				case "4":
					var name string = "ZikaVirusBrazil"
					 FileReadFromDisk(name);
				case "5":
					var name string = "MersSaudiaArabia"
					 FileReadFromDisk(name);
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
			hashRet := MessageToHash(matchE,matchP,user); if hashRet == false{
				fmt.Fprintf(w, "Sorry provided data must not match with rules\n. Email must be in Upper or Lower case or some digits, while password must contain Uppercase Letter , lowercase letter")
				temp.Execute(w,"Regsiter")
			}


			println("Gender:" , user.sir)
			println("Gender2:", user.madam)
			if r.FormValue("check") == "on"{
					user.check_me_out = true
			}else{
				user.check_me_out = false
			}
			
			println("check:" , user.check_me_out)
			println("User record:" , user.name, user.email)
			temp.Execute(w,"Regsiter")
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
	r.ParseMultipartForm(10 << 50)
			file , handler, err := r.FormFile("fileSeq"); if err != nil{
				fmt.Println("Error failed.... retry",err)
				return nil
			}
			defer file.Close()
				if(handler.Size <= (50 * 1024)){
					fmt.Println("File name:" + handler.Filename)
					if _, err := os.Stat(handler.Filename);os.IsExist(err){
						fmt.Println("FILE exist ", err)
					}
					upldFile , err := ioutil.TempFile("user_data", handler.Filename+"-*.txt"); if err != nil{
					fmt.Println("Error received!", err)
				}
				defer upldFile.Close()
				// file convert into bytes
				bytesFile , err := ioutil.ReadAll(file); if err != nil{
					fmt.Println("Error received!", err)
				}
				
				upldFile.Write(bytesFile)
				fmt.Println("File added on server")
					return upldFile
				}
				return nil
}

func FileReadFromDisk(filename string){
	f , err := os.OpenFile(filename + ".txt", os.O_RDWR | os.O_CREATE, 0755); if err != nil{
		println("FILE Open Error ... " , err)
	}
	println("File Exist...", f)
	finfo, err := f.Stat() ; if err != nil{
		println("File Info not found" , err)
	}
	println("File Info" , finfo.Name())
}

func MessageToHash(matchE, matchP bool , user Create_User) (bool){
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
				return true
	}
	return false
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
	return fmt.Sprintf("0x%x\n", r), fmt.Sprintf("0x%x\n", s)
	
}

func DB_Client(){
	client , err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://enigma:wisdom@enigma-hwykp.mongodb.net/test?retryWrites=true&w=majority")); if err != nil{
		log.Fatal(err)
	}

	ctx , _ := context.WithTimeout(context.Background() , 10 * time.Second)
	err = client.Connect(ctx); if err != nil{
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary()); if err != nil{
	 	log.Fatal(err)
	}
	database , err := client.ListDatabaseNames(ctx, bson.M{}); if err != nil{
		log.Fatal(err)
	}
	println("Welcome to MongoDB....", database)
	// collection := con.Database("AppsDB").Collection("Users")
	// println("collection name : " , collection)
}





