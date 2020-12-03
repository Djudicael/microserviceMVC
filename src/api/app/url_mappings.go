package app

import (
	"github.com/Djudicael/microserviceMVC/src/api/controllers/polo"
	"github.com/Djudicael/microserviceMVC/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
