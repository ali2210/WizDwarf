package structs


import (
	"net/http"
	"text/template"
	"log"
)


type Response struct {
	Flag bool
	Message string
	Links string
	cResponse http.ResponseWriter
	cRequest *http.Request
	t *template.Template
}

type clientHandleInterface interface{
	ClientHTMLRequest(file string) *template.Template
	ClientRequestHandle(flag bool, l, m string, w http.ResponseWriter, r *http.Request) *Response
	ClientLogs()
	Run() error
}


func (*Response) ClientHTMLRequest(file string)*template.Template  {

		temp := template.Must(template.ParseFiles(file+".html"))
		return temp

}

func (a *Response) ClientRequestHandle(f bool, l, m string, w http.ResponseWriter, r *http.Request) *Response{
	(*a).Flag = f
	(*a).Message = m
	(*a).Links = l
	(*a).cResponse = w
	(*a).cRequest = r

	return a

}

func (a *Response) ClientLogs(){

			log.Println("[Ok] logs ...", (*a))

}

func (a *Response) Run(t *template.Template)error{
		(*a).t = t
		err := (*a).t.Execute((*a).cResponse, (*a))
		return err
}
