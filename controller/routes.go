package controller

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"

	"github.com/hachibeeDI/tiny-akasha/controller/answer"
	"github.com/hachibeeDI/tiny-akasha/controller/oauth"
	"github.com/hachibeeDI/tiny-akasha/controller/question"
	"github.com/hachibeeDI/tiny-akasha/helper"
)

func v1API(uri string) string {
	return fmt.Sprintf("/api/v1%s", uri)
}

func InitRoute() {
	fmt.Println("route init")

	goji.Get("/", question.Index)
	goji.Get("/login", question.Login)
	goji.Get("/oauth/github/callback", oauth.GithubCallback)

	goji.Get("/view/*", question.Index)

	staticFs := http.FileServer(http.Dir(helper.DirName + "/template/static"))
	goji.Get("/static/*", http.StripPrefix("/static", staticFs))

	goji.Post(v1API("/question"), question.Create)
	goji.Get(v1API("/question"), question.Get)
	goji.Get(v1API("/question/id/:id"), question.GetById)
	goji.Delete(v1API("/question/id/:id"), question.Delete)

	goji.Post(v1API("/question/search"), question.QueryByWords)

	// 解答一覧
	goji.Post(v1API("/question/id/:questionId/answer"), answer.Create)
	goji.Get(v1API("/question/id/:questionId/answer"), answer.GetByQuestionId)
	goji.Delete(v1API("/question/id/:questionId/answer"), answer.DeleteByQuestionID) // 特権つけないと危険かも

	goji.Post(v1API("/question/id/:questionId/answer/:answerId"), answer.Update)

	// 質問に紐付いた操作
	goji.Get(v1API("/answer/search/q/:id"), answer.GetByQuestionId)
	goji.Delete(v1API("/answer/search/q/:id"), answer.DeleteByQuestionID)

	// 回答そのものの操作
	goji.Delete(v1API("/answer/id/:id"), answer.Delete)

	goji.Post(v1API("/answer/search/id/:id"), answer.Update)
	goji.Get(v1API("/answer/search/id/:id"), answer.GetById)

	// goji.Post(v1API("/question/answer"), question.Ansewer)
	// goji.Get(v1API("/question/answer"), question.Get)
}
