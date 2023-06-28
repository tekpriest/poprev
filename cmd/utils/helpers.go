package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func PanicOnError(err error, msg ...string) {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	if err != nil {
		log.Fatalf("%s %s", err, m)
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func GenerateString(length int) string {
	l := make([]byte, length)

	rand.Read(l)
	return base64.StdEncoding.EncodeToString(l)
}

type PaginationQuery struct {
	Size int `json:"size" default:"20"`
	Page int `json:"page" default:"1"`
}

type Pagination struct {
	TotalPages  int   `json:"totalPages"`
	TotalItems  int64 `json:"totalItems"`
	NextPage    int   `json:"nextPage"`
	CurrentPage int   `json:"currentPage"`
} //	@Name	Pagination

type Paginator interface {
	Paginate(totalItems, totalData, page, size int) Pagination
}

func Paginate(totalItems int64, totalData, page, size int) Pagination {
	nextPage := page + 1
	pages := float64(totalItems / int64(size))

	if nextPage > int(pages) {
		nextPage = int(pages)
	}

	if page-1 == 0 {
		page = 1
	}

	return Pagination{
		TotalPages:  int(pages),
		TotalItems:  int64(totalData),
		NextPage:    nextPage,
		CurrentPage: page,
	}
}

func FormatDBURL(url string) string {
	var dbURL string
	s := strings.Split(url, "/")
	creds := strings.Split(s[2], "@")[0]
	hostname := fmt.Sprintf("tcp(%s)", strings.Split(s[2], "@")[1])

	dbName := s[len(s)-1]
	dbName = strings.Split(dbName, "?")[0]

	dbURL = fmt.Sprintf("%s@%s/%s?parseTime=true", creds, hostname, dbName)

	return dbURL
}
