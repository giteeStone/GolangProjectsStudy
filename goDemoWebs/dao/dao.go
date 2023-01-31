package dao

import (
	"golang/goDemoWebs/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	RegisterUser(user *model.User)
	Login(username string) model.User
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:960690@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
}

func (mgr *manager) RegisterUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}
