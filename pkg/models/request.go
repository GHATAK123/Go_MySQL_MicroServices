package models

import (
	"github.com/ghatu/go-requestManagement/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Request struct {
	gorm.Model
	Request      string `gorm:"" json:"request"`
	RequestedBy  string `json:"requestedBy"`
	RequestedFor string `json:"requestedFor"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Request{})

}

func (r *Request) CreateRequest() *Request {
	db.NewRecord(r)
	db.Create(&r)
	return r
}

func GetAllRequest() []Request {
	var Requests []Request
	db.Find(&Requests)
	return Requests
}

func GetRequestById(Id int64) (*Request, *gorm.DB) {
	var getRequest Request
	db := db.Where("id=?", Id).Find(&getRequest)
	return &getRequest, db
}

func DeleteRequest(ID int64) Request {
	var request Request
	db.Where("id=?", ID).Delete(request)
	return request
}

func DeleteAllRequest() []Request {
	var Requests []Request
	db.Delete(&Requests)
	return Requests
}
