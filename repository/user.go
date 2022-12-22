package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	result := entity.User{}
	response := r.db.Where("id = ?", id).Model(&entity.User{}).First(&result)

	if response.Error != nil {
		return entity.User{}, response.Error
	}

	if response.RowsAffected == 0 {
		return entity.User{}, nil
	}

	return result, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	result := entity.User{}
	response := r.db.Where("email = ?", email).Model(&entity.User{}).First(&result)

	if response.RowsAffected == 0 {
		return entity.User{}, nil
	}

	if response.Error != nil {
		fmt.Println(response.Error)
		return entity.User{}, response.Error
	}

	return result, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	result := r.db.Model(&user).Updates(user)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
