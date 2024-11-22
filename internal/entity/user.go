package entity

import "github.com/google/uuid"


type User struct {
	ID	uuid.UUID	`gorm:"column:id;primaryKey"`
	FirstName	string	`gorm:"column:first_name"`
	LastName	string	`gorm:"column:last_name"`
	Username	string	`gorm:"column:username;not null;unique;index"`
	Password	string	`gorm:"column:password"`
}