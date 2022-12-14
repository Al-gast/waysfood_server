package repositories

import (
	"waysfood/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	FindPartners(Role string) ([]models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Products").Preload("Carts.Order").Preload("Transactions").Find(&users).Error

	return users, err
}

func (r *repository) FindPartners(Role string) ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Products").Where(" role = ?", Role).Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Products").Preload("Carts").Preload("Transactions").First(&user, ID).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}
