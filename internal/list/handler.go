package list

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"tg-bot-learning-english/internal/handlers"
)

const (
	listsURL           = "/lists"
	listsRepetitionURL = listsURL + "/repetition"
)

const (
	intervalRepeatHours = 6
)

type handler struct {
	listService Service
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, listsURL, h.Create)
	router.HandlerFunc(http.MethodGet, listsRepetitionURL, h.Repeat)
}

func NewHandler(listService Service) handlers.Handler {
	return &handler{
		listService: listService,
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

	err = h.listService.Create(cl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) Repeat(w http.ResponseWriter, _ *http.Request) {
	dateRepeat := time.Now().Add(time.Hour * intervalRepeatHours)

	list, err := h.listService.Repeat(dateRepeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	listBytes, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(listBytes)
}
