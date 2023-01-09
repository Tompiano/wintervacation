package model

type User struct {
	ID                int    `json:"ID"`
	UserName          string `json:"userName"`
	Password          string `json:"password"`
	Phone             string `json:"phone"`
	PersonInformation string `json:"personInformation"`
}
