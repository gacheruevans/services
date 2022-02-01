package routers

import (
	"github.com/egacheru/services/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/products/", controllers.GetProducts).Methods("GET")
	// router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	// router.HandleFunc("/product/{id}", controllers.GetProductById).Methods("GET")
	// router.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	// router.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
}
