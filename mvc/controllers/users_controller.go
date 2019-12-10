package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Djudicael/microserviceMVC/mvc/services"
	"github.com/Djudicael/microserviceMVC/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	//log.Printf("About to process user_id %v", userId)
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))

	if err != nil {
		// Just return the bad resquest to the client
		apiErr := &utils.ApplicationError{Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		//handle the err and return to the client
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	//return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
