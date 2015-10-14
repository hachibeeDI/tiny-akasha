package question

// TODO: entityを直接読んでるのをどうにかアレする

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	// "github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/hachibeeDI/tiny-akasha/model/entity"
	"github.com/hachibeeDI/tiny-akasha/model/entity/question"
	"github.com/hachibeeDI/tiny-akasha/view/helper"
)

type NewArticleBody struct {
	Title   string `json:title`
	Name    string `json:name`
	Content string `json:content`
}

// params:
//    title, name, content
func Create(c web.C, w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var posted NewArticleBody
	json.Unmarshal(body, &posted)
	fmt.Println("Create called")
	fmt.Println(string(body))
	if len(posted.Title) < 0 || len(posted.Content) < 0 {
		helper.RenderJson(map[string]interface{}{"error": "empty data"}, w)
		return
	}
	db := entity.Db
	if err := question.Init(posted.Title, posted.Name, posted.Content).Insert(db); err != nil {
		panic(err)
	}
	// TODO: return created id
	w.WriteHeader(http.StatusCreated)
	helper.RenderJson(map[string]interface{}{"status": "success"}, w)
}

func Get(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("question Get called")
	db := entity.Db
	ques := question.SelectAll(db)
	if ques == nil {
		w.WriteHeader(http.StatusNoContent)
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
