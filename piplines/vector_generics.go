/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// Package or module
package piplines

import (
	"reflect"

	user "github.com/ali2210/wizdwarf/other/users/register"
)

func GetID[T user.Updated_User, U user.New_User](x *T, y *U) (*T, *U) {

	var t T
	var u U

	if !reflect.DeepEqual(x, user.Updated_User{}) {
		return x, &u
	}
	return &t, y
}
