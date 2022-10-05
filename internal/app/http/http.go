package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"tg-bot-learning-english/internal/config"
	"tg-bot-learning-english/internal/word"
	"tg-bot-learning-english/pkg/client/postgresql"
)

func Run(cfg *config.Config) error {
	router := httprouter.New()
	postgresqlClient, err := postgresql.NewClient(context.Background(), cfg.Client.Postgresql)
	if err != nil {
		return err
	}

	wordStore := word.NewStore(postgresqlClient)
	wordHandler := word.NewHandler(wordStore)
	wordHandler.Register(router)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), router)
	if err != nil {
		return err
	}

	return nil
}
