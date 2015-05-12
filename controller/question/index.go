package question

import (
	"fmt"
	"io/ioutil"
	"net/http"

	// "golang.org/x/net/context"
	// "github.com/zenazn/goji/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./template/index.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(content[:]))
}
