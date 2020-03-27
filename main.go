package main



import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"	
	"log"
	"io/ioutil"
	"os"
)

	type Response struct{
		id int 
		flag bool
	}

func main(){

	routing := mux.NewRouter()

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
				println(file)
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
	temp.Execute(w,"Regsiter")
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
					upldFile , err := ioutil.TempFile("user_data", "myFiles-*.txt"); if err != nil{
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
