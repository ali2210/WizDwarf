package structs

// libaraies

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type RouteParameter struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type UsersTokenization interface {
	ExpireToken(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) error
	NewToken(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) error
}

func (p *RouteParameter) ExpireToken(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) error {
	(*p).Request = r
	(*p).Response = w
	sessId, _ := user.Get((*p).Request, "session-name")
	sessId.Values["authenticated"] = false
	err := sessId.Save((*p).Request, (*p).Response)
	if err != nil {
		return err
	}
	return nil
}

func (p *RouteParameter) NewToken(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) error {
	(*p).Request = r
	(*p).Response = w
	sessId, _ := user.Get((*p).Request, "session-name")
	sessId.Values["authenticated"] = true
	err := sessId.Save((*p).Request, (*p).Response)
	if err != nil {
		return err
	}
	return nil
}
