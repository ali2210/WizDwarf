package piplines

//  one-way call

import(
	//firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"github.com/ali2210/wizdwarf/structs/users"
	bio "github.com/ali2210/wizdwarf/structs/bioinformatics"
	info "github.com/ali2210/wizdwarf/structs/bioinformatics/model"
	"reflect"
)

var(
	AppName           *firestore.Client
	Cloud             users.DBFirestore
	Firestore_Rf      string
	Edit              bio.LevenTable
	Algo              info.Levenshtein
)

const(
	ConfigFilename string = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	ProjectID      string = "htickets-cb4d0"
)

func SetDBClientRef() *firestore.Client {
	AppName = Firestore_Reference()
	return AppName
}

func GetDBClientRef() *firestore.Client{ 
	if AppName == (&firestore.Client{}) {return (&firestore.Client{})}
	 return AppName
}

func SetDBCollect() users.DBFirestore {
	Cloud = users.NewCloudInstance()
	return Cloud
}

func GetDBCollect() users.DBFirestore {
	var collect users.DBFirestore 
	if reflect.DeepEqual(Cloud, collect){ return collect }
	 return Cloud
}

func GetKeyFile() string{ 
	if ConfigFilename == ""{ return ""} 
	return ConfigFilename
}

func GetProjectID() string{
	if ProjectID == "" { return ""}
	return ProjectID 
}

func GetEditParameters() bio.LevenTable{
	var matx bio.LevenTable
	if reflect.DeepEqual(Edit, matx){
		return matx
	}
	return Edit
}

func SetEditParameters() bio.LevenTable{
	Edit = bio.NewMatch()
	return Edit
}

func SetBioAlgoParameters(prob float32, pattern_name string, percent float32){
	Algo = info.Levenshtein{Probablity : prob, Name : pattern_name, Percentage : percent}
}

func GetBioAlgoParameters() info.Levenshtein{
	if Algo == (info.Levenshtein{}){ return info.Levenshtein{}}
	return Algo	
}
