package user

import (
	"errors"

	"github.com/adiet95/go-order-api/src/database/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (re *user_repo) FindByEmail(email string) (*models.User, error) {
	var data *models.User
	var datas *models.Users

	res := re.db.Model(&datas).Where("email = ?", email).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("email not found")
	}
	return data, nil
}

func (re *user_repo) RegisterEmail(data *models.User) (*models.User, error) {
	var datas *models.Users

	res := re.db.Model(&datas).Where("email = ?", data.Email).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected > 0 {
		return nil, errors.New("email registered, go to login")
	}

	data.Role = "user"
	r := re.db.Create(data)
	if r.Error != nil {
		return nil, errors.New("failled to obtain data")
	}
	return data, nil
}
