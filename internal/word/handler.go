package word

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"tg-bot-learning-english/internal/handlers"
)

const (
	wordsURL = "/words"
)

type handler struct {
	wordStore Store
}

func (h handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, wordsURL, h.Create)
}

func NewHandler(wordStore Store) handlers.Handler {
	return handler{
		wordStore: wordStore,
	}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	var cw CreateWord
	err = json.Unmarshal(body, &cw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	word := Create(cw, time.Now())
	err = h.wordStore.Create(context.Background(), word)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
