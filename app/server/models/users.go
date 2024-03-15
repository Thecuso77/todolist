package models

type Users struct {
	Id       int `gorm:"primary_key"`
	Login    string
	Password string
}
