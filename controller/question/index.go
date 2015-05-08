package question

import (
	"fmt"
	"net/http"

	// "golang.org/x/net/context"
	// "github.com/zenazn/goji/web"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index called")
	fmt.Fprintf(w, "create faq success !")
}
