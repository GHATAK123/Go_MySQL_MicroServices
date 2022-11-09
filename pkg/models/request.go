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
