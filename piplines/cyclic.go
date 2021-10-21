package piplines

//  one-way call

import (
	//firebase "firebase.google.com/go"
	"reflect"

	"cloud.google.com/go/firestore"
	bio "github.com/ali2210/wizdwarf/other/bioinformatics"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/users"
)

var (
	appName *firestore.Client
	cloud   users.DBFirestore
	//Firestore_Rf      string
	Edit bio.LevenTable
	Algo info.Levenshtein
)

const (
	ConfigFilename string = "htickets-cb4d0-firebase-adminsdk-orfdf-b3528d7d65.json"
	ProjectID      string = "htickets-cb4d0"
)

func SetDBClientRef() *firestore.Client {
	appName = Firestore_Reference()
	return appName
}

func GetDBClientRef() *firestore.Client {
	if appName == (&firestore.Client{}) {
		return (&firestore.Client{})
	}
	return appName
}

func SetDBCollect() users.DBFirestore {
	cloud = users.NewCloudInstance()
	return cloud
}

func GetDBCollect() users.DBFirestore {
	var collect users.DBFirestore
	if reflect.DeepEqual(cloud, collect) {
		return collect
	}
	return cloud
}

func GetKeyFile() string {
	if ConfigFilename == "" {
		return ""
	}
	return ConfigFilename
}

func GetProjectID() string {
	if ProjectID == "" {
		return ""
	}
	return ProjectID
}

func GetEditParameters() bio.LevenTable {
	var matx bio.LevenTable
	if reflect.DeepEqual(Edit, matx) {
		return matx
	}
	return Edit
}

func SetEditParameters() bio.LevenTable {
	Edit = bio.NewMatch()
	return Edit
}

func SetBioAlgoParameters(prob float32, pattern_name string, percent float32) {
	Algo = info.Levenshtein{Probablity: prob, Name: pattern_name, Percentage: percent}
}

func GetBioAlgoParameters() info.Levenshtein {
	if Algo == (info.Levenshtein{}) {
		return info.Levenshtein{}
	}
	return Algo
}
