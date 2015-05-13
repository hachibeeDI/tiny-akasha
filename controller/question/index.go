package question

import (
	"fmt"
	"io/ioutil"
	"net/http"

	// "golang.org/x/net/context"
	// "github.com/zenazn/goji/web"
	"github.com/hachibeeDI/tiny-akasha/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile(helper.DirName + "/template/index.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(content[:]))
}
