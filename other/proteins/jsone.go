/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

package proteins

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ali2210/wizdwarf/other/jsonpb/jsonledit"
)

func CreateNewJSONFile(jf *jsonledit.FileDescriptor) {

	// create new json file

	// application store user genome heatmap in the app_data directory
	path, err := os.Stat("app_data/")
	if err != nil {
		log.Println(" Error file reading :", err.Error())
		return
	}

	// Application storage path
	if !path.IsDir() {
		log.Println(" No such file or directory!", err.Error())
		return
	}

	intent, err := json.MarshalIndent(jf.Molecule, "", " ")
	if err != nil {
		log.Println(" Error marshalling intent ", err.Error())
		return
	}

	err = ioutil.WriteFile("app_data/"+jf.Names+jf.Types, intent, 0644)
	if err != nil {
		log.Println(" file have no permissions", err.Error())
		return
	}

}
