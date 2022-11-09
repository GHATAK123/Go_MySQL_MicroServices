package routes

import (
	"github.com/ghatu/request_management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRequestRoutes = func(route *mux.Router) {
	route.HandleFunc("/request/", controllers.CreateRequest).Methods("POST")
	route.HandleFunc("/request/", controllers.GetRequest).Methods("GET")
	route.HandleFunc("/request/{requestId}", controllers.GetRequestById).Methods("GET")
	route.HandleFunc("/request/{requestId}", controllers.UpdateRequestById).Methods("PUT")
	route.HandleFunc("/request/{requestId}", controllers.DeleteRequestById).Methods("DELETE")
	route.HandleFunc("/request/", controllers.DeleteRequest).Methods("DELETE")

}
