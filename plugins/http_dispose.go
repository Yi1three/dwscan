package plugins

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

//识别html中get/post方法和form表单
func recognize (url string) string{
	response,err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	doc,err := goquery.NewDocumentFromReader(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	form := doc.Find("form")
	return form.Text()

}
