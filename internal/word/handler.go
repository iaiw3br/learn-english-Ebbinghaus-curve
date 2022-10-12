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
	wordService Service
}

func (h handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, wordsURL, h.Create)
}

func NewHandler(wordService Service) handlers.Handler {
	return handler{
		wordService: wordService,
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

	ctx := context.Background()
	now := time.Now()

	err = h.wordService.Create(ctx, cw, now)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
