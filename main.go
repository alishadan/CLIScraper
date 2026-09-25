package main
import (
	"github.com/alishadan/sscraper"
	"CLIScraper/codes/goquery"
	"CLIScraper/codes/mysqlite"

)
func main(){
	url:="https://www.noon.com/uae-en/books/health-and-personal-development/mind-body-and-spirit/mind-body-spirit-self-help/?f%5BisCarousel%5D=true"
	body,err:=sscraper.Myhttp(url)
	if err!=nil{
		panic("happend error in connect to site \n")
	}
	defer body.Close()
	
	Books:=goquery.MyQuery(body,url)
	if Books!=nil{

		mysqlite.Sq(Books)

	}

}