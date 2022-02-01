package models

import (
	"github.com/egacheru/services/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Title         string  `gorm:"" json:"title"`
	Supplier      string  `json:"supplier"`
	Price         float64 `json:"price"`
	ModeOfPayment string  `json:"modeofpayment"`
	TransactionID string  `json:"transactionid"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}

func (p *Product) CreateProduct() *Product {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllProducts() []Product {
	var Products []Product
	db.Find(&Products)
	return Products
}

func GetProductById(Id int64) (*Product, *gorm.DB) {
	var getProduct Product
	db := db.Where("ID=?", Id).Find(&getProduct)
	return &getProduct, db
}

func DeleteProduct(Id int64) Product {
	var product Product
	db.Where("ID=?", Id).Delete(product)
	return product
}
