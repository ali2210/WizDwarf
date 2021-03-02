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
		ReadContent(pageURL string) ([]HtmlContent, error)
	}

	ParserObject struct{}
)

func (*ParserObject) ReadContent(pageURL string) ([]HtmlContent, error) {

	file, err := os.Open(pageURL)
	if err != nil {
		log.Fatalln("Error", err, file)
		return []HtmlContent{}, err
	}
	stats, err := file.Stat()
	if err != nil {
		return []HtmlContent{}, err
	}
	fmt.Println("@param_file:", stats.IsDir(), file.Name())

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("Error", err)
		return []HtmlContent{}, err
	}

	transactPage := make([]HtmlContent, 0)
	parse := [4]HtmlContent{}
	for i := range parse {
		parse[i].Typeclass = doc.Find("h3").Text()
		parse[i].Price = doc.Find("i").Text()
		parse[i].Description = doc.Find("h6").Text()
		parse[i].Paypal = doc.Find("i").Text()
		parse[i].Token = doc.Find("i").Text()
		fmt.Println("@Param:", parse[i])
		transactPage = append(transactPage, parse[i])
	}

	fmt.Println("@Param:", transactPage)

	return transactPage, nil
}
