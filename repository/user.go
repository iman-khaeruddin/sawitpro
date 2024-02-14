package repository

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sawitpro/dto"
	"sawitpro/entity"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{
		db: db,
	}
}

type UserInterface interface {
	FindByPhoneNumberPassword(ctx context.Context, phoneNumber string) (entity.User, error)
	AddLoginAttempt(ctx context.Context, phoneNumber string) error
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindByID(ctx context.Context, ID int) (dto.ProfileResponse, error)
	UpdateSelectedFields(ctx context.Context, maintenanceTrx *entity.User, fields ...string) (*entity.User, error)
}

func (repo User) FindByPhoneNumberPassword(ctx context.Context, phoneNumber string) (entity.User, error) {
	var user entity.User
	err := repo.db.WithContext(ctx).
		Raw("SELECT id, full_name, phone_number, password FROM users WHERE phone_number = ?", phoneNumber).
		Scan(&user).
		Error
	return user, err
}

func (repo User) AddLoginAttempt(ctx context.Context, phoneNumber string) error {
	err := repo.db.WithContext(ctx).Exec("UPDATE users SET login_attempt = login_attempt + 1 WHERE phone_number = ?", phoneNumber).Error
	return err
}

func (repo User) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := repo.db.WithContext(ctx).Model(&entity.User{}).Omit(clause.Associations).Create(user).Error
	return user, err
}

func (repo User) FindByID(ctx context.Context, ID int) (dto.ProfileResponse, error) {
	var user dto.ProfileResponse
	err := repo.db.WithContext(ctx).
		Raw("SELECT id, full_name, phone_number FROM users WHERE id = ?", ID).
		Scan(&user).
		Error
	return user, err
}

func (repo User) UpdateSelectedFields(ctx context.Context, user *entity.User, fields ...string) (*entity.User, error) {
	err := repo.db.WithContext(ctx).Model(user).Select(fields).Updates(*user).Error
	return user, err
}
