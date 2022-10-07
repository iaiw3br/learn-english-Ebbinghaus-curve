package list

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"tg-bot-learning-english/internal/handlers"
	"tg-bot-learning-english/internal/word"
)

const (
	listsURL = "/lists"
)

type handler struct {
	listStore Store
	wordStore word.Store
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, listsURL, h.Create)
}

func NewHandler(listStore Store, wordStore word.Store) handlers.Handler {
	return &handler{
		listStore: listStore,
		wordStore: wordStore,
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cl CreateList
	err = json.Unmarshal(body, &cl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	list := ConvertCreateToList(cl)
	listID, err := h.listStore.Create(ctx, list)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now()
	for _, cw := range cl.Words {
		nw := word.Create(cw, now)
		nw.ListID = listID

		err = h.wordStore.Create(ctx, nw)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
