package main

import (
	"log"

	"github.com/dzemildupljak/go_app_tools/internal/presentation"
	"github.com/dzemildupljak/go_app_tools/utils"
)

func main() {
	utils.InitLogger()

	server := presentation.NewServer()
	log.Fatal(server.ListenAndServe())
}
