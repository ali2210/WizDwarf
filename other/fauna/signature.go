package fauna

import (
	"errors"
	"math"
	"reflect"

	"github.com/ali2210/wizdwarf/other/fauna/store"
	openledge "github.com/fauna/faunadb-go/v4/faunadb"
)

const WizNetHook string = "fnAEq5bDlFACSTAKfEKmy1fDdwVblUOYvSYO04pF"
const DATABASE_NAME string = "OpenNetLedger"
const COLLECTION_NAME string = "OpenLedger"
const Index_NAME string = "Owners"
const DOC_NAME string = "OweAssets"

var Document_Size int64 = int64(math.Pow(2, 5))
var _ledge_conn *openledge.FaunaClient
var count int64 = -1

type Digital_Signature_Service interface {
	ConnectFaunaLedger(endpoint *store.Endpoint_Info) store.Error
}
type Digital_Signature struct {
	Public_Address  string
	Private_Address string
}

func Content_X(addr, _private_addr string) Digital_Signature_Service {
	return &Digital_Signature{Public_Address: addr, Private_Address: _private_addr}
}

func (s *Digital_Signature) ConnectFaunaLedger(endpoint *store.Endpoint_Info) store.Error {

	client := openledge.NewFaunaClient((*endpoint).API_Key, openledge.Endpoint((*endpoint).Address))
	_set_Client_Ref(client)

	if _led_Ref := _get_Client_Ref(); reflect.DeepEqual(_led_Ref, openledge.FaunaClient{}) {
		return store.Error{Description: "_ledge_conn failed to connect with" + (*endpoint).Address, State: store.Error_ERROR}
	}

	_create_db(endpoint)

	clientSession, err := _new_Session(endpoint)
	if reflect.DeepEqual(err.State, store.Error_ERROR) {
		return store.Error{Description: err.Description, State: store.Error_ERROR}
	}

	ok, errs := _new_collection(clientSession)
	if !ok && errs != nil {
		return store.Error{Description: errs.Error(), State: store.Error_ERROR}
	}

	if err := _new_index(clientSession); reflect.DeepEqual(err.State, store.Error_ERROR) {
		return store.Error{Description: err.Description, State: store.Error_ERROR}
	}

	index := func(c int64) int64 {
		if c != Document_Size {
			c += 1
		}
		return c
	}

	if err := _new_data(clientSession, &store.Digital_Signature{Private_Address: s.Private_Address, Public_Address: s.Public_Address}, index(count)); reflect.DeepEqual(err.State, store.Error_ERROR) {
		return store.Error{Description: err.Description, State: store.Error_ERROR}
	}

	if err := _new_doc(clientSession); reflect.DeepEqual(err.State, store.Error_ERROR) {
		return store.Error{Description: err.Description, State: store.Error_ERROR}
	}

	return store.Error{Description: "Operation successful!, Check out your content signature wallet...", State: store.Error_OK}
}

func _set_Client_Ref(conn *openledge.FaunaClient) {

	_ledge_conn = conn
}

func _get_Client_Ref() *openledge.FaunaClient { return _ledge_conn }

func _create_db(endpoint *store.Endpoint_Info) (openledge.Value, store.Error) {

	var value openledge.Value
	if _led_Ref := _get_Client_Ref(); reflect.DeepEqual(_led_Ref, openledge.FaunaClient{}) {
		return value, store.Error{Description: "Application Ledger is not connected with" + (*endpoint).Address, State: store.Error_ERROR}
	}

	value, err := _get_Client_Ref().Query(openledge.If(openledge.Exists(openledge.Database(DATABASE_NAME)), true, openledge.CreateDatabase(openledge.Obj{
		"name": DATABASE_NAME,
	})))
	if err != nil {
		return value, store.Error{Description: DATABASE_NAME + "is not created on the machine" + (*endpoint).Address, State: store.Error_ERROR}
	}

	if value != openledge.BooleanV(true) {
		return value, store.Error{Description: "Connection handshake is not possible", State: store.Error_ERROR}
	}

	return value, store.Error{Description: DATABASE_NAME + "is created on the machine", State: store.Error_OK}
}

func _new_Session(endpoint *store.Endpoint_Info) (*openledge.FaunaClient, store.Error) {

	objv, err := _get_Client_Ref().Query(openledge.CreateKey(openledge.Obj{"name": openledge.Database(DATABASE_NAME), "role": "server"}))
	if err != nil {
		return &openledge.FaunaClient{}, store.Error{Description: COLLECTION_NAME + "is not created on the machine... make sure application is connected with a secure connection" + (*endpoint).Address, State: store.Error_ERROR}
	}

	object := ""
	const key string = "secret"

	err = objv.At(openledge.ObjKey(key)).Get(&object)
	if err != nil {
		return &openledge.FaunaClient{}, store.Error{Description: err.Error(), State: store.Error_ERROR}
	}

	return _get_Client_Ref().NewSessionClient(object), store.Error{Description: DATABASE_NAME + "secure connection ", State: store.Error_OK}
}

func _new_collection(client *openledge.FaunaClient) (bool, error) {

	collect, err := client.Query(openledge.If(openledge.Exists(openledge.Collection(COLLECTION_NAME)), true, openledge.CreateCollection(openledge.Obj{
		"name": COLLECTION_NAME,
	})))
	if err != nil {
		return false, err
	}

	if collect != openledge.BooleanV(false) {
		return true, nil
	}

	return false, errors.New("collection is not initialized")
}

func _new_index(client *openledge.FaunaClient) store.Error {

	_, err := client.Query(openledge.If(openledge.Exists(openledge.Index(Index_NAME)), true, openledge.CreateIndex(openledge.Obj{
		"name":   Index_NAME,
		"source": openledge.Collection(COLLECTION_NAME),
		// "unique":     "true",
		// "serialized": "true",
		// "terms":      []string{"data", "name"},
	})))

	if err != nil {
		return store.Error{Description: err.Error(), State: store.Error_ERROR}
	}

	// if reflect.DeepEqual(openledge.BooleanV(false), _index) {
	// 	return store.Error{Description: "index is not created ", State: store.Error_ERROR}
	// }

	return store.Error{Description: Index_NAME + "created on the machine", State: store.Error_OK}
}

func _new_data(client *openledge.FaunaClient, attached *store.Digital_Signature, _index_value int64) store.Error {

	attachment_key, err := client.Query(openledge.Create(openledge.Ref(openledge.Collection(COLLECTION_NAME), _index_value), openledge.Obj{
		"data": attached,
	}))
	if err != nil {
		return store.Error{Description: err.Error(), State: store.Error_ERROR}
	}

	_, err = client.Query(openledge.Do(attachment_key))
	if err != nil {
		return store.Error{Description: err.Error(), State: store.Error_ERROR}
	}
	return store.Error{Description: "new attachment added", State: store.Error_OK}
}

func _new_doc(client *openledge.FaunaClient) store.Error {

	_, err := client.Query(openledge.If(openledge.Exists(openledge.Documents(openledge.Collection(COLLECTION_NAME))), true, openledge.Map(openledge.Paginate(openledge.Documents(openledge.Collection(COLLECTION_NAME)), openledge.Size(Document_Size)), openledge.Lambda(DOC_NAME+"Ref", openledge.Get(openledge.Var(DOC_NAME+"Ref"))))))
	if err != nil {
		return store.Error{Description: err.Error(), State: store.Error_ERROR}
	}

	return store.Error{Description: DOC_NAME + "is created", State: store.Error_OK}
}
