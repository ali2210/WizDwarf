package structs

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type (
	HtmlContent struct {
		Typeclass   string
		Price       string
		Description string
		Token       string
		Paypal      string
	}

	HtmlContentInterface interface {
		ReadContent(pageURL string) (HtmlContent, error)
	}

	ParserObject struct{}
)

func (*ParserObject) ReadContent(pageURL string) (HtmlContent, error) {

	doc, err := goquery.NewDocument(pageURL)
	if err != nil {
		fmt.Println("Error", err)
		return HtmlContent{}, err
	}

	transactPage := HtmlContent{}

	doc.Find("typeClass").Each(func(index int, item *goquery.Selection) {
		transactPage.Description = item.Text()
	})
	doc.Find("price").Each(func(index int, item *goquery.Selection) {
		transactPage.Price = item.Text()
	})
	doc.Find("paypal").Each(func(index int, item *goquery.Selection) {
		transactPage.Paypal = item.Text()
	})
	doc.Find("token").Each(func(index int, item *goquery.Selection) {
		transactPage.Token = item.Text()
	})
	doc.Find("describe").Each(func(index int, item *goquery.Selection) {
		transactPage.Description = item.Text()
	})

	// transactPage.Typeclass = doc.Find("typeClass").Text()
	// transactPage.Price = doc.Find("price").Text()
	// transactPage.Description = doc.Find("describe").Text()
	// transactPage.Paypal = doc.Find("paypal").Text()
	// transactPage.Token = doc.Find("token").Text()

	return transactPage, nil
}
