package question

import (
	"fmt"
	"net/http"

	// "github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/hachibeeDI/tiny-akasha/model/entity"
	"github.com/hachibeeDI/tiny-akasha/model/entity/question"
	"github.com/hachibeeDI/tiny-akasha/view/helper"
)

// params:
//    title, name, content
func Create(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Create called")
	title := c.URLParams["title"]
	name := c.URLParams["name"]
	content := c.URLParams["content"]
	db := entity.Db
	if err := question.Init(title, name, content).Insert(db); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "create faq success !")
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("question Get called")
	db := entity.Db
	ques := question.SelectAll(db)
	if ques == nil {
		helper.RenderJson(map[string]interface{}{"error": "no data"}, w)
		return
	}
	obj := map[string]interface{}{"questions": ques}
	// for _, que := range ques {
	// 	fmt.Fprintf(w, "%s", que)
	// }
	helper.RenderJson(obj, w)
}

func GetById(w http.ResponseWriter, r *http.Request) {
}
