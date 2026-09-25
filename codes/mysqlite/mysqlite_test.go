package mysqlite

import (
	"testing"
	"CLIScraper/codes/kind"
)

func Test_sq(t *testing.T) {
	books:=make([]kind.Book,1)
	books[0].Title="ha ha ha"
	books[0].Price=150000
	books[0].UrlImage="https://example.com"

	sq(books)

}
