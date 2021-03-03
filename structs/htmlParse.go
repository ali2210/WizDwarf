package structs

import (
	"fmt"
	"log"
	"os"
	"strings"

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
		SplitContent(content HtmlContent) []HtmlContent
	}

	ParserObject struct{}
)

func (p *ParserObject) ReadContent(pageURL string) ([]HtmlContent, error) {

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

	transactPage := make([]HtmlContent, 3)
	parse := HtmlContent{}
	parse.Typeclass = doc.Find("h3").Text()
	parse.Price = doc.Find("i").Text()
	parse.Description = doc.Find("h6").Text()
	parse.Paypal = doc.Find("i").Text()
	parse.Token = doc.Find("i").Text()
	t := p.SplitContent(parse)
	for i := range transactPage {
		transactPage = append(transactPage, t[i])
	}

	fmt.Println("@Param_Data:", t)

	return transactPage, nil
}

func (*ParserObject) SplitContent(content HtmlContent) []HtmlContent {
	parser := make([]HtmlContent, 3)
	for i := range parser {
		if strings.Contains(content.Typeclass, "Kernel") && parser[0].Typeclass == "" {
			parser[i].Typeclass = "Kernel"
			parser[i].Price = "$50"
			parser[i].Description = "2D Molecular structure"
			parser[i].Paypal = "Paypal"
			parser[i].Description = "Crypto-Tokens"
		} else {
			if strings.Contains(content.Typeclass, "Cluster") && (parser[1].Typeclass == "") {
				parser[i].Typeclass = "Cluster"
				parser[i].Price = "$100"
				parser[i].Description = "3D Molecular structure"
				parser[i].Paypal = "Paypal"
				parser[i].Token = "Crypto-Tokens"
			} else {
				if strings.Contains(content.Typeclass, "Multi-Cluster") && (parser[2].Typeclass == "") {
					parser[i].Typeclass = "Multi-Cluster"
					parser[i].Price = "$500"
					parser[i].Paypal = "Paypal"
					parser[i].Token = "Crypto-Tokens"
					parser[i].Description = "Automated argumented 4D+ Molecular structure analysis."
				}
			}
		}
	}
	return parser
}
