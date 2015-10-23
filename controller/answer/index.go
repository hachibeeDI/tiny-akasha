package answer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/hachibeeDI/tiny-akasha/model/entity"
	"github.com/hachibeeDI/tiny-akasha/model/entity/answer"
	"github.com/hachibeeDI/tiny-akasha/view/helper"
)

type NewAnswerBody struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// url_params:
//    questionId
// params:
//    question_id, name, content
func Create(c web.C, w http.ResponseWriter, r *http.Request) {
	s_question_id := c.URLParams["questionId"]
	// TODO: Content-Typeがapplication/jsonじゃない時の分岐も必要
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	var posted NewAnswerBody
	json.Unmarshal(body, &posted)
	fmt.Println(posted)
	name := posted.Name
	content := posted.Content
	question_id, err := strconv.Atoi(s_question_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "question_id type should be int"}, w)
		return
	}
	if name == "" || content == "" {
		helper.RenderJson(map[string]interface{}{"error": "empty data"}, w)
		return
	}
	db := entity.Db
	ans := answer.Init(question_id, name, content)
	if err := ans.Insert(db); err != nil {
		panic(err)
	}
	helper.RenderJson(ans, w)
}

// url_params:
//   questionId, answerId
// params:
//    username, content
func Update(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("answer Update called")
	// s_question_id := c.URLParams["questionId"]
	// question_id, err := strconv.Atoi(s_question_id)
	s_answer_id := c.URLParams["answerId"]
	answer_id, err := strconv.Atoi(s_answer_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "question_id type should be int"}, w)
		return
	}
	db := entity.Db
	ques := answer.Update(db, answer_id, r.FormValue("username"), r.FormValue("content"))
	if ques == nil {
		helper.RenderJson(map[string]interface{}{"error": "no data"}, w)
		return
	}
	obj := map[string]interface{}{"status": "success"}
	helper.RenderJson(obj, w)
}

func Get(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("answer Get called")
	db := entity.Db
	ques := answer.SelectAll(db)
	if ques == nil {
		helper.RenderJson(map[string]interface{}{"error": "no data"}, w)
		return
	}
	obj := map[string]interface{}{"answers": ques}
	helper.RenderJson(obj, w)
}

func GetByQuestionId(c web.C, w http.ResponseWriter, r *http.Request) {
	s_id := c.URLParams["questionId"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "id type should be int"}, w)
		return
	}
	answers := map[string]interface{}{"answers": answer.SelectByQuestionId(entity.Db, id)}
	helper.RenderJson(answers, w)
}

func GetById(c web.C, w http.ResponseWriter, r *http.Request) {
	s_id := c.URLParams["id"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "id type should be int"}, w)
		return
	}
	helper.RenderJson(answer.SelectById(entity.Db, id), w)
}

func DeleteByQuestionID(c web.C, w http.ResponseWriter, r *http.Request) {
	s_id := c.URLParams["answerId"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "id type should be int"}, w)
		return
	}
	err = answer.DeleteByQuestionID(entity.Db, id)
	if err != nil {
		helper.RenderJson(err, w)
		return
	}
	helper.RenderJson(map[string]interface{}{"status": "success"}, w)
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	s_id := c.URLParams["id"]
	id, err := strconv.Atoi(s_id)
	if err != nil {
		helper.RenderJson(map[string]interface{}{"error": "id type should be int"}, w)
		return
	}
	err = answer.Delete(entity.Db, id)
	if err != nil {
		helper.RenderJson(err, w)
		return
	}
	helper.RenderJson(map[string]interface{}{"status": "success"}, w)
}
