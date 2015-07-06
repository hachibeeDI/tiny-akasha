package question

// TODO: entityを直接読んでるのをどうにかアレする

import (
	"fmt"
	"net/http"
	"strconv"

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
	title := r.FormValue("title")
	name := r.FormValue("name")
	content := r.FormValue("content")
	if len(title) < 0 || len(content) < 0 {
		helper.RenderJson(map[string]interface{}{"error": "empty data"}, w)
		return
	}
	db := entity.Db
	if err := question.Init(title, name, content).Insert(db); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "create faq success !")
}

func Get(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("question Get called")
	db := entity.Db
	ques := question.SelectAll(db)
	if ques == nil {
		helper.RenderJson(map[string]interface{}{"error": "no data"}, w)
		return
	}
	obj := map[string]interface{}{"questions": ques}
	helper.RenderJson(obj, w)
}

func GetById(c web.C, w http.ResponseWriter, r *http.Request) {
	s_id := c.URLParams["id"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "id type should be int"}, w)
		return
	}
	helper.RenderJson(question.SelectById(entity.Db, id), w)
}

func QueryByWords(c web.C, w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	// 検索語彙がなければデフォルト表示
	if word == "" {
		Get(c, w, r)
		return
	}
	questions := question.SelectByWord(entity.Db, word)
	if questions == nil {
		helper.RenderJson(map[string]interface{}{"error": "no data matched"}, w)
		return
	}
	obj := map[string]interface{}{"questions": questions}
	helper.RenderJson(obj, w)
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	s_id := c.URLParams["id"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "id type should be int"}, w)
		return
	}
	err = question.Delete(entity.Db, id)
	if err != nil {
		helper.RenderJson(err, w)
		return
	}
	helper.RenderJson(map[string]interface{}{"status": "success"}, w)
}
