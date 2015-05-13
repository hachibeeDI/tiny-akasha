package controller

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"

	"github.com/hachibeeDI/tiny-akasha/controller/question"
	"github.com/hachibeeDI/tiny-akasha/helper"
)

func v1API(uri string) string {
	return fmt.Sprintf("/api/v1%s", uri)
}

func InitRoute() {
	fmt.Println("route init")

	goji.Get("/", question.Index)

	staticFs := http.FileServer(http.Dir(helper.DirName + "/template/static"))
	goji.Get("/static/*", http.StripPrefix("/static", staticFs))

	goji.Post(v1API("/question"), question.Create)
	goji.Get(v1API("/question"), question.Get)
	goji.Get(v1API("/question/id/:id"), question.GetById)
	goji.Delete(v1API("/question/id/:id"), question.Delete)
	// goji.Get(v1API("/question/user/:username"), nil)

	// goji.Post(v1API("/question/answer"), question.Ansewer)
	// goji.Get(v1API("/question/answer"), question.Get)
}
