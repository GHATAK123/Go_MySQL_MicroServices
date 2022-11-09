package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ghatu/go-requestManagement/pkg/models"
	"github.com/ghatu/go-requestManagement/pkg/utils"
	"github.com/gorilla/mux"
)

var newRequest models.Request

func GetAllRequest(w http.ResponseWriter, r *http.Request) {
	newRequests := models.GetAllRequest()
	res, _ := json.Marshal(newRequests)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRequestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	ID, err := strconv.ParseInt(requestId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	requestDetails, _ := models.GetRequestById(ID)
	res, _ := json.Marshal(requestDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	CreateRequest := &models.Request{}
	utils.ParseBody(r, CreateRequest)
	b := CreateRequest.CreateRequest()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRequestById(w http.ResponseWriter, r *http.Request) {
	var updateRequest = &models.Request{}
	utils.ParseBody(r, updateRequest)
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	ID, err := strconv.ParseInt(requestId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	requestDetails, db := models.GetRequestById(ID)
	if updateRequest.Request != "" {
		requestDetails.Request = updateRequest.Request
	}
	if updateRequest.RequestedBy != "" {
		requestDetails.RequestedBy = updateRequest.RequestedBy
	}
	if updateRequest.RequestedFor != "" {
		requestDetails.RequestedFor = updateRequest.RequestedFor
	}
	db.Save(&requestDetails)
	res, _ := json.Marshal(requestDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRequestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["requestId"]
	ID, err := strconv.ParseInt(requestId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	request := models.DeleteRequest(ID)
	res, _ := json.Marshal(request)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAllRequest(w http.ResponseWriter, r *http.Request) {
	deleteRequest := models.DeleteAllRequest()
	res, _ := json.Marshal(deleteRequest)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
