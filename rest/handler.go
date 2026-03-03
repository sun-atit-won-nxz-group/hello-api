package rest

import (
	"hello-api/translation"
	"net/http"
)

type TranslateHandler struct {
	service translation.Translator // ฉีด Interface เข้ามาที่นี่
}

func NewTranslateHandler(service translation.Translator) *TranslateHandler {
	return &TranslateHandler{service: service}
}

func (t *TranslateHandler) TranslateHandler(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	lang := r.URL.Query().Get("language")

	res := t.service.Translate(word, lang)
	w.Write([]byte(res))
}
