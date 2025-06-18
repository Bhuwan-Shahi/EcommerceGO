package repository

import (
	"ecommerceGO/internal/domain"

	"gorm.io/gorm"
)

type CatalogRepository interface {
}

type catalogRepository struct {
	db *gorm.DB
}

// CreateBankAccount implements UserRepository.
func (r *userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}

func NewCatalogRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
