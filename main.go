package main

import (
	"github.com/eovinicius/go-jobs/config"
	"github.com/eovinicius/go-jobs/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errf("initialization error: %v", err)
		return
	}

	router.Initialize()
}
