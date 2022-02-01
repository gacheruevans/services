package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/egacheru/services/pkg/models"
)

var NewProduct models.Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	Products := models.GetAllProducts
	res, _ := json.Marshal(Products)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetProductById(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	productId := params["id"]
// 	ID, err := strconv.ParseInt(productId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	productDetails, _ := models.GetProductById(ID)
// 	res, _ := json.Marshal(productDetails)
// 	w.Header().Set("content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func CreateProduct(w http.ResponseWriter, r *http.Request) {
// 	CreateProduct := &models.Product{}
// 	utils.ParseBody(r, CreateProduct)
// 	p := CreateProduct.CreateProduct()
// 	res, _ := json.Marshal(p)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	productId := params["id"]
// 	ID, err := strconv.ParseInt(productId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	product := models.DeleteProduct(ID)
// 	res, _ := json.Marshal(product)
// 	w.Header().Set("Contetn-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	var updateProduct = &models.Product{}
// 	utils.ParseBody(r, updateProduct)
// 	params := mux.Vars(r)
// 	productId := params["id"]
// 	ID, err := strconv.ParseInt(productId, 0, 0)
// 	if err != nil {
// 		fmt.Println("error while parsing")
// 	}
// 	prodDetails, db := models.GetProductById(ID)
// 	if updateProduct.Title != "" {
// 		prodDetails.Title = updateProduct.Title
// 	}
// 	if updateProduct.Supplier != "" {
// 		prodDetails.Supplier = updateProduct.Supplier
// 	}
// 	if updateProduct.Price != 0 {
// 		prodDetails.Price = updateProduct.Price
// 	}
// 	if updateProduct.ModeOfPayment != "" {
// 		prodDetails.ModeOfPayment = updateProduct.ModeOfPayment
// 	}
// 	if updateProduct.TransactionID != "" {
// 		prodDetails.TransactionID = updateProduct.TransactionID
// 	}
// 	db.Save(&prodDetails)
// 	res, _ := json.Marshal(prodDetails)
// 	w.Header().Set("Content-type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
