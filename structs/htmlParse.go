package structs

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type (
	HtmlContent struct {
		Typeclass   *goquery.Selection
		Price       *goquery.Selection
		Description *goquery.Selection
		Token       *goquery.Selection
		Paypal      *goquery.Selection
	}

	HtmlContentInterface interface {
		ReadContent(filename string) (*HtmlContent, error)
	}

	ParserObject struct{}
)

func (*ParserObject) ReadContent(filename string) (HtmlContent, error) {

	file, err := os.Open(filename)
	if err != nil {
		return HtmlContent{}, err
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return HtmlContent{}, err
	}
	transactPage := HtmlContent{}
	transactPage.Typeclass = doc.Find("typeClass")
	transactPage.Price = doc.Find("price")
	transactPage.Description = doc.Find("describe")
	transactPage.Paypal = doc.Find("paypal")
	transactPage.Token = doc.Find("token")
	fmt.Println(transactPage)
	return transactPage, nil
}
