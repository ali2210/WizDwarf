/* This codebase desgin according to mozilla open source license.
Redistribution , contribution and improve codebase under license
convensions. @contact Ali Hassan AliMatrixCode@protonmail.com */

// package
package session

// Libraries
import (
	"net/http"

	"github.com/gorilla/sessions"
)

//  Cookies struct

type Cookies struct {
	Response http.ResponseWriter
	Request  *http.Request
	cookies  *sessions.CookieStore
}

// Cookies Services
type UsersTokenization interface {
	ExpireToken() error
	NewToken() error
	SetContextSession(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request)
	GetContextSession() Cookies
}

// ExpireToken is a helper function that will expire the user session. After that user will relogin the application
// @Receiver cookies
// @return error message
func (c *Cookies) ExpireToken() error {

	param := c.GetContextSession()

	sessID, _ := param.cookies.Get(param.Request, "session-name")
	sessID.Values["authenticated"] = false
	err := sessID.Save(param.Request, param.Response)

	if err != nil {
		panic(err.Error())
	}
	return nil
}

// Whenever user will login ; application will create user sesson.
// @Reciver Cookies
// @return error message
func (c *Cookies) NewToken() error {

	param := c.GetContextSession()

	sessID, _ := param.cookies.Get(param.Request, "session-name")
	sessID.Values["authenticated"] = true
	err := sessID.Save(param.Request, param.Response)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

// @param sessions , responsewriter, request

func (c *Cookies) SetContextSession(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	(*c).Request = r
	(*c).Response = w
	(*c).cookies = user
}

// @return cookies
func (c *Cookies) GetContextSession() Cookies {
	return *c
}
