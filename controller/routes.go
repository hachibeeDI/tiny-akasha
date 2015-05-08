package controller

import (
	"fmt"

	"github.com/zenazn/goji"

	"github.com/hachibeeDI/tiny-akasha/controller/question"
)

func v1API(uri string) string {
	return fmt.Sprintf("/api/v1%s", uri)
}

func InitRoute() {
	fmt.Println("route init")

	goji.Get("/", question.Index)

	goji.Post(v1API("/question"), question.Create)
	goji.Get(v1API("/question"), question.Get)
	// goji.Get(v1API("/question/show"), question.Get)
	goji.Get(v1API("/question/id/:id"), question.GetById)
	goji.Delete(v1API("/question/id/:id"), question.Delete)
	// goji.Get(v1API("/question/user/:username"), nil)

	// http.Handle("/static/*", http.FileServer(http.Dir("/template/static/")))
	// goji.Post(v1API("/question/answer"), question.Ansewer)
	// goji.Get(v1API("/question/answer"), question.Get)
}
