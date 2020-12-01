package app

import (
	"github.com/Djudicael/microserviceMVC/mvc/controllers"
)

func mapUrls() {

	// http.HandleFunc("/users", controllers.GetUser)
	router.GET("/users/:user_id", controllers.GetUser)

}
