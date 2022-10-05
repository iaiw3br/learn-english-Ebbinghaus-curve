package app

import (
	"log"

	"tg-bot-learning-english/internal/app/http"
	"tg-bot-learning-english/internal/config"
)

func Start() {
	cfg := config.GetConfig()

	//err := telegram.Run(cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	err := http.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
