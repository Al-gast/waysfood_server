package models

import "time"

type User struct {
	ID           int           `json:"id" gorm:"primary_key:auto_increment"`
	Fullname     string        `json:"fullname" gorm:"type: varchar(255)"`
	Email        string        `json:"email" gorm:"type: varchar(255)"`
	Password     string        `json:"password" gorm:"type: varchar(255)"`
	Image        string        `json:"image" gorm:"type: varchar(255)"`
	Gender       string        `json:"gender" gorm:"type: varchar(255)"`
	Phone        string        `json:"phone" gorm:"type: varchar(255)"`
	Role         string        `json:"role" gorm:"type: varchar(255)"`
	Location     string        `json:"location" gorm:"type: varchar(255)"`
	Products     []ProductUser `json:"products"`
	Carts        []Cart        `json:"carts"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:BuyerID"`
	Incomes      []Transaction `json:"incomes" gorm:"foreignKey:SellerID"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
}

type UserProfile struct {
	ID       int    `json:"id"  `
	Fullname string `json:"fullname" `
	Email    string `json:"email" `
	Image    string `json:"image" `
	Gender   string `json:"gender" `
	Phone    string `json:"phone" `
	Role     string `json:"role" `
	Location string `json:"location" `
}

func (UserProfile) TableName() string {
	return "users"
}
