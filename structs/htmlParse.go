package structs

import (
	"fmt"
	"log"
	"os"

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

	file, err := os.Open(pageURL)
	if err != nil {
		log.Fatalln("Error", err, file)
		return HtmlContent{}, err
	}
	stats, err := file.Stat()
	if err != nil {
		return HtmlContent{}, err
	}
	fmt.Println("@param_file:", stats.IsDir(), file.Name())

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("Error", err)
		return HtmlContent{}, err
	}

	transactPage := HtmlContent{}

	var cont string = ""
	cont = doc.Find("h3 ").Text()
	tag := doc.Find("i").Text()
	fmt.Println("@param_content:", cont, "@param_tag:", tag)

	return transactPage, nil
}
