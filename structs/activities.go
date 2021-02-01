package structs

// libaraies

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type RouteParameter struct {
	Response http.ResponseWriter
	Request  *http.Request
	cookies  *sessions.CookieStore
}

type UsersTokenization interface {
	ExpireToken() error
	NewToken() error
	SetContextSession(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request)
	GetContextSession() RouteParameter
}

func (p *RouteParameter) ExpireToken() error {

	param := p.GetContextSession()
	sessId, _ := param.cookies.Get(param.Request, "session-name")
	sessId.Values["authenticated"] = false
	err := sessId.Save(param.Request, param.Response)
	if err != nil {
		return err
	}
	return nil
}

func (p *RouteParameter) NewToken() error {

	param := p.GetContextSession()
	sessId, _ := param.cookies.Get(param.Request, "session-name")
	sessId.Values["authenticated"] = true
	err := sessId.Save(param.Request, param.Response)
	if err != nil {
		return err
	}
	return nil
}

func (p *RouteParameter) SetContextSession(user *sessions.CookieStore, w http.ResponseWriter, r *http.Request) {
	(*p).Request = r
	(*p).Response = w
	(*p).cookies = user
}
func (p *RouteParameter) GetContextSession() RouteParameter {
	return *p
}
