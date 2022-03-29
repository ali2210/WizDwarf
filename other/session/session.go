package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Cookies struct {
	Response http.ResponseWriter
	Request  *http.Request
	cookies  *sessions.CookieStore
}

type UsersTokenization interface {
	ExpireToken() error
	NewToken() error
	SetContextSession(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request)
	GetContextSession() Cookies
}

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

func (c *Cookies) SetContextSession(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	(*c).Request = r
	(*c).Response = w
	(*c).cookies = user
}
func (c *Cookies) GetContextSession() Cookies {
	return *c
}
