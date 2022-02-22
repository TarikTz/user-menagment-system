package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TarikTz/user-menagment-system/helpers"
	"github.com/TarikTz/user-menagment-system/models"
	"github.com/julienschmidt/httprouter"
)

type Users struct {
	model models.User
}

// GetUsers ...
func (u Users) GetUsers(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var filters models.FilterUser
	var err error
	queryValues := req.URL.Query()

	filters.Limit, err = strconv.Atoi(queryValues.Get("limit"))

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	filters.Page, err = strconv.Atoi(queryValues.Get("page"))

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	filters.Order = queryValues.Get("order")
	filters.Status = queryValues.Get("status")

	if filters.Page == 0 {
		filters.Page = 1
	}
	if filters.Limit == 0 {
		filters.Limit = 10
	}

	users, count, err := models.ListUsers(filters)
	if err != nil {
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(res, http.StatusOK, map[string]interface{}{"users": users, "users_count": count})
}

// GetUser ...
func (u Users) GetUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user models.User

	userID := params.ByName("id")

	err := user.User(userID)
	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(res, http.StatusOK, user)
}

func (u Users) CreateUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	err = user.CreateUser()

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(res, http.StatusOK, user)
}

func (u Users) UpdateUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)

	userID := params.ByName("id")

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	err = user.Update(userID)

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(res, http.StatusOK, map[string]interface{}{"status": "User updated"})
}

func (u Users) DeleteUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user models.User
	userID := params.ByName("id")

	err := user.Delete(userID)

	if err != nil {
		fmt.Println(err)
		helpers.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(res, http.StatusOK, map[string]interface{}{"status": "User deleted"})
}


func (u Users) RenderAccess(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	helpers.RespondWithJSON(res, http.StatusOK, map[string]interface{}{"status": "UMS API"})
}