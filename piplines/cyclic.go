package piplines

//  one-way call

import (
	//firebase "firebase.google.com/go"

	"errors"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	vault_wraper "github.com/WisdomEnigma/vault-keygen/vault"
	bio "github.com/ali2210/wizdwarf/other/bioinformatics"
	info "github.com/ali2210/wizdwarf/other/bioinformatics/model"
	"github.com/ali2210/wizdwarf/other/users"
	"github.com/hashicorp/vault/api"
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

	if reflect.DeepEqual((info.Levenshtein{}), Algo) {
		return info.Levenshtein{}
	}
	return Algo
}

func Extractor(in, pattern string) string {

	return strings.Trim(in, pattern)
}

type HCLDeclaration struct {
	Weatherapi  string `hcl:"Weatherapi"`
	Channel_key string `hcl:"Channel_key"`
	Channel_id  string `hcl:"Channel_id"`
	Secret      string `hcl:"Secret"`
	Cluster_ID  string `hcl:"Cluster_ID"`
	Token_Auth  string `hcl:"Token_Auth"`
}

var v_wrapper vault_wraper.Vault_Services

func PutKV(object *HCLDeclaration, path string, client *api.Client) error {

	if reflect.DeepEqual(object, &HCLDeclaration{}) {
		return errors.Unwrap(errors.New("object field must not be empty"))
	}

	if reflect.DeepEqual(client, &api.Client{}) {
		return errors.Unwrap(errors.New("vault client is not running in background or vault credentials had not submit yet"))
	}

	v_wrapper = vault_wraper.NewClient(client)

	_, err := v_wrapper.SaveKeygen(vault_wraper.Keygen{
		Vault_path: path,
		Vault_record: map[string]interface{}{
			"data": map[string]interface{}{
				"Weatherapi":  object.Weatherapi,
				"Channel_key": object.Channel_key,
				"Channel_id":  object.Channel_id,
				"Secret":      object.Secret,
				"Cluster_ID":  object.Cluster_ID,
			},
		},
	})

	return err
}

func GetKV(path string) (interface{}, error) {

	keygen, err := v_wrapper.GetKeygen(vault_wraper.Keygen{
		Vault_path:   path,
		Vault_record: nil,
	})
	if err != nil {
		return nil, err
	}

	if reflect.DeepEqual(keygen.Data["data"], map[string]interface{}{}) {
		return map[string]interface{}{}, errors.Unwrap(errors.New("secrets vault is empty"))
	}

	return keygen.Data["data"], nil
}
